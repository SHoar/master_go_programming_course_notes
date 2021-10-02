package main

import "fmt"

type km = int64
type distance km
type mps = int64
type speed mps

func main() {
	const SunToEarth distance = 149_600_000 * 1000
	const lightSpeed speed = 299_792_458

	// how long does it take for sunlight to reach the earth?
	// time = distance / speed

	time := speed(SunToEarth) / lightSpeed

	minutes := time / 60

	fmt.Printf("Supposedly, it would take %v seconds for sunlight to reach Earth, or %v minutes", time, minutes)

}
