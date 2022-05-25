package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"Weather-Checker-Go/utils"
)

func main() {
	// Logging to a file
	filehandler, err := os.OpenFile("weatherlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer filehandler.Close() // Closes file at the end of execution
	log.SetOutput(filehandler)

	var apikey string
	var option string
	var optionpin string
	var cityname string
	var pincode string
	var apiresponse utils.WeatherResponse

	fmt.Print("Enter your API key: ")
	fmt.Scan(&apikey)
	apikey = strings.TrimSpace(apikey)
	fmt.Println("The API Key being used is:", apikey)
	log.Println("API Key:", apikey)

	fmt.Print("Continue (Y/N): ")
	fmt.Scan(&option)
	option = strings.TrimSpace(strings.ToUpper(option))

	if option == "N" {
		os.Exit(0)
	} else if option != "Y" {
		fmt.Println("Enter a valid option!")
		os.Exit(0)
	}

	fmt.Print("Do you want to search by City/Location/PIN (city/location/pin): ")
	fmt.Scan(&optionpin)
	optionpin = strings.TrimSpace(strings.ToLower(optionpin))

	if optionpin == "city" {
		fmt.Print("Enter city name: ")
		fmt.Scan(&cityname)
		apiresponse = utils.GetWeatherByCity(apikey, cityname)
	} else if optionpin == "pin" {
		fmt.Print("Enter PIN code: ")
		fmt.Scan(&pincode)
		apiresponse = utils.GetWeatherByPincode(apikey, pincode)
	} else if optionpin == "location" {
		ip := utils.GetDeviceIP()
		lat, lon := utils.GetLocation(ip)
		apiresponse = utils.GetWeatherByLatLon(apikey, lat, lon)
	} else {
		fmt.Println("Enter a valid option!")
		os.Exit(0)
	}

	header := fmt.Sprintf("%s - %.2f°C - %s", apiresponse.Name, apiresponse.Main.Temp, utils.GetConditionIcon(apiresponse))
	length := len(header)

	fmt.Println()
	fmt.Println(strings.Repeat("-", length))
	fmt.Println(header)
	fmt.Println(strings.Repeat("-", length))
	fmt.Println()

	fmt.Printf("Conditions: %s\n", apiresponse.Weather[0].Description)
	fmt.Printf("Minimum Temperature: %.2fC\n", apiresponse.Main.TempMin)
	fmt.Printf("Maximum Temperature: %.2fC\n", apiresponse.Main.TempMax)
	fmt.Printf("Feel's Like: %.2fC\n", apiresponse.Main.FeelsLike)
	fmt.Printf("Pressure: %dhPa\n", apiresponse.Main.Pressure)
	fmt.Printf("Humidity: %d%s\n", apiresponse.Main.Humidity, "%")

	sunrise := time.Unix(int64(apiresponse.Sys.Sunrise), 0).Local().Format(time.Kitchen)
	sunset := time.Unix(int64(apiresponse.Sys.Sunset), 0).Local().Format(time.Kitchen)
	fmt.Printf("Sunrise: %s\n", sunrise)
	fmt.Printf("Sunset: %s\n", sunset)
	fmt.Println()
}
