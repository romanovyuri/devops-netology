package main

import "fmt"
import "math"


func main() {
	fmt.Print("Enter meters to convert to feet: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input / 0.3048)
	fmt.Println("Feet ",math.Round(output*100)/100)
}
