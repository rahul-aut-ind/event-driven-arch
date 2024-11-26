package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	idea1()
	fmt.Println("- - - - - ")
	idea2()
}

func idea1() {
	prefix := "Buffered Channel > "
	fmt.Println(prefix, "begin")
	var wg sync.WaitGroup
	nums := make(chan int, 3)

	wg.Add(3)

	go func() {
		fmt.Println(prefix, 1)
		defer wg.Done()
		time.Sleep(5 * time.Second)
		nums <- 1
		fmt.Println(prefix, "1 done")
	}()
	go func() {
		fmt.Println(prefix, 2)
		defer wg.Done()
		time.Sleep(3 * time.Second)
		nums <- 2
		fmt.Println(prefix, "2 done")
	}()
	go func() {
		fmt.Println(prefix, 3)
		defer wg.Done()
		time.Sleep(1 * time.Second)
		nums <- 3
		fmt.Println(prefix, "3 done")
	}()

	wg.Wait()
	close(nums)
	/*
		This approach blocks the main execution till all go routines are finished
	*/

	for i := range nums {
		fmt.Println(prefix, "received :: ", i)
	}
	fmt.Println(prefix, "end")
}

func idea2() {
	prefix := "Unbuffered Channel > "
	fmt.Println(prefix)
	var wg sync.WaitGroup
	nums := make(chan int)

	wg.Add(3)

	go func() {
		fmt.Println(prefix, 3)
		defer wg.Done()
		time.Sleep(1 * time.Second)
		nums <- 3
		fmt.Println(prefix, "3 done")
	}()
	go func() {
		fmt.Println(prefix, 2)
		defer wg.Done()
		time.Sleep(3 * time.Second)
		nums <- 2
		fmt.Println(prefix, "2 done")
	}()
	go func() {
		fmt.Println(prefix, 1)
		defer wg.Done()
		time.Sleep(5 * time.Second)
		nums <- 1
		fmt.Println(prefix, "1 done")
	}()
	go func() {
		wg.Wait()
		close(nums)
	}()
	/*
		If we were to call wg.Wait() directly in the main function before ranging over the channel,
		the program would get stuck, as follows:
		- As the channel is unbuffered, unless a value is read, another value to channel cannot be written
		and this will lead to the first goroutine to not able to exit as it will wait for value to read and
		in main we are waiting for the go routines to finish. The other go routines cannot write as value is not read.
		- Channel Reads Require a Closed Channel (or Sender Signal): For a range loop to exit,
		the channel needs to be closed. Without closing it, range would keep waiting indefinitely for
		new values.
	*/

	for i := range nums {
		fmt.Println(prefix, "received :: ", i)
	}
	fmt.Println(prefix, "end")
}

/* output
Bufferred Channel >  begin
Bufferred Channel >  3
Bufferred Channel >  2
Bufferred Channel >  1
Bufferred Channel >  3 done
Bufferred Channel >  2 done
Bufferred Channel >  1 done
Bufferred Channel >  received ::  3
Bufferred Channel >  received ::  2
Bufferred Channel >  received ::  1
Bufferred Channel >  end
- - - - -
Unbufferred Channel >
Unbufferred Channel >  3
Unbufferred Channel >  2
Unbufferred Channel >  1
Unbufferred Channel >  received ::  3
Unbufferred Channel >  3 done
Unbufferred Channel >  received ::  2
Unbufferred Channel >  2 done
Unbufferred Channel >  1 done
Unbufferred Channel >  received ::  1
Unbufferred Channel >  end
*/
