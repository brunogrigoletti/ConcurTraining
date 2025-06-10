package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://sapit-home-prod-004.launchpad.cfapps.eu10.hana.ondemand.com/site#my-events&/weeklyReservationTimetable/14955650"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	content := string(body)

	fmt.Println(content)
}
