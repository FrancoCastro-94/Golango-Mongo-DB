package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/FrancoCastro-94/Go-Mongo-practice/middlew"
	"github.com/FrancoCastro-94/Go-Mongo-practice/routers"
)

//Handlers set port and listen routers
func Handlers() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on port: " + port)
	http.HandleFunc("/", middlew.CheckDB(routers.ShowIndexTemplate))
	http.HandleFunc("/show", middlew.CheckDB(routers.ShowViewTemplate))
	http.HandleFunc("/new", routers.ShowNewTemplate)
	http.HandleFunc("/edit", middlew.CheckDB(routers.ShowEditTemplate))
	http.HandleFunc("/insert", middlew.CheckDB(routers.SavePlayer))
	http.HandleFunc("/update", middlew.CheckDB(routers.UpdatePlayer))
	http.HandleFunc("/delete", middlew.CheckDB(routers.DeletePlayer))
	http.ListenAndServe(":"+port, nil)

}
