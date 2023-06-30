package openweather

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	ErrorCityNotFound   = errors.New("kota tidak ditemukan")
	ErrorAPIKeyNotFound = errors.New("api key not found")
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
		} else if err == ErrorAPIKeyNotFound {
			fmt.Println("Mohon Maaf, API Key tidak ditemukan")
			fmt.Println("Silahkan lihat README.md untuk tutorial set API Key")
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

func getAPIKey() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")

	if apiKey == "" {
		return "", ErrorAPIKeyNotFound
	}

	return apiKey, nil

}

// Fungsi untuk mengambil data cuaca 5 hari kedepan.
// Limitasi:
// - Waktu cuaca yagn diambil adalah cuaca pukul 13.00 (jam 1 siang).
// - Limitasi API yang hanya bisa melihat forecast menyebabkan jam eksekusi program mempengaruhi hasil.
// - Apabila program dieksekusi sebelum jam 13.00 maka forecast hari ini akan tercantum.
// - Apabila program dieksekusi setelah jam 13.00 maka forecast akan dimulai dari hari besoknya.
func getFiveDayForecast(city, unit string) (*Weather, error) {

	apiKey, err := getAPIKey()
	if err != nil {
		return nil, err
	}

	if unit == "" {
		unit = "metric"
	}
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&units=%s&appid=%s", city, unit, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
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
