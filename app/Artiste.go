package app

// Artiste represents an owner of a song
type Artiste struct {
	ISNI string
	Name string
}

// Save saves a new artiste to the database
func (a *Artiste) Save() error {
	return nil
}

// Update updates the model
func (a *Artiste) Update() error {
	return nil
}
