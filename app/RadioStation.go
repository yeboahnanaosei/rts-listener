package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// RadioStation represents a single radio station
type RadioStation struct {
	ID        int
	Name      string
	Frequency string
	StreamURL string
}

// Save saves a new radio station to the database
func (r *RadioStation) Save() error {
	return nil
}

// Update updates the model
func (r *RadioStation) Update() error {
	return nil
}

// GetRecording gets a recording from the radio station
func (r *RadioStation) GetRecording(seconds string) (Recording, error) {
	rec := new(Recording)

	stationDirName := strings.Replace(r.Name, " ", "", -1)
	baseDir := os.Getenv("RECORDINGS_LOCATION")
	recordingLocation := filepath.Join(baseDir, stationDirName)

	_, err := os.Stat(recordingLocation)
	if os.IsNotExist(err) {
		err := os.MkdirAll(recordingLocation, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	outputFile := fmt.Sprint(
		strings.Replace(r.Name, " ", "", -1),
		"_",
		time.Now().Format("20060102_150405"),
		".mp3",
	)

	// recTitle := fmt.Sprint(r.Name, time.Now().Format("2006-01-02_15:04:05"))

	cmd := exec.Command(
		"ffmpeg",
		"-t", seconds,
		"-i", r.StreamURL,
		// "-metadata", fmt.Sprint("title=", recTitle),
		filepath.Join(recordingLocation, outputFile),
		"-y",
	)

	if err := cmd.Run(); err != nil {
		return *rec, err
	}

	rec.Filename = outputFile
	rec.RecordDateUnix = time.Now().Unix()
	rec.RecordDateHuman = time.Now().Format("2006-01-02 15:04:05")
	rec.RadioStation = *r
	return *rec, nil
}
