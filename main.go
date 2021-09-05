package main

import (
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

}
