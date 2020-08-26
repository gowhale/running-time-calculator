package main //1

import (
	"fmt"
	"math"
)

func main() { //3

	// fmt.Println("Hello World")

	fmt.Println(calculateAveragePace(5, 24, 1))

}

func calculateAveragePace(kilometers float64, timeMinutes float64, timeSeconds float64) string {

	totalSeconds := (timeMinutes * 60) + timeSeconds

	secondsPerKilometer := totalSeconds / kilometers / 60

	secondsPerKilometerFloor := math.Floor(secondsPerKilometer)

	decimalSeconds := math.Mod(secondsPerKilometer, 1.0)

	averagePaceSeconds := decimalSeconds * 60

	averagePaceSecondsString := fmt.Sprintf("%02.0f", averagePaceSeconds)
	averagePaceMinutesString := fmt.Sprintf("%.0f", secondsPerKilometerFloor)

	return averagePaceMinutesString + ":" + averagePaceSecondsString + " /km"

}
