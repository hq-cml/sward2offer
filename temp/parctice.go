package temp

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"math"
	"sort"
	"strconv"
)

// 二叉排序树转换成双向链表
func Convert(root *common.TreeNode) *common.TreeNode {
	var currTail *common.TreeNode
	convert(root, &currTail)

	//此时nowTail指向了队列尾部，需要不断向前移动到头部
	head := currTail
	for head != nil && head.Left != nil {
		head = head.Left
	}
	return head
}

// 本质上，这是一个中序遍历递归
// 但是，由于指针比较多，所以比较抽象（尤其是root.Right的没有做处理，交给递归过程了）
// 第二个参数currTail，即当前已经处理好的链表尾部，通过这个二级指针，将整个链表顺直
// 难度：5*
func convert(root *common.TreeNode, currTail **common.TreeNode) {
	if root == nil {
		return
	}

	//递归左子树
	if root.Left != nil {
		convert(root.Left, currTail)
	}

	//当前节点是root，处理左节点善后
	root.Left = *currTail
	if (*currTail) != nil {
		(*currTail).Right = root
	}

	//currTail后移成为当前节点（root.Right没有特殊处理，交给下面的递归过程了）
	*currTail = root

	//递归处理右子树
	if root.Right != nil {
		convert(root.Right, currTail)
	}
}

// 检验栈的合法顺序
func CheckValidStackOrder(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	stk := common.NewStackInt()
	for len(a) > 0 {
		tmp, ok := stk.Top()
		if ok && tmp == b[0] {
			stk.Pop()
			b = b[1:]
			continue
		}

		if a[0] == b[0] {
			a = a[1:]
			b = b[1:]
		} else {
			stk.Push(a[0])
			a = a[1:]
		}
	}

	if stk.Len() == 0 {
		return true
	}

	for stk.Len() > 0 {
		v, _ := stk.Pop()
		if v == b[0] {
			b = b[1:]
		} else {
			return false
		}
	}
	return true
}

// 链表倒数第k个节点
func FindBackN(head *common.ListNode, k int) (v int, err error) {
	if head == nil {
		return 0, errors.New("invalid")
	}

	if k <= 0 {
		return 0, errors.New("invalid")
	}

	p1 := head
	for p1 != nil && k > 0 {
		p1 = p1.Next
		k--
	}

	if k > 0 {
		return 0, errors.New("invalid")
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2.Val, nil
}

// 简单的urlenocde替换，' '=> "%20"
func Rep(str []byte) []byte {
	if len(str) == 0 {
		return nil
	}

	var cnt int
	for _, c := range str {
		if c == ' ' {
			cnt++
		}
	}

	l := len(str)
	tmp := make([]byte, 2*cnt)
	str = append(str, tmp...)
	j := len(str) - 1
	for i := l - 1; i >= 0; i-- {
		if str[i] != ' ' {
			str[j] = str[i]
			j--
		} else {
			str[j-2] = '%'
			str[j-1] = '2'
			str[j] = '0'
			j -= 3
		}
	}
	return str
}
// 正则表达式
// 模式中的字符'.'表示任意一个字符
// 而'*'表示它前面的字符可以出现任意次（含0次）
func Regexp(str, pattern string) bool {
	if len(pattern) == 0 {
		if len(str) != 0 {
			return false
		}
		return true
	}

	if pattern[0] == '.' {
		if len(pattern) == 1 {
			if len(str) == 1 {
				return true
			}
			return false
		} else if pattern[1] == '*' {
			return Regexp(str[1:], pattern[2:]) ||
				Regexp(str, pattern[2:]) ||
				Regexp(str[1:], pattern)
		} else {
			if len(str) == 0 {
				return false
			}
			if str[0] == pattern[0] {
				return Regexp(str[1:], pattern[1:])
			}
			return false
		}
	} else if pattern[0] == '*' {
		return false
	} else {
		if len(pattern) == 1 {
			if len(str) == 1 && str[0] == pattern[0] {
				return true
			}
			return false
		} else if pattern[1] == '*' {
			if str[0] == pattern[0] {
				return Regexp(str[1:], pattern[2:]) ||
					Regexp(str[1:], pattern)
			} else {
				return Regexp(str, pattern[2:])
			}
		} else {
			if str[0] == pattern[0] {
				return Regexp(str[1:], pattern[1:])
			}
			return false
		}
	}
	return false
}

// 最大公共子串长度
func MaxCommonSubLen(s1, s2 []byte) int {
	len1 := len(s1)
	len2 := len(s2)
	if len1 == 0 || len2 == 0 {
		return 0
	}
	var dp [][]int
	// len1行，len2列
	for i := 0; i < len1; i++ {
		dp = append(dp, make([]int, len2))
	}
	max := 0
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if s1[i] == s2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}

				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max
}

