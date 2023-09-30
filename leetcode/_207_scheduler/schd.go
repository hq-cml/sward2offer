/*
 * 207. 课程表
 *
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，
 * 其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
 *
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 *
 * 示例 1：
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：true
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
 * 示例 2：
 *
 * 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
 * 输出：false
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 */
package _207_scheduler

// 思路：这题需要抽象成一个图形式来解决
//
//	很显然，所有的课程节点如果是一个有向无环图，那么题目有解
//	如果存在环，则无解。综上，题目转换成了判断图是否存在环
//	如何判断环？
//	结合两个概念：出度和入度，分别表示节点出去的箭头数和入的箭头数
func CanFinish(numCourses int, prerequisites [][]int) bool {
	var allCourse = map[int]int{}
	for c := 0; c < numCourses; c++ {
		allCourse[c] = 0 // 所有课程的入度
	}

	// 分析所有课程的出度，以及得到课程间的映射关系
	m := map[int][]int{}
	for _, pair := range prerequisites {
		s, d := pair[1], pair[0]
		allCourse[d]++
		m[s] = append(m[s], d)
	}

	for len(allCourse) > 0 {
		findZero := false
		for c, indegree := range allCourse {
			if indegree == 0 {
				// 入度为0，则可以学习
				delete(allCourse, c)

				// 学习完成之后，则当前课程的下游课程的入度应该递减
				for _, t := range m[c] {
					allCourse[t]--
				}
				findZero = true
			}
		}

		// 如果一轮下来，没在发现入度为0的课程且此时还存在没学完的课程
		// 则说明存在环
		if !findZero && len(allCourse) != 0 {
			return false
		}
	}
	return true
}
