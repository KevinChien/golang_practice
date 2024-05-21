package main

import "fmt"

func advanced_slice() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {
	//array
	//var a [5]int                 //宣告一個大小5的 int array
	//var b []int                  //宣告空的 int array
	var c = []int{1, 2, 3, 4, 5} //宣告一個大小5的 int array, 並給初始值

	/*
		initialization array could be:
		var a = [4]int {
			1,
			2,
			3,
			4,
		}
	*/
	fmt.Println("Array.....")
	for i := range c {
		fmt.Println(c[i])
	}

	// slice
	/*

		    ptr
		    len
		    cap
			a := make([]int,5,10)
			len =>已使用的長度
			cap =>可利用的長度, 可以當作已經要到的記憶體區塊大小
	*/

	fmt.Println("slice....")
	var a []int
	a = append(a, 1)
	a = append(a, 2)
	fmt.Println("a loop....")
	for i := range a {
		fmt.Println(a[i])
	}

	fmt.Println("a1 loop....")
	a1 := make([]int, 5, 10)

	for i := range a1 {
		fmt.Println(a1[i])
	}

	fmt.Println("a2 loop....")
	a2 := make([]int, 5)
	for i := range a2 {
		fmt.Println(" ", a2[i])
	}

	var a3 = []int{1, 2, 3, 4, 5}
	b := a3[2:4]
	d := a3[:]
	fmt.Println("b loop....")
	/*
		for i := range b {
			fmt.Println(" ", b[i])
		}
	*/
	fmt.Println(a3)

	fmt.Println("d loop....")
	for i := range d {
		fmt.Println(" ", d[i])
	}

	copy(b, d) //copy d data to b, base on slice size.

	var numbers = make([]int, 3, 5)

	printSlice(numbers)

	advanced_slice()
	// map
	/*
		m := make(map[string]int)
		等同
		m := map[string]int{}
	*/
	fmt.Println("map....")
	m := make(map[string]int) //建立一個 map
	//m := map[string]int{}   //另外的建立方法, 和上面結果相同

	m["age"] = 16 //定義 age ＝ 16

	x := m["age"]
	fmt.Println(x)

	y := m["height"] //y 指定給一個沒有宣告過的 map index
	fmt.Println(y)   //print 的值為 0

	m["height"] = 180 //此時 map 總共有兩個 index "age" "height"

	y = m["height"]
	fmt.Println(y)

	value, ok := m["age2"] //回傳值, 和 map 是否存在
	fmt.Println("age2:", value, ok)

	value, ok = m["age"] //回傳值, 和 map 是否存在
	fmt.Println("age:", value, ok)

	//利用 range 回傳 key value
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println("")
	//移除 map 中的值
	delete(m, "age")

	fmt.Println("移除 key 'age' ")
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
	/*
		//如果要一次宣告多個, 可以用以下方式
		person := map[string]int{
		       "age" : 16,
		       "height" : 180,
		       "weight" : 6,
		}  */

}