// 字典序有序排列，下一个值
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,3,2 → 2,1,3
// 4,2,5,1->4 1 2 5
// 5,3,6,2->4 1 2 5
// 2,4,5,1->2 1 4 5
func NextPermutation(arr []int) {
	if len(arr) <= 1 {
		return
	}
	first := -1
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			first = i - 1
			break
		}
	}
	if first == -1 {
		// reverse
		sort.Ints(arr)
		return
	} else {
		spIdx := first + 1
		min := arr[spIdx]
		for i := spIdx + 1; i < len(arr); i++ {
			if arr[i] > arr[first] && min > arr[i] {
				spIdx = i
				min = arr[spIdx]
			}
		}
		arr[first], arr[spIdx] = arr[spIdx], arr[first]
		// 将first后面的值，升序排列
		sort.Ints(arr[first+1:])
	}
}

// 全排列
func Permutation(arr []byte) []string {
	var ret []string
	var dfs func(arr []byte, s string)
	dfs = func(arr []byte, s string) {
		if len(arr) == 0 {
			ret = append(ret, s)
		}
		for i := 0; i < len(arr); i++ {
			c := arr[i]
			var left []byte
			left = append(left, arr[:i]...)
			left = append(left, arr[i+1:]...)
			dfs(left, s+string(c))
		}
	}
	dfs(arr, "")
	return ret
}

// https://zhuanlan.zhihu.com/p/441410615
// 入参是一个二维数组，分别是X和O，将不在边上的O，全部替换成X
func Solve(arr [][]byte) {
	rows := len(arr)
	if rows < 3 {
		return
	}
	cols := len(arr[0])
	if cols < 3 {
		return
	}
	var visited [][]bool
	for i := 0; i < rows; i++ {
		visited = append(visited, make([]bool, cols))
	}
	// 四周替换
	for i := 0; i < rows; i++ {
		dfs1(arr, rows, cols, i, 0, visited)
	}
	for j := 0; j < cols; j++ {
		dfs1(arr, rows, cols, 0, j, visited)
	}
	for i := 0; i < rows; i++ {
		dfs1(arr, rows, cols, i, cols-1, visited)
	}
	for j := 0; j < cols; j++ {
		dfs1(arr, rows, cols, rows-1, j, visited)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if arr[i][j] == 'T' {
				arr[i][j] = 'O'
			} else {
				arr[i][j] = 'X'
			}
		}
	}
}

func dfs1(arr [][]byte, rows, cols, i, j int, visited [][]bool) {
	if i < 0 || i >= rows {
		return
	}
	if j < 0 || j >= cols {
		return
	}
	if visited[i][j] {
		return
	}
	visited[i][j] = true
	if arr[i][j] != 'O' {
		return
	}
	arr[i][j] = 'T'
	dfs1(arr, rows, cols, i+1, j, visited)
	dfs1(arr, rows, cols, i, j+1, visited)
	dfs1(arr, rows, cols, i-1, j, visited)
	dfs1(arr, rows, cols, i, j-1, visited)
}

