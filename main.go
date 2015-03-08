package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	tmdb "github.com/ryanbradynd05/go-tmdb"
)

var (
	client *tmdb.TMDb
)

func init() {
	api_key := os.Getenv("TMDB_API_KEY")

	if api_key == "" {
		log.Fatal("Missing API key")
	}

	client = tmdb.Init(api_key)
}

func getFilm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	film, err := client.GetMovieInfo(id, nil)

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
