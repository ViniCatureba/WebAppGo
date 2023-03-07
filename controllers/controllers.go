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
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	porductId := r.URL.Query().Get("id")
	models.DeleteProduct(porductId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	temp.ExecuteTemplate(w, "update", product)

}

func ConfirmedUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		cnvId, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}

		cnvPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}

		cnvAmount, err := strconv.Atoi(amount)
		if err != nil {
			panic(err.Error())
		}

		models.UpdateProduct(cnvId, cnvAmount, name, description, cnvPrice)

	}
}
