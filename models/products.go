package models

import (
	"webappgo/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func SearchForProducts() []Product {
	db := db.DataBaseConection()
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
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount
		products = append(products, p)
	}

	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, amount int) {
	db := db.DataBaseConection()

	insertData, err := db.Prepare("Insert into shoes(name, description, price, amount) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}
	insertData.Exec(name, description, price, amount)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DataBaseConection()
	deleteProduct, err := db.Prepare("delete from shoes where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}
