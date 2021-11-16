package main

import (
	"encoding/json"
	"log"
	Rockpaperscissors "myapp/rps"
	"net/http"
	"strconv"
	"text/template"
)

func homepage(w http.ResponseWriter, r *http.Request) {

	Template(w, "index.html")

}

func playRound(w http.ResponseWriter, r *http.Request) {
	player_Choice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := Rockpaperscissors.Playround(player_Choice)

	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "json")
	w.Write(out)

}

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/play", playRound)
	http.ListenAndServe(":8080", nil)
}
func Template(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return

	}
}
