package _38_permutation

import "fmt"

//难度：5*
func Permutation(str string) {
    if len(str) == 0 {
        return
    }

    permutaion([]byte(str), 0)
}

//本质上是递归，很抽象
//首字母和后面的字母逐个替换
func permutaion(org []byte, beginIdx int) {
    if beginIdx >= len(org)-1 {
        fmt.Println(string(org))
        return
    }

    for i:=beginIdx; i<len(org); i++ {
        //swap，首字母和后面的字母逐个替换
        org[beginIdx], org[i] = org[i], org[beginIdx]

        //递归
        permutaion(org, beginIdx+1)

        //swap，换回来，如果不换回来，就乱了，就无法达到逐个替换的目的，
        org[i], org[beginIdx] = org[beginIdx], org[i]
    }

}