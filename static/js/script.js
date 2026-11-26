
// Récupération des éléments HTML.
const searchInput = document.getElementById("searchInput");
const weaponFilter = document.getElementById("weaponFilter");
const resultCount = document.getElementById("resultCount");

// Récupération de toutes les cartes des personnages.
const cards = Array.from(document.querySelectorAll(".legend-card"));

// Normalise le texte pour faciliter les recherches.
function normalizeText(text) {
    return text
        .normalize("NFD")
        .replace(/[\u0300-\u036f]/g, "")
        .toLowerCase()
        .trim();
}

// Création automatique de la liste des armes disponibles.
const weapons = new Map();

cards.forEach(card => {
    const weaponElements = card.querySelectorAll(".weapons span");

    weaponElements.forEach(element => {
        const weaponName = element.textContent.trim();
        const normalizedName = normalizeText(weaponName);

        // Évite d'ajouter plusieurs fois la même arme.
        if (!weapons.has(normalizedName)) {
            weapons.set(normalizedName, weaponName);
        }
    });
});

// Ajout des armes dans le menu déroulant.
weapons.forEach((name, value) => {
    const option = document.createElement("option");

    option.value = value;
    option.textContent = name;

    weaponFilter.appendChild(option);
});

// Fonction qui applique les filtres.
function filterLegends() {

    const searchValue = normalizeText(searchInput.value);
    const selectedWeapon = weaponFilter.value;

    let visibleCount = 0;

    cards.forEach(card => {

        // Récupération du nom du personnage.
        const name = normalizeText(
            card.querySelector("h3").textContent
        );

        // Récupération des deux armes.
        const cardWeapons = Array.from(
            card.querySelectorAll(".weapons span")
        ).map(weapon => normalizeText(weapon.textContent));

        // Vérification de la recherche par nom.
        const matchesName = name.includes(searchValue);

        // Vérification du filtre par arme.
        const matchesWeapon = selectedWeapon === "" ||
            cardWeapons.includes(normalizeText(selectedWeapon));

        // Affichage uniquement des personnages correspondants.
        const isVisible = matchesName && matchesWeapon;

        card.hidden = !isVisible;

        if (isVisible) {
            visibleCount++;
        }
    });

    // Mise à jour du nombre de personnages affichés.
    resultCount.textContent =
        `${visibleCount} personnage(s) affiché(s) sur ${cards.length}`;
}

// Déclenchement des filtres lors des interactions.
searchInput.addEventListener("input", filterLegends);
weaponFilter.addEventListener("change", filterLegends);

// Affichage initial.
filterLegends();

// Gestion de l'affichage des détails de chaque légende.
cards.forEach(card => {

    const detailsButton = card.querySelector(".details-button");
    const detailsContainer = card.querySelector(".legend-details");

    // Vérification de la présence des éléments.
    if (!detailsButton || !detailsContainer) {
        return;
    }

    // Affichage ou masquage des détails au clic.
    detailsButton.addEventListener("click", () => {

        const isHidden = detailsContainer.hidden;

        detailsContainer.hidden = !isHidden;

        detailsButton.textContent = isHidden
            ? "Masquer les détails"
            : "Voir les détails";
    });
});
