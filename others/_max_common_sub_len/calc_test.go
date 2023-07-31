package _max_common_sub_len

import "testing"

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