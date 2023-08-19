package _bracket_match

import (
	"github.com/hq-cml/sward2offer/temp"
	"testing"
)

func TestMatch(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				str: "({{",
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				str: ")",
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				str: "()",
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				str: "{()[]}",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := temp.Match(tt.args.str); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
