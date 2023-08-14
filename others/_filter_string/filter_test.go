package _filter_string

import "testing"

func TestFilter(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				str: "acbdc",
			},
			want: "dc",
		},
		{
			name: "case2",
			args: args{
				str: "abcd",
			},
			want: "d",
		},
		{
			name: "case3",
			args: args{
				str: "aabccx",
			},
			want: "x",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.str); got != tt.want {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
