package main //1

import "fmt"

func main() { //3

	// fmt.Println("Hello World")

	// calculateAveragePacego(1, 1, 1)

}

func calculateAveragePace(kilometers float32, timeMinutes float32, timeSeconds float32) string {

	averagePace := 1000 / (kilometers * 1000 / timeMinutes)

	averagePaceString := fmt.Sprintf("%.0f", averagePace)

	averagePaceStringFormatted := averagePaceString + ":00 /km"

	return averagePaceStringFormatted

}
