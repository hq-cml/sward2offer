package _43_numberof1_between_1_n

import (
    "math"
    "strconv"
)

//复杂度O(logn)
//这种题，非常抽象，只能欣赏欣赏了
//难度:5*
func Calc(n int) int {
    if n <= 0 {
        return 0
    }

    return numberOf1([]byte(strconv.Itoa(n)))
}

func numberOf1(str []byte) int {
    if len(str) == 0 {
        return 0
    }

    first := str[0] - '0'
    length := len(str)

    //递归退出条件
    if length == 1 && first == 0 {
        return 0
    }

    if length == 1 && first > 0 {
        return 1
    }

    numFirstDigit := 0
    if first == 1 {
        t, _ := strconv.Atoi(string(str[1:]))
        numFirstDigit = t + 1
    } else { //first > 1
        numFirstDigit = int(math.Pow10(length-1))
    }

    numOtherDigit := int(first) * (length-1) * int(math.Pow10(length - 2))

    numRecurse := numberOf1(str[1:])

    return numFirstDigit + numOtherDigit + numRecurse
}

//复杂度为O(n*logn)，用于验证Calc1
func Calc2(n int) int {
    cal := func(n int) int { //匿名小函数
        num := 0
        for n != 0 {
            if n % 10 == 1 {
                num ++
            }
            n = n/10
        }
        return num
    }

    num := 0
    for n != 0 {
        num += cal(n)
        n--
    }
    return num
}