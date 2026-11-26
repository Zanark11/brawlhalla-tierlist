
package models

// Tournament représente un tournoi récupéré depuis BrawlTools.
type Tournament struct {
	ID            string `json:"id"`
	TournamentName string `json:"tournamentName"`
	Year          int    `json:"year"`
	IsOfficial    bool   `json:"isOfficial"`
	IsTwos        bool   `json:"isTwos"`
	StartTime     int64  `json:"startTime"`
	Host          string `json:"host"`
}
