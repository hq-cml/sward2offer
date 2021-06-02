package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Hello world")

    tmp := strings.Trim("$aa,aab,", ",")
    fmt.Println(tmp)

    a := []int{1,2,3}
    length := len(a)
    fmt.Println(a[0:length])
    fmt.Println(a[length:])
    fmt.Println(a[0:0])
    //fmt.Println(a[length+1:])
}
