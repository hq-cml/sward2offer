package _42_max_sum_sub_arr

import (
    "errors"
    "math"
)

//举一个列子，比如{1,2,-4,3,1,2,-4}，列出一张表格
//可以看出，一旦累计和出现了负数，那还不如从当前开始，从头开始更划算
func Calc(arr []int) (int, error) {
    if len(arr) == 0 {
        return 0, errors.New("Invalid input")
    }

    max := math.MinInt64     //最大值
    accumulate := 0          //累计值
    for _, v := range arr {
        if accumulate <= 0 {
            accumulate = v   //从当前开始，从头开始算
        } else {
            accumulate += v
        }

        if accumulate > max {
            max = accumulate
        }
    }
    return max, nil
}