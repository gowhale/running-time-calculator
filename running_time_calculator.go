package main //1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Run struct {
	kilometers          int
	time                string
	paceStruct          Pace
	paceFormatted       string
	kilometerBenchmarks []string
}

type Pace struct {
	minutes float64
	seconds float64
}

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

func convertStringToPaceStruct(pace string) Pace {
	// takes a string input to struct
	// mm:ss /km

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

func (r *Run) setRunTime() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What time would you like to achieve: ")
	text, _ := reader.ReadString('\n')

	strippedText := (strings.Replace(text, "\n", "", 2))

	r.time = strippedText

}

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
