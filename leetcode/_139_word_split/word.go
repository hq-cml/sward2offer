/*
 * 单词拆分
 * 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
 * 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
 *
 * 示例 1：
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
 *
 * 示例 2：
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
 *
 * 示例 3：
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 */
package _139_word_split

// 思路：这题明显的就应该用DFS
// 难点：如果简单的递归，则会在一些特殊case下面陷入深度循环
//      问题的根源是因为存在大量重复的计算
//      所以引入记忆递归，将一些存在重复的情况，进行快速查表返回
func WordBreak(s string, wordDict []string) bool {
	failed := map[string]struct{}{}
	return wordbreak(s, wordDict, failed)
}

// 每一次递归都是新的，因为wordDict可以复用，所以只需要一个map，就可以大量的减少重复计算！！
func wordbreak(s string, wordDict []string, failed map[string]struct{}) bool {
	if len(s) == 0 {
		return true
	}

	// 查表快速返回
	if _, ok := failed[s]; ok {
		return false
	}

	for _, word := range wordDict {
		length := len(word)
		if len(s) < length || string(s[:length]) != word {
			continue
		}

		if wordbreak(s[length:], wordDict, failed) {
			return true
		}

		// 最最精髓的在此，将失败的结果存入map表！！
		failed[s] = struct{}{}
	}

	return false
}
