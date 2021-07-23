package main

import (
    "fmt"
)

func tt(a []int) {
    a = append(a, 1)
    fmt.Println(a)
}
//父亲
func main() {
    a  := []int{1,2,3}
    fmt.Println(a)
    tt(a)
    fmt.Println(a) //外层感知不到
}
