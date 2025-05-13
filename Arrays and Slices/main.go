package main

import (
	"fmt"
	"sort"
)


func main() {
    // Array have fixed size and complete remain same at compile time 
    var numbers [5]int
    fmt.Println("Empty array:", numbers)

    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30
    numbers[3] = 40
    numbers[4] = 50
    fmt.Println("Array after assigning values:", numbers)

    fmt.Println("The first element is:", numbers[0])

	//Also covering Basic level of loops within-it

    fmt.Println("All array elements:")
    for i := 0; i < len(numbers); i++ {
        fmt.Printf("Index %d = %d\n", i, numbers[i])
    }

    fruits := [3]string{"Apple", "Banana", "Cherry"}
    fmt.Println("Fruits:", fruits)

    //Using range loop 
    fmt.Println("Fruits using range:")
    for index, value := range fruits {
        fmt.Printf("Index %d = %s\n", index, value)
    }

	//Slices have dynamic size and grow and shrink with it

	var Universities = []string{"LGU", "PU", "ITU", "GCU"}
	fmt.Printf("The type of Slice are %T\n", Universities)
	Universities = append(Universities, "FAST", "NCA")
	fmt.Println("The Total Universities are :", Universities)

	//append 
	Universities = Universities[1:4]
	fmt.Println(Universities)

	//best way to use a slice Is to Create a slice with make function with Length and Capicity
	lang := make([]int, 5)
	lang[0] = 22
	lang[1] = 223
	lang[2] = 2233
	lang[3] = 22333
	lang[4] = 223333

	lang = append(lang, 33, 334, 3344)
	fmt.Println(lang)

	fmt.Println(len(lang))
	//sorting Slices in the order
	sort.Ints(lang)
	fmt.Println(lang)
	//removing values from the slices
	var index int = 2
	lang = append(lang[:index], lang[index+1:]...)
	fmt.Println(lang)
}
