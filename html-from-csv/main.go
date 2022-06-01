package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Define the struct for each record (line) from the CSV
type Record struct {
	Date                                     time.Time
	Open, High, Low, Close, Volume, AdjClose float64
}

// Define type that holds a slice of records
type Records []Record

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// Parse a CSV file
	data, err := os.ReadFile("data.csv")
	if err != nil {
		panic(err)
	}

	var records Records

	// Read each line of the CSV and create a struct with the fields
	for i, line := range strings.Split(string(data), "\n") {
		// Skip the first line
		if i == 0 {
			continue
		}

		var record Record

		if date, err := time.Parse("2006-01-02", strings.Split(line, ",")[0]); err == nil {
			record.Date = date
		}
		if open, err := strconv.ParseFloat(strings.Split(line, ",")[1], 64); err == nil {
			record.Open = open
		}
		if high, err := strconv.ParseFloat(strings.Split(line, ",")[2], 64); err == nil {
			record.High = high
		}
		if low, err := strconv.ParseFloat(strings.Split(line, ",")[3], 64); err == nil {
			record.Low = low
		}
		if close, err := strconv.ParseFloat(strings.Split(line, ",")[4], 64); err == nil {
			record.Close = close
		}
		if volume, err := strconv.ParseFloat(strings.Split(line, ",")[5], 64); err == nil {
			record.Volume = volume
		}
		if adjClose, err := strconv.ParseFloat(strings.Split(line, ",")[6], 64); err == nil {
			record.AdjClose = adjClose
		}

		records = append(records, record)
	}

	for _, record := range records {
		fmt.Println(record.Date, record.Open)
	}

	// Pass the data structure to the template
	err = tpl.Execute(os.Stdout, records)
	if err != nil {
		log.Fatalln(err)
	}
}
