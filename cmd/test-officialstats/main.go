package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"brawlhalla-tierlist/api"
)

// LegendStat représente les statistiques d'un personnage.
type LegendStat struct {
	LegendID int `json:"legend_id"`
	Games    int `json:"games"`
	Wins     int `json:"wins"`
}

// PlayerStats représente les statistiques d'un joueur.
type PlayerStats struct {
	BrawlhallaID int          `json:"brawlhalla_id"`
	Name         string       `json:"name"`
	Games        int          `json:"games"`
	Wins         int          `json:"wins"`
	Region       string       `json:"region"`
	Legends      []LegendStat `json:"legends"`
}

func main() {
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	// Identifiant Brawlhalla de Zivai.
	playerID := "5156845"

	params := url.Values{}
	params.Set("brawlhalla_id", playerID)
	params.Set("mode", "ranked_1v1")

	apiURL := "https://api.brawlhalla.com/v1/player/stats?" + params.Encode()

	// Récupération des statistiques du joueur.
	response, err := client.Get(apiURL)
	if err != nil {
		fmt.Println("Erreur lors de la récupération :", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Erreur API :", response.Status)
		return
	}

	var player PlayerStats

	err = json.NewDecoder(response.Body).Decode(&player)
	if err != nil {
		fmt.Println("Erreur JSON :", err)
		return
	}

	// Récupération de la liste officielle des personnages.
	allLegends, err := api.GetLegends()
	if err != nil {
		fmt.Println("Erreur récupération des personnages :", err)
		return
	}

	// Création d'une correspondance entre ID et nom.
	legendNames := make(map[int]string)

	for _, legend := range allLegends {
		legendNames[legend.ID] = legend.Name
	}

	// Affichage du rapport.
	fmt.Println("======================================")
	fmt.Println("    STATISTIQUES OFFICIELLES")
	fmt.Println("======================================")
	fmt.Println("Joueur :", player.Name)
	fmt.Println("Brawlhalla ID :", player.BrawlhallaID)
	fmt.Println("Région :", player.Region)
	fmt.Println("Parties jouées :", player.Games)
	fmt.Println("Victoires :", player.Wins)

	fmt.Println("--------------------------------------")
	fmt.Println("     STATISTIQUES PAR PERSONNAGE")
	fmt.Println("--------------------------------------")

	for _, legend := range player.Legends {
		name := legendNames[legend.LegendID]

		if name == "" {
			name = fmt.Sprintf("ID %d", legend.LegendID)
		}

		winRate := 0.0

		if legend.Games > 0 {
			winRate = float64(legend.Wins) / float64(legend.Games) * 100
		}

		fmt.Printf(
			"%-15s | %3d parties | %3d victoires | %.2f%%\n",
			name,
			legend.Games,
			legend.Wins,
			winRate,
		)
	}

	fmt.Println("======================================")
}
