package app

import (
	"log"
)

// Recording represents a single recording that the system has made
type Recording struct {
	Filename        string
	RecordDateUnix  int64
	RecordDateHuman string
	RadioStation
}

// Save saves a new recording to the database
func (r *Recording) Save() error {
	stmt, err := DB.Prepare(
		`INSERT INTO recordings (
			date_recorded, date_recorded_unix, radio_station_id
		) VALUES (
			?, ?, ?
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(r.RecordDateHuman, r.RecordDateUnix, r.RadioStation.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates the model
func (r *Recording) Update() error {
	return nil
}
