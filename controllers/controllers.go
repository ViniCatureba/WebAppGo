package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"webappgo/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchForProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)

}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}
		convAmount, err := strconv.Atoi(amount)
		if err != nil {
			panic(err.Error())
		}
		models.CreateNewProduct(name, description, convPrice, convAmount)
	}
	http.Redirect(w, r, "/", 301)
}
