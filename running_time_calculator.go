package main //1

import (
	"fmt"
	"math"
	"strconv"
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
	// fmt.Println(kilometers)
	// fmt.Println(pace)
	strippedPace := (strings.Replace(pace, " /km", "", 2))
	splitPace := strings.Split(strippedPace, ":")
	// fmt.Println()

	minutesFloat, err := strconv.ParseFloat(splitPace[0], 64)
	secondsFloat, err := strconv.ParseFloat(splitPace[1], 64)

	// fmt.Println(minutesFloat)
	// fmt.Println(secondsFloat)
	fmt.Println(err)

	totalSeconds := (minutesFloat * 60) + secondsFloat

	// totalSeconds := splitPace[0]*60 + splitPace[1]

	benchMarks := make([]string, kilometers)

	for i := 0; i < kilometers; i++ {

		totalSeconds := float64((i + 1) * int(totalSeconds))
		decimalTime := totalSeconds / 60

		decimalMinutes := math.Floor(decimalTime)
		decimalSeconds := math.Mod(decimalTime, 1.0) * 60
		fmt.Println(decimalSeconds)

		// totalSecondsString := strconv.Itoa(decimalTime)

		averagePaceSecondsString := fmt.Sprintf("%02.0f", decimalSeconds)
		averagePaceMinutesString := fmt.Sprintf("%.0f", decimalMinutes)

		currentTime := averagePaceMinutesString + ":" + averagePaceSecondsString

		benchMarks[i] = currentTime
	}

	return benchMarks
}
