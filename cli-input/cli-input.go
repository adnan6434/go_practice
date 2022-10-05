package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "

	/*	for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	*/
	for idx, arg := range os.Args[1:] {
		s += sep + fmt.Sprint(idx) + sep + arg
	}
	fmt.Println(s)
}
