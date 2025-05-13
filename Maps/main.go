package main

import "fmt"

func main(){

	fmt.Println("Studing about maps in golang")

	key_data := make(map[int]float64)
	key_data[22023686] = 100.22
	key_data[22003686] = 200.22
	key_data[22023233] = 3300.22
	key_data[22023679] = 3300.22
	key_data[37323683] = 3300.22

	for key, value := range key_data {
		fmt.Printf("for hash %v, the Amount is %v\n ", key, value)
	}

	//now doing the advance version of maps to store The transaction data in the hash
	//Example 
	
	hash_data := make(map[int]float32)
	var n int

	fmt.Println("How many Hashes you want to store")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		var key int
		var value float32

		fmt.Printf("Entre the Hash number %d (int)", i+1)
		fmt.Scan(&key)

		fmt.Printf("Entre the Amount in the bank %d (Float)", i+1)
		fmt.Scan(&value)

		hash_data[key] = value
	}

	for key, value := range hash_data {
		fmt.Printf("The hash number is %v, and the Amount is %v\n", key, value)
	}

	var searchkey int
	fmt.Println("Entre the hash amount by copy it if you want to find the ammount:")
	fmt.Scan(&searchkey)

	if val, exists := hash_data[searchkey]; exists {
		fmt.Printf("The Transaction of this hash %d is %v ", searchkey, val)
	} else {
		fmt.Printf("The transaction of this hash %d is not found", searchkey)
	}
}
