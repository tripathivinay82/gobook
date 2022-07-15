package main 

import  ( 
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Enter the string: ")

    var str string
    fmt.Scanln(&str)
    str = strings.TrimSpace(str)

    nstr := strings.Split(str,"")
    var fstr  string

    c := len (nstr) - 1
    fmt.Println(c)

    for _, s := range nstr {
	fstr += nstr[c] 
        fmt.Println(s)
	c -= 1
    }

    fmt.Println("Palindrome string is: ", fstr)
}
