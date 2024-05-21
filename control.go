package main

import "fmt"

func switch_condition(i int) { //function內的參數不需要加上var宣告, 直接後修飾型態即可
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}
}

func main() {

	// for-loop
	/*
		in c:
		int i = 0;
		int sum = 0;
		for( i  = 0 ; i < 10 ; ++i )
		{
			sum += i;
		}
	*/
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// do while
	/*
		in c:
		int i = 0;

		do{
			i += 1
		}while( i < 10 );
	*/
	i := 0
	for i < 100 {
		i += 1
	}
	fmt.Println(i)
	// while 1
	/*
		in c:
		while(1)
		{

		}
	*/
	/*
		for {

		}
	*/

	// if condition
	var a int = 10
	if a > 0 {
		fmt.Println(a)
	}

	// switch condition
	switch_condition(0)
	switch_condition(1)
	switch_condition(2)
}
