/*
 * 给定一个二维数组, 一个坐标(x, y), src, dst, 要求实现一个函数:
 * 如果数字==src的替换为dst, 并且如果发生了替换, 则周围(上下左右）也会发生该替换行为.
 * 例子
 * 给定一个数组input如下，参数: src = 2 , dst = 3, xy = (2,2):
    0 2 0 0 0
    0 2 2 2 0
    0 1 2 2 0
    0 2 0 0 0
 * 输出以下数组output:
    0 3 0 0 0
    0 3 3 3 0
    0 1 3 3 0
    0 2 0 0 0
 */
package replace_double_arr

import (
	"errors"
	"fmt"
)

type Pair struct {
	X int
	Y int
}

// 二维数组循环替换
// 这个题不是特别难，关键是各种边界条件的判断
// 不能越界，并且，同时要考虑检测过得元素不能二次检测
func Replace(arr []int, row, col int, src, dst int, x, y int) error {
	if len(arr) != row * col {
		return errors.New("Invalid arr")
	}

	if x < 0 || x >= row || y < 0 || y >= col {
		return errors.New("Invalid input")
	}

	v, err := get(arr, row, col, x, y)
	if err != nil {
		return err
	}

	if v != src {
		return nil
	}

	s := []Pair{}
	s = append(s, Pair{x,y})

	//防止二次检测
	uniq := map[string]struct{}{
		fmt.Sprintf("%v_%v", x, y): {},
	}

	for len(s) > 0 {
		item := s[len(s)-1]
		arr [item.X * col + item.Y] = dst
		s = s[:len(s)-1]
		//上
		s = checkAndFix(arr, row, col, item.X-1, item.Y, src, uniq, s)

		//下
		s = checkAndFix(arr, row, col, item.X+1, item.Y, src, uniq, s)

		//左
		s = checkAndFix(arr, row, col, item.X, item.Y-1, src, uniq, s)

		//右
		s = checkAndFix(arr, row, col, item.X, item.Y+1, src, uniq, s)
	}

	return nil
}

//检测上下左右
func checkAndFix(arr []int, row, col int, x, y int, src int, uniq map[string]struct{}, s []Pair) []Pair {
	v, err := get(arr, row, col, x, y)
	_, exist := uniq[fmt.Sprintf("%v_%v", x, y)]
	if err == nil && v == src && !exist {
		s = append(s, Pair{x,y})
		uniq[fmt.Sprintf("%v_%v", x, y)] = struct{}{}
	}
	return s
}

func get(arr []int, row, col int, x, y int) (int, error){
	//越界检测
	if x < 0 || x >= row || y < 0 || y >= col {
		return 0, errors.New("Invalid input")
	}

	return arr[x*col + y], nil
}