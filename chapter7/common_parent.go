package chapter7

import (
    "github.com/hq-cml/sward2offer/basic"
    "github.com/hq-cml/sward2offer/common"
)

/*
 * 树节点的公共节点
 * 1. 如果是二叉搜索树，这是一个递归的思路。如果两个节点都比根小，则递归左子树；如果都比根大，递归右子树；如果一大一小，答案就是根。（利用了二叉搜索树的性质）
 * 2. 如果是包含指向父节点的指针，那么这个题就平移成了求两个链表的首个汇合节点。这个前面有
 * 3. 如果上述特殊条件都没有，那么就分别求出从根到两个节点各自的路径，然后从头开始找到第一个共同元素即可。
 */
//求公共节点
func FindCommonParent(root *common.TreeNode, num1, num2 int) (int, bool) {
    path1 := []int{}
    path2 := []int{}

    path1, ok1 := basic.FindPath(root, num1, path1)
    path2, ok2 := basic.FindPath(root, num2, path2)

    if !ok1 || !ok2 {
        return 0, false
    }

    for k, v := range path1 {
        if k > len(path2)-1 {
            //paht1长
            return path1[k-1], true;
        }
        //第一个公共节点
        if v != path2[k] {
            return path1[k-1], true;
        }
    }
    return 0, false
}