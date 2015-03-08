package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	tmdb "github.com/ryanbradynd05/go-tmdb"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
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

func getFilm(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])

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
	goji.Get("/film/:id/", getFilm)
	goji.Serve()
}
