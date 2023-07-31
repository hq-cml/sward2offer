/*
 * N字形打印字符串
 * 输入：PAYPALISHIRING
 * 转化成N字形矩阵：
 *  P   A   H   N
 *  A P L S I I G
 *  Y   I   R
 * 输出：PAHNAPLSIIGYIR
 */
package _006_n_style_string

import "fmt"
// 思路：
// 首先计算出矩阵的容量
// 将字符串逐个字符计算坐标位置，填入矩阵
// 逐行扫描填写后的矩阵，所有非空的字符拼成字符串返回
// 难点：
// 逐个字符的坐标位置计算，比较烧脑
func Convert(s string, numRows int) string {
	l := len(s)
	if l <= numRows {
		return s
	}
	// 计算矩阵容量，初始化矩阵
	groupSize := numRows*2 - 2
	rows := numRows
	groupCnt := l/groupSize
	if l % groupSize != 0 {
		groupCnt++
	}
	cols := (numRows-1)* groupCnt
	var arr [][]byte
	for i:=0; i<rows; i++ {
		arr = append(arr, make([]byte, cols))
	}

	// 逐个字符计算位置，填入矩阵
	for idx, c := range s {
		// 计算横坐标i
		temp := idx % groupSize
		i := -1
		if temp < rows {
			i = temp
		} else {
			i = rows - (temp-rows) - 2
		}
		// 计算纵坐标j
		temp = idx / groupSize
		j := (rows-1)*temp
		temp = idx % groupSize
		if temp < rows {
			j = j
		} else {
			j = j + (temp-rows + 1)
		}
		// 填入
		arr[i][j] = byte(c)
	}

	// 扫描矩阵得到最终结果
	ret := ""
	for _, line := range arr {
		for _, c := range line {
			if c == 0 {
				fmt.Print(" ", " ")
				continue
			}
			fmt.Print(string(c), " ")
			ret += string(c)
		}
		fmt.Println()
	}
	return ret
}
