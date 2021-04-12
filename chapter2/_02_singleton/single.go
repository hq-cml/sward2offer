package _02_singleton
import (
    "sync"
)

type Singleton struct {
}

var ins *Singleton
var mutex sync.Mutex
var once sync.Once

func GetIns1() *Singleton {
    mutex.Lock()
    defer mutex.Unlock()

    if ins == nil {
        ins = &Singleton{}
    }
    return ins
}

func GetIns2() *Singleton {
    if ins != nil {
        return ins
    }

    mutex.Lock()
    defer mutex.Unlock()
    if ins == nil {
        ins = &Singleton{}
    }
    return ins
}

func GetIns3() *Singleton {
    once.Do(func(){
        ins = &Singleton{}
    })

    return ins
}