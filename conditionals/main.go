package main

import "fmt"

func main() {
	a := 5
	b := 5
	//if else
	if a > b {
		fmt.Println(a)
	} else if a == b {
		fmt.Println("a is equal to b")
	} else {
		fmt.Println(b)
	}
	// switch
	switch a {
	case 5:
		fmt.Println("a is equal to b")
	}
}
