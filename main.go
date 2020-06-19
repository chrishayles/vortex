package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(VortexStringWithNumbers("Test string, yo."))
	fmt.Println(VortexNumbersFromString("06/18/2020"))
}

//VortexAdd processes a multi-digit int down to a single digit int
func VortexAdd(number int) int {
	s := strconv.Itoa(number)
	total := 0

	for _,c := range s {
		x, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println(err)
		}
		total = total + x
	}

	if total > 9 {
		total = VortexAdd(total)
	}

	return total
}

//VortexString returns single-digit length of string minus punctuation, numbers, and spaces.
func VortexString(s string) int {
	clean := removePuncAndNumbers(s)
	l := len(clean)

	return VortexAdd(l)
}

//VortexStringWithNumbers returns single-digit length of string minus punctuation and spaces.
func VortexStringWithNumbers(s string) int {
	clean := removePunc(s)
	l := len(clean)

	return VortexAdd(l)
}


func removePunc(s string) string {
	
	output := ""
	
	for _,c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			output = fmt.Sprintf("%s%s", output, string(c))
		}
	}

	return output
}

func removePuncAndNumbers(s string) string {
	
	output := ""
	
	for _,c := range s {
		if unicode.IsLetter(c) {
			output = fmt.Sprintf("%s%s", output, string(c))
		}
	}

	return output
}

//VortexNumbersFromString returns single-digit sum of all numbers within a string. Ignores non-number characters.
func VortexNumbersFromString(s string) int {
	
	output := 0
	
	for _,c := range s {
		if unicode.IsNumber(c) {
			output = output + int(c - '0')
		}
	}

	return VortexAdd(output)
}