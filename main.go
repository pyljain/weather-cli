package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/logrusorgru/aurora/v3"
)

const WEATHER_API = "https://api.openweathermap.org/data/2.5/weather"

func main() {

	city := flag.String("city", "london", "City for which you would like to see the weather")
	apiKey := flag.String("key", "", "Provide your API Key")
	flag.Parse()

	weather, err := getWeather(*city, *apiKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Description: %s \n", aurora.Bold(aurora.Cyan(weather.Weather[0].Description)))
	fmt.Printf("Min: %v \n", aurora.Bold(aurora.Red(ktoc(weather.Main.TempMin))))
	fmt.Printf("Max: %v \n", aurora.Bold(aurora.Green(ktoc(weather.Main.TempMax))))
}

func getWeather(city, apiKey string) (*WeatherInfo, error) {

	weatherEndpoint := fmt.Sprintf("%s?q=%s&appid=%s", WEATHER_API, city, apiKey)
	res, err := http.Get(weatherEndpoint)

	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var weatherInfo WeatherInfo
	err = json.Unmarshal(resBody, &weatherInfo)

	if err != nil {
		return nil, err
	}

	return &weatherInfo, nil
}

// WeatherInfo contains weather info
type WeatherInfo struct {
	Weather []WeatherDetail `json:"weather"`
	Main    WeatherMain     `json:"main"`
}

// WeatherDetail contains description
type WeatherDetail struct {
	Description string `json:"description"`
	Main        string `json:"main"`
}

// WeatherMain contains temperature details
type WeatherMain struct {
	TempMin float32 `json:"temp_min"`
	TempMax float32 `json:"temp_max"`
}

func ktoc(temp float32) int32 {
	tempc := temp - 273.15
	return int32(tempc)
}
