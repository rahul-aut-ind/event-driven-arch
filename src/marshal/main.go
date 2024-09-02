package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	jsonString := `{"name":"John Doe","age":30,"address":{"street":"123 Elm St","city":"Gotham"}}`

	var person Person

	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Printf("Unmarshalled Struct with Nested Address: %+v\n", person)
	fmt.Printf("Unmarshalled Struct %+v\n", person.Name)
}

/****
Marshalling is the process of converting a Go object
(like a struct) into a format that can be easily stored or
transmitted, such as JSON or XML.

Unmarshalling is the reverse process: converting data from
formats like JSON or XML back into Go objects.

****/
