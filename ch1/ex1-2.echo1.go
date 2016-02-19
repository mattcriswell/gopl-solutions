// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
//add
        "strconv"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
//		s += sep + os.Args[i]
		s += sep + strconv.Itoa(i) + " " + os.Args[i]
//		sep = " "
		sep = "\n"
	}
	fmt.Println(s)
}
