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

func EditProduct(id string) Product {
	db := db.DataBaseConection()

	productFromDb, err := db.Query("select * from shoes where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	productToEdit := Product{}
	for productFromDb.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = productFromDb.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		productToEdit.Id = id
		productToEdit.Name = name
		productToEdit.Description = description
		productToEdit.Price = price
		productToEdit.Amount = amount

	}
	defer db.Close()
	return productToEdit
}

func UpdateProduct(id, amount int, name, description string, price float64) {
	db := db.DataBaseConection()

	UpdateProduct, err := db.Prepare("update shoes set name=$1, description=$2, price=$3, amount=$4,id=$5")
	if err != nil {
		panic(err.Error())
	}
	UpdateProduct.Exec(name, description, price, amount, id)
	defer db.Close()
}
