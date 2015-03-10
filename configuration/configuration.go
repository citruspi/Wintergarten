package configuration

type Configuration struct {
	Web struct {
		Port    int    `json:"port"`
		Address string `json:"address"`
	} `json:"web"`
	Cache struct {
		TTL int `json:"ttl"`
	} `json:"cache"`
	TMDb struct {
		APIKEY string `json:"api_key"`
	} `json:"tmdb"`
}
