package model

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// Magazine represents a magazine
type Magazine struct {
	MagazineID uuid.UUID
	Title      string
	Issue      string
}

// NewMagazine create a new magazine instance
func NewMagazine(magazineID uuid.UUID, title string, issue string) *Magazine {
	return &Magazine{magazineID, title, issue}
}

// NewMagazineFromJSON parses a json byte array in order to create a new instance of Magazine
func NewMagazineFromJSON(bytes []byte) (Magazine, error) {
	var magazine Magazine
	err := json.Unmarshal(bytes, &magazine)
	return magazine, err
}

func (m *Magazine) String() string {
	return fmt.Sprintf("Magazine[MagazineID:%s,Title:%s,Issue:%s]", m.MagazineID, m.Title, m.Issue)
}

// JSON converts this object to json
func (m *Magazine) JSON() ([]byte, error) {
	return json.Marshal(m)
}