// * 岛屿数量
// * 入参是一个二维数组，分别是1和0，1标识陆地，0标识水
// * 所有连成片（上下左右）的陆地，算一块岛屿
func IslandNum(arr [][]byte) int {
	// 初步校验
	rows := len(arr)
	if rows < 1 {
		return 0
	}
	cols := len(arr[0])
	if cols < 1 {
		return 0
	}

	// 已访问标记
	var visited [][]bool
	for i := 0; i < rows; i++ {
		visited = append(visited, make([]bool, cols))
	}

	// 逐个探测
	cnt := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 奔着新陆地
			dsfIsland(arr, rows, cols, i, j, visited, &cnt, true)
		}
	}
	return cnt
}

func dsfIsland(arr [][]byte, rows, cols, i, j int, visited [][]bool, cnt *int, new bool) {
	// 边界
	if i < 0 || i >= rows {
		return
	}
	if j < 0 || j >= cols {
		return
	}

	// 已访问
	if visited[i][j] {
		return
	}

	// 访问标记
	visited[i][j] = true

	// 遇到水
	if arr[i][j] == 0 {
		return
	}

	// 遇到陆地
	if new {
		*cnt++ // 是新陆地，岛屿自增
	}
	dsfIsland(arr, rows, cols, i+1, j, visited, cnt, false) // 非新陆地
	dsfIsland(arr, rows, cols, i, j+1, visited, cnt, false)
	dsfIsland(arr, rows, cols, i-1, j, visited, cnt, false)
	dsfIsland(arr, rows, cols, i, j-1, visited, cnt, false)
}

/*
 * 计算一个二叉树路径数字之和
 * 下面数的路径组成12和13，所以最终返回25（12+13）
 *    1
 *  2   3
 */
func DfsCalc(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	var total int
	dfsCalc(root, 0, &total)
	return total
}

func dfsCalc(root *common.TreeNode, curr int, total *int) {
	if root == nil {
		return
	}
	curr = curr*10 + root.Val
	// 是叶子
	if root.Left == nil && root.Right == nil {
		*total = *total + curr
	}
	dfsCalc(root.Left, curr, total)
	dfsCalc(root.Right, curr, total)
}

// 回文子字符串数量
// aaa => 6
// abc => 3
//
// 回文判断：
// 一个子字符串str[i,j]是回文，必然是str[i] == str[j]，且满足下面三种情况之一：
// 单个字母：i==j
// 相邻字母：i==j+1
// 不相邻字母：str[i+1, j-1]是回文
// 通过dp二维数组，预存储[i,j]的子串是否是回文子串，布尔类型
func CalcCnt(s string) int {
	if len(s) == 0 {
		return 0
	}
	length := len(s)
	// dp[i][j]表示区间范围[i,j]的子串是否是回文子串，布尔类型
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	var ret int
	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if s[i] != s[j] {
				continue
			}
			// s[i] == s[j]，到达此处i和j指向的字母已经是相同的了
			if (j-i <= 1) || // 如果i和j指向同一个字母或者i和j相邻，则必然是回文
				dp[i+1][j-1] { // 如果i和j不相邻，则查看i+1到j-1是否是回文，如果是则这里也是回文
				ret++
				dp[i][j] = true
			}
		}
	}

	return ret
}

// Good节点个数
func Good(root *common.TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *common.TreeNode, preMax int) int {
	if root == nil {
		return 0
	}
	num := 0
	if preMax <= root.Val {
		num = 1
		preMax = root.Val
	}
	return num + dfs(root.Left, preMax) + dfs(root.Right, preMax)
}

// 数组对角乘法，不包含自己
func Multiply(src []int) []int {
	l := len(src)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return []int{1}
	}
	left := make([]int, l)
	right := make([]int, l)
	left[0] = 1
	right[l-1] = 1
	for i := 1; i < l; i++ {
		left[i] = left[i-1] * src[i-1]
	}
	for i := l - 2; i >= 0; i-- {
		right[i] = right[i+1] * src[i+1]
	}
	ret := make([]int, l)
	for i := 0; i < l; i++ {
		ret[i] = left[i] * right[i]
	}
	return ret
}

// 不用加法的加法
func Add(i, j int) int {
	for i&j != 0 {
		x := i ^ j
		i = (i & j) << 1
		j = x
	}
	return i | j
}

