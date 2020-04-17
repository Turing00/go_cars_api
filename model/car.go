package model

import (
	"time"
)

// Car describe car struct
type Car struct {
	ID           int       `json:"id"`
	Manufacturer string    `json:"manufacturer"`
	Design       string    `json:"design"`
	Style        string    `json:"style"`
	Doors        uint8     `json:"doors"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Cars describe a list of car
type Cars []Car
