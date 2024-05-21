package main

import "fmt"

func single_chan() {
	s := make(chan string)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("sender hello", i)
			s <- fmt.Sprintf("receiver hello %d", i)
		}
	}()

	for i := 0; i < 6; i++ {
		val := <-s
		fmt.Println(val)
	}
}

func mutil_chan() {
	s := make(chan string, 2)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("sender hello", i)
			s <- fmt.Sprintf("receiver hello %d", i)
		}
	}()

	for i := 0; i < 6; i++ {
		val := <-s
		fmt.Println(val)
	}
}

func main() {

	fmt.Println("========single_chan=========")

	single_chan()

	fmt.Println("========mutil_chan=========")

	mutil_chan()
}
