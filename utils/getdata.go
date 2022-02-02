package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetWeatherByCity(key, cityname string) WeatherResponse{
	apiurl := "http://api.openweathermap.org/data/2.5/weather?units=metric&"

	urltocall := fmt.Sprintf("%sq=%s&appid=%s", apiurl, cityname, key)
	log.Println("Sending API call to:", urltocall)

	response, err := http.Get(urltocall) // Sending GET request
	if err != nil {
		fmt.Println("An error has occurred during the API call:", err.Error())
		log.Println(err)
		os.Exit(1)
	}

	if response.StatusCode == 401{
		fmt.Println("API key entered is invalid!")
		log.Println(err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("An error has occurred when reading the response call:", err.Error())
		log.Println(err)
		os.Exit(1)
	}

	var apiresponse WeatherResponse

	err = json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println("An error has occurred when processing the response json:", err.Error())
		log.Println(err)
		os.Exit(1)
	}
	return apiresponse
}

func GetWeatherByPincode(key, pincode string) WeatherResponse{
	apiurl := "http://api.openweathermap.org/data/2.5/weather?units=metric&"

	urltocall := fmt.Sprintf("%szip=%s,in&appid=%s", apiurl, pincode, key)
	log.Println("Sending API call to:", urltocall)

	response, err := http.Get(urltocall) // Sending GET request
	if err != nil {
		fmt.Println("An error has occurred during the API call:", err.Error())
		log.Println(err)
		os.Exit(1)
	}

	if response.StatusCode == 401{
		fmt.Println("API key entered is invalid!")
		log.Println(err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("An error has occurred when reading the response call:", err.Error())
		log.Println(err)
		os.Exit(1)
	}

	var apiresponse WeatherResponse

	err = json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println("An error has occurred when processing the response json:", err.Error())
		log.Println(err)
		os.Exit(1)
	}

	return apiresponse
}

func GetWeatherByLatLon(key string, lat float64, lon float64) WeatherResponse{
	apiurl := "http://api.openweathermap.org/data/2.5/weather?units=metric&"

	urltocall := fmt.Sprintf("%slat=%f&lon=%f&appid=%s", apiurl, lat, lon, key)
	log.Println("Sending API call to:", urltocall)

	response, err := http.Get(urltocall) // Sending GET request
	if err != nil {
		fmt.Println("An error has occurred during the API call:", err.Error())
		log.Println(err)
		os.Exit(1)
	}

	if response.StatusCode == 401{
		fmt.Println("API key entered is invalid!")
		log.Println(err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("An error has occurred when reading the response call:", err.Error())
		log.Println(err)
		os.Exit(1)
	}

	var apiresponse WeatherResponse

	err = json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println("An error has occurred when processing the response json:", err.Error())
		log.Println(err)
		os.Exit(1)
	}

	return apiresponse
}

func GetConditionIcon(apiresponse WeatherResponse) string{
	id := apiresponse.Weather[0].ID
	if id >= 200 && id <300 {
		// Thundestorm
		return "⛈"
	} else if id >= 300 && id < 400 {
		// Drizzle
		return "☔️"
	} else if id >= 500 && id < 600 {
		// Rain
		return "🌧"
	}else if id >= 600 && id < 700 {
		// Snow
		return "❄️"
	} else if id >= 700 && id < 800 {
		// Atmosphere
		return "🌫"
	} else if id == 800 {
		// Clear
		return "☀️"
	} else if id > 800 && id < 800{
		// Cloudy
		return "🌤"
	} else {
		return "❓"
	}
}