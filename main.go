package main

import (
	"log"
	"net/http"

	"brawlhalla-tierlist/handlers"
)

func main() {

	// Mise à disposition des fichiers CSS, JS et images.
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	// Gestion de la page d'accueil.
	http.HandleFunc("/", handlers.Home)

	// Démarrage du serveur.
	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
