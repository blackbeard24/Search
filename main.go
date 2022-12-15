package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	var num int
	fmt.Println("User please enter an sorted array")
	for i := 0; i < 5; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println("Find which number?")
	fmt.Scanln(&num)
	fmt.Println("Linear Search")
	linear(arr, num)
	fmt.Println("Binary Search")
	binary(arr, num)
}

func linear(arr [5]int, num int) {
	fl := false
	for i := 0; i < 5; i++ {
		if arr[i] == num {
			fmt.Println("Found your number")
			fl = true
			break
		}
	}
	check(fl)

}

func binary(arr [5]int, num int) {
	final := len(arr)
	init := 0
	mid := (init + final) / 2
	fl := false
	i := 0
	for i < 5 {
		if arr[mid] == num {
			fmt.Println("Found your number")
			fl = true
			break
		} else if num > arr[mid] {
			init = mid
			mid = (init + final) / 2
		} else {
			final = mid
			mid = (init + final) / 2
		}
	}
	check(fl)
}

func check(fl bool) {
	if fl == false {
		fmt.Println("Number not found")
	}
}
