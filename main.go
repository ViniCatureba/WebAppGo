package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func DataBaseConection() *sql.DB {
	connection := "user=postgres dbname=Shoe_Store password=Vinivini14! host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := DataBaseConection()
	selectProducts, err := db.Query("select * from shoes")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}
	for selectProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount
		products = append(products, p)
	}

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
