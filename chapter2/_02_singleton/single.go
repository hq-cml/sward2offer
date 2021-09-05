/*
 * 单例模式
 */
package _02_singleton
import (
    "sync"
)

/*
 * 用三种方式实现单例，性能逐步提高
 *
 * 性能测试：go test -bench=. -run=^$
 * -bench=.表示执行所有的性能测试  -run=^&表示过滤掉单元测试（功能测试）
 * 默认情况下go test会优先跑单测，所以需要过滤掉单测
 *
 * 结果：
   goos: linux
   goarch: amd64
   pkg: github.com/hq-cml/sward2offer/chapter2/_02_singleton
   BenchmarkGetIns1-4      88454830                13.8 ns/op
   BenchmarkGetIns2-4      584879067                2.13 ns/op
   BenchmarkGetIns3-4      711746034                1.60 ns/op
   PASS
   ok      github.com/hq-cml/sward2offer/chapter2/_02_singleton    4.014s
*/
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
//once.Do内部使用原子操作，性能进一步提升
func GetIns3() *Singleton {
    once.Do(func(){
        ins = &Singleton{}
    })

    return ins
}