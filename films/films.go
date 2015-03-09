package films

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/fzzy/radix/redis"
)

type Film struct {
	Adult             bool `json:"adult"`
	AlternativeTitles struct {
		Titles []struct {
			Iso31661 string `json:"iso_3166_1"`
			Title    string `json:"title"`
		} `json:"titles"`
	} `json:"alternative_titles"`
	Availability        *FilmAvailability `json:"availability,omitempty"`
	BackdropPath        string            `json:"backdrop_path"`
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

type FilmAvailability struct {
	Purchase *struct {
		AmazonVideo *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"amazon_video_purchase,omitempty"`
		GooglePlay *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"android_purchase,omitempty"`
		AppleiTunes *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"apple_itunes_purchase,omitempty"`
		SonyEntertainment *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"sony_purchase,omitempty"`
		Vudu *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"vudu_purchase,omitempty"`
		Youtube *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"youtube_purchase,omitempty"`
		XBOXMarketplace *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"xbox_purchase,omitempty"`
	} `json:"purchase,omitempty"`
	Rental *struct {
		AmazonVideo *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"amazon_video_rental,omitempty"`
		GooglePlay *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"android_rental,omitempty"`
		AppleiTunes *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"apple_itunes_rental,omitempty"`
		SonyEntertainment *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"sony_rental,omitempty"`
		Vudu *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"vudu_rental,omitempty"`
		Youtube *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"youtube_rental,omitempty"`
	} `json:"rental,omitempty"`
	Streaming *struct {
		NetflixInstant *struct {
			URL     string `json:"direct_url,omitempty"`
			Checked int64  `json:"date_checked,omitempty"`
		} `json:"netflix_instant,omitempty"`
	} `json:"streaming,omitempty"`
}

var (
	api_key string
)

func init() {
	api_key = os.Getenv("TMDB_API_KEY")

	if api_key == "" {
		log.Fatal("Missing API key")
	}
}

func Get(film_id string) (Film, error) {
	var film Film

	client, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	err = client.Cmd("PING").Err
	if err != nil {
		log.Fatal(err)
	}

	s, err := client.Cmd("get", film_id).Str()

	if err == nil {
		err = json.Unmarshal([]byte(s), &film)

		if err == nil {
			return film, nil
		}
	}

	film, err = queryTMDb(film_id)

	if err != nil {
		return film, err
	}

	availability, err := determineAvailability(film)

	if err == nil {
		film.Availability = &availability
	}

	marshalled, err := json.Marshal(film)

	if err == nil {
		_ = client.Cmd("set", film_id, marshalled)
	}

	return film, nil
}

func queryTMDb(film_id string) (Film, error) {
	var url string
	var film Film

	var buffer bytes.Buffer

	buffer.WriteString("http://api.themoviedb.org/3/movie/")
	buffer.WriteString(film_id)
	buffer.WriteString("?api_key=")
	buffer.WriteString(api_key)
	buffer.WriteString("&append_to_response=")
	buffer.WriteString("alternative_titles,credits,images,keywords,releases,videos,similar")

	url = string(buffer.Bytes())

	resp, err := http.Get(url)

	if err != nil {
		return film, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return film, err
	}

	err = json.Unmarshal(body, &film)

	if err != nil {
		return film, err
	}

	return film, nil
}

func determineAvailability(film Film) (FilmAvailability, error) {
	type searchResult struct {
		Id    string `json:"_id"`
		Links struct {
			IMDb string `json:"imdb"`
		} `json:"links"`
	}

	var availability FilmAvailability
	var results []searchResult

	var cisi_url *url.URL

	cisi_url, err := url.Parse("http://www.canistream.it/services/search")

	if err != nil {
		return availability, err
	}

	parameters := url.Values{}
	parameters.Add("movieName", film.Title)

	cisi_url.RawQuery = parameters.Encode()

	resp, err := http.Get(cisi_url.String())

	if err != nil {
		return availability, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return availability, err
	}

	err = json.Unmarshal(body, &results)

	if err != nil {
		return availability, err
	}

	var canIStreamItID string

	for _, result := range results {
		var imdb_id string

		if result.Links.IMDb != "" {
			imdb_id = strings.Split(result.Links.IMDb, "/")[4]
		}

		if imdb_id == film.ImdbID {
			canIStreamItID = result.Id
		}
	}

	if canIStreamItID != "" {
		mediaTypes := []string{"rental", "purchase", "streaming"}

		for _, mediaType := range mediaTypes {
			var buffer bytes.Buffer

			buffer.WriteString("http://www.canistream.it/services/query?movieId=")
			buffer.WriteString(canIStreamItID)
			buffer.WriteString("&attributes=1&mediaType=")
			buffer.WriteString(mediaType)

			resp, err = http.Get(string(buffer.Bytes()))

			if err != nil {
				return availability, err
			}

			defer resp.Body.Close()

			body, err = ioutil.ReadAll(resp.Body)

			if err != nil {
				return availability, err
			}

			switch mediaType {
			case "rental":
				err = json.Unmarshal(body, &availability.Rental)

				if err != nil {
					return availability, err
				}
			case "purchase":
				err = json.Unmarshal(body, &availability.Purchase)

				if err != nil {
					return availability, err
				}
			case "streaming":
				err = json.Unmarshal(body, &availability.Streaming)

				if err != nil {
					return availability, err
				}
			}
		}
	}

	return availability, nil
}
