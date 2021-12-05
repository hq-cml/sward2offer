package _20_string_is_numeric

import (
	"reflect"
	"testing"
)

func Test_scanUnSignedInt(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 bool
	}{
		{
			name:  "case1",
			args:  args{
				str: []byte("0123"),
			},
			want:  []byte{},
			want1: true,
		},
		{
			name:  "case2",
			args:  args{
				str: []byte("012a3"),
			},
			want:  []byte("a3"),
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := scanUnSignedInt(tt.args.str)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanUnSignedInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("scanUnSignedInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
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
				str :"+100",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				str :"5e2",
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				str :"-123",
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				str :"3.142",
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				str :"-1E-16",
			},
			want: true,
		},
		{
			name: "case6",
			args: args{
				str :"12e",
			},
			want: false,
		},
		{
			name: "case7",
			args: args{
				str :"1a3.15",
			},
			want: false,
		},
		{
			name: "case8",
			args: args{
				str :"1.2.3",
			},
			want: false,
		},
		{
			name: "case9",
			args: args{
				str :"+-5",
			},
			want: false,
		},
		{
			name: "case10",
			args: args{
				str :"12e+5.4",
			},
			want: false,
		},
		{
			name: "case11",
			args: args{
				str: "12.3e+5",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.args.str); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}