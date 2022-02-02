package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetWeatherByCity(key, cityname string) {
	apiurl := "http://api.openweathermap.org/data/2.5/weather?units=metric&"

	urltocall := fmt.Sprintf("%sq=%s&appid=%s", apiurl, cityname, key)
	fmt.Println("Sending API call to:", urltocall)

	response, err := http.Get(urltocall) // Sending GET request
	if err != nil {
		fmt.Println("An error has occurred during the API call:", err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("An error has occurred when reading the response call:", err.Error())
		os.Exit(1)
	}

	var apiresponse WeatherResponse

	err = json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println("An error has occurred when processing the response json:", err.Error())
		os.Exit(1)
	}

	fmt.Println("\nWeather data for:", cityname)
	fmt.Printf("Current Temperature: %.2fC\n", apiresponse.Main.Temp)
	fmt.Printf("Minimum Temperature: %.2fC\n", apiresponse.Main.TempMin)
	fmt.Printf("Maximum Temperature: %.2fC\n", apiresponse.Main.TempMax)
	fmt.Printf("Feel's Like: %.2fC\n", apiresponse.Main.FeelsLike)
}

func GetWeatherByPincode(key, pincode string) {
	apiurl := "http://api.openweathermap.org/data/2.5/weather?units=metric&"

	urltocall := fmt.Sprintf("%szip=%s,in&appid=%s", apiurl, pincode, key)
	fmt.Println("Sending API call to:", urltocall)

	response, err := http.Get(urltocall) // Sending GET request
	if err != nil {
		fmt.Println("An error has occurred during the API call:", err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("An error has occurred when reading the response call:", err.Error())
		os.Exit(1)
	}

	var apiresponse WeatherResponse

	err = json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println("An error has occurred when processing the response json:", err.Error())
		os.Exit(1)
	}

	fmt.Println("\nWeather data for:", apiresponse.Name)
	fmt.Printf("Current Temperature: %.2fC\n", apiresponse.Main.Temp)
	fmt.Printf("Minimum Temperature: %.2fC\n", apiresponse.Main.TempMin)
	fmt.Printf("Maximum Temperature: %.2fC\n", apiresponse.Main.TempMax)
	fmt.Printf("Feel's Like: %.2fC\n", apiresponse.Main.FeelsLike)
}
