package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rts/listener/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	radioStations, err := app.DB.Query("SELECT * from radio_stations")
	if err != nil {
		log.Fatal(err)
	}

	for radioStations.Next() {
		station := new(app.RadioStation)
		err := radioStations.Scan(&station.ID, &station.Name, &station.Frequency, &station.StreamURL)
		if err != nil {
			log.Fatal(err)
		}

		recording, err := station.GetRecording("5")
		if err != nil {
			log.Fatalf("There was an error trying to get recording from: %s\n %v", station.Name, err)
		}

		fmt.Println("\n", recording)

		if err = recording.Save(); err != nil {
			fmt.Printf("There was an error trying to save a recording to the database %v", err)
		}
	}
	defer radioStations.Close()
	defer app.DB.Close()
}
