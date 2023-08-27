/**
 *
 */
package _394_string_decode

import "testing"

func TestDecodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{
		//	name: "case0",
		//	args: args{
		//		s: "aaa2[bc]",
		//	},
		//	want: "aaabcbc",
		//},
		//{
		//	name: "case1",
		//	args: args{
		//		s: "3[a]2[bc]",
		//	},
		//	want: "aaabcbc",
		//},
		//{
		//	name: "case2",
		//	args: args{
		//		s: "3[a2[c]]",
		//	},
		//	want: "accaccacc",
		//},
		{
			name: "case4",
			args: args{
				s: "100[leetcode]",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeString(tt.args.s); got != tt.want {
				t.Errorf("DecodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
