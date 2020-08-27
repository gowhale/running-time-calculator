package main //1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() { //3

	var currentRun Run

	currentRun.setRunDistance()
	currentRun.setRunPace()

	fmt.Println("Run attributes:")
	fmt.Println(currentRun.kilometers)
	fmt.Println(currentRun.pace)

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

type Pace struct {
	minutes float64
	seconds float64
}

func convertStringToPaceStruct(pace string) Pace {
	// takes a string input to struct
	// mm:ss /km

	var customPace Pace

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

	customPace.minutes = minutesFloat
	customPace.seconds = secondsFloat

	return customPace

}

type Run struct {
	kilometers int
	pace       string
}

func (r *Run) setRunDistance() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many kilometers would you like to go: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	strippedText := (strings.Replace(text, "\n", "", 2))

	k, err := strconv.Atoi(strippedText)

	if err != nil {
		fmt.Println(k)
		fmt.Println("INVALID INPUT PLEASE TRY AGAIN")
		// r.setRunDistance()
	} else {
		r.kilometers = k
	}

}

func (r *Run) setRunPace() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many kilometers would you like to go: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	strippedText := (strings.Replace(text, "\n", "", 2))

	r.pace = strippedText

}

// func setRunKilometer() Run {

// 	var r Run

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("How many kilometers would you like to go: ")
// 	text, _ := reader.ReadString('\n')
// 	fmt.Println(text)

// 	strippedText := (strings.Replace(text, "\n", "", 2))

// 	k, err := strconv.Atoi(strippedText)

// 	if err != nil {
// 		fmt.Println("INVALID INPUT PLEASE TRY AGAIN")
// 		return getRun()
// 	}

// 	r.kilometers = k

// 	fmt.Println("Kilometers entered")
// 	fmt.Println(r.kilometers)

// 	return r

// }
