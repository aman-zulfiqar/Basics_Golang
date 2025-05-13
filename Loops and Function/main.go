package main

import (
	"fmt"
)

// Function to separate even and odd numbers
func separateEvenOdd(numbers []int) ([]int, []int) {
	var evens []int
	var odds []int

	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}

	return evens, odds
}

func main() {
	var n int
	fmt.Print("How many numbers do you want to enter? ")
	fmt.Scan(&n)

	var numbers []int

	// Loop to input numbers
	for i := 0; i < n; i++ {
		var val int
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&val)
		numbers = append(numbers, val)
	}

	evens, odds := separateEvenOdd(numbers)

	fmt.Println("\nEven numbers:")
	for _, e := range evens {
		fmt.Printf("%d ", e)
	}

	fmt.Println("\n\nOdd numbers:")
	for _, o := range odds {
		fmt.Printf("%d ", o)
	}
}
