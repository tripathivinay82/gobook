package main 

import (
	"fmt"
	"strings"
)

func main() {

    str := "it was the best of times it was the worst of times it was the age of wisdom it was the age of foolishness"
    ds := make(map[string]int)

    strslc := strings.Split(str," ")
    fmt.Println(strslc)

    for i, w := range strslc {
        fmt.Println(i, w)

	wc := 0
        for _, ww := range strslc {
		if w == ww {
		    wc += 1
		}
	}
	ds[w] = wc
    }
    fmt.Println(ds)
}