func Sum1ton_ArithmeticProgression(n int) int {
	return int(int(math.Pow(float64(n), 2))+n) >> 1
}

// 求1+2+…+n，只能使用加法，不使用基本运算符和其他逻辑控制
func Sum1ton_Recurse(n int) int {
	sum := 0
	add(n, &sum)
	return sum
}

func add(n int, sum *int) bool {
	*sum = *sum + n
	return (n == 0) || add(n-1, sum)
}

// 9, 11, 8, 5, 7, 12, 16, 14
func MaxStock(price []int) int {
	if len(price) < 2 {
		return 0
	}
	currMin := price[0]
	maxDiff := 0
	for i := 1; i < len(price); i++ {
		if (price[i] - currMin) > maxDiff {
			maxDiff = price[i] - currMin
		}
		if price[i] < currMin {
			currMin = price[i]
		}
	}
	return maxDiff
}

// 递增序列1,2,3... ，给定一个和
func FindSequence(sum int) [][]int {
	if sum < 1 || sum == 2 {
		return nil
	}
	if sum == 1 {
		return [][]int{{1}}
	}
	i := 1
	j := 2
	curr := i + j
	var ret [][]int

	for i != j {
		if curr < sum {
			j++
			curr += j
		} else if curr == sum {
			var s []int
			for x := i; x <= j; x++ {
				s = append(s, x)
			}
			ret = append(ret, s)
			j++
			curr += j
		} else {
			curr -= i
			i++
		}
	}
	return ret
}

// 最长不重复子串
func FindMaxNoRep(str string) (int, int) {
	if len(str) == 0 {
		return 0, 0
	}

	uniq := map[byte]int{}
	idx := 0     // 最终返回值，最长子串的idx
	max := -1    // 最终返回值，最长子串的长度
	currIdx := 0 // 当前子串的开始idx
	for i, c := range []byte(str) {
		if repIdx, ok := uniq[c]; !ok {
			uniq[c] = i
			length := i - currIdx + 1
			if length > max {
				max = length
				idx = currIdx // 将当前开始idx，作为返回结果
			}
		} else {
			// 将重复的子串部分，全部清除掉
			for j := currIdx; j <= repIdx; j++ {
				delete(uniq, str[j])
			}
			uniq[c] = i          // 重复部分已清除，新字符入uniq
			currIdx = repIdx + 1 // 当前子串的起始位置从清除后的位置开始
		}
	}

	return idx, max
}

// f(n, n) = max(f(n-1, n), f(n, n-1)) + v(n,n)
func MaxGiftValue(values []int, rows, cols int) int {

	return maxGiftValue(values, rows, cols, rows-1, cols-1)
}

func maxGiftValue(values []int, rows, cols int, x, y int) int {
	if x == 0 && y == 0 {
		return values[0]
	}
	if x < 0 || y < 0 {
		return 0
	}
	return values[x*cols+y] + int(
		math.Max(float64(maxGiftValue(values, rows, cols, x-1, y)),
			float64(maxGiftValue(values, rows, cols, x, y-1))))
}

func TranslateRecurse(org string) int {
	if len(org) == 0 {
		return 0
	}
	if len(org) == 1 {
		return 1
	}
	if len(org) == 2 {
		if check(org) {
			return 2
		}
		return 1
	}

	if !check(org[0:2]) {
		return TranslateRecurse(org[1:])
	} else {
		return TranslateRecurse(org[1:]) +
			TranslateRecurse(org[2:])
	}
}

// 判断一个2字符的串，是否在[10, 25]之间
func check(str string) bool {
	if len(str) != 2 {
		return false
	}
	i, _ := strconv.Atoi(str)
	if i >= 10 && i <= 25 {
		return true
	} else {
		return false
	}
}

// 连续子数组最大值
func CalcDynamic(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, nil
	}
	max := arr[0]
	sum := arr[0]
	for j := 1; j < len(arr); j++ {
		if sum < 0 {
			sum = 0
		}
		sum = sum + arr[j]
		if sum > max {
			max = sum
		}
	}
	return max, nil
}

