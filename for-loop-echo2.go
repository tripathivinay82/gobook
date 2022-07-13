package main

import (
    "fmt"
    "os"
)

func main() {

    s, sep := "", ""

    for _, arg := range os.Args[1:] { // range function unpacks the index and value oer iteration; index is stored in '_' (unused) and value in arg
	s += sep + arg
	sep = " "
    }
    // fmt.Println(os.Args[0])
    // fmt.Println(os.Args[1])
    // fmt.Println(os.Args[2])
    fmt.Println(s)
}
