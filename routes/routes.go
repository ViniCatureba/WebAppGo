package routes

import (
	"net/http"
	"webappgo/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/update", controllers.Edit)
	http.HandleFunc("/confirmedupdate", controllers.ConfirmedUpdate)
}
