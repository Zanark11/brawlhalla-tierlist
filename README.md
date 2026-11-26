# Brawlhalla Tier List

## Présentation

Brawlhalla Tier List est un site web permettant de classer les légendes du jeu **Brawlhalla** en fonction de leurs statistiques.

Le classement repose sur plusieurs critères : le taux de victoire, le taux de sélection et les performances des armes.

L'objectif est de proposer un classement dynamique permettant de comparer les différentes légendes à partir de données statistiques réelles.

## Fonctionnalités

- Affichage des 71 légendes de Brawlhalla.
- Classement automatique des légendes selon un score global.
- Recherche d'une légende par son nom.
- Filtrage des légendes par arme.
- Affichage des statistiques principales : taux de victoire, taux de sélection et nombre de parties.
- Affichage des performances individuelles des deux armes.
- Affichage des statistiques officielles des légendes : force, dextérité, défense et vitesse.
- Affichage et masquage des statistiques détaillées.
- Interface responsive avec thème sombre.

## Technologies utilisées

### Backend
- **Go** : récupération, traitement et calcul des statistiques.
- **API officielle Brawlhalla** : récupération des informations des légendes.
- **API Brawlmance** : récupération des statistiques de jeu.

### Frontend
- **HTML5** : structure des pages.
- **CSS3** : mise en forme et responsive design.
- **JavaScript** : recherche, filtrage et interactions.

## Méthode de calcul du classement

Chaque légende obtient un score global sur 100 points, calculé à partir de trois critères :

| Critère | Pondération |
|---|---:|
| Taux de victoire (Winrate) | 50 % |
| Taux de sélection (Pickrate) | 30 % |
| Performance des armes | 20 % |

### Performance des armes

La performance de chaque arme est évaluée à partir des dégâts moyens.

Les comparaisons sont effectuées uniquement entre les légendes qui partagent la même arme, indépendamment de sa position (première ou deuxième arme).

La performance globale des armes correspond à la moyenne des scores obtenus avec les deux armes.

### Calcul du score final

Le score final est calculé avec la formule suivante :

`Score final = (Winrate × 0,50) + (Pickrate × 0,30) + (Performance des armes × 0,20)`

Les différents critères sont normalisés sur 100 avant le calcul du score final.

## Sources des données

### API officielle Brawlhalla

Utilisée pour récupérer les informations officielles des légendes, notamment :
- Leur nom.
- Leur identifiant.
- Leurs deux armes.
- Leurs statistiques : force, dextérité, défense et vitesse.

### API Brawlmance

Utilisée pour récupérer les statistiques de jeu :
- Nombre de parties.
- Taux de victoire.
- Taux de sélection.
- Dégâts infligés.
- Dégâts associés à chaque arme.
- Temps passé avec chaque arme.

## Configuration des données

Le classement utilise la configuration suivante :
- **Patch :** 11.01
- **Niveau :** Diamond
- **Région :** toutes les régions disponibles (données mondiales).

**Limitation :** les statistiques Brawlmance utilisées ne garantissent pas une séparation stricte des parties classées en 1v1. Les données doivent donc être interprétées comme un indicateur statistique global, et non comme un classement officiel compétitif.

## Installation et lancement

### Prérequis

- Go 1.27.1 ou version compatible.
- Une connexion Internet pour récupérer les données des API.

### Cloner le projet

```bash
git clone https://github.com/TON-PSEUDO/brawlhalla-tierlist.git
cd brawlhalla-tierlist
```

### Lancer le serveur

```bash
go run .
```

Le site sera accessible à l'adresse :

http://localhost:8080

## Structure du projet

```text
brawlhalla-tierlist/
│
├── api/                    # Communication avec les API
│
├── cmd/                    # Programmes de test et outils
│
├── handlers/               # Gestion des requêtes HTTP
│
├── models/                 # Structures de données
│
├── services/               # Calcul et logique du classement
│
├── static/
│   ├── css/                # Feuilles de style
│   └── js/                 # Scripts JavaScript
│
├── templates/              # Pages HTML
│
├── main.go                 # Point d'entrée du serveur
├── go.mod                  # Configuration du module Go
└── README.md               # Documentation du projet
```

---

*Brawlhalla Tier List — 2026*