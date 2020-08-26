package main //1

import "fmt"

func main() { //3

	// fmt.Println("Hello World")

	// calculate_average_pace(1, 1, 1)

}

func calculate_average_pace(kilometers float32, time_minutes float32, time_seconds float32) string {

	fmt.Println(kilometers)
	fmt.Println(time_minutes)
	fmt.Println(time_seconds)

	average_pace := 1000 / (kilometers * 1000 / time_minutes)

	average_pace_string := fmt.Sprintf("%.0f", average_pace)

	average_pace_string_formatted := average_pace_string + ":00 /km"
	print(average_pace)

	return average_pace_string_formatted

}
