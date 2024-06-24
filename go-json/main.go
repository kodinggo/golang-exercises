package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Phone  string `json:"hp"`
	Email  string `json:"email_address"`
	Gender string `json:"-"`
}

func main() {
	// Encode struct & map ke JSON
	person := Person{
		ID:     0,
		Name:   "",
		Phone:  "08123456",
		Email:  "john@gmail.com",
		Gender: "male",
	}
	byteJSON, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteJSON))

	userMap := map[string]any{
		"id":   1,
		"name": "David",
	}
	byteJSON, err = json.Marshal(userMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteJSON))

	// Decode JSON string ke struct/map
	jsonStr := `{"id":1,"name":"John","hp":"08123456","email_address":"john@gmail.com","gender":"male"}`

	var p Person
	err = json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		panic(err)
	}
	fmt.Println("email", p.Email)
	fmt.Println("gender", p.Gender)
}
