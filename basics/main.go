package main

import (
	"fmt"
	"reflect"
)

func main() {
	var word string = "hello"
	word2 := "haha"
	number := 2858
	number2 := 2858.5
	flag := false
	fmt.Println(word, word2, number, number2, flag)
	// one way of getting type of variable
	fmt.Printf("%T\n", number)
	// another way of getting type of variable
	fmt.Println(reflect.TypeOf(number2), reflect.TypeOf(flag))
}
