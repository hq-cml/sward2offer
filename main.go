package main

import (
    "fmt"
    "math/rand"
    "time"
)

func shuffle(arr []int) {
    length := len(arr)
    rand.Seed(time.Now().Unix())
    for i:=length-1; i>=1; i-- {
        idx := rand.Int63n(int64(i))
        arr[i], arr[idx] = arr[idx], arr[i]
    }
}

type A struct {
    a *int
    b string
}

type B struct {
    a *A
    b string
    c []string
}

func main() {
    s := [10]int{1,2,3,4,5,6,7,8,9,10}
    s1 := s[:5]
    fmt.Println(s1)
    s2 := s[3:]
    fmt.Println(s2)
    s2[0] = 100
    s2[1] = 200
    s2[2] = 300
    fmt.Println(s1)
    fmt.Println(s2)
}
