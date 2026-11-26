
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"brawlhalla-tierlist/models"
)

// Structure représentant la réponse de BrawlTools.
type tournamentResponse struct {
	Tournaments []models.Tournament `json:"tournaments"`
	NextToken   string              `json:"nextToken"`
}

// GetEUTournaments récupère les tournois officiels européens en 1v1.
func GetEUTournaments(year int) ([]models.Tournament, error) {

	client := http.Client{
		Timeout: 15 * time.Second,
	}

	var allTournaments []models.Tournament
	nextToken := ""

	// Récupération de toutes les pages disponibles.
	for {
		params := url.Values{}

		params.Set("gameMode", "1")
		params.Set("isOfficial", "true")
		params.Set("year", fmt.Sprintf("%d", year))
		params.Set("maxResults", "50")

		// Récupération de la page suivante si nécessaire.
		if nextToken != "" {
			params.Set("nextToken", nextToken)
		}

		apiURL := "https://api.brawltools.com/v2/event?" + params.Encode()

		// Envoi de la requête HTTP.
		response, err := client.Get(apiURL)
		if err != nil {
			return nil, fmt.Errorf("erreur API BrawlTools : %w", err)
		}

		if response.StatusCode != http.StatusOK {
			response.Body.Close()
			return nil, fmt.Errorf("erreur serveur : %s", response.Status)
		}

		// Conversion du JSON en structure Go.
		var result tournamentResponse

		err = json.NewDecoder(response.Body).Decode(&result)
		response.Body.Close()

		if err != nil {
			return nil, fmt.Errorf("erreur JSON : %w", err)
		}

		// Conservation uniquement des tournois européens officiels en 1v1.
		// Stocke les identifiants déjà récupérés pour éviter les doublons.
		seenTournaments := make(map[string]bool)

		for _, tournament := range result.Tournaments {
		name := strings.ToLower(tournament.TournamentName)

		// Vérifie que le tournoi correspond à nos critères.
		if strings.Contains(name, "europe") &&
		strings.Contains(name, "1v1") &&
		tournament.IsOfficial {

		// Ajoute uniquement les tournois jamais récupérés.
		if !seenTournaments[tournament.ID] {
			allTournaments = append(allTournaments, tournament)
			seenTournaments[tournament.ID] = true
		}
	}
}

		// Arrêt lorsqu'il n'y a plus de page suivante.
		if result.NextToken == "" || result.NextToken == nextToken {
			break
		}

		nextToken = result.NextToken
	}

	return allTournaments, nil
}
