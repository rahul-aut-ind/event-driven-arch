package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
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
	module4Assignment()
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
