package main

import (
	"fmt"
	"time"
)

func backgroundTask(c func()) {
	ticker := time.NewTicker(1 * time.Second)
	a := 0
	for _ = range ticker.C {
		a += 1
		fmt.Println("Tock 1 ->", a)
		if a > 10 {
			c()
			break
		}
	}
}

func backgroundTask2() {
	ticker := time.NewTicker(2 * time.Second)
	a := 0
	for _ = range ticker.C {
		a += 1
		fmt.Println("Tock 2 ->", a)
		if a > 10 {
			break
		}
	}
}

func callback() {
	fmt.Println("callback...")
}

func main() {
	fmt.Println("prueba go...")

	go backgroundTask(callback)
	go backgroundTask2()

	// This print statement will be executed before
	// the first `tock` prints in the console
	fmt.Println("The rest of my application can continue")
	// here we use an empty select{} in order to keep
	// our main function alive indefinitely as it would
	// complete before our backgroundTask has a chance
	// to execute if we didn't.
	select {}
}
