package _33_bst_valid_order

func CheckValidOrder(order []int) bool {
    //如果空树或者仅一个节点，认为合法
    if len(order) <= 1 {
        return true
    }

    //
    rootIdx := len(order) - 1
    rootVal := order[rootIdx]

    //找到左右节点分界线
    idx := 0
    for ; idx < len(order)-1; idx ++ {
        if order[idx] < rootVal {
            continue
        }
        break
    }

    //检查分界线之后，是否存在违背定义的节点
    for i:=idx; i<len(order); i++ {
        if order[i] < rootVal {
            return false
        }
    }

    //递归校验左右
    return CheckValidOrder(order[0:idx]) &&
        CheckValidOrder(order[idx: rootIdx])
}