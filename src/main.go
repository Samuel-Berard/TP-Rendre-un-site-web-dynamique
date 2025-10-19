package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Produit struct {
	Name     string
	Price    string
	Img      string
	OldPrice string
}

var produits = []Produit{
	{Name: "PALACE PULL A CAPUCHE UNISEXE CHASSEUR", Price: "148", Img: "/static/img/products/19A.webp"},
	{Name: "PALACE PULL A CAPUCHON MARINE", Price: "138", Img: "/static/img/products/21A.webp"},
	{Name: "PALACE PULL CREW PASSEPOSE NOIR", Price: "90", OldPrice: "128", Img: "/static/img/products/22A.webp"},
	{Name: "PALACE WASHED TERRY 1/4 PLACKET HOOD MOJITO", Price: "168", Img: "/static/img/products/16A.webp"},
	{Name: "PALACE PANTALON BOSSY JEAN STONE", Price: "125", Img: "/static/img/products/34B.webp"},
	{Name: "PALACE PANTALON CARGO GORE-TEX R-TEK NOIR", Price: "110", Img: "/static/img/products/33B.webp"},
}

func main() {

	temp, errtemp := template.ParseGlob("./src/template/*.html")
	if errtemp != nil {
		fmt.Println(errtemp)
		os.Exit(1)
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "index", produits)
	})

	http.HandleFunc("/consult", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "consult", produits)
	})

	chemin, _ := os.Getwd()
	fmt.Println(chemin)
	fileserver := http.FileServer(http.Dir(chemin + "/src/assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe(":8000", nil)
}
