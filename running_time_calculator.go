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

	strippedPace := (strings.Replace(pace, " /km", "", 2))
	splitPace := strings.Split(strippedPace, ":")

	minutesFloat, err := strconv.ParseFloat(splitPace[0], 64)
	if err != nil {
		fmt.Println(err)
	}

	secondsFloat, err := strconv.ParseFloat(splitPace[1], 64)
	if err != nil {
		fmt.Println(err)
	}

	secondsPerKilometer := (minutesFloat * 60) + secondsFloat

	benchMarks := make([]string, kilometers)

	for i := 0; i < kilometers; i++ {

		totalSeconds := float64((i + 1) * int(secondsPerKilometer))
		totalHoursDecimal := totalSeconds / 60

		onlyMinutes := math.Floor(totalHoursDecimal)
		onlySeconds := math.Mod(totalHoursDecimal, 1.0) * 60

		averagePaceSecondsString := fmt.Sprintf("%02.0f", onlySeconds)
		averagePaceMinutesString := fmt.Sprintf("%.0f", onlyMinutes)

		currentTime := averagePaceMinutesString + ":" + averagePaceSecondsString

		benchMarks[i] = currentTime
	}

	return benchMarks

}
