package temp

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: []int{5, 1, 3, 4, 2, 7, 8, 6},
				//arr: []int{7,8,6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Quick(tt.args.arr)
			//fmt.Println(tt.args.arr)
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case-1",
			args: args{
				arr: []int{1, 2},
			},
			want: 0,
		},
		{
			name: "case-2",
			args: args{
				arr: []int{2, 1},
			},
			want: 1,
		},
		{
			name: "case1",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				arr: []int{8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: 7,
		},
		{
			name: "case3",
			args: args{
				arr: []int{5, 1, 3, 4, 2, 7, 8, 6},
			},
			want: 4,
		},
		{
			name: "case4",
			args: args{
				arr: []int{5, 1, 3, 4, 2, 7, 6},
			},
			want: 4,
		},
		{
			name: "case5",
			args: args{
				arr: []int{5, 3, 7, 6, 4, 1, 0, 2, 9, 10, 8},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.arr); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.args.arr)
		})
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: []int{5, 1, 3, 4, 2, 7, 8, 6},
			},
		},
		{
			name: "case2",
			args: args{
				arr: []int{5, 1, 3, 4, 2, 7, 6},
			},
		},
		{
			name: "case3",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
		{
			name: "case4",
			args: args{
				arr: []int{8, 7, 6, 5, 4, 3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}

func TestReverseList(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	var l *common.ListNode
	l = l.PushNode(2)
	l = l.PushNode(4)
	l = l.PushNode(5)
	l = l.PushNode(8)

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{
				head: l,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseList(tt.args.head)
			got.Foreach(common.NodePrint)
		})
	}
}

func TestPostTree(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				root: common.NewNodeWithChild(1,
					common.NewNodeWithChild(2,
						common.NewNode(4),
						common.NewNode(5)),
					common.NewNodeWithChild(3,
						nil,
						common.NewNode(6))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostTree(tt.args.root)
		})
	}
}

func TestChildBrother(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode
	}{
		{
			name: "case1",
			args: args{
				root: common.NewNodeWithChild(1,
					common.NewNodeWithChild(2,
						common.NewNode(4),
						common.NewNode(5),
					),
					common.NewNodeWithChild(3,
						common.NewNode(6),
						common.NewNode(7),
					),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			common.Pre(tt.args.root)
			fmt.Println()
			common.Mid(tt.args.root)
			fmt.Println()
			common.Post(tt.args.root)
			fmt.Println()

			fmt.Println()

			ChildBrother(tt.args.root)
			common.Pre(tt.args.root)
			fmt.Println()
			common.Mid(tt.args.root)
			fmt.Println()
			common.Post(tt.args.root)
			fmt.Println()
		})
	}
}

func TestCrossReverse(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	l1 := common.NewList()
	l1 = l1.PushNode(1)
	l1 = l1.PushNode(2)

	l2 := common.NewList()
	l2 = l2.PushNode(1)
	l2 = l2.PushNode(2)
	l2 = l2.PushNode(3)

	l3 := common.NewList()
	l3 = l3.PushNode(1)
	l3 = l3.PushNode(2)
	l3 = l3.PushNode(3)
	l3 = l3.PushNode(4)
	l3 = l3.PushNode(5)

	l4 := common.NewList()
	l4 = l4.PushNode(1)
	l4 = l4.PushNode(2)
	l4 = l4.PushNode(3)
	l4 = l4.PushNode(4)
	l4 = l4.PushNode(5)
	l4 = l4.PushNode(6)
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "case1",
			args: args{
				head: l1,
			},
			want: nil,
		},
		{
			name: "case2",
			args: args{
				head: l2,
			},
			want: nil,
		},
		{
			name: "case3",
			args: args{
				head: l3,
			},
			want: nil,
		},
		{
			name: "case4",
			args: args{
				head: l4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CrossReverse(tt.args.head)
			got.Foreach(common.NodePrint)
		})
	}
}

