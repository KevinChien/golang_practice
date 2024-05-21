package main

import "fmt"

func main() {

	var int_x int
	int_x = 10

	fmt.Println("x is", int_x)

	var float_y float64 = 3.1415

	fmt.Println("y is", float_y)

	var s string = "哈囉 go lang"

	fmt.Println("s:", s)

	var test bool = true

	fmt.Println("test is", test)

	var c rune = 'b'
	fmt.Println("c is", c)

	// initialization value
	var x int = 10

	var y = 20

	var z int
	z = 30

	k := 40

	// const value need not define type, just value
	const PI = 3.1415
	const NAME = "Nelson"

}
