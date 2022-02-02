package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetDeviceIP() string {
	apiurl := "https://api.ipify.org?format=text"

	fmt.Println("Sending API call to:", apiurl)

	response, err := http.Get(apiurl) // Sending GET request
	if err != nil {
		fmt.Println("An error has occurred during the API call:", err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("An error has occurred when reading the response call:", err.Error())
		os.Exit(1)
	}

	ipaddress := string(data)
	return ipaddress
}

func GetLocation(ipaddress string) (float64, float64) {
	apiurl := "http://ip-api.com/json/"

	urltocall := fmt.Sprintf("%s%s", apiurl, ipaddress)
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

	var apiresponse LocationResponse

	err = json.Unmarshal(data, &apiresponse)
	if err != nil {
		fmt.Println("An error has occurred when processing the response json:", err.Error())
		os.Exit(1)
	}

	return apiresponse.Lat, apiresponse.Lon
}
