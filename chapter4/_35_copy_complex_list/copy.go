/*
 * 面试题35：复杂链表的复制
 * 题目：请实现函数ComplexListNode* Clone(ComplexListNode* pHead)，复
 * 制一个复杂链表。在复杂链表中，每个结点除了有一个m_pNext指针指向下一个
 * 结点外，还有一个m_pSibling 指向链表中的任意结点或者nullptr。
 */
package _35_copy_complex_list

type CompNode struct {
    Val int
    Next  *CompNode
    Radom *CompNode
}

//思路：这个题目作者给了几个思路，
//第一个最朴素的就是多次循环操作。
//第二个思路是空间换时间，就是利用一个多出来的hash表。存储的是老节点=>新节点的指针映射关系。
//     这样重建之后可以以O(N)的复杂度来实现m_pSibling指针的
//     这需要简单画个图，相对会清晰很多
//第三个思路比较巧妙，分别在节点后面插入节点，然后最后在切分链表。。。但是这个思路感觉在面试中不可能搞出来。
func Copy(src *CompNode) *CompNode {
    if src == nil {
        return nil
    }

    pSrc := src
    var dst *CompNode
    var pDst *CompNode
    pMap := map[*CompNode]*CompNode{} //指针=>指针，old=>new

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
            pDst.Radom = pMap[pSrc.Radom] //牛逼
        }
        pSrc = pSrc.Next
        pDst = pDst.Next
    }

    return dst
}