package openweather

type ForecastResp struct {
	Count uint32         `json:"cnt"`
	List  []ForecastItem `json:"list"`
	City  CityInfo       `json:"city"`
}

type ForecastItem struct {
	DateTime int64           `json:"dt"`
	Main     MainTemperature `json:"main"`
}

type MainTemperature struct {
	Temp        float64 `json:"temp"`
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	Pressure    float64 `json:"pressure"`
	SeaLevel    float64 `json:"sea_level"`
	GroundLevel float64 `json:"grnd_level"`
}

type CityInfo struct {
	Id         uint32 `json:"id"`
	Name       string `json:"name"`
	Coordinate LatLng `json:"coord"`
	Country    string `json:"country"`
}

type LatLng struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lon"`
}
