package main

/************
* author : Rahul Upadhyay
* Peer-graded Assignment: Module 4 Activity
*************/

import (
	"fmt"
	"strings"
)

type (
	Animal interface {
		Eat()
		Move()
		Speak()
	}

	Cow struct {
		food       string
		locomotion string
		noise      string
	}

	Bird struct {
		food       string
		locomotion string
		noise      string
	}

	Snake struct {
		food       string
		locomotion string
		noise      string
	}
)

func (a *Cow) Eat() {
	println(a.food)
}

func (a *Cow) Move() {
	println(a.locomotion)
}

func (a *Cow) Speak() {
	println(a.noise)
}

func (a *Bird) Eat() {
	println(a.food)
}

func (a *Bird) Move() {
	println(a.locomotion)
}

func (a *Bird) Speak() {
	println(a.noise)
}
func (a *Snake) Eat() {
	println(a.food)
}

func (a *Snake) Move() {
	println(a.locomotion)
}

func (a *Snake) Speak() {
	println(a.noise)
}

func NewCow() *Cow {
	return &Cow{food: "grass", locomotion: "walk", noise: "moo"}
}
func NewBird() *Bird {
	return &Bird{food: "worms", locomotion: "fly", noise: "peep"}
}

func NewSnake() *Snake {
	return &Snake{food: "mice", locomotion: "slither", noise: "hsss"}
}

func (a *Cow) processInfo(input string) {
	switch strings.ToLower(input) {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		println("unknown information requested for Cow. Try again!! ( valid values > eat move speak")
	}
}

func (b *Bird) processInfo(input string) {
	switch strings.ToLower(input) {
	case "eat":
		b.Eat()
	case "move":
		b.Move()
	case "speak":
		b.Speak()
	default:
		println("unknown information requested for Bird. Try again!! ( valid values > eat move speak")
	}
}

func (s *Snake) processInfo(input string) {
	switch strings.ToLower(input) {
	case "eat":
		s.Eat()
	case "move":
		s.Move()
	case "speak":
		s.Speak()
	default:
		println("unknown information requested for Snake. Try again!! ( valid values > eat move speak")
	}
}

func main() {
	/*******
	Write a program which allows the user to create a set of animals and to get information about those animals.
	Each animal has a name and can be either a cow, bird, or snake.
	With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created.
	Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type,
	but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

	Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
	Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
	Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

	Each “newanimal” command must be a single line containing three strings.
	The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal.
	The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
	Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

	Each “query” command must be a single line containing 3 strings.
	The first string is “query”. The second string is the name of the animal.
	The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
	Your program should process each query command by printing out the requested data.

	Define an interface type called Animal which describes the methods of an animal.
	Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
	The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
	Define three types Cow, Bird, and Snake.
	For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
	When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.
	*******/

	println("Please enter your query. example `newanimal milka cow` or `query milka speak`")

	animals := make(map[string]Animal, 0)

	for {
		var command, secondArg, thirdArg string
		fmt.Println(">")
		fmt.Scanln(&command, &secondArg, &thirdArg)

		msg := ""
		switch strings.ToLower(command) {
		case "newanimal":
			msg = createAnimal(secondArg, thirdArg, animals)
		case "query":
			msg = queryAnimal(secondArg, thirdArg, animals)
		default:
			println("Unknown command entered. Try again!!")
			continue
		}
		println(msg)
		// println("map ",len(animals))
	}

}

func createAnimal(name, animalType string, aMap map[string]Animal) string {

	//example `newanimal milka cow`

	switch strings.ToLower(animalType) {
	case "cow":
		aMap[name] = NewCow()
	case "bird":
		aMap[name] = NewBird()
	case "snake":
		aMap[name] = NewSnake()
	default:
		return ("Unknown animal entered. Try again!!")
	}
	return "Created it!"
}

func queryAnimal(name, animalInfo string, aMap map[string]Animal) string {

	// example `query milka speak`

	v, ok := aMap[name]
	if !ok {
		return fmt.Sprintf("No animals found by name %v", name)
	}

	switch v.(type) {
	case *Cow:
		c, _ := v.(*Cow)
		c.processInfo(animalInfo)
	case *Bird:
		b, _ := v.(*Bird)
		b.processInfo(animalInfo)
	case *Snake:
		s, _ := v.(*Snake)
		s.processInfo(animalInfo)
	default:
		return fmt.Sprintf("No known types found for name %v", name)
	}
	return ""
}
