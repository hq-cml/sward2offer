/*
 * 面试题58（一）：翻转单词顺序
 * 题目：输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
 * 为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student."，
 * 则输出"student. a am I"。
 *
 * 面试题58（二）：左旋转字符串
 * 题目：字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
 * 请定义一个函数实现字符串左旋转操作的功能。比如输入字符串"abcdefg"和数
 * 字2，该函数将返回左旋转2位得到的结果"cdefgab"。
 */
package _58_reverse_string

//题目1：翻转每个单词
//思路：比较常见，就是两次翻转：先整体翻转，然后每个单词翻转。
//难度：2*
func ReverseWord(str string) string {
	if len(str) == 0 {
		return str
	}
	//整体翻转
	b := []byte(str)
	reverse(b, 0, len(b)-1)

	//逐个单词翻转
	start := 0
	end := 0
	for end < len(b) {
		if b[end] == ' ' {
			reverse(b, start, end-1)
			start = end + 1
			end = end + 1
		} else {
			end++
		}
	}

	return string(b)
}

//题目2：坐旋转字符串
//思路1：我个人的思路，拼接字符串，然后直接后移。这种方式需要消耗O(n)空间
//思路2：仍然利用翻转思路，三次：先翻转0-(k-1)，再翻转k-(len-1)，然后整体翻转
//难度：3*
func SwapLeft(str string, n int) string {
	if len(str) < n {
		return str
	}

	if n > len(str) {
		n = n % len(str)
	}

	//先翻转局部
	b := []byte(str)
	reverse(b, 0, n-1)
	reverse(b, n, len(b)-1)

	//整体翻转
	reverse(b, 0, len(b)-1)

	return string(b)
}

//反转一段序列
func reverse(str []byte, beg, end int) {
	for beg < end {
		str[beg], str[end] = str[end], str[beg]
		beg++
		end--
	}
}
