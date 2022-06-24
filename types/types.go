package types

type LivesItem struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	Winddirection string `json:"winddirection"`
	Windpower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	Reporttime    string `json:"reporttime"`
}

type Casts struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	Dayweather   string `json:"dayweather"`
	Nightweather string `json:"nightweather"`
	Daytemp      string `json:"daytemp"`
	Nighttemp    string `json:"nighttemp"`
	Daywind      string `json:"daywind"`
	Nightwind    string `json:"nightwind"`
	Daypower     string `json:"daypower"`
	Nightpower   string `json:"nightpower"`
}

type Forecasts struct {
	City       string  `json:"city"`
	Adcode     string  `json:"adcode"`
	Province   string  `json:"province"`
	Reporttime string  `json:"reporttime"`
	Casts      []Casts `json:"casts"`
}

type WeatherResponse struct {
	Status    string      `json:"status"`
	Count     string      `json:"count"`
	Info      string      `json:"info"`
	Infocode  string      `json:"infocode"`
	Lives     []LivesItem `json:"lives,omitempty"`
	Forecasts []Forecasts `json:"forecasts,omitempty"`
}
