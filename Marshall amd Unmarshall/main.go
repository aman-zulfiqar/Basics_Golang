package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Password   string   `json:"-"`
	Permission []string `json:"roles"`
}

func main() {

	//marshalling
	u := User{
		ID:         2,
		Name:       "Aman",
		Age:        21,
		Password:   "356900",
		Permission: []string{"student", "swimmer"},
	}

	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Ther Error occure in This Program:", err)
		panic(err)
	}
	fmt.Println(string(b))

	//Unmarshalling

	jsonData := `{"id":34 , "name":"Muhamad_Aman" , "age":23 , "roles":["MMA","Developer"]}`

	var user User

	error := json.Unmarshal([]byte(jsonData), &user)
	if error != nil {
		fmt.Println("The error which is happening is:", error)
		panic(error)
	}
	fmt.Println("ID", user.ID)
	fmt.Println("Name", user.Name)
	fmt.Println("Age", user.Age)
	fmt.Println("Skills", user.Permission)
	fmt.Println("Password", user.Password)
}
