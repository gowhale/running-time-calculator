package main //1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Run is the 'object' which holds all needed attributes about the run
type Run struct {
	kilometers          int
	time                string
	paceStruct          Pace
	paceFormatted       string
	kilometerBenchmarks []string
}

// Pace is used within the Run struct to help calculate wanted pace etc
type Pace struct {
	minutes float64
	seconds float64
}

// Main will ask the user for inputs then print out metrics to achieve the running goal
func main() {

	var currentRun Run

	currentRun.setRunDistance()
	currentRun.setRunTime()

	currentRun.paceStruct = convertStringToPaceStruct(currentRun.time)

	kilometerFloat := float64(currentRun.kilometers)

	currentRun.paceFormatted = (calculateAveragePace(kilometerFloat, currentRun.paceStruct.minutes, currentRun.paceStruct.seconds))

	currentRun.kilometerBenchmarks = calculateKilometerTimes(currentRun.kilometers, currentRun.paceFormatted)

	displayMetrics(currentRun)

}

// calculateAveragePace takes in desired metrics and returns an average pace for the user to follow to achieve thier dreams
func calculateAveragePace(kilometers float64, timeMinutes float64, timeSeconds float64) string {

	totalSeconds := (timeMinutes * 60) + timeSeconds

	decimalMinutesPerKilometer := totalSeconds / kilometers / 60
	decimalSeconds := math.Mod(decimalMinutesPerKilometer, 1.0)

	averagePaceMinutes := math.Floor(decimalMinutesPerKilometer)
	averagePaceSeconds := decimalSeconds * 60

	averagePaceSecondsString := fmt.Sprintf("%02.0f", averagePaceSeconds)
	averagePaceMinutesString := fmt.Sprintf("%.0f", averagePaceMinutes)

	return averagePaceMinutesString + ":" + averagePaceSecondsString + " /km"

}

// calculateKilometerTimes takes in the amount of km the user wants to run and then creates a list of times for each km
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

	decimalMinutesPerKilometer := (minutesFloat * 60) + secondsFloat

	benchMarks := make([]string, kilometers)

	for i := 0; i < kilometers; i++ {

		totalSeconds := float64((i + 1) * int(decimalMinutesPerKilometer))
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

// convertStringToPaceStruct takes a string of pace and converts it to the struct Pace
func convertStringToPaceStruct(pace string) Pace {

	var customPace Pace

	splitPace := strings.Split(pace, ":")

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

// setRunDistance asks the user to set a Run's distance attribute
func (r *Run) setRunDistance() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many kilometers would you like to go: ")
	text, _ := reader.ReadString('\n')

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

// setRunDistance asks the user to set a Run's time attribute
func (r *Run) setRunTime() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What time would you like to achieve: ")
	text, _ := reader.ReadString('\n')

	strippedText := (strings.Replace(text, "\n", "", 2))

	r.time = strippedText

}

// displayMetrics takes in a run struct and then prints out metrics in a friendly format
func displayMetrics(r Run) {
	fmt.Println("-----------------------------------------")
	fmt.Println()

	concatanatedLine := fmt.Sprint("If you wish to run ", r.kilometers, "KM in ", r.time)
	fmt.Println(concatanatedLine)

	concatanatedLine = fmt.Sprint("Then your average pace should be ", r.paceFormatted)
	fmt.Println(concatanatedLine)
	fmt.Println()

	for km := 0; km < len(r.kilometerBenchmarks); km++ {
		concatanatedLine = fmt.Sprint(km+1, " - ", r.kilometerBenchmarks[km])
		fmt.Println(concatanatedLine)
	}
	fmt.Println()
}
