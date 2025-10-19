package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Produit struct {
	ID          int
	Name        string
	Description string
	Price       string
	OldPrice    string
	Stock       int
	Img         string
}

var (
	produits = []Produit{
		{
			ID:          1,
			Name:        "PALACE PULL A CAPUCHE UNISEXE CHASSEUR",
			Description: "Sweat à capuche confortable en coton premium, idéal pour un look décontracté. Coupe unisexe avec logo brodé Palace. Disponible en coloris chasseur exclusif.",
			Price:       "148",
			OldPrice:    "",
			Stock:       12,
			Img:         "/static/img/products/19A.webp",
		},
		{
			ID:          2,
			Name:        "PALACE PULL A CAPUCHON MARINE",
			Description: "Pull à capuchon marine élégant et intemporel. Matière douce et respirante avec finitions soignées. Parfait pour les journées fraîches.",
			Price:       "138",
			OldPrice:    "",
			Stock:       7,
			Img:         "/static/img/products/21A.webp",
		},
		{
			ID:          3,
			Name:        "PALACE PULL CREW PASSEPOSE NOIR",
			Description: "Crewneck noir avec détails passepoisés signature. Design minimaliste et moderne, confectionné en molleton de qualité. Coupe ajustée pour un style urbain affirmé.",
			Price:       "90",
			OldPrice:    "128",
			Stock:       3,
			Img:         "/static/img/products/22A.webp",
		},
		{
			ID:          4,
			Name:        "PALACE WASHED TERRY 1/4 PLACKET HOOD MOJITO",
			Description: "Hoodie 1/4 zip en éponge lavée, coloris mojito vibrant. Texture douce et confortable avec finition vintage. Idéal pour un look streetwear décontracté et original.",
			Price:       "168",
			OldPrice:    "",
			Stock:       5,
			Img:         "/static/img/products/16A.webp",
		},
		{
			ID:          5,
			Name:        "PALACE PANTALON BOSSY JEAN STONE",
			Description: "Jean stone lavé coupe Bossy, silhouette droite et confortable. Denim robuste avec détails Palace subtils. Un essentiel pour toutes les saisons.",
			Price:       "125",
			OldPrice:    "",
			Stock:       9,
			Img:         "/static/img/products/34B.webp",
		},
		{
			ID:          6,
			Name:        "PALACE PANTALON CARGO GORE-TEX R-TEK NOIR",
			Description: "Cargo technique Gore-Tex R-Tek imperméable et résistant. Multiple poches utilitaires avec fermetures YKK. Performance et style pour les aventuriers urbains.",
			Price:       "110",
			OldPrice:    "",
			Stock:       4,
			Img:         "/static/img/products/33B.webp",
		},
	}
)

func main() {

	temp, errtemp := template.ParseGlob("./src/template/*.html")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/index", http.StatusFound)
	})
	if errtemp != nil {
		fmt.Println(errtemp)
		os.Exit(1)
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "index", produits)
	})

	http.HandleFunc("/consult", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "ID pas trouvé", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID non répertorié", http.StatusBadRequest)
			return
		}

		var prod *Produit
		for i := range produits {
			if produits[i].ID == id {
				prod = &produits[i]
				break
			}
		}

		if prod == nil {
			http.NotFound(w, r)
			return
		}

		temp.ExecuteTemplate(w, "consult", prod)
	})

	chemin, _ := os.Getwd()
	fmt.Println(chemin)
	fileserver := http.FileServer(http.Dir(chemin + "/src/assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe(":8000", nil)
}
