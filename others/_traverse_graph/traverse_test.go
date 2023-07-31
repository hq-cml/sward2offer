package _traverse_graph

import (
	"fmt"
	"testing"
)

// 普遍情况，单向调用，不存在环
// 1 -> 2 -> 3
// \
//  \-> 4 -> 6 -> 7
//       \
//        \-> 5
func TestTraverse1(t *testing.T) {
	node := &TopologyNode{
		Val: 1,
		DownStreas: []*TopologyNode{
			{
				Val: 2,
				DownStreas: []*TopologyNode{
					{
						Val: 3,
					},
				},
			},
			{
				Val: 4,
				DownStreas: []*TopologyNode{
					{
						Val: 6,
						DownStreas: []*TopologyNode{
							{
								Val: 7,
							},
						},
					},
					{
						Val: 5,
					},
				},
			},
		},
	}
	passedBy := map[int]struct{}{}
	help := NewStackInt()
	exist, circle := Traverse(node, passedBy, help)
	fmt.Printf("Exist? : %v, %v\n", exist, circle)
}

// 存在公共节点，但不存在环
// 1 -> 2 -> 3
// \
//  \-> 4 -> 6 -> 7 ---> 5
//       \          /
//        \________/
func TestTraverse2(t *testing.T) {
	node5 := &TopologyNode{
		Val: 5,
	}
	node := &TopologyNode{
		Val: 1,
		DownStreas: []*TopologyNode{
			{
				Val: 2,
				DownStreas: []*TopologyNode{
					{
						Val: 3,
					},
				},
			},
			{
				Val: 4,
				DownStreas: []*TopologyNode{
					{
						Val: 6,
						DownStreas: []*TopologyNode{
							{
								Val: 7,
								DownStreas: []*TopologyNode{
									node5,
								},
							},
						},
					},
					node5,
				},
			},
		},
	}
	passedBy := map[int]struct{}{}
	help := NewStackInt()
	exist, circle := Traverse(node, passedBy, help)
	fmt.Printf("Exist? : %v, %v\n", exist, circle)
}

// 存在环的情况：
// 1 -> 2 -> 3
//  \           _________
//   \    4 <--/         \
//    \-> 4 -----> 6 ->  7
//        \
//         \---> 5
func TestTraverse3(t *testing.T) {
	node4 := &TopologyNode{
		Val: 4,
		DownStreas: []*TopologyNode{
			{
				Val: 6,
				DownStreas: []*TopologyNode{
					{
						Val: 7,
					},
				},
			},
			{
				Val: 5,
			},
		},
	}
	node4.DownStreas[0].DownStreas[0].DownStreas = append(node4.DownStreas[0].DownStreas[0].DownStreas, node4)
	node := &TopologyNode{
		Val: 1,
		DownStreas: []*TopologyNode{
			{
				Val: 2,
				DownStreas: []*TopologyNode{
					{
						Val: 3,
					},
				},
			},
			node4,
		},
	}
	passedBy := map[int]struct{}{}
	help := NewStackInt()
	exist, circle := Traverse(node, passedBy, help)
	fmt.Printf("Exist? : %v, %v\n", exist, circle)
}