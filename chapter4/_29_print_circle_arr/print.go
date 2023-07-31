/*
 * 面试题29：顺时针打印矩阵
 * 题目：输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
 */
package _29_print_circle_arr

import (
	"errors"
	"fmt"
)

// 思路：
// 这个题比较抽象借，借助画图，发现规律
// 从外到内，一圈圈的打印，每一圈都是顺时针的顺序
// 难度：5*
func Print(arr []int, rows, cols int) error {
	if len(arr) != rows*cols || len(arr) == 0 {
		return errors.New("Wrong arr!")
	}

	//每一圈左上角的坐标，都是从(idx, idx)开始
	idx := 0

	//退出边界，就是idx>rows/2或者idx>cols/2
	for rows > idx*2 && cols > idx*2 {
		printCircle(arr, idx, rows, cols)
		idx++
	}
	return nil
}

//从矩阵的左上角，顺时针打印一个圈，关键在于细节
func printCircle(arr []int, idx int, rows, cols int) {
	endRows := rows - idx - 1 //rows范围：[idx, endRows]
	endCols := cols - idx - 1 //cols范围：[idx, endCols]

	//向右
	for c := idx; c <= endCols; c++ {
		fmt.Print(arr[idx*cols+c], " ")
	}

	//向下，初始行数要加1，因为已经打印过了
	for r := idx + 1; r <= endRows; r++ {
		fmt.Print(arr[r*cols+endCols], " ")
	}

	//向左，初始列减1，因为已经打印过
	for c := endCols - 1; c >= idx; c-- {
		fmt.Print(arr[endRows*cols+c], " ")
	}

	//向上，初始行减1，且最后一行idx不能算，因为一开始打印过了
	for r := endRows - 1; r > idx; r-- {
		fmt.Print(arr[r*cols+idx], " ")
	}
}
