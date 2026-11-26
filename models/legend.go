
package models

// Legend représente un personnage de Brawlhalla.
type Legend struct {
	ID        int    `json:"legend_id"`
	Name      string `json:"bio_name"`
	Weapon1   string `json:"weapon_one"`
	Weapon2   string `json:"weapon_two"`
	Strength  int    `json:"strength"`
	Dexterity int    `json:"dexterity"`
	Defense   int    `json:"defense"`
	Speed     int    `json:"speed"`
}
