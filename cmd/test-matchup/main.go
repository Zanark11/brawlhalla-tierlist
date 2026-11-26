package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Match représente une confrontation entre deux joueurs.
type Match struct {
	MatchID string     `json:"matchId"`
	Scores  []int      `json:"scores"`
	Legends [][]string `json:"legends"`
	Maps    []string   `json:"maps"`
}

// MatchupResponse représente la réponse de l'API.
type MatchupResponse struct {
	MatchupMatches []Match `json:"matchupMatches"`
}

func main() {
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	// Identifiant du tournoi Winter Championship Europe 2026.
	tournamentID := "a253e32c-0942-40d7-6d0a-08de30f37986"

	// Deux joueurs de notre Power Ranking EU.
	player1 := "147982" // Zivai
	player2 := "75405"  // Ahmet_

	// Préparation des paramètres de la requête.
	params := url.Values{}
	params.Set("tournamentId", tournamentID)
	params.Set("entrant1PlayerIds", player1)
	params.Set("entrant2PlayerIds", player2)

	apiURL := "https://api.brawltools.com/v2/matchup/match?" + params.Encode()

	// Envoi de la requête à BrawlTools.
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

	// Décodage de la réponse JSON.
	var result MatchupResponse

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("Erreur JSON :", err)
		return
	}

	// Affichage des résultats.
	fmt.Println("======================================")
	fmt.Println("       TEST MATCHUP BRAWLHALLA")
	fmt.Println("======================================")
	fmt.Println("Joueur 1 : Zivai")
	fmt.Println("Joueur 2 : Ahmet_")
	fmt.Println("Nombre de confrontations :", len(result.MatchupMatches))

	for i, match := range result.MatchupMatches {
		fmt.Println("--------------------------------------")
		fmt.Println("Match :", i+1)
		fmt.Println("ID :", match.MatchID)
		fmt.Println("Scores :", match.Scores)
		fmt.Println("Personnages :", match.Legends)
		fmt.Println("Cartes :", match.Maps)
	}

	fmt.Println("======================================")
}
