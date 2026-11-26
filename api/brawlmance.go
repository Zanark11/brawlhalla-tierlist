package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"brawlhalla-tierlist/models"
)

// Configuration des statistiques utilisées pour notre tierlist.
const (
	BrawlmanceBaseURL = "https://brawlmance.com/api/v1"
	SelectedPatch     = "11.01"
	SelectedTier      = "Diamond"
)

// GetLegendStats récupère les statistiques de toutes les légendes.
func GetLegendStats() ([]models.LegendWithStats, error) {

	// Création du client HTTP avec un délai maximum de 15 secondes.
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	// Construction de l'URL avec les filtres souhaités.
	url := fmt.Sprintf(
		"%s/legends?patch=%s&tier=%s",
		BrawlmanceBaseURL,
		SelectedPatch,
		SelectedTier,
	)

	// Envoi de la requête HTTP.
	response, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erreur API Brawlmance : %w", err)
	}
	defer response.Body.Close()

	// Vérification du statut HTTP.
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"erreur serveur Brawlmance : %s",
			response.Status,
		)
	}

	// Décodage des données JSON.
	var result models.LegendStatsResponse

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("erreur JSON Brawlmance : %w", err)
	}

	return result.Legends, nil
}
