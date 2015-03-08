package films

import (
	"log"
	"os"
	"strconv"

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

func Get(film_id string) (*tmdb.Movie, error) {
	var film *tmdb.Movie
	id, err := strconv.Atoi(film_id)

	if err != nil {
		return film, err
	}

	film, err = client.GetMovieInfo(id, nil)

	if err != nil {
		return film, err
	}

	return film, nil
}
