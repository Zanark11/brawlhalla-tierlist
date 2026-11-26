package handlers

import (
	"html/template"
	"log"
	"net/http"

	"brawlhalla-tierlist/api"
	"brawlhalla-tierlist/models"
	"brawlhalla-tierlist/services"
)

// Données transmises au template HTML.
type HomeData struct {
	Rankings []models.LegendRanking
}

// Template de la page d'accueil.
var homeTemplate = template.Must(
	template.ParseFiles("templates/index.html"),
)

// Home affiche la page d'accueil avec le classement.
func Home(w http.ResponseWriter, r *http.Request) {

	// Vérification de l'URL.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Récupération des statistiques officielles.
	legends, err := api.GetLegends()
	if err != nil {
		http.Error(w, "Erreur API Brawlhalla", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Récupération des statistiques Brawlmance.
	stats, err := api.GetLegendStats()
	if err != nil {
		http.Error(w, "Erreur API Brawlmance", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Association des statistiques officielles aux statistiques Brawlmance.
	officialLegends := make(map[int]models.Legend)

	for _, legend := range legends {
		officialLegends[legend.ID] = legend
	}

	for i := range stats {
		if official, exists := officialLegends[stats[i].ID]; exists {
			stats[i].Legend = official
		}
	}

	// Calcul du classement.
	rankings := services.CalculateRanking(stats)

	// Préparation des données pour le HTML.
	data := HomeData{
		Rankings: rankings,
	}

	// Configuration de la réponse HTML.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Affichage du template.
	err = homeTemplate.Execute(w, data)
	if err != nil {
		log.Println("Erreur d'affichage :", err)
	}
}
