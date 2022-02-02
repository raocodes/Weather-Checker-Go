package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/raocodes/WeatherCheckerGo/utils"
)

func main() {
	var apikey string
	var option byte
	var optionpin string
	var cityname string
	var pincode string

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

	fmt.Print("Do you want to search by City/Location/Indian PIN code (city/location/pin): ")
	fmt.Scan(&optionpin)
	optionpin = strings.ToLower(optionpin)

	if optionpin == "city" {
		fmt.Print("Enter city name: ")
		fmt.Scan(&cityname)
		utils.GetWeatherByCity(apikey, cityname)
	} else if optionpin == "pin" {
		fmt.Print("Enter PIN code: ")
		fmt.Scan(&pincode)
		utils.GetWeatherByPincode(apikey, pincode)
	} else if optionpin == "location" {
		ip := utils.GetDeviceIP()
		lat, lon := utils.GetLocation(ip)
		utils.GetWeatherByLatLon(apikey, lat, lon)
	} else {
		fmt.Println("Enter a valid option!")
		os.Exit(0)
	}
}
