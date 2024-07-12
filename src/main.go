package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	runThis()

}

func runThis() {
	AnimalTypeAssignment()
}

func AnimalTypeAssignment() {
	/*******
	Write a program which allows the user to get information about a predefined set of animals.
	Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
	The user can issue a request to find out one of three things about an animal:
		1) the food that it eats,
		2) its method of locomotion, and
		3) the sound it makes when it speaks.
	The following table contains the three animals and their associated data which should be hard-coded into your program.

	Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
	Your program accepts one request at a time from the user, prints out the answer to the request,
	and prints out a new prompt. Your program should continue in this loop forever.
	Every request from the user must be a single line containing 2 strings.
	The first string is the name of an animal, either “cow”, “bird”, or “snake”.
	The second string is the name of the information requested about the animal,
		either “eat”, “move”, or “speak”.
	Your program should process each request by printing out the requested data.

	You will need a data structure to hold the information about each animal.
	Make a type called Animal which is a struct containing three fields:food,
	locomotion, and noise, all of which are strings. Make three methods called Eat(),
	Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
	The Eat() method should print the animal’s food, the Move() method should print the animal’s
	locomotion, and the Speak() method should print the animal’s spoken sound.
	Your program should call the appropriate method when the user makes a request.

	Submit your Go program source code.
	*******/

	println("Initializing data . . .")
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	println("Please enter your query. example `cow move`")

	// for {
	var animalType, animalInfo string
	fmt.Println(">")
	fmt.Scanln(&animalType, &animalInfo)

	println(animalType, animalInfo)

	switch strings.ToLower(animalType) {
	case "cow":
		cow.processInfo(animalInfo)
	case "bird":
		bird.processInfo(animalInfo)
	case "snake":
		snake.processInfo(animalInfo)
	default:
		println("Unknown animal entered. Try again!!")

	}
	// }

}

func (a *Animal) processInfo(input string) {
	switch strings.ToLower(input) {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		println("unknown information requested. Try again!!")
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	println(a.food)
}

func (a *Animal) Move() {
	println(a.locomotion)
}

func (a *Animal) Speak() {
	println(a.noise)
}

func displacementAssignment() {
	/*********
	Let us assume the following formula for
	displacement s as a function of time t, acceleration a, initial velocity vo,
	and initial displacement so.

	s = ½ a t2 + vot + so

	Write a program which first prompts the user
	to enter values for acceleration, initial velocity, and initial displacement.
	Then the program should prompt the user to enter a value for time and the
	program should compute the displacement after the entered time.

	You will need to define and use a function
	called GenDisplaceFn() which takes three float64
	arguments, acceleration a, initial velocity vo, and initial
	displacement so. GenDisplaceFn()
	should return a function which computes displacement as a function of time,
	assuming the given values acceleration, initial velocity, and initial
	displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
	float64 argument which is the displacement travelled after time t.

	For example, let’s say that I want to assume
	the following values for acceleration, initial velocity, and initial
	displacement: a = 10, vo = 2, so = 1. I can use the
	following statement to call GenDisplaceFn() to
	generate a function fn which will compute displacement as a function of time.

	fn := GenDisplaceFn(10, 2, 1)

	Then I can use the following statement to
	print the displacement after 3 seconds.

	fmt.Println(fn(3))

	And I can use the following statement to print
	the displacement after 5 seconds.

	fmt.Println(fn(5))

	Submit your Go program source code.

	*********/

	var input string
	fmt.Println("Enter acceleration")
	fmt.Scan(&input)

	accl, _ := strconv.ParseFloat(input, 64)

	fmt.Println("Enter initial velocity")
	fmt.Scan(&input)

	inVel, _ := strconv.ParseFloat(input, 64)

	fmt.Println("Enter initial displacement")
	fmt.Scan(&input)

	inDisp, _ := strconv.ParseFloat(input, 64)

	fn := GenDisplaceFn(accl, inVel, inDisp)

	//println("Displacement after 0 secs is ", fn(0))
	// println("Displacement after 5 secs is ", fn(5))
	// println("Displacement after 10 secs is ", fn(10))

	fmt.Println("Enter time (in sec) at which displacement has to be calculated")
	fmt.Scan(&input)

	time, _ := strconv.ParseFloat(input, 64)

	fmt.Printf("Displacement after %v secs is %v", time, fn(time))

}

func GenDisplaceFn(accl, velo, disp float64) func(float64) float64 {
	//s = ½ a t2 + vot + so

	return func(time float64) float64 {
		return 0.5*accl*math.Pow(time, 2) + velo*time + disp
	}

}

func bubbleSortAssignment() {
	/****
	Write a Bubble Sort program in Go. The program
	should prompt the user to type in a sequence of up to 10 integers. The program
	should print the integers out on one line, in sorted order, from least to
	greatest. Use your favorite search tool to find a description of how the bubble
	sort algorithm works.

	As part of this program, you should write a
	function called BubbleSort() which
	takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
	order.

	A recurring operation in the bubble sort algorithm is
	the Swap operation which swaps the position of two adjacent elements in the
	slice. You should write a Swap() function which performs this operation. Your Swap()
	function should take two arguments, a slice of integers and an index value i which
	indicates a position in the slice. The Swap() function should return nothing, but it should swap
	the contents of the slice in position i with the contents in position i+1.

	Submit your Go program source code.
	*****/

	// 23,67,2,89,324,879,234,65,98,2,6,23,76,-23,67,54,90,-45,-2
	var input string
	fmt.Println("Enter Integers to be sorted seperated by a comma ','. Do not put space in between numbers!!")
	fmt.Scan(&input)

	// 23,67,2,89,324,879,234,65,98,2,6,23,76,-23,67,54,90,-45,-2
	sArr := strings.Split(input, ",")
	iArr := make([]int, 0)

	for _, val := range sArr {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			fmt.Println("Found an improper Number in the input >> ", val)
			continue
		}
		iArr = append(iArr, int(i))
	}

	fmt.Println("Input  Integer Array >> ", iArr)
	BubbleSort(iArr)
	fmt.Println("Sorted Integer Array >> ", iArr)
}