// 二叉树中和为某一值的路径
func FindSumPath(root *common.TreeNode, num int) []int {
	var path []int
	if root == nil {
		return nil
	}
	path, ok := findSumPath(root, num, 0, path)
	if !ok {
		return nil
	}
	return path
}

func findSumPath(root *common.TreeNode, num, curr int, path []int) ([]int, bool) {
	curr += root.Val
	path = append(path, root.Val)
	if curr == num {
		return path, true
	}
	if root.Left != nil {
		pathLeft, ok := findSumPath(root.Left, num, curr, path)
		if ok {
			return pathLeft, ok
		}
	}
	if root.Right != nil {
		pathRight, ok := findSumPath(root.Right, num, curr, path)
		if ok {
			return pathRight, ok
		}
	}

	//path = path[0 : len(path)-1]
	return path, false
}

func PrintTreeInLine(root *common.TreeNode) error {
	if root == nil {
		return nil
	}
	l := list.New()
	l.PushBack(root)
	curLineCnt := 1
	nextLineCnt := 0
	for l.Len() > 0 {
		e := l.Front()
		node := e.Value.(*common.TreeNode)
		l.Remove(e)
		if node == nil {
			continue
		}
		fmt.Print(node.Val, " ")
		curLineCnt--

		l.PushBack(node.Left)
		nextLineCnt++
		l.PushBack(node.Right)
		nextLineCnt++
		if curLineCnt == 0 {
			curLineCnt = nextLineCnt
			nextLineCnt = 0
			fmt.Println()
		}
	}
	return nil
}

// 判断二叉树是否对称
func Symmetry(root *common.TreeNode) bool {
	return symmetry(root, root)
}

func symmetry(root1, root2 *common.TreeNode) bool {
	// 结束条件
	if root1 == nil && root2 == nil {
		return true
	}
	// 结束条件：有一个非空，必然不是镜像
	if root1 == nil || root2 == nil {
		return false
	}

	// 结束条件：根不相同，必然不是镜像
	if root1.Val != root2.Val {
		return false
	}

	//递归比较
	//注意顺序：左子树和右子树对比
	return symmetry(root1.Left, root2.Right) &&
		symmetry(root1.Right, root2.Left)
}

// 求镜像
func Mirror(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return root
	}

	root.Left = Mirror(root.Right)
	root.Right = Mirror(root.Left)
	return root
}

func CheckSubtree(root1, root2 *common.TreeNode) bool {
	if checkSubtree(root1, root2) {
		return true
	}
	return CheckSubtree(root1.Left, root2) ||
		CheckSubtree(root1.Right, root2)
}

func checkSubtree(root1, root2 *common.TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return checkSubtree(root1.Left, root2.Left) &&
		checkSubtree(root1.Right, root2.Right)
}

func Merge(head1, head2 *common.ListNode) *common.ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	var head *common.ListNode
	if head1.Val < head2.Val {
		head = head1
		head1 = head1.Next
	} else {
		head = head2
		head2 = head2.Next
	}
	head.Next = Merge(head1, head2)
	return head
}

func Generate(n int) []string {
	result := []string{}
	// 初始字符串为空，且左Remain和右Remain都是n
	gen(n, n, "", &result)
	return result
}

func gen(leftRemain, rightRemain int, str string, s *[]string) {
	if leftRemain == 0 && rightRemain == 0 {
		*s = append(*s, str)
		return
	}
	if leftRemain > rightRemain {
		return
	}
	if leftRemain > 0 {
		gen(leftRemain-1, rightRemain, str+"(", s)
	}
	if rightRemain > 0 {
		gen(leftRemain, rightRemain-1, str+")", s)
	}
}

func MaxCut(length int) int {
	//一些特殊情况
	if length <= 1 {
		return 0
	}
	if length == 2 { //1*1
		return 1
	}
	if length == 3 { //1*2
		return 2
	}

	maxArr := make([]int64, length+1)
	maxArr[1] = 1
	maxArr[2] = 2
	maxArr[3] = 3
	var max int64
	for i := 4; i <= length; i++ {
		max = math.MinInt64
		for j := 1; j < i; j++ {
			if max < maxArr[j]*maxArr[i-j] {
				max = maxArr[j] * maxArr[i-j]
			}
		}
		maxArr[i] = max
	}
	return int(maxArr[length])
}

