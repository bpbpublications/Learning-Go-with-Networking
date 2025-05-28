package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeviceInfo represents the structured data we expect to receive from the API.
type DeviceInfo struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Status   string `json:"status"`
}

func main() {
	// URL of the device's API endpoint
	url := "http://192.168.1.1/api/info"

	// Send an HTTP GET request to the device's API
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		// Read and parse the response body
		data, _ := ioutil.ReadAll(response.Body)
		var deviceInfo DeviceInfo
		json.Unmarshal(data, &deviceInfo)

		// Print the device information
		fmt.Println("Device Hostname:", deviceInfo.Hostname)
		fmt.Println("Device IP:", deviceInfo.IP)
		fmt.Println("Device Status:", deviceInfo.Status)
	}
}
