package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Holidays struct {
	EnglandAndWales struct {
		Events []struct {
			Title string `json:"title"`
			Date  string `json:"date"`
		} `json:"events"`
	} `json:"england-and-wales"`
}

func main() {
	res, err := http.Get("https://www.gov.uk/bank-holidays.json")

	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		panic("Non 200 status code")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var holidays Holidays
	err = json.Unmarshal(body, &holidays)
	if err != nil {
		fmt.Println(err)
	}

	events := holidays.EnglandAndWales.Events

	for _, e := range events {
		fmt.Printf(
			"%s - %s\n",
			e.Title,
			e.Date,
		)
	}
}