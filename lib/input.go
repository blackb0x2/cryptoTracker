package lib

import (
	"fmt"
	"strconv"
	"strings"
)

// Function to ask boolean questions with user input
func BooleanQuestion(question string) bool {
	for {
		fmt.Print(question, " [y/N] ")
		fmt.Scan(&input)

		input = strings.ToLower(input)

		if input == "y" || input == "yes" {
			answer = true
		} else if input == "n" || input == "no" {
			answer = false
		} else {
			fmt.Println("Invalid input")
			continue
		}
		break
	}
	return answer
}

// Function to print slice with two columns
func PrintSlice(slice []string) {
	n = len(slice)

	for x = 0; x != n; x += 2 {
		a = fmt.Sprintf("%d %s", x+1, slice[x])

		if (n - x) == 1 {
			fmt.Printf("%-20s\n", a)
			break
		} else {
			b = fmt.Sprintf("%d %s", x+2, slice[x+1])
			fmt.Printf("%-20s%-20s\n", a, b)
		}
	}
}

// Function for multiple selection with integer input
func MultipleSelection(slice []string, question string) int {
	n = len(slice)
	PrintSlice(slice)

	for {
		x = 0
		fmt.Print(question, " [int] ")
		fmt.Scan(&x)

		if x <= n && x > 0 {
			break
		}
	}
	return x
}

// Function to get multiple string input values from the user
func MultipleInputString(petition string, minLength int, maxLength int) (inputList []string) {
	fmt.Println("Type 'x' or 'exit' when done, and 'r' or 'reset' to clear")

	for {
		fmt.Print(petition)
		fmt.Scan(&input)

		n = len(input)
		input = strings.ToLower(input)

		if input == "x" || input == "exit" {
			break
		} else if input == "r" || input == "repeat" {
			inputList = []string{}
			continue
		} else {
			if n < minLength || n > maxLength {
				PrintError("Invalid input")
				continue
			}
		}

		_, err := strconv.Atoi(input)
		if err == nil {
			PrintError("Invalid input")
			continue
		}
		inputList = append(inputList, input)
	}
	return
}

// Function to get multiple int input values from the user
func MultipleInputInt(petition string, slice []string) (inputList []string) {
	n = len(slice)

	fmt.Println("Type 'x' or 'exit' when done, and 'r' or 'reset' to clear")
	for {
		fmt.Print(petition)
		fmt.Scan(&input)

		input = strings.ToLower(input)

		if input == "x" || input == "exit" {
			break
		} else if input == "r" || input == "repeat" {
			inputList = []string{}
			continue
		}

		y, err := strconv.Atoi(input)
		if err != nil {
			PrintError("Invalid input")
			continue
		}

		if y > 0 && y < n+1 {
			inputList = append(inputList, input)
		} else {
			PrintError("Invalid input")
		}
	}
	return
}

// Function with the properties of MultipleSelection and MultipleInputInt
func MultipleInputSelection(slice []string, petition string) (sliceSelected []string) {
	PrintSlice(slice)

	inputList := MultipleInputInt(petition, slice)
	n = len(inputList)

	for x = 0; x != n; x++ {
		a = inputList[x]
		y, _ := strconv.Atoi(a)
		sliceSelected = append(sliceSelected, slice[y-1])
	}
	return
}
