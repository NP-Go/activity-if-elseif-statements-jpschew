package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//Insert your code here
	//Hint if a:= ?? ; condition {  }
	if number1 := 17; number1%2 == 0 {
		fmt.Println(strconv.Itoa(number1) + " is even.")
		if math.Abs(float64(number1)) < 9 {
			fmt.Println(strconv.Itoa(number1) + " has one digit.")
		} else {
			fmt.Println(strconv.Itoa(number1) + " has more than one digit.")
		}
	} else {
		fmt.Println(strconv.Itoa(number1) + " is odd.")
		if math.Abs(float64(number1)) < 9 {
			fmt.Println(strconv.Itoa(number1) + " has one digit.")
		} else {
			fmt.Println(strconv.Itoa(number1) + " has more than one digit.")
		}
	}
}
