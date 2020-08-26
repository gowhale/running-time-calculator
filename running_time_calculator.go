package main //1

import (
	"fmt"
	"math"
	"strings"
)

func main() { //3

	// fmt.Println("Hello World")

	fmt.Println(calculateKilometerTimes(5, "5:00 /km"))

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

func calculateKilometerTimes(kilometers int, pace string) []string {
	fmt.Println(kilometers)
	fmt.Println(pace)
	strippedPace := (strings.Replace(pace, " /km", "", 2))
	splitPace := strings.Split(strippedPace, ":")
	fmt.Println(splitPace[1])

	// totalSeconds := splitPace[0]*60 + splitPace[1]

	benchMarks := make([]string, kilometers)

	for i := 0; i < kilometers; i++ {
		benchMarks[i] = "epic"
	}

	return benchMarks
}
