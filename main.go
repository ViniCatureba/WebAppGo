package main

import (
	"net/http"
	"webappgo/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8000", nil)
}
