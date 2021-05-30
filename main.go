package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Hello world")

    tmp := strings.Trim("$aa,aab,", ",")
    fmt.Println(tmp)
   // fmt.Println(len(tmp))
}
