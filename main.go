package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Events []struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}

type Holidays struct {
	EnglandAndWales struct {
		Events `json:"events"`
	} `json:"england-and-wales"`
}

func main() {
	refreshData := false
	var events Events
	if len(os.Args) >= 2 && os.Args[1] == "--refresh" {
		refreshData = true
	}

	if refreshData == true {
		res, err := http.Get("https://www.gov.uk/bank-holidays.json")

		if err != nil {
			panic(err)
		}

		if res.StatusCode != 200 {
			panic("Non 200 status code, something's wrong")
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

		events = holidays.EnglandAndWales.Events
	} else {
		file, err := os.Open("events.json")
		if err != nil {
			panic(err)
		}

		jsonBytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		var eventsCopy Events
		err = json.Unmarshal(jsonBytes, &eventsCopy)
		if err != nil {
			panic(err)
		}
		events = eventsCopy
	}

	today := time.Now()
	limit, err := time.Parse("2006-01-02", "2025-01-01")

	if err != nil {
		panic(err)
	}

	for _, e := range events {
		parsed, err := time.Parse("2006-01-02", e.Date)

		if err != nil {
			panic(err)
		}

		if parsed.After(today) && parsed.Before(limit) {
			fmt.Printf(
				"%s - %s\n",
				e.Title,
				e.Date,
			)
		}
	}
}
