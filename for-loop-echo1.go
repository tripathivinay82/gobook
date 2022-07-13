package main

import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i ++ {
	s += sep + os.Args[i]
	sep = " "
	fmt.Println(i)
    }
    // fmt.Println(os.Args[0])
    // fmt.Println(os.Args[1])
    // fmt.Println(os.Args[2])
    fmt.Println(s)
}