func BubbleSort(sl []int) {
	len := len(sl)
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if sl[j] > sl[j+1] {
				Swap(sl, j)
			}
		}
	}
}

func Swap(s []int, swapIndex int) {
	temp := s[swapIndex]
	s[swapIndex] = s[swapIndex+1]
	s[swapIndex+1] = temp
}

type (
	Chopsticks struct {
		sync.Mutex
	}
	Philsopher struct {
		ID     int
		LChops *Chopsticks
		RChops *Chopsticks
		Times  int
	}
	Host struct {
		CurrentDiners int
	}
)

var wg sync.WaitGroup

const (
	MAX_TIMES_EAT      = 3
	MAX_PHILOSOPHERS   = 5
	MAX_CONCURRENT_EAT = 2
)

func (h *Host) serve(p []*Philsopher, wg *sync.WaitGroup) {

	h.CurrentDiners = 0
	fmt.Println("Host is starting to serve.")
	defer wg.Done()

	for i := 0; i < MAX_PHILOSOPHERS*MAX_TIMES_EAT; i++ {
		if h.CurrentDiners < MAX_CONCURRENT_EAT {
			j := rand.Intn(MAX_PHILOSOPHERS)
			fmt.Printf("Serving philosopher %d for the %d time\n", j, (p[j].Times + 1))
			if p[j].Times < MAX_TIMES_EAT {
				h.CurrentDiners += 1
				wg.Add(1)
				go p[j].eat(wg, &h.CurrentDiners)

			} else {
				fmt.Printf("\nPhilosopher %d is done eating %d times\t\tChoosing different Philosopher..\n", j, p[j].Times)
				i--
			}

		} else {
			// fmt.Printf("\n%d philosophers eating concurrently. Trying again\n", h.CurrentDiners)
			i--
		}
	}
}

func (p *Philsopher) eat(wg *sync.WaitGroup, counter *int) {
	defer wg.Done()
	p.LChops.Lock()
	p.RChops.Lock()
	p.Times += 1

	fmt.Println("starting to eat ", p.ID)
	time.Sleep(time.Second)
	fmt.Println("finishing eating ", p.ID)

	p.LChops.Unlock()
	p.RChops.Unlock()
	*counter--
}

func module4Assignment() {
	/*******
		Implement the dining philosopher’s problem with the following constraints/modifications.
	    There should be 5 philosophers sharing chopsticks,
		with one chopstick between each adjacent pair of philosophers.
	    Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	    The philosophers pick up the chopsticks in any order,
		not lowest-numbered first (which we did in lecture).
	    In order to eat, a philosopher must get permission from a host which executes in its
		own goroutine.
	    The host allows no more than 2 philosophers to eat concurrently.
	    Each philosopher is numbered, 1 through 5.
	    When a philosopher starts eating (after it has obtained necessary locks) it prints
		“starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
	    When a philosopher finishes eating (before it has released its locks) it prints
		“finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
		*******/

	CSticks := make([]*Chopsticks, MAX_PHILOSOPHERS)
	for i := 0; i < MAX_PHILOSOPHERS; i++ {
		CSticks[i] = new(Chopsticks)
	}

	philos := make([]*Philsopher, MAX_PHILOSOPHERS)
	for i := 0; i < MAX_PHILOSOPHERS; i++ {
		philos[i] = &Philsopher{
			ID:     i + 1,
			LChops: CSticks[i],
			RChops: CSticks[(i+1)%MAX_PHILOSOPHERS],
			Times:  0,
		}
	}
	host := new(Host)
	wg.Add(1)
	go host.serve(philos, &wg)
	wg.Wait()
	fmt.Println("...")

	for i := 0; i < MAX_PHILOSOPHERS; i++ {
		fmt.Printf("\n Philospher %d ate %d times\n", philos[i].ID, philos[i].Times)
	}
}

