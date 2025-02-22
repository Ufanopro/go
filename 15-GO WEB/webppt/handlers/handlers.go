package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"
	rps "webppt/logic"
)

type Player struct {
	Name string
}

var player Player

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	restartVaue()
	renderTemplate(w, templateBase, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartVaue()
	renderTemplate(w, templateBase, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al parsear el formulario", http.StatusBadRequest)
			return
		}

		player.Name = r.Form.Get("name")
	}

	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}

	renderTemplate(w, templateBase, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartVaue()
	renderTemplate(w, templateBase, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, base, page string, data any) {
	tpl, err := template.ParseFiles(base, templateDir+page)
	if err != nil {
		http.Error(w, "Error al cargar las plantillas", http.StatusInternalServerError)
		log.Println("Error al cargar plantillas:", err)
		return
	}

	err = tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		log.Println("Error al ejecutar plantilla:", err)
		return
	}
}

// Reiniciar valores
func restartVaue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
