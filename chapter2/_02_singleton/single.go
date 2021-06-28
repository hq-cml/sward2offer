/*
 * 单例
 */
package _02_singleton
import (
    "sync"
)

type Singleton struct {
}

var ins *Singleton
var mutex sync.Mutex
var once sync.Once

//常规思路：效率低
func GetIns1() *Singleton {
    mutex.Lock()
    defer mutex.Unlock()

    if ins == nil {
        ins = &Singleton{}
    }
    return ins
}

//改进思路，效率更高
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

//直接利用sync.once包
func GetIns3() *Singleton {
    once.Do(func(){
        ins = &Singleton{}
    })

    return ins
}