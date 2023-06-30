package openweather

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Weather struct {
	city     CityInfo
	weathers []WeatherPerDay
}

type WeatherPerDay struct {
	Temperature string
	DateTime    time.Time
}

var (
	ErrorCityNotFound = errors.New("kota tidak ditemukan")
)

func RunForecast() {
	fmt.Println("Silahkan masukan nama kota yang akan dicari cuacanya (contoh: Jakarta)")
	reader := bufio.NewReader(os.Stdin)
	city, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}

	weather, err := getFiveDayForecast(strings.TrimSpace(city), "metric")
	if err != nil {
		if err == ErrorCityNotFound {
			fmt.Println("Mohon Maaf, data Kota tidak ditemukan!")
		} else {
			log.Fatal(err)
		}
		return
	}

	fmt.Println("=========")
	fmt.Printf("Data Ramalan Cuaca untuk Kota: %s (Lat:%f, Long: %f)\n\n", weather.city.Name, weather.city.Coordinate.Lat, weather.city.Coordinate.Long)
	fmt.Println("Weather Forecast")
	for _, wth := range weather.weathers {
		timeString := wth.DateTime.Format("Mon, 02 Jan 2006")
		fmt.Printf("%s: %s\n", timeString, wth.Temperature)
	}
}

func getFiveDayForecast(city, unit string) (*Weather, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	if unit == "" {
		unit = "metric"
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&units=%s&appid=%s", city, unit, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ForecastResp
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if len(response.List) < 1 {
		return nil, ErrorCityNotFound
	}

	var weathers []WeatherPerDay
	dates := map[string]*time.Time{}

	for _, forecast := range response.List {

		time := time.Unix(int64(forecast.DateTime), 0)

		if dates[time.Format("02-01-2006")] != nil {
			continue
		}

		if time.Hour() != 13 {
			continue
		}

		weathers = append(weathers, WeatherPerDay{
			Temperature: fmt.Sprintf("%.2f°C", forecast.Main.Temp),
			DateTime:    time,
		})
		dates[time.Format("02-01-2006")] = &time
	}

	weather := Weather{
		city:     response.City,
		weathers: weathers,
	}

	return &weather, nil
}
