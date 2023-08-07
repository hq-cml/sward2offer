/*
 * mt面试题：一个字符串，递归的去除b，再递归去除'ac'
 * 比如：
 * acbdc => acdc => dc;
 * abcd => acd => d;
 * aabccx => aaccx => acx => x;。
 */
package _filter_string

// 利用栈，来实现递归的消除效果
func Filter(str string) string {
	var stk []byte
	for _, c := range []byte(str) {
		if c == 'b' { //直接去除
			continue
		}
		if len(stk) == 0 {
			stk = append(stk, c)
			continue
		}

		// 利用top操作判断去除
		if stk[len(stk)-1] == 'a' && c == 'c' {
			stk = stk[:len(stk)-1] //pop
		} else {
			stk = append(stk, c)
		}
	}
	return string(stk)
}
