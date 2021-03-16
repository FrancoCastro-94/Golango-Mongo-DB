package routers

import (
	"net/http"

	"github.com/FrancoCastro-94/Go-Mongo-practice/db"
	"github.com/FrancoCastro-94/Go-Mongo-practice/models"
	"github.com/FrancoCastro-94/Go-Mongo-practice/templates"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ShowIndexTemplate return a template index
func ShowIndexTemplate(w http.ResponseWriter, r *http.Request) {

	listPlayers := db.GetAllPlayers()

	templates.Tmpl.ExecuteTemplate(w, "Index", listPlayers)
}

//ShowNewTemplate return New template
func ShowNewTemplate(w http.ResponseWriter, r *http.Request) {

	templates.Tmpl.ExecuteTemplate(w, "New", nil)
}

//ShowViewTemplate return Show template
func ShowViewTemplate(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	var player models.Player
	player = db.GetOnePlayer(id)

	templates.Tmpl.ExecuteTemplate(w, "View", player)
}

//ShowEditTemplate return a template New
func ShowEditTemplate(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	var player models.Player
	player = db.GetOnePlayer(id)

	templates.Tmpl.ExecuteTemplate(w, "Edit", player)
}

//UpdatePlayer update by id
func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	if r.Method == "POST" {
		player.UUID = r.FormValue("uuid")
		player.Name = r.FormValue("name")
		player.LastName = r.FormValue("lastName")
		player.Number = r.FormValue("number")
		player.Foto = r.FormValue("foto")
	}

	status, err := db.UpdatePlayer(player)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	http.Redirect(w, r, "/", 301)
}

//DeletePlayer and return template index
func DeletePlayer(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	_, err := db.DeletePlayer(id)
	if err != nil {
		http.Error(w, "Error", 401)
	}

	http.Redirect(w, r, "/", 301)
}

//SavePlayer in db
func SavePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	if r.Method == "POST" {
		player.ID = primitive.NewObjectID()
		player.UUID = player.ID.Hex()
		player.Name = r.FormValue("name")
		player.LastName = r.FormValue("lastName")
		player.Number = r.FormValue("number")
		player.Foto = r.FormValue("foto")
	}

	status, err := db.InsertPlayer(player)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	http.Redirect(w, r, "/", 301)
}
