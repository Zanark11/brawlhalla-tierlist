package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"brawlhalla-tierlist/api"
)

// Player représente un joueur récupéré depuis BrawlTools.
type Player struct {
	PlayerID    int    `json:"playerId"`
	BrawlhallaID int   `json:"brawlhallaId"`
	Name        string `json:"name"`
}

// PlayersResponse représente la réponse de BrawlTools.
type PlayersResponse struct {
	Players []Player `json:"players"`
}

// LegendStat représente les statistiques d'un personnage.
type LegendStat struct {
	LegendID int `json:"legend_id"`
	Games    int `json:"games"`
	Wins     int `json:"wins"`
}

// PlayerStats représente les statistiques officielles d'un joueur.
type PlayerStats struct {
	Name    string       `json:"name"`
	Games   int          `json:"games"`
	Wins    int          `json:"wins"`
	Region  string       `json:"region"`
	Legends []LegendStat `json:"legends"`
}

// getBrawlhallaID récupère l'identifiant officiel d'un joueur.
func getBrawlhallaID(client *http.Client, playerID int) (Player, error) {
	apiURL := fmt.Sprintf(
		"https://api.brawltools.com/v2/players?playerIds=%d",
		playerID,
	)

	response, err := client.Get(apiURL)
	if err != nil {
		return Player{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Player{}, fmt.Errorf("erreur BrawlTools : %s", response.Status)
	}

	var result PlayersResponse

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return Player{}, err
	}

	if len(result.Players) == 0 {
		return Player{}, fmt.Errorf("joueur introuvable")
	}

	return result.Players[0], nil
}

// getPlayerStats récupère les statistiques Ranked 1v1.
func getPlayerStats(client *http.Client, brawlhallaID int) (PlayerStats, error) {
	params := url.Values{}
	params.Set("brawlhalla_id", fmt.Sprint(brawlhallaID))
	params.Set("mode", "ranked_1v1")

	apiURL := "https://api.brawlhalla.com/v1/player/stats?" + params.Encode()

	response, err := client.Get(apiURL)
	if err != nil {
		return PlayerStats{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return PlayerStats{}, fmt.Errorf("erreur API officielle : %s", response.Status)
	}

	var result PlayerStats

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return PlayerStats{}, err
	}

	return result, nil
}

func main() {
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	// Les cinq joueurs de notre classement EU.
	playerIDs := []int{
		147982,
		75405,
		169858,
		79853,
		169211,
	}

	// Récupération de tous les personnages officiels.
	allLegends, err := api.GetLegends()
	if err != nil {
		fmt.Println("Erreur récupération des personnages :", err)
		return
	}

	// Association entre ID et nom du personnage.
	legendNames := make(map[int]string)

	for _, legend := range allLegends {
		legendNames[legend.ID] = legend.Name
	}

	fmt.Println("======================================")
	fmt.Println("  STATISTIQUES DES 5 JOUEURS EU")
	fmt.Println("======================================")

	// Parcours des cinq joueurs.
	for _, playerID := range playerIDs {
		player, err := getBrawlhallaID(&client, playerID)
		if err != nil {
			fmt.Println("Erreur joueur", playerID, ":", err)
			continue
		}

		stats, err := getPlayerStats(&client, player.BrawlhallaID)
		if err != nil {
			fmt.Println("Erreur statistiques", player.Name, ":", err)
			continue
		}

		fmt.Println("--------------------------------------")
		fmt.Println("Joueur :", player.Name)
		fmt.Println("Nom officiel :", stats.Name)
		fmt.Println("Brawlhalla ID :", player.BrawlhallaID)
		fmt.Println("Région :", stats.Region)
		fmt.Println("Parties :", stats.Games)
		fmt.Println("Victoires :", stats.Wins)

		fmt.Println("Personnages :")

		for _, legend := range stats.Legends {
			name := legendNames[legend.LegendID]

			if name == "" {
				name = fmt.Sprintf("ID %d", legend.LegendID)
			}

			winRate := 0.0

			if legend.Games > 0 {
				winRate = float64(legend.Wins) / float64(legend.Games) * 100
			}

			fmt.Printf(
				"  %-12s | %3d parties | %3d victoires | %.2f%%\n",
				name,
				legend.Games,
				legend.Wins,
				winRate,
			)
		}
	}

	fmt.Println("======================================")
}
