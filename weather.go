package main

import (
	"fmt"
	"os"

	"github.com/raocodes/WeatherCheckerGo/helpers"
)

func main() {
	var apikey string
	var option byte
	var cityname string

	fmt.Print("Enter your API key: ")
	fmt.Scan(&apikey)
	fmt.Println("The API Key being used is:", apikey)

	fmt.Print("Continue (Y/N): ")
	fmt.Scanf("%c\n", &option)

	if option == 'N' || option == 'n' {
		os.Exit(0)
	} else if option != 'Y' && option != 'y' {
		fmt.Println("Enter a valid option!")
		os.Exit(0)
	}

	fmt.Print("Enter city name: ")
	fmt.Scan(&cityname)

	helpers.GetWeatherByCity(apikey, cityname)
}
