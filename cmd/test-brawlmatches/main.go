package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// Match représente un match récupéré depuis BrawlTools.
type Match struct {
	MatchID   string     `json:"matchId"`
	Scores    []int      `json:"scores"`
	Legends   [][]string `json:"legends"`
	Maps      []string   `json:"maps"`
	StartTime int64      `json:"startTime"`
}

// MatchResponse représente la réponse de l'API.
type MatchResponse struct {
	PlayerMatches []Match `json:"playerMatches"`
}

// LegendStat représente les statistiques d'un personnage.
type LegendStat struct {
	Name    string
	Picks   int
	Sets    int
	Wins    int
	Losses  int
	WinRate float64
}

// getConstantLegend vérifie qu'un joueur a utilisé le même personnage
// pendant toutes les parties du set.
func getConstantLegend(legends []string) (string, bool) {
	if len(legends) == 0 {
		return "", false
	}

	firstLegend := strings.ToUpper(strings.TrimSpace(legends[0]))

	if firstLegend == "" {
		return "", false
	}

	for _, legend := range legends {
		name := strings.ToUpper(strings.TrimSpace(legend))

		if name == "" || name != firstLegend {
			return "", false
		}
	}

	return firstLegend, true
}

func main() {
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	// Identifiant du tournoi Winter Championship Europe 2026.
	tournamentID := "a253e32c-0942-40d7-6d0a-08de30f37986"

	// Les cinq joueurs récupérés dans le Power Ranking EU.
	playerIDs := []string{
		"147982",
		"75405",
		"169858",
		"79853",
		"169211",
	}

	// On utilise une map pour éviter les doublons de matchs.
	uniqueMatches := make(map[string]Match)

	// Récupération des matchs de chaque joueur.
	for _, playerID := range playerIDs {
		params := url.Values{}
		params.Set("tournamentId", tournamentID)
		params.Set("playerIds", playerID)

		apiURL := "https://api.brawltools.com/v2/player/match?" + params.Encode()

		response, err := client.Get(apiURL)
		if err != nil {
			fmt.Println("Erreur lors de la récupération :", err)
			return
		}

		if response.StatusCode != http.StatusOK {
			response.Body.Close()
			fmt.Println("Erreur API :", response.Status)
			return
		}

		var result MatchResponse

		err = json.NewDecoder(response.Body).Decode(&result)
		response.Body.Close()

		if err != nil {
			fmt.Println("Erreur JSON :", err)
			return
		}

		// On conserve uniquement les matchs qui ne sont pas déjà présents.
		for _, match := range result.PlayerMatches {
			if _, exists := uniqueMatches[match.MatchID]; !exists {
				uniqueMatches[match.MatchID] = match
			}
		}
	}

	// Préparation des statistiques.
	totalMatches := len(uniqueMatches)
	matchesWithLegends := 0
	usableSets := 0
	totalPicks := 0

	legendCounts := make(map[string]int)
	legendStats := make(map[string]*LegendStat)

	// Analyse de chaque match.
	for _, match := range uniqueMatches {
		if len(match.Legends) < 2 {
			continue
		}

		// Comptage des picks, même pour les sets avec changement de personnage.
		for _, playerLegends := range match.Legends[:2] {
			for _, legend := range playerLegends {
				name := strings.ToUpper(strings.TrimSpace(legend))

				if name != "" {
					legendCounts[name]++
					totalPicks++
				}
			}
		}

		// Vérification des personnages utilisés par les deux joueurs.
		legend1, player1Valid := getConstantLegend(match.Legends[0])
		legend2, player2Valid := getConstantLegend(match.Legends[1])

		if player1Valid && player2Valid {
			matchesWithLegends++
		}

		// Pour calculer le winrate, il faut deux personnages constants
		// et un score final valide pour les deux joueurs.
		if !player1Valid || !player2Valid || len(match.Scores) != 2 {
			continue
		}

		// Un score égal ne permet pas de déterminer un vainqueur.
		if match.Scores[0] == match.Scores[1] {
			continue
		}

		usableSets++

		// Création des statistiques si nécessaire.
		if legendStats[legend1] == nil {
			legendStats[legend1] = &LegendStat{Name: legend1}
		}

		if legendStats[legend2] == nil {
			legendStats[legend2] = &LegendStat{Name: legend2}
		}

		// Chaque personnage participe à un set.
		legendStats[legend1].Sets++
		legendStats[legend2].Sets++

		// Attribution de la victoire et de la défaite.
		if match.Scores[0] > match.Scores[1] {
			legendStats[legend1].Wins++
			legendStats[legend2].Losses++
		} else {
			legendStats[legend2].Wins++
			legendStats[legend1].Losses++
		}
	}

	// Calcul du winrate de chaque personnage.
	var winrateStats []LegendStat

	for _, stat := range legendStats {
		if stat.Sets > 0 {
			stat.WinRate = float64(stat.Wins) / float64(stat.Sets) * 100
		}

		winrateStats = append(winrateStats, *stat)
	}

	// Tri par nombre de victoires, puis par nom.
	sort.Slice(winrateStats, func(i, j int) bool {
		if winrateStats[i].WinRate == winrateStats[j].WinRate {
			return winrateStats[i].Name < winrateStats[j].Name
		}

		return winrateStats[i].WinRate > winrateStats[j].WinRate
	})

	// Tri des personnages par nombre de picks.
	var pickStats []LegendStat

	for name, picks := range legendCounts {
		pickStats = append(pickStats, LegendStat{
			Name:  name,
			Picks: picks,
		})
	}

	sort.Slice(pickStats, func(i, j int) bool {
		if pickStats[i].Picks == pickStats[j].Picks {
			return pickStats[i].Name < pickStats[j].Name
		}

		return pickStats[i].Picks > pickStats[j].Picks
	})

	// Affichage du rapport.
	fmt.Println("======================================")
	fmt.Println("     RAPPORT BRAWLHALLA - 2026")
	fmt.Println("======================================")

	fmt.Println("Nombre total de matchs :", totalMatches)
	fmt.Println("Matchs avec personnages :", matchesWithLegends)
	fmt.Println("Nombre total de picks :", totalPicks)
	fmt.Println("Sets exploitables :", usableSets)
	fmt.Println("Sets exclus :", totalMatches-usableSets)

	fmt.Println("--------------------------------------")
	fmt.Println("       UTILISATION DES PERSONNAGES")
	fmt.Println("--------------------------------------")

	for _, stat := range pickStats {
		fmt.Printf("%-15s : %d picks\n", stat.Name, stat.Picks)
	}

	fmt.Println("--------------------------------------")
	fmt.Println("       WINRATE PAR PERSONNAGE")
	fmt.Println("--------------------------------------")

	for _, stat := range winrateStats {
		fmt.Printf(
			"%-15s : %.2f%% (%d V / %d D) - %d sets\n",
			stat.Name,
			stat.WinRate,
			stat.Wins,
			stat.Losses,
			stat.Sets,
		)
	}

	fmt.Println("======================================")
}
