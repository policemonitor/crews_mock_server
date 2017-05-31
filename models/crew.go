package models

import "math/rand"


type Crew struct {
	VinNumber string  `json:"vin_number"`
	Latitude  float32 `json:"latitude,string"`
	Longitude float32 `json:"longitude,string"`
}

func newMovement() float32 {
	return (0.5 - rand.Float32()) / 200
}

func (c *Crew) Move() {
	c.Latitude += newMovement()
	c.Longitude += newMovement()
}