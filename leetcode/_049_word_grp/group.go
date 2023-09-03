/*
 * 给你一个字符串数组，请你将字母异位词组合在一起。可以按任意顺序返回结果列表。
 * 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
 *
 * 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
 *
 * 输入: strs = [""]
 * 输出: [[""]]
 *
 * 输入: strs = ["a"]
 * 输出: [["a"]]
 */
package _049_word_grp

import "sort"

// 思路：并不难，很容易想到字符串先排序然后利用map
//
//	优化的点是map的设计，这里设计成了string=>idx
//	这样只需要一次性的遍历，即解决了所有的问题
func Group(strs []string) [][]string {
	uniq := map[string]int{}
	idx := 0
	var ret [][]string
	for _, s := range strs {
		t := sortStr(s)
		if _, ok := uniq[t]; !ok {
			uniq[t] = idx // 分组的索引
			idx++
			ret = append(ret, []string{})
		}
		// 找到对应的分组
		ret[uniq[t]] = append(ret[uniq[t]], s)
	}
	return ret
}

func sortStr(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
