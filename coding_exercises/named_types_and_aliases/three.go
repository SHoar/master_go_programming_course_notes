package main

import "fmt"

type mile = float64
type kilometer = float64

const m2km = 1.609 // mile to km

func main() {
	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = mileBerlinToParis * m2km

	fmt.Printf(" The distance from Berlin To Paris is %.3f kilometers\n", kmBerlinToParis)
}
