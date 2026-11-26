package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Structure représentant la réponse de l'API.
type MatchResponse struct {
	Matches json.RawMessage `json:"matches"`
}

func main() {
	// Identifiant du Winter Championship Europe 2026.
	tournamentID := "a253e32c-0942-40d7-6d0a-08de30f37986"

	// Identifiants des 5 joueurs récupérés.
	playerIDs := []string{
		"147982",
		"75405",
		"169858",
		"79853",
		"169211",
	}

	// Prépare les paramètres de la requête.
	params := url.Values{}
	params.Set("tournamentId", tournamentID)
	params.Set("playerIds", strings.Join(playerIDs, ","))

	// Construit l'URL de l'API.
	apiURL := "https://api.brawltools.com/v2/player/match?" + params.Encode()

	client := http.Client{Timeout: 15 * time.Second}

	// Envoie la requête.
	response, err := client.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println("Statut HTTP :", response.Status)

	// Vérifie la réponse.
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Erreur API : %s", response.Status)
	}

	// Affiche les données récupérées.
	var result json.RawMessage

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))
}