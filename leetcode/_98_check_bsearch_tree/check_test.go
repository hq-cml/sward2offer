package _98_check_bsearch_tree

import "testing"

func TestCheck(t *testing.T) {
	type args struct {
		node []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				node: []int{2, 1, 3},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				node: []int{5, 1, 4, 0, 0, 3, 6},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check(tt.args.node); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
