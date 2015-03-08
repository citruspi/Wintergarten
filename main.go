package main

import (
	"encoding/json"
	"net/http"

	"github.com/citruspi/wintergarten/films"
	"github.com/gorilla/mux"
)

func getFilm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	film, err := films.Get(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshalled, err := json.Marshal(film)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/film/{id:[0-9]+}/", getFilm)

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
