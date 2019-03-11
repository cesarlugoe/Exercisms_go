package space

import (
	"math"
)

// Planet type to determine where the age is calculated
type Planet string

// Age Function to determine the age
func Age(s float64, p Planet) float64 {
	ey := s / 31557600
	var a float64

	if p == "Earth" {
		a = math.Round(ey*100) / 100
	}
	if p == "Mercury" {
		a = ey / 0.2408467
	}

	if p == "Venus" {
		a = ey / 0.61519726
	}

	if p == "Mars" {
		a = ey / 1.8808158
	}

	if p == "Jupiter" {
		a = ey / 11.862615
	}

	if p == "Saturn" {
		a = ey / 29.447498
	}

	if p == "Uranus" {
		a = ey / 84.016846
	}

	if p == "Neptune" {
		a = ey / 164.79132
	}

	return a
}
