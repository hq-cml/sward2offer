package _35_copy_complex_list

type CompNode struct {
    Val int
    Next  *CompNode
    Radom *CompNode
}

func Copy(src *CompNode) *CompNode {
    if src == nil {
        return nil
    }

    pSrc := src
    var dst *CompNode
    var pDst *CompNode
    pMap := map[*CompNode]*CompNode{}

    //复制基本字段
    for pSrc != nil {
        node := &CompNode{
            Val:   pSrc.Val,
        }
        if dst == nil {
            dst = node
            pDst = dst
        } else {
            pDst.Next = node
            pDst = pDst.Next
        }
        pMap[pSrc] = node //空间换时间，备份一个old=>new的逐个节点
        pSrc = pSrc.Next

    }

    //复制随机字段
    pSrc = src
    pDst = dst
    for pSrc != nil {
        if pSrc.Radom != nil {
            pDst.Radom = pMap[pSrc.Radom]
        }
        pSrc = pSrc.Next
        pDst = pDst.Next
    }

    return dst
}