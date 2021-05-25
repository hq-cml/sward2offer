package _29_print_circle_arr

import (
	"errors"
	"fmt"
)

func Print(arr []int, rows, cols int) error {
	if len(arr) != rows*cols || len(arr) == 0 {
		return errors.New("Wrong arr!")
	}

	idx := 0
	for rows > idx*2 && cols > idx*2 {
		printCircle(arr, idx, rows, cols)
		idx ++
	}
	return nil
}

func printCircle(arr []int, idx int, rows, cols int) {
	endRows := rows - idx -1
	endCols := cols - idx -1

	//向右
	for c := idx; c <= endCols; c ++ {
		fmt.Print(arr[idx*cols + c], " ")
	}

	//向下，初始行数要加1，因为已经打印过了
	for r := idx+1; r <= endRows; r++ {
		fmt.Print(arr[r*cols + endCols], " ")
	}

	//向左，初始列减1，因为已经打印过
	for c:=endCols-1; c >= idx; c-- {
		fmt.Print(arr[endRows*cols + c], " ")
	}

	//向上，初始行减1，且最后一行idx不能算，因为一开始打印过了
	for r := endRows-1; r > idx; r-- {
		fmt.Print(arr[r*cols + idx], " ")
	}
}
