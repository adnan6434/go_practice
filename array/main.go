package main

import "fmt"

func main() {
	var arr [10]int
	arr[1] = 5
	fmt.Print(arr[5:])
	arr2 := []string{"one", "two"}
	fmt.Print(arr2, len(arr2))
}
