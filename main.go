package main

import (
	"fmt"
	"strconv"
)

// Calculator is a calculator.
// it can calculate simple addition, subtraction, multiplication and division.
func main () {
fmt.Println("Welcome to Isabella's calculator.")
	fmt.Println("It can add, subtract, multiply and divide whole numbers")
	fmt.Print("Enter your first number: ")
	var inputA string
	fmt.Scanln(&inputA)
	fmt.Println("Now choose what you would like to do with that number from the following options: + - * / ")
	var howAreWeCalculating string
	fmt.Scanln(&howAreWeCalculating)
	fmt.Print("Now please enter your final number: ")
	var inputB string
	fmt.Scanln(&inputB)
	fmt.Println(inputA, howAreWeCalculating, inputB)

	convertedinputA, _ := strconv.Atoi(inputA)
	convertedinputB, _ := strconv.Atoi(inputB)

	if howAreWeCalculating == "+" {
		fmt.Println(convertedinputA + convertedinputB)
	}
	if howAreWeCalculating == "-" {
	 	fmt.Println(convertedinputA - convertedinputB)
	}
	if howAreWeCalculating == "*" {
		fmt.Println(convertedinputA * convertedinputB)
	}
	if howAreWeCalculating == "/" {
		fmt.Println(convertedinputA / convertedinputB)
	}
}