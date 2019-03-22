package baidu_map

type BDConfig struct {
	ApiUrl string `yaml:"api_url"`
	AppKey string `yaml:"app_key"`
	QueryLog string `yaml:"query_log"`
}

type location struct {
	Lat float64
	Lng float64
}

type BMapItem struct {
	Name         string   `json:"name"`
	Location     location `json:"location"`
	LocationInfo string
	Address      string `json:"address"`
	Province     string `json:"province"`
	City         string `json:"city"`
	Area         string `json:"area"`
	StreetId     string `json:"street_id"`
	Telephone    string `json:"telephone"`
	Detail       int    `json:"detail"`
	Uid          string `json:"uid"`
}

type BMapResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Results []BMapItem `json:"results"`
}