func TestBuildTree(t *testing.T) {
	type args struct {
		pre []int
		mid []int
	}
	tests := []struct {
		name    string
		args    args
		want    *common.TreeNode
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				pre: []int{1, 2, 4, 7, 3, 5, 6, 8},
				mid: []int{4, 7, 2, 1, 5, 3, 8, 6},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildTree(tt.args.pre, tt.args.mid)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			common.Pre(got)
			fmt.Println()
			common.Mid(got)
			fmt.Println()
			common.Post(got)
			fmt.Println()
		})
	}
}

func Test_fibo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibo(tt.args.n); got != tt.want {
				t.Errorf("fibo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPath(t *testing.T) {
	type args struct {
		arr  []byte
		rows int
		cols int
		find string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				arr:  []byte("abtgcfcsjdeh"),
				rows: 3,
				cols: 4,
				find: "bfce",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				arr:  []byte("abtgcfcsjdeh"),
				rows: 3,
				cols: 4,
				find: "abfb",
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				arr:  []byte("abtgcfcsjdeh"),
				rows: 3,
				cols: 4,
				find: "abtgshedjc",
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				arr:  []byte("abtgcfcsjdeh"),
				rows: 3,
				cols: 4,
				find: "abtgshedjca",
			},
			want: false,
		},
		{
			name: "case5",
			args: args{
				arr:  []byte("abtgcfcsjdeh"),
				rows: 3,
				cols: 4,
				find: "q",
			},
			want: false,
		},
		{
			name: "case6",
			args: args{
				arr:  []byte("abtgcfcsjdeh"),
				rows: 3,
				cols: 4,
				find: "g",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPath(tt.args.arr, tt.args.rows, tt.args.cols, tt.args.find); got != tt.want {
				t.Errorf("FindPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxCut(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				length: 8,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxCut(tt.args.length); got != tt.want {
				t.Errorf("MaxCut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	type args struct {
		root1 *common.TreeNode
		root2 *common.TreeNode
	}
	tree1 := common.NewNodeWithChild(8,
		common.NewNode(9),
		common.NewNode(2),
	)

	tree2 := common.NewNodeWithChild(8,
		common.NewNode(9),
		common.NewNode(2),
	)

	tree3 := common.NewNodeWithChild(8,
		common.NewNode(9),
		common.NewNode(3),
	)

	tree4 := common.NewNodeWithChild(8,
		common.NewNodeWithChild(8,
			common.NewNode(9),
			common.NewNodeWithChild(2,
				common.NewNode(4),
				common.NewNode(7),
			),
		),
		common.NewNode(7),
	)

	tree5 := common.NewNodeWithChild(1,
		common.NewNodeWithChild(2,
			common.NewNode(4),
			common.NewNodeWithChild(5,
				common.NewNode(6),
				common.NewNode(7),
			),
		),
		common.NewNode(3),
	)

	tree6 := common.NewNodeWithChild(1,
		common.NewNodeWithChild(5,
			common.NewNode(6),
			common.NewNode(7),
		),
		common.NewNode(3),
	)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				root1: tree1,
				root2: tree2,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				root1: tree1,
				root2: tree3,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				root1: tree4,
				root2: tree1,
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				root1: tree4,
				root2: tree3,
			},
			want: false,
		},
		{
			name: "case5",
			args: args{
				root1: tree5,
				root2: tree6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSubtree(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				n: 3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Generate(tt.args.n)
			fmt.Println(got)
		})
	}
}

func TestPrintTreeIn1Line(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				common.NewNodeWithChild(1,
					common.NewNodeWithChild(2,
						common.NewNode(4),
						common.NewNode(5)),
					common.NewNodeWithChild(3,
						nil,
						common.NewNode(6))),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrintTreeInLine(tt.args.root); (err != nil) != tt.wantErr {
				t.Errorf("PrintTreeIn1Line() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				arr: []int{1, 2, -4, 3, 1, 2, -4},
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				arr: []int{1, -2, 3, 10, -4, 7, 2, -5},
			},
			want:    18,
			wantErr: false,
		},
		{
			name: "case3",
			args: args{
				arr: []int{1, -2, 3, 5, -3, 2},
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "case4",
			args: args{
				arr: []int{1, -2, 3, 5, -1, 2},
			},
			want:    9,
			wantErr: false,
		},
		{
			name: "case5",
			args: args{
				arr: []int{-9, -2, -3, -5, -3},
			},
			want:    -2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcDynamic(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSumPath(t *testing.T) {
	type args struct {
		root *common.TreeNode
		num  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				root: common.NewNodeWithChild(10,
					common.NewNodeWithChild(5,
						common.NewNode(4),
						common.NewNode(7)),
					common.NewNode(12)),
				num: 22,
			},
		},
	}
	for _, tt := range tests {
		path := FindSumPath(tt.args.root, tt.args.num)
		fmt.Println(path)
	}
}

func TestTranslateRecurse(t *testing.T) {
	type args struct {
		org string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "case1",
		//	args: args{
		//		org: "12258",
		//	},
		//	want: 5,
		//},
		{
			name: "case2",
			args: args{
				org: "12",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TranslateRecurse(tt.args.org); got != tt.want {
				t.Errorf("TranslateRecurse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxGiftValue(t *testing.T) {
	type args struct {
		values []int
		rows   int
		cols   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case2",
			args: args{
				values: []int{1, 10, 3, 8, 12, 2, 9, 6, 5, 7, 4, 11, 3, 7, 16, 5},
				rows:   4,
				cols:   4,
			},
			want: 53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxGiftValue(tt.args.values, tt.args.rows, tt.args.cols); got != tt.want {
				t.Errorf("MaxGiftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxNoRepititionStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "case1",
			args: args{
				str: "arabcacfr",
			},
			want:  1,
			want1: 4,
		},
		{
			name: "case2",
			args: args{
				str: "abcdeefghij",
			},
			want:  5,
			want1: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindMaxNoRepititionStr(tt.args.str)
			if got != tt.want {
				t.Errorf("FindMaxNoRepititionStr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindMaxNoRepititionStr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindSequence(t *testing.T) {
	type args struct {
		sum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				sum: 9,
			},
		},
		{
			name: "case2",
			args: args{
				sum: 15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindSequence(tt.args.sum)
			fmt.Println(got)
		})
	}
}

func TestMaxStock(t *testing.T) {
	type args struct {
		price []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				price: []int{9, 11, 8, 5, 7, 12, 16, 14},
			},
			want: 11,
		},
		{
			name: "case2",
			args: args{
				price: []int{9, 11, 8, 5, 5, 5, 6},
			},
			want: 2,
		},
		{
			name: "case3",
			args: args{
				price: []int{11, 9, 8, 7, 6, 5},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxStock(tt.args.price); got != tt.want {
				t.Errorf("MaxStock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum1ton_Recurse(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum1ton_Recurse(tt.args.n); got != tt.want {
				t.Errorf("Sum1ton_Recurse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				i: 101,
				j: 202,
			},
			want: 303,
		},
		{
			name: "case2",
			args: args{
				i: -5,
				j: 11,
			}, want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDfs(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				root: common.NewNodeWithChild(3,
					common.NewNodeWithChild(1,
						common.NewNode(3),
						nil),
					common.NewNodeWithChild(4, common.NewNode(1),
						common.NewNode(5))),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Good(tt.args.root); got != tt.want {
				t.Errorf("Dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				src: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckRoot(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				root: common.NewNodeWithChild(3,
					common.NewNodeWithChild(1,
						common.NewNode(3),
						nil),
					common.NewNodeWithChild(4, common.NewNode(1),
						common.NewNode(5))),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := Dfs(tt.args.root); got != tt.want {
			//	t.Errorf("CheckRoot() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestCalcCnt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "abc",
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				s: "aba",
			},
			want: 4,
		},
		{
			name: "case3",
			args: args{
				s: "aaa",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcCnt(tt.args.s); got != tt.want {
				t.Errorf("CalcCnt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIslandNum(t *testing.T) {
	type args struct {
		arr [][]byte
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr: [][]byte{
					{
						1, 1, 1, 1, 0,
					},
					{
						1, 1, 0, 1, 0,
					},
					{
						1, 1, 0, 0, 0,
					},
					{
						0, 0, 0, 0, 0,
					},
				},
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				arr: [][]byte{
					{
						1, 1, 0, 0, 0,
					},
					{
						1, 1, 0, 0, 0,
					},
					{
						0, 0, 1, 0, 0,
					},
					{
						0, 0, 0, 1, 1,
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IslandNum(tt.args.arr); got != tt.want {
				t.Errorf("IslandNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestDfsCalc(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				root: common.NewNodeWithChild(1,
					common.NewNode(3),
					common.NewNode(2)),
			},
			want: 25,
		},

		{
			name: "case2",
			args: args{
				root: common.NewNodeWithChild(3,
					common.NewNodeWithChild(1,
						common.NewNode(3),
						nil),
					common.NewNodeWithChild(4, common.NewNode(1),
						common.NewNode(5))),
			},
			want: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DfsCalc(tt.args.root); got != tt.want {
				t.Errorf("DfsCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	type args struct {
		arr [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: [][]byte{
					{
						'X', 'X', 'X', 'X',
					},
					{
						'X', 'O', 'O', 'X',
					},
					{
						'X', 'X', 'O', 'X',
					},
					{
						'X', 'O', 'X', 'X',
					},
				},
			},
		},
		{
			name: "case1",
			args: args{
				arr: [][]byte{
					{
						'X', 'X', 'X', 'X', 'X', 'X',
					},
					{
						'X', 'O', 'O', 'X', 'O', 'X',
					},
					{
						'X', 'X', 'O', 'X', 'O', 'X',
					},
					{
						'X', 'O', 'X', 'X', 'O', 'X',
					},
				},
			},
		},
	}
	for _, tt := range tests {
		Solve(tt.args.arr)
		for _, line := range tt.args.arr {
			fmt.Println(string(line))
		}
	}
}

func TestPermutation(t *testing.T) {
	type args struct {
		arr []byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				arr: []byte("ab"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permutation(tt.args.arr)
			fmt.Println(got)
		})
	}
}

func TestNextPermutation(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: []int{1, 2, 3},
			},
		},
		{
			name: "case2",
			args: args{
				arr: []int{3, 2, 1},
			},
		},
		{
			name: "case3",
			args: args{
				arr: []int{1, 3, 2},
			},
		},
		{
			name: "case4",
			args: args{
				arr: []int{4, 2, 5, 1},
			},
		},
		{
			name: "case5",
			args: args{
				arr: []int{4, 2, 5, 3, 1},
			},
		},
	}
	for _, tt := range tests {
		NextPermutation(tt.args.arr)
		fmt.Println(tt.args.arr)
	}
}

func TestMaxCommonSubLen(t *testing.T) {
	type args struct {
		s1 []byte
		s2 []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s1: []byte("abcdkkk"),
				s2: []byte("baabcdadabc"),
			},
			want: 4,
		},
		{
			name: "case2",
			args: args{
				s1: []byte("abcdhuvwxyzcd"),
				s2: []byte("efabcnabcdefuvwxyz"),
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxCommonSubLen(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("MaxCommonSubLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcMax(t *testing.T) {
	type args struct {
		price []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				price: []int{9, 11, 8, 5, 7, 12, 16, 14},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := CalcMax(tt.args.price); got != tt.want {
			//	t.Errorf("CalcMax() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestRegexp(t *testing.T) {
	type args struct {
		str     string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				str:     "a",
				pattern: "a.",
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				str:     "ab",
				pattern: "a.",
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				str:     "aaaabc",
				pattern: "a*bc",
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				str:     "aaaabc",
				pattern: "a*abc",
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				str:     "abc",
				pattern: "a*bc",
			},
			want: true,
		},
		{
			name: "case6",
			args: args{
				str:     "baaabc",
				pattern: "a*aaabc",
			},
			want: false,
		},
		{
			name: "case7",
			args: args{
				str:     "baaabc",
				pattern: "*aaabc",
			},
			want: false,
		},
		{
			name: "case8",
			args: args{
				str:     "baaabc",
				pattern: ".*bc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Regexp(tt.args.str, tt.args.pattern); got != tt.want {
				t.Errorf("Regexp() = %v, want %v", got, tt.want)
			}
		})
	}
}
