package main

import "fmt"

//main gets inputs
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

	fmt.Print("your BMI is : ", bmiCalculator(height, weight))

}

//bmiCalculator gets inputs and calculate the BMI
func bmiCalculator(height, weight float32) (bmi float32) {
	bmi = weight / (height * height)
	return
}
