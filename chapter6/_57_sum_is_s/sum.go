/*
// 面试题57（一）：和为s的两个数字
// 题目：输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们
// 的和正好是s。如果有多对数字的和等于s，输出任意一对即可。

// 面试题57（二）：为s的连续正数序列
// 题目：输入一个正数s，打印出所有和为s的连续正数序列（至少含有两个数）。
// 例如输入15，由于1+2+3+4+5=4+5+6=7+8=15，所以结果打印出3个连续序列1～5、
// 4～6和7～8。

比较巧妙的思路，利用一头一尾两个指针，看加起来之和和s比较，大了则前移尾指针，小了则后移头指针。
 */

package _57_sum_is_s

//题目1:
//给一个递增数组，然后求出和为给定数字的任意两个成员
//两个指针，从两头不断逼近，调整
//难度：3*
func FindTwoNum(arr []int, sum int) (int, int, bool) {
    //异常输入
    if sum < 0 {
        return 0, 0, false
    }

    length := len(arr)

    //滑动窗口
    i := 0
    j := length - 1
    curr := arr[i] + arr[j]
    for i < j {
        if curr == sum {
            return arr[i], arr[j] , true
        }

        //左边界右移
        if curr < sum {
            i++
            if i > (length - 1) {
                break
            }
            curr = arr[i] + arr[j]
        }

        //右边界左移
        if curr > sum {
            j --
            if j < 0 {
                break
            }
            curr = arr[i] + arr[j]
        }
    }
    return 0, 0, false
}

//题目2:
//和为s的连续递增序列，列举出全部
//思路和题目1类似，也是不断试探，但是区别在于要列举出全部组合
//所以这里要明确退出条件，很穷表明，如果small>sum/2，则序列之后必然已经超过sum，可以退出
//难度：3*
func FindSequence(sum int) [][]int {
    if sum < 3 {
        return nil
    }

    mid := (1+sum)/2 //作为退出条件
    small := 1
    big := 2
    curr := small + big

    ret := [][]int{}
    for small < mid {
        if curr == sum {
            t := []int{}
            for i:=small; i<=big; i++ {
                t = append(t, i)
            }
            ret = append(ret, t)

            //二选一均可，打破循环
            //big ++
            //curr += big
            curr -= small
            small ++
        } else if curr < sum {
            big ++
            curr += big
        } else {
            curr -= small
            small ++
        }
    }

    return ret
}