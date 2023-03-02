package routes

import (
	"net/http"
	"webappgo/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
}
