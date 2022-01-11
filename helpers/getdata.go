package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetWeatherByCity(key, cityname string) {
	apiurl := "http://api.openweathermap.org/data/2.5/weather?"

	urltocall := fmt.Sprintf("%sq=%s&appid=%s", apiurl, cityname, key)
	fmt.Println("Sending API call to:", urltocall)

	response, error := http.Get(urltocall) // Sending GET request
	if error != nil {
		fmt.Println("An error has occurred during the API call:", error)
		os.Exit(1)
	}

	data, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("An error has occurred when reading the response call:", error)
		os.Exit(1)
	}

	var apiresponse ByCityResponse

	error = json.Unmarshal(data, &apiresponse)
	if error != nil {
		fmt.Println("An error has occurred when processing the response json:", error)
		os.Exit(1)
	}

	fmt.Println("\nWeather data for", cityname)
	fmt.Println("Current Temperature:", apiresponse.Main.Temp)
	fmt.Println("Minimum Temperature:", apiresponse.Main.TempMin)
	fmt.Println("Maximum Temperature:", apiresponse.Main.TempMax)
	fmt.Println("Feel's Like:", apiresponse.Main.FeelsLike)

}
