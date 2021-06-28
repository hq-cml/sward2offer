package main

import (
    "context"
    "fmt"
    "time"
)
type Pet interface {
    Name() string
    Category() string
}
type Dog struct {
    name string // 名字。
}
func (dog *Dog) SetName(name string) {
    dog.name = name
}
func (dog Dog) Name() string {
    return dog.name
}
func (dog Dog) Category() string {
    return "dog"
}
//父亲
func main() {
    ctx, cancel := context.WithCancel(context.Background())

    //儿子
    go func(ctx context.Context) {

        //ctxSon, _ := context.WithCancel(ctx)
        ctxSon := context.WithValue(ctx, "A", 1)

        //孙子
        go func(ctxSon context.Context) {
            fmt.Println("K--------", ctxSon.Value("A"))

            //ctxSonSon, cancel := context.WithCancel(ctxSon)
            ctxSonSon, _ := context.WithCancel(ctxSon)
            //重孙子
            go func(ctxSonSon context.Context) {
                fmt.Println("L--------", ctxSonSon.Value("A"))
                select {
                case <-ctxSonSon.Done():
                    fmt.Println("SonSon exit!")
                }
            }(ctxSonSon)

            select {
            case <-ctxSon.Done():
                fmt.Println("Son exit!")
            }
        }(ctxSon)

        select {
        case <-ctx.Done():
            fmt.Println("Father exit!")
        }

    }(ctx)

    time.Sleep(5 * time.Second)
    cancel()
    time.Sleep(5 * time.Second)
}
