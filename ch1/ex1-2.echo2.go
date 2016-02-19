// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
//add
	"strconv"
)

func main() {
	s, sep := "", ""
//	for _, arg := range os.Args[1:] {
	for i, arg := range os.Args[1:] {
//		s += sep + arg
		s += sep + strconv.Itoa(i + 1) + " " + arg
//		sep = " "
		sep = "\n" 
	}
	fmt.Println(s)
}
