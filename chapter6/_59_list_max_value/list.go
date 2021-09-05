/*
 * 实现Max为O(1)复杂度的List
 */
package _59_list_max_value

import "fmt"

//O(1)复杂度实现一个List，支持Max操作
//方法1：
//利用Maxstack结构来实现，即两个栈实现一个队列
type MaxList struct {
	stack1 *MaxStack
	stack2 *MaxStack
}

func NewMaxList() *MaxList {
	return &MaxList{
		stack1: NewMaxStack(),
		stack2: NewMaxStack(),
	}
}

//队尾插入
func (ml *MaxList) Insert(value int) {
	fmt.Println("Insert", value)
	ml.stack1.Push(value)
}

//队头弹出
func (ml *MaxList) Pop() (int, bool) {
	if ml.stack2.Len() > 0 {
		return ml.stack2.Pop()
	}
	if ml.stack1.Len() > 0 {
		for ml.stack1.Len() > 0 {
			v, _ := ml.stack1.Pop()
			fmt.Println("Swap", v)
			ml.stack2.Push(v)
		}
		return ml.stack2.Pop()
	}
	return 0, false
}

func (ml *MaxList) Max() (int, bool) {
	mx1, ok1 := ml.stack1.Max()
	mx2, ok2 := ml.stack2.Max()

	//两个栈均为空
	if !ok1 && !ok2 {
		return 0, false
	}

	//两个栈均非空
	if ok1 && ok2 {
		if mx1 > mx2 {
			return mx1, true
		} else {
			return mx2, true
		}
	}

	//有一个为空
	if ok1 {
		return mx1, true
	} else {
		return mx2, true
	}
}

//方法2：
//思路：参考O(1)复杂度获取Max的Stack的实现
//设置两个队列，一个主队列，一个辅助队列，关键在辅助队列
//如果新增的元素，大于辅助队列的尾巴则不断清空尾巴，直到不满足之后，入辅助队列尾巴（因为被清理掉的那些，不可能作为最大值）
//如果删除的元素，要删除的正好是辅助队列的头相等，则辅助队列和主队列都清空头；否则主队列自己删除
type MaxListArr struct {
	Base []int
	Help []int //辅助队列
}

func NewMaxListArr() *MaxListArr {
	return &MaxListArr{
		Base: []int{},
		Help: []int{},
	}
}
func (l *MaxListArr) PushBack(i int) {
	l.Base = append(l.Base, i)
	//
	if len(l.Help) == 0 {
		l.Help = append(l.Help, i)
	} else {
		//如果新加入的元素，比help的尾巴小，则需要顶替
		for len(l.Help) > 0 && i > l.Help[len(l.Help)-1] {
			l.Help = l.Help[0 : len(l.Help)-1]
		}
		l.Help = append(l.Help, i)
	}
}

func (l *MaxListArr) PopFront() (int, bool) {
	if l.Len() == 0 {
		return 0, false
	}

	head := l.Base[0]
	if head == l.Help[0] {
		l.Help = l.Help[1:]
	}
	l.Base = l.Base[1:]
	return head, true
}

func (l *MaxListArr) Max() (int, bool) {
	if l.Len() == 0 {
		return 0, false
	}
	return l.Help[0], true
}

func (l *MaxListArr) Len() int {
	return len(l.Base)
}
