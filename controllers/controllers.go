package controllers

import (
	"html/template"
	"net/http"
	"webappgo/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchForProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)

}
