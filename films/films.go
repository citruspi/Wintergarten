package films

import (
	"log"
	"os"
	"strconv"

	tmdb "github.com/ryanbradynd05/go-tmdb"
)

type Film struct {
	Adult             bool `json:"adult"`
	AlternativeTitles struct {
		Titles []struct {
			Iso31661 string `json:"iso_3166_1"`
			Title    string `json:"title"`
		} `json:"titles"`
	} `json:"alternative_titles"`
	BackdropPath        string `json:"backdrop_path"`
	BelongsToCollection struct {
		BackdropPath string `json:"backdrop_path"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
	} `json:"belongs_to_collection"`
	Budget  int `json:"budget"`
	Credits struct {
		Cast []struct {
			CastID      int    `json:"cast_id"`
			Character   string `json:"character"`
			CreditID    string `json:"credit_id"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Order       int    `json:"order"`
			ProfilePath string `json:"profile_path"`
		} `json:"cast"`
		Crew []struct {
			CreditID    string `json:"credit_id"`
			Department  string `json:"department"`
			ID          int    `json:"id"`
			Job         string `json:"job"`
			Name        string `json:"name"`
			ProfilePath string `json:"profile_path"`
		} `json:"crew"`
	} `json:"credits"`
	Genres []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage string `json:"homepage"`
	ID       int    `json:"id"`
	Images   struct {
		Backdrops []struct {
			AspectRatio float64     `json:"aspect_ratio"`
			FilePath    string      `json:"file_path"`
			Height      int         `json:"height"`
			Iso6391     interface{} `json:"iso_639_1"`
			VoteAverage float64     `json:"vote_average"`
			VoteCount   int         `json:"vote_count"`
			Width       int         `json:"width"`
		} `json:"backdrops"`
		Posters []struct {
			AspectRatio float64 `json:"aspect_ratio"`
			FilePath    string  `json:"file_path"`
			Height      int     `json:"height"`
			ID          string  `json:"id"`
			Iso6391     string  `json:"iso_639_1"`
			VoteAverage float64 `json:"vote_average"`
			VoteCount   int     `json:"vote_count"`
			Width       int     `json:"width"`
		} `json:"posters"`
	} `json:"images"`
	ImdbID   string `json:"imdb_id"`
	Keywords struct {
		Keywords []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"keywords"`
	} `json:"keywords"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float64 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ProductionCompanies []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1"`
		Name     string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate string `json:"release_date"`
	Releases    struct {
		Countries []struct {
			Certification string `json:"certification"`
			Iso31661      string `json:"iso_3166_1"`
			Primary       bool   `json:"primary"`
			ReleaseDate   string `json:"release_date"`
		} `json:"countries"`
	} `json:"releases"`
	Revenue int `json:"revenue"`
	Runtime int `json:"runtime"`
	Similar struct {
		Page    int `json:"page"`
		Results []struct {
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
		} `json:"results"`
		TotalPages   int `json:"total_pages"`
		TotalResults int `json:"total_results"`
	} `json:"similar"`
	SpokenLanguages []struct {
		Iso6391 string `json:"iso_639_1"`
		Name    string `json:"name"`
	} `json:"spoken_languages"`
	Status  string `json:"status"`
	Tagline string `json:"tagline"`
	Title   string `json:"title"`
	Video   bool   `json:"video"`
	Videos  struct {
		Results []struct {
			ID      string `json:"id"`
			Iso6391 string `json:"iso_639_1"`
			Key     string `json:"key"`
			Name    string `json:"name"`
			Site    string `json:"site"`
			Size    int    `json:"size"`
			Type    string `json:"type"`
		} `json:"results"`
	} `json:"videos"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}

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
