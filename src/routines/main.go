package main

import (
	"fmt"
	"sync"
)

func main() {
	idea1()
	fmt.Println("- - - - - ")
	idea2()
}

func idea1() {
	prefix := "idea 1> "
	fmt.Println(prefix, "begin")
	var wg sync.WaitGroup
	nums := make(chan int, 3)

	wg.Add(3)

	go func() {
		fmt.Println(prefix, 1)
		defer wg.Done()
		nums <- 1
		fmt.Println(prefix, "1 done")
	}()
	go func() {
		fmt.Println(prefix, 2)
		defer wg.Done()
		nums <- 2
		fmt.Println(prefix, "2 done")
	}()
	go func() {
		fmt.Println(prefix, 3)
		defer wg.Done()
		nums <- 3
		fmt.Println(prefix, "3 done")
	}()

	wg.Wait()
	close(nums)

	for i := range nums {
		fmt.Println(prefix, "received :: ", i)
	}
	fmt.Println(prefix, "end")
}

func idea2() {
	prefix := "idea 2> "
	fmt.Println(prefix)
	var wg sync.WaitGroup
	nums := make(chan int)

	wg.Add(3)

	go func() {
		fmt.Println(prefix, 3)
		defer wg.Done()
		nums <- 3
		fmt.Println(prefix, "3 done")
	}()
	go func() {
		fmt.Println(prefix, 2)
		defer wg.Done()
		nums <- 2
		fmt.Println(prefix, "2 done")
	}()
	go func() {
		fmt.Println(prefix, 1)
		defer wg.Done()
		nums <- 1
		fmt.Println(prefix, "1 done")
	}()
	go func() {
		wg.Wait()
		close(nums)
	}()

	for i := range nums {
		fmt.Println(prefix, "received :: ", i)
	}
	fmt.Println(prefix, "end")
}

/* output
idea 1>  begin
idea 1>  3
idea 1>  3 done
idea 1>  1
idea 1>  1 done
idea 1>  2
idea 1>  2 done
idea 1>  received ::  3
idea 1>  received ::  1
idea 1>  received ::  2
idea 1>  end
- - - - -
idea 2>
idea 2>  3
idea 2>  3 done
idea 2>  received ::  3
idea 2>  1
idea 2>  1 done
idea 2>  2
idea 2>  received ::  1
idea 2>  received ::  2
idea 2>  2 done
idea 2>  end
*/
