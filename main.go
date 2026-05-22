package main

import (
	"fmt"
	"time"
)

// sub 1
func sendGreeting(a chan string) {
	a <- "Привет из горутины!"

}

// sub 2
func squareWorker(f int, c chan int) {

	fmt.Println(f * f)

	c <- f

}

// sub 3
func emitNumbers(c chan int) {
	for {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		c <- 5
		close(c)
		break
	}

}

//sub 4

func sumReader(c1 chan int) {
	for v := range c1 {
		summ := v
		fmt.Println(summ)

	}

}

// sub 5
func filterEven(input chan int, output chan int) {
	for v := range input {
		if v%2 == 0 {
			output <- v
		}

	}
	close(output)

}

// sub 6

func checkChannel(c chan string) {
	val, ok := <-c
	if ok == true {
		fmt.Println(val)
	} else {
		fmt.Println("Канал закрыт")
	}

}

func main() {
	b := make(chan string)

	go sendGreeting(b)

	fmt.Println(<-b)

	d := make(chan int)

	go squareWorker(9, d)

	c := make(chan int)
	go emitNumbers(c)

	for v := range c {

		fmt.Println(v)

	}

	f := make(chan int)

	go sumReader(f)

	f <- 1
	f <- 2
	f <- 3

	close(f)

	k := make(chan int)
	l := make(chan int)

	go filterEven(k, l)

	go func() {
		for v := 0; v <= 10; v++ {
			k <- v
		}
		close(k)
	}()

	for v := range l {
		fmt.Println(v)

	}

	o := make(chan string)
	go func() {
		o <- "SSS"

	}()

	go checkChannel(o)
	time.Sleep(2 * time.Second)

	close(o)
	go checkChannel(o)
	time.Sleep(2 * time.Second)

	r := make(chan string, 3)

	r <- "asd"
	r <- "asdasd"
	r <- "asdasdasd"
	r <- "asdasdasdasd"

	fmt.Println(<-r)
	fmt.Println(<-r)
	fmt.Println(<-r)

}
