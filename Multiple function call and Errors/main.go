package main

import (
    "errors"
    "fmt"
)

//with Making Of Function to Multiple Function Return it all practice Occure
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    // Example 1: Successful division

    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("10 / 2 = %.2f\n", result)
    }

    // Example 2: Division by zero (shows the power of multiple returns)
    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("10 / 0 = %.2f\n", result)
    }
}