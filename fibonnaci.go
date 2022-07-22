package main

import (
    "fmt"
)

func main() {

    var n int
    fmt.Print("Type the number: ")
    fmt.Scan(&n)
    fmt.Println("User Entered: ", n)

    var l []int

    for i := 0; i < n; i++ {
	if i == 0 {
	    l = append(l, i)
	} else if i == 1 {
	    l = append(l, i)
	} else {
	    l = append(l, l[i-1] + l[i-2])
	}
    }
    fmt.Println(l)
}

// [LOAS] tripathivinay@tripathivinay:~/gobook$ go run fibonnaci.go
// Type the number: 10
// User Entered:  10
//[0 1 1 2 3 5 8 13 21 34]
