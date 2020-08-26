package main //1

import "fmt"

func main() { //3

	// fmt.Println("Hello World")

	// calculateAveragePacego(1, 1, 1)

}

func calculateAveragePace(kilometers float32, timeMinutes float32, timeSeconds float32) string {

	averagePaceMinutes := 1000 / (kilometers * 1000 / timeMinutes)
	averagePaceSeconds := 1000 / (kilometers * 1000 / timeSeconds)

	averagePaceMinutesString := fmt.Sprintf("%.0f", averagePaceMinutes)
	averagePaceSecondsString := fmt.Sprintf("%.0f", averagePaceSeconds)

	averagePaceStringFormatted := averagePaceMinutesString + ":0" + averagePaceSecondsString + " /km"

	return averagePaceStringFormatted

}
