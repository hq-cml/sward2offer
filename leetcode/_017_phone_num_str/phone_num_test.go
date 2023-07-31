package _017_phone_num_str

import (
	"fmt"
	"testing"
)

func TestPhoneNumStr(t *testing.T) {
	type args struct {
		phoneNum string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				phoneNum: "23",
			},
			want: nil,
		},
		{
			name: "case2",
			args: args{
				phoneNum: "2",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PhoneNumStr(tt.args.phoneNum)
			fmt.Println(got)
		})
	}
}