func assignment_mergeSubArray() {
	/**********
	Write a program to sort an array of integers.
	The program should partition the array into 4 parts,
	each of which is sorted by a different goroutine.
	Each partition should be of approximately equal size.
	Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

	The program should prompt the user to input a series of integers.
	Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
	When sorting is complete, the main goroutine should print the entire sorted list.
	********/

	var input string
	fmt.Println("Enter Integers to be sorted seperated by a comma ','. Do not put space in between numbers!!")
	fmt.Scan(&input)

	// 23,67,2,89,324,879,234,65,98,2,6,23,76,-23,67,54,90,-45,-2
	sArr := strings.Split(input, ",")
	iArr := make([]int, 0)

	for _, val := range sArr {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			fmt.Println("Found an improper Number in the input >> ", val)
			continue
		}
		iArr = append(iArr, int(i))
	}

	fmt.Println("Input integers to sort are >> ", iArr)
	length := len(iArr)
	partSize := (length + 3) / 4 // Round up to ensure all elements are covered

	var wg sync.WaitGroup
	subArrays := make([][]int, 4)

	// Partition the array into 4 subarrays and sort each in a separate goroutine
	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if end > length {
			end = length
		}
		subArrays[i] = append([]int{}, iArr[start:end]...)
		wg.Add(1)
		go sortSubArray(&wg, subArrays[i], i)
	}

	wg.Wait()

	// Merge the sorted subarrays into one sorted array
	sortedArray := mergeSortedArrays(subArrays...)
	fmt.Println("Final sorted array:", sortedArray)

}

// Function to sort a portion of the array and print it
func sortSubArray(wg *sync.WaitGroup, arr []int, index int) {
	defer wg.Done()
	fmt.Printf("Goroutine %d sorting subarray: %v\n", index+1, arr)
	sort.Ints(arr)
	fmt.Printf("Goroutine %d sorted subarray: %v\n", index+1, arr)
}

func mergeSortedArrays(arrays ...[]int) []int {
	var merged []int
	for len(arrays) > 1 {
		var newArrays [][]int
		for i := 0; i < len(arrays)-1; i += 2 {
			merged = mergeTwoSortedArrays(arrays[i], arrays[i+1])
			newArrays = append(newArrays, merged)
		}
		if len(arrays)%2 == 1 {
			newArrays = append(newArrays, arrays[len(arrays)-1])
		}
		arrays = newArrays
	}
	return arrays[0]
}

func mergeTwoSortedArrays(a, b []int) []int {
	result := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result[k] = a[i]
			i++
		} else {
			result[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		result[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		result[k] = b[j]
		j++
		k++
	}
	return result
}

func slice() {
	var v, cont string
	enterValue := true

	sl := make([]int, 0)

	for {
		fmt.Println("Enter a number> ")
		fmt.Scan(&v)

		iInput, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Println("You have entered an improper Number. Try again!!")
			continue
		}
		sl = append(sl, int(iInput))

		fmt.Print("do you want to enter another number (Y/N)? > ")
		fmt.Scan(&cont)
		enterValue = !strings.ContainsAny(cont, "Nn")
		if !enterValue {
			break
		}
	}
	fmt.Println("----------------------------------")

	// sorting the int slice
	sort.Ints(sl)

	fmt.Println("Sorted Slice of integers entered is :: ")
	for _, v := range sl {
		fmt.Print(v, " ")
	}
}

