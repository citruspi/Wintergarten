package collections

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/citruspi/wintergarten/configuration"
	"github.com/citruspi/wintergarten/films"
)

type TMDbResponse struct {
	Dates struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	} `json:"dates"`
	Results      []Film `json:"results"`
	Page         int    `json:"page"`
	TotalPages   int    `json:"total_pages"`
	TotalResults int    `json:"total_results"`
}
type Film struct {
	Adult         bool    `json:"adult"`
	BackdropPath  string  `json:"backdrop_path"`
	ID            int     `json:"id"`
	OriginalTitle string  `json:"original_title"`
	Popularity    float64 `json:"popularity"`
	PosterPath    string  `json:"poster_path"`
	ReleaseDate   string  `json:"release_date"`
	Title         string  `json:"title"`
	Video         bool    `json:"video"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int     `json:"vote_count"`
}

var (
	conf configuration.Configuration
)

func init() {
	conf = configuration.Init()
}

func GetFilms(collectionName string) ([]Film, error) {
	var collection []Film

	response, err := queryTMDb(collectionName)

	if err != nil {
		return collection, err
	}

	collection = response.Results

	for _, film := range collection {
		filmID := strconv.Itoa(film.ID)

		go films.Prepare(filmID)
	}

	return collection, nil
}

func queryTMDb(collectionName string) (TMDbResponse, error) {
	var url string
	var response TMDbResponse
	var buffer bytes.Buffer

	buffer.WriteString("http://api.themoviedb.org/3/movie/")
	buffer.WriteString(collectionName)
	buffer.WriteString("?api_key=")
	buffer.WriteString(conf.TMDb.APIKEY)

	url = string(buffer.Bytes())

	resp, err := http.Get(url)

	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}
