
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"brawlhalla-tierlist/models"
)

// Structure représentant la réponse de l'API.
type legendsResponse struct {
	Legends    []models.Legend `json:"legends"`
	TotalPages int             `json:"total_pages"`
}

// GetLegends récupère tous les personnages de Brawlhalla.
func GetLegends() ([]models.Legend, error) {

	// Création du client HTTP avec un délai maximum de 10 secondes.
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	var allLegends []models.Legend
	totalPages := 1

	// Parcours de toutes les pages de l'API.
	for page := 1; page <= totalPages; page++ {

		// Demande jusqu'à 100 personnages par page.
		url := fmt.Sprintf(
			"https://api.brawlhalla.com/v1/static/legends?page=%d&max_results=100",
			page,
		)

		// Envoi de la requête HTTP.
		response, err := client.Get(url)
		if err != nil {
			return nil, fmt.Errorf("erreur API : %w", err)
		}

		// Vérification du statut HTTP.
		if response.StatusCode != http.StatusOK {
			response.Body.Close()
			return nil, fmt.Errorf("erreur serveur : %s", response.Status)
		}

		// Décodage des données JSON.
		var result legendsResponse

		err = json.NewDecoder(response.Body).Decode(&result)
		response.Body.Close()

		if err != nil {
			return nil, fmt.Errorf("erreur JSON : %w", err)
		}

		// Récupération du nombre total de pages.
		if page == 1 {
			totalPages = result.TotalPages
		}

		// Ajout des personnages récupérés à notre liste.
		allLegends = append(allLegends, result.Legends...)
	}

	return allLegends, nil
}
