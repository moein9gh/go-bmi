package main

import "fmt"

// get inputs and calculate BMI
func main() {

	fmt.Println("BMI calculator")
	fmt.Println("--------------------------")

	//getting the inputs
	var weight float32
	var height float32

	fmt.Print("Please enter your weight (kg) : ")
	fmt.Scan(&weight)

	fmt.Print("Please enter your height (m) :")
	fmt.Scan(&height)

}
