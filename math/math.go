package main

import (
	"fmt"
	"math"
	"time"
	"log"
)


func main() {
	i1, i2, i3 := 12, 41, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum is: ", intSum)

	f1, f2, f3 := 24.6, 67.5, 43.2
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum is: ", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Float sum is now: ", floatSum)

	floatSum = math.Round(floatSum)
	fmt.Println("Float sum is now: ", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)

	n := time.Now()
	fmt.Println("This time right now is", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go was launched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, err := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	if err != nil {
		log.Fatal(err) // Handle the error appropriately
	}
	fmt.Printf("The type of parsed time is %T\n", parsedTime)

}
