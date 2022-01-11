package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetWeatherByCity(key, cityname string) {
	apiurl := "http://api.openweathermap.org/data/2.5/weather?"

	urltocall := fmt.Sprintf("%sq=%s&appid=%s", apiurl, cityname, key)
	fmt.Println(urltocall)

	response ,error := http.Get(urltocall) // Sending GET request
	if error != nil {
		fmt.Println("An error has occurred during the API call:", error)
		os.Exit(1)
	}

	data, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("An error has occurred when reading the response call:", error)
		os.Exit(1)
	}

	fmt.Println(data)
}