func FindPath(arr []byte, rows, cols int, find string) bool {
	if len(arr) == 0 {
		return false
	}
	if find == "" {
		return true
	}
	visited := make([]bool, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if findPath(arr, rows, cols, i, j, find, 0, visited) {
				return true
			}
		}
	}
	return false
}

func findPath(arr []byte, rows, cols int, x, y int, find string, strIdx int, visited []bool) bool {
	if strIdx == len(find) {
		return true
	}

	if x < 0 || x > rows-1 || y < 0 || y > cols-1 {
		return false
	}

	if arr[x*cols+y] != find[strIdx] || visited[x*cols+y] {
		return false
	}

	visited[x*cols+y] = true

	if findPath(arr, rows, cols, x+1, y, find, strIdx+1, visited) ||
		findPath(arr, rows, cols, x, y+1, find, strIdx+1, visited) ||
		findPath(arr, rows, cols, x-1, y, find, strIdx+1, visited) ||
		findPath(arr, rows, cols, x, y-1, find, strIdx+1, visited) {
		return true
	} else {
		visited[x*cols+y] = false
		return false
	}
}

func HeapSort(arr []int) {
	l := len(arr)
	if l == 0 || l == 1 {
		return
	}

	for bondIdx := l - 1; bondIdx > 0; bondIdx-- {
		HeapFy(arr, bondIdx)
		arr[0], arr[bondIdx] = arr[bondIdx], arr[0]
	}
}

func HeapFy(arr []int, bondIdx int) {
	for i := len(arr) - 1; i >= 0; i-- {
		heapfy(arr, i, bondIdx)
		fmt.Println(arr)
	}
}

func heapfy(arr []int, currIdx, bondIdx int) {
	if currIdx >= bondIdx {
		return
	}
	leftIdx := currIdx*2 + 1
	rightIdx := currIdx*2 + 2
	maxIdx := currIdx
	if leftIdx <= bondIdx && arr[leftIdx] > arr[maxIdx] {
		maxIdx = leftIdx
	}
	if rightIdx <= bondIdx && arr[rightIdx] > arr[maxIdx] {
		maxIdx = rightIdx
	}
	if currIdx != maxIdx {
		arr[currIdx], arr[maxIdx] = arr[maxIdx], arr[currIdx]
		heapfy(arr, maxIdx, bondIdx)
	}
}

func fibo(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	n1 := 0
	n2 := 1
	for i := 2; i <= n; i++ {
		tmp := n2
		n2 = n1 + n2
		n1 = tmp
	}
	return n2
}

func BuildTree(pre, mid []int) (*common.TreeNode, error) {
	if len(pre) != len(mid) {
		return nil, errors.New("Error")
	}
	if len(pre) == 0 {
		return nil, nil
	}
	// 找到根
	rootVal := pre[0]
	midIdx := 0
	for ; midIdx < len(mid); midIdx++ {
		if mid[midIdx] == rootVal {
			break
		}
	}
	root := common.NewNode(rootVal)
	root.Left, _ = BuildTree(pre[1:midIdx+1], mid[0:midIdx])
	root.Right, _ = BuildTree(pre[midIdx+1:], mid[midIdx+1:])
	return root, nil
}

func CrossReverse(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := head.Next
	p1 := head
	p2 := head.Next
	for p1 != nil && p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		if p3 == nil {
			p1.Next = nil
		} else if p3.Next == nil {
			p1.Next = p3
		} else {
			p1.Next = p3.Next
		}

		p1 = p3
		if p1 != nil {
			p2 = p1.Next
		}
	}

	return ret
}

func ChildBrother(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return root
	}
	right := ChildBrother(root.Right)

	if root.Left == nil {
		root.Left = right
	} else {
		left := ChildBrother(root.Left)
		left.Right = right
	}
	root.Right = nil
	return root
}