func learning() {
	// slice
	sl := make([]int, 0, 3)
	sl = append(sl, 2, 5, 10)
	for _, v := range sl {
		fmt.Println(v)
	}
	fmt.Println("----------------------------------")

	// hash table
	// suitable for key-value pairs
	// maps are teh implementation of hash table in GO
	idMap := make(map[string]int)
	idMap["A"] = 10
	idMap["B"] = 15
	for k, v := range idMap {
		fmt.Println(k, " - ", v)
	}
	_, isPres := idMap["C"]
	if v, ok := idMap["A"]; ok && v == 10 {
		fmt.Printf("Found %s in idMap with value as %d\n", "A", 10)
	}
	fmt.Println(isPres)
	fmt.Println("----------------------------------")

	// struct
	type Recipe struct {
		Title string `json:title,omitempty`
		ID    int    `json:id,omitempty`
		Desc  string `json:desc,omitempty`
	}

	r := make([]Recipe, 0, 2)
	r1 := Recipe{
		Title: "Egg omlette",
		ID:    1,
		Desc:  "dsfasdf",
	}
	r2 := Recipe{
		Title: "Maggi",
		ID:    3,
		Desc:  "asdfasdf",
	}
	r = append(r, r1, r2)
	r = append(r, *new(Recipe))

	r3, err := json.Marshal(r2)
	if err != nil {
		fmt.Println("error marshalling ", err)
	}
	r4 := &Recipe{}
	err = json.Unmarshal(r3, r4)
	if err != nil {
		fmt.Println("error un-marshalling ", err)
	}
	r4.ID = 15
	r4.Title = "New Maggi"

	r = append(r, *r4)

	for _, v := range r {
		fmt.Printf("Recipe %d is %s\n", v.ID, v.Title)
	}
	fmt.Println("----------------------------------")
}

func getMaxSizeName(v string) string {
	if len(v) > 20 {
		return v[:20]
	}
	return v

}

func assignment4() {
	/***
	Write a program which reads information from a file and represents it in a
	slice of structs. Assume that there is a text file which contains a series of
	names. Each line of the text file has a first name and a last name, in that order,
	separated by a single space on the line.

	Your program will define a name struct which has two fields,
	fname for the first name, and lname for the last name.
	Each field will be a string of size 20 (characters).

	Your program should prompt the user for the name of the text file.
	Your program will successively read each line of the text file and
	create a struct which contains the first and last names found in the file.
	Each struct created will be added to a slice, and after all lines have
	been read from the file, your program will have a slice containing one
	struct for each line in the file. After reading all lines from the file,
	your program should iterate through your slice of structs and print the first
	and last names found in each struct.
	***/

	// define the struct
	type Name struct {
		fname string
		lname string
	}

	// define the slice of structs
	slNames := make([]Name, 0)

	// instructions to run the code.
	fmt.Println("----  INFO  ------")
	fmt.Println("Please place the `file.txt` in the path where the `go run main.go` command will be triggerred from the CLI.")
	fmt.Println("----x  END  x------")
	fmt.Println()

	// take the file name as input from user
	var v string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name of the file > ")
	text, _ := reader.ReadString('\n')
	v = strings.TrimSpace(text)

	// read the file
	readFile, err := os.Open(v)
	if err != nil {
		fmt.Println("unable to read file. Check the path. Filenames are case sensitive.\nErr ::  ", err)
		return
	}

	// read contents of the file line by line
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Print(". ")
		namesFromLine := strings.Split(line, " ")
		name := Name{
			fname: getMaxSizeName(namesFromLine[0]),
			lname: getMaxSizeName(namesFromLine[0]),
		}
		slNames = append(slNames, name)

	}

	readFile.Close()
	fmt.Println("succesfully read the the file.")
	fmt.Println("Here are the contents of the file : ")

	//iterate over the slice of struct and print the names
	for _, item := range slNames {
		fmt.Printf("Fname: %s  Lname: %s\n", item.fname, item.lname)
	}
}

func assignmentGoRoutineRaceCondition() {

	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	x := 0
	ch1 := make(chan int)
	ch2 := make(chan int)

	g.Go(func() error {
		defer close(ch1)
		fmt.Println("CH1 | Working with x as ", x)
		x += 9
		select {
		// case <-ctx.Done():
		// 	return ctx.Err()
		case ch1 <- (x):
		}
		return nil
	})
	g.Go(func() error {
		defer close(ch2)
		fmt.Println("CH2 | Working with x as ", x)
		x += 10
		select {
		// case <-ctx.Done():
		// 	return ctx.Err()
		case ch2 <- (x):
		}
		return nil
	})

	x = 1

	fmt.Printf("x:= %d\tch1Int:= %d\tch2Int:= %d", x, <-ch1, <-ch2)
	g.Wait()

}

func assignment3() {
	var v string
	jMap := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a name > \n")
	text, _ := reader.ReadString('\n')
	v = strings.TrimSpace(text)

	if v == "" {
		fmt.Println("You did not enter anything. Try again!!")
		return
	}
	jMap["name"] = v

	fmt.Print("Enter an address > \n")
	text, _ = reader.ReadString('\n')
	v = strings.TrimSpace(text)

	if v == "" {
		fmt.Println("You did not enter anything. Try again!!")
		return
	}
	jMap["address"] = v

	fmt.Println("----------------------------------")

	js, err := json.Marshal(jMap)
	if err != nil {
		fmt.Println("error marshalling ", err)
	}
	//print the print the JSON object in string format
	fmt.Println(string(js))

	// fmt.Println("----------------------------------")
	// print the print the JSON object as bytearray
	// fmt.Println(js)
}
