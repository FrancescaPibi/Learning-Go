package main

import "fmt"

func main() {
	var example string
	example = "best example"
	fmt.Println(example, example == "best example", example == "not the best example")
	example = "not the best example"
	fmt.Println(example, example == "best example", example == "not the best example")

	var firstInput int
	fmt.Println("Please give me a number:")
	fmt.Scan(&firstInput)
	fmt.Println("The number you put in is", firstInput)

	var secondInput int
	fmt.Println("Please give me a number:")
	fmt.Scan(&secondInput)
	fmt.Println("The number you put in is", secondInput)
	
	var thirdInput string
	fmt.Println("Please choose between + and -:")
	fmt.Scan(&thirdInput)
	fmt.Println("You chose", thirdInput) 

	if (thirdInput == "+") {
		fmt.Println(firstInput, "+", secondInput, "=", firstInput + secondInput)
	} else if (thirdInput == "-") {
		fmt.Println(firstInput, "-", secondInput, "=", firstInput - secondInput)
	} else {
		fmt.Println("Invalid Operator")
	}
}