func PostTree(root *common.TreeNode) {
	if root == nil {
		return
	}
	stk := common.NewStack(false)
	p := root
	var pre *common.TreeNode
	for p != nil || stk.Len() > 0 {
		for p != nil {
			stk.Push(p)
			p = p.Left
		}
		top, _ := stk.Top()
		node := top.(*common.TreeNode)
		if node.Right != nil && node.Right != pre {
			p = node.Right
		} else {
			fmt.Println(node.Val)
			stk.Pop()
			pre = node
		}

	}
}

func MidTree(root *common.TreeNode) {
	if root == nil {
		return
	}
	stk := common.NewStack(false)
	p := root
	for p != nil || stk.Len() > 0 {
		for p != nil {
			stk.Push(p)
			p = p.Left
		}
		top, _ := stk.Pop()
		node := top.(*common.TreeNode)
		fmt.Println(node.Val)
		if node.Right != nil {
			p = node.Right
		}
	}
}

func PreTree(root *common.TreeNode) {
	if root == nil {
		return
	}
	stk := common.NewStack(false)
	//stk.Push(root)
	p := root
	for p != nil || stk.Len() > 0 {
		for p != nil {
			stk.Push(p)
			fmt.Println(p.Val)
			p = p.Left
		}
		top, _ := stk.Pop()
		node := top.(*common.TreeNode)
		if node.Right != nil {
			p = node.Right
		}
	}
}

// reverse list
func ReverseList(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1, p2 := head, head.Next
	head.Next = nil
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}

// tree height
func TreeHeight(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(TreeHeight(root.Left)),
		float64(TreeHeight(root.Right))))
}

// binary search
func BinarySearch(arr []int, need int) int {
	if len(arr) == 0 {
		return -1
	}
	beg := 0
	end := len(arr) - 1

	for beg <= end {
		midIdx := beg + (end-beg)/2
		if arr[midIdx] == need {
			return midIdx
		} else if arr[midIdx] > need {
			end = midIdx - 1
		} else {
			beg = midIdx + 1
		}
	}
	return -1
}

// Quick Sort
func QuickSort(arr []int) {
	length := len(arr)
	if length == 0 || length == 1 {
		return
	}

	// 对arr进行左右的分区
	idx := partition(arr)
	if idx == -1 {
		// 不应该出现的异常情况
		return
	}

	// 分别递归左右
	if idx != 0 {
		QuickSort(arr[:idx])
	}
	if idx != length-1 {
		QuickSort(arr[idx+1:])
	}
}

func partition(arr []int) int {
	// 边界校验
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return 0
	}

	// 第一个值作为基准
	val := arr[0]
	beg := 0
	end := len(arr) - 1

	// 进行多轮次向心逼近，直到相遇
	for beg < end {
		// 从后向前逼近
		for beg < end && arr[end] >= val {
			end--
		}
		arr[beg], arr[end] = arr[end], arr[beg]
		// 从前向后逼近
		for beg < end && arr[beg] <= val {
			beg++
		}
		arr[beg], arr[end] = arr[end], arr[beg]
	}
	// beg和end已经相遇
	return beg
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	midIdx := length / 2
	arr1 := MergeSort(arr[0:midIdx])
	arr2 := MergeSort(arr[midIdx:])

	var ret []int
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			ret = append(ret, arr1[0])
			arr1 = arr1[1:]
		} else {
			ret = append(ret, arr2[0])
			arr2 = arr2[1:]
		}
	}
	if len(arr1) > 0 {
		ret = append(ret, arr1...)
	}
	if len(arr2) > 0 {
		ret = append(ret, arr2...)
	}
	return ret
}

// select sort
func SelectSort(arr []int) {
	length := len(arr)
	if length == 0 || length == 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		minIdx := i
		for j := i; j < length; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// insert sort
func InsertSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// bubble sort
func BubbleSort(arr []int) {
	length := len(arr)
	if length == 0 || length == 1 {
		return
	}

	for i := 1; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
