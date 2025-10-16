package controller

import (
	"fmt"
	"html/template"
	"net/http"
	crane "power4/fonctions"
	"strconv"
)

func init() {
    fmt.Println("Controller chargé !")
}


func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	tmpl.Execute(w, data)
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":   "Accueil",      
		"Message": "Bienvenue sur la page d'accueil 🎉", 
	}
	renderTemplate(w, "index.html", data)
}

func About(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":   "Jouer",
		"Message": "caca ✨",
	}
	renderTemplate(w, "about.html", data)
}

func GamePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		col := (r.FormValue("colonne"))
		colInt, _ := strconv.Atoi(col)
		crane.PlacerPièce(colInt)
	}
	gridForTemplate := make([][]int, len(crane.Grid))
		for i := range crane.Grid {
			gridForTemplate[i] = crane.Grid[i][:]
		}

	tmpl := template.Must(template.ParseFiles("template/tableau.html"))
	tmpl.Execute(w, gridForTemplate)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { 
		name := r.FormValue("name") 
		msg := r.FormValue("msg") 

		data := map[string]string{
			"Title":   "Contact",
			"Message": "Merci " + name + " pour ton message : " + msg, 
		}
		renderTemplate(w, "contact.html", data)
		return 
	}
	data := map[string]string{
		"Title":   "Contact",
		"Message": "Envoie-nous un message 📩",
	}
	renderTemplate(w, "contact.html", data)
}
