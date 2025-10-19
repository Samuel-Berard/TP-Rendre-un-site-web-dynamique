# TP-Rendre-un-site-web-dynamique
Rendre un site web dynamique en Go

L'objectif est de construire un mini site e-commerce dynamique en trois étapes : affichage d'articles, consultation des détails et ajout de nouveaux produits via formulaire.

## Fonctionnalités

### Liste des articles
Page d'accueil avec en-tête (logo + navigation).
Affichage de 6 produits (image, nom, prix, réduction optionnelle).
Utilisation d'un template Go avec boucles, variables et conditions.
Bouton "Consulter" pour accéder aux détails.

### Détails d'un article
Page dédiée à un produit sélectionné.
Affiche : nom, description, image, prix, stock, réduction éventuelle.
Gestion des erreurs si l'ID n'existe pas.
Navigation cohérente (retour à la liste).

### Ajout d'un produit
Formulaire pour ajouter un nouvel article.
Vérification des données saisies (stock >= 0, champs obligatoires).
Ajout du produit à la liste globale en mémoire.
Redirection automatique vers la page détail du produit ajouté.
Image par défaut si aucune n'est fournie.

## Technologies utilisées

Go (net/http, html/template) - gestion des routes et rendu dynamique
HTML5 / CSS3 - mise en forme et structure

## Installation et exécution

### Prérequis
Go 1.16 ou plus récent
Navigateur web

### Étapes
```bash
# Cloner le dépôt
git clone https://github.com/Samuel-Berard/TP-Rendre-un-site-web-dynamique.git
cd TP-Rendre-un-site-web-dynamique

# Lancer le serveur
cd src
go run main.go
```

Ouvrir http://localhost:8000 dans le navigateur.