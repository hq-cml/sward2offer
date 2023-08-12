/**
 * 一群朋友组队玩游戏，至少有5组人，一组至少2人，要求：
 * 1.每2个人组一队或者3个人组一队，每个人只能加到一个队伍里，不能落单
 * 2.2人队和3人队各自的队伍数均不得少于1，队伍中的人不能来自相同组
 * 3.随机组队，重复执行程序得到的结果不一样，总队伍数也不能一样
 * 4.必须有注释
 * 注：要同时满足条件1-4
 *
 * 举例：
 * GroupList=[#小组列表
 * ['少华','少平','少军','少安','少康'],
 * ['福军','福堂','福民','福平','福心']
 * ['小明','小红','小花','小丽','小强'],
 * ['大壮','大力','大1','大2','大3'],
 * ['阿花','阿朵','阿蓝','阿紫','阿红'],
 * ['A','B','C','D','E'],
 * ['一','二','三','四','五'],
 * ]
 *
 * 输入：GroupList
 * 示例输出:[小强 大3][阿红 E][五 少华][福军 小明][大壮 阿花][A 一][少平 福堂][小红 大力][阿朵 B][二 少军]
 * [福民 小花][大1 阿蓝][C 三][少安 福平 小丽][大2 阿紫 D][四 少康 福心]
 */
package devide_group

import (
	"errors"
	"math/rand"
	"time"
)

// 随机种子初始化
func init() {
	rand.Seed(time.Now().Unix())
}

func DevideGroup(groupList [][]string) ([][]string, error) {
	// 入参校验
	err := checkParams(groupList)
	if err != nil {
		return nil, err
	}

	// 将每个初始队列洗牌，保证充分的随机性

	// 进行分组
	grpPtr := make([]int, len(groupList)) // 每个输入组的下一个候选人指针
	var result [][]string
	devide(groupList, grpPtr, &result)

	return result, nil
}

func devide(grpList [][]string, grpPtr []int, result *[][]string) bool {
	// 随机确定队人数，每2个人组一队或者3个人组一队
	memberCnt := 2 + rand.Intn(2)

	// 查看还有哪些队列非空
	var candidateGrp []int
	for idx, grp := range grpList {
		if grpPtr[idx] == len(grp) { // 当前组已经空了
			continue
		} else {
			candidateGrp = append(candidateGrp, idx)
		}
	}
	if memberCnt > len(candidateGrp) {
		// 组数不够了
		return false
	}

	// 随机选出候选组，然后选出组首成员，组队
	pickedGrp := randPick(candidateGrp, memberCnt)
	currGrp := []string{}
	for _, idx := range pickedGrp {
		currGrp = append(currGrp, grpList[idx][grpPtr[idx]])
		grpPtr[idx]++
	}

	// 将组队结果暂存
	*result = append(*result, currGrp)

	// 进行完美结局的判断
	if succOver(grpList, grpPtr, result) {
		return true
	}

	// DFS深度探测剩下的
	if devide(grpList, grpPtr, result) {
		// 探测成功！
		return true
	}

	// 探测失败，则当前状态要回溯，然后重新探测
	*result = (*result)[:len(*result)-1]
	for _, idx := range pickedGrp {
		grpPtr[idx]--
	}
	return devide(grpList, grpPtr, result)
}

// 完美结局判断，所有的候选组，全部用完 & 组队同时存在2个和3个成员
func succOver(grpList [][]string, grpPtr []int, result *[][]string) bool {
	for idx, grp := range grpList {
		if grpPtr[idx] < len(grp) {
			return false
		}
	}

	has2 := false
	has3 := false
	for _, grp := range *result {
		if len(grp) == 2 {
			has2 = true
		}
		if len(grp) == 3 {
			has3 = true
		}
	}
	return has2 && has3
}

// 从一个slice中随机选择cnt个数
func randPick(a []int, cnt int) []int {
	rand.Shuffle(len(a), func(i int, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a[:cnt]
}

// 入参校验：至少有5组人，一组至少2人
func checkParams(groupList [][]string) error {
	if len(groupList) < 5 {
		return errors.New("Numbers of group must >= 5")
	}
	for _, grp := range groupList {
		if len(grp) < 2 {
			return errors.New("Numbers of each group must >= 2")
		}
	}
	return nil
}
