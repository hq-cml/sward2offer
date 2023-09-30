package _010_regexp

import "testing"

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
