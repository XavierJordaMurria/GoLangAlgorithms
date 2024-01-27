package main

import (
	"fmt"
	"os"
	"strconv"
)

func smallerThan13(theNumber int) (theLetters string) {
	switch theNumber {
	case 0:
		theLetters = "zero"
	case 1:
		theLetters = "one"
	case 2:
		theLetters = "two"
	case 3:
		theLetters = "three"
	case 4:
		theLetters = "four"
	case 5:
		theLetters = "five"
	case 6:
		theLetters = "six"
	case 7:
		theLetters = "seven"
	case 8:
		theLetters = "eight"
	case 9:
		theLetters = "nine"
	case 10:
		theLetters = "ten"
	case 11:
		theLetters = "eleven"
	case 12:
		theLetters = "twelve"
	case 13:
		theLetters = "thirteen"
	}
	return
}

func twoDigits(theNumber int) (theLetters string) {
	
	if (theNumber >= 20) && (theNumber <= 29) {
		theLetters = "twenty"
	} else if (theNumber >= 30) && (theNumber <= 39) {
		theLetters = "thirty"
	} else if (theNumber >= 40) && (theNumber <= 49) {
		theLetters = "forty"
	} else if (theNumber >= 50) && (theNumber <= 59) {
		theLetters = "fifty"
	} else if (theNumber >= 60) && (theNumber <= 99) {
		theLetters = smallerThan13(theNumber/10) + "ty"
	}

	if theNumber%10 != 0 {
		theLetters += "-" + smallerThan13(theNumber%10)
	}
	return
}

func numberToLetter(theNumber int)(theLetters string) {
	
	if (theNumber <= 13) {
		theLetters = smallerThan13(theNumber)
	} else if (theNumber >= 14) && (theNumber <= 19) {
		theLetters = smallerThan13(theNumber%10) + "teen"
	} else if (theNumber >= 20) && (theNumber <= 99) {
		theLetters = twoDigits(theNumber)
	} else if (theNumber >= 100) && (theNumber <= 999) {
		theLetters = smallerThan13(theNumber/100) + " hundred"
		if theNumber%100 != 0 {
			theLetters += " " + numberToLetter(theNumber%100)
		}
	} else if (theNumber >= 1000) && (theNumber <= 999999) {
		theLetters = numberToLetter(theNumber/1000) + " thousand"
		if theNumber%1000 != 0 {
			theLetters += " " + numberToLetter(theNumber%1000)
		}
	} else if (theNumber >= 1000000) && (theNumber <= 999999999) {
		theLetters = numberToLetter(theNumber/1000000) + " million"
		if theNumber%1000000 != 0 {
			theLetters += " " + numberToLetter(theNumber%1000000)
		}
	} else if (theNumber >= 1000000000) && (theNumber <= 999999999999) {
		theLetters = numberToLetter(theNumber/1000000000) + " billion"
		if theNumber%1000000000 != 0 {
			theLetters += " " + numberToLetter(theNumber%1000000000)
		}
	}

	return
}
func main() {
	fmt.Println("Hello World")

	argsWithoutProg := os.Args[1:]

	// string to int
	theNumber, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}

	// var under10 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// for i, v := range under10 {
	// 	fmt.Println(i, v)
	// 	fmt.Printf("%v/%v = %v\n", theNumber, v, theNumber/v)
	// 	fmt.Printf("%v%%%v = %v\n", theNumber, v, theNumber%v)
	// 	fmt.Println("--")
	// }

	var theLetters string = numberToLetter(theNumber)
	fmt.Println(theLetters)
}
