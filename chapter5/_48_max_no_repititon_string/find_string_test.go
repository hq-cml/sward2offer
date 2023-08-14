package _48_max_no_repititon_string

import (
	"fmt"
	"testing"
)

func TestFindMaxNoRepititionString(t *testing.T) {
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
			name:  "case1",
			args:  args{
				str: "arabcacfr",
			},
			want:  1,
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindMaxNoRepititionString(tt.args.str)
			if got != tt.want {
				t.Errorf("FindMaxNoRepititionString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindMaxNoRepititionString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindMaxNoRepititionStrSlideWindow(t *testing.T) {
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
			name:  "case1",
			args:  args{
				str: "arabcacfr",
			},
			want:  1,
			want1: 4,
		},
		{
			name:  "case2",
			args:  args{
				str: "abcdeefghij",
			},
			want:  5,
			want1: 6,
		},
		{
			name:  "case3",
			args:  args{
				str: "pwwkew",
			},
			want:  2,
			want1: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindMaxNoRepititionStrSlideWindow(tt.args.str)
			if got != tt.want {
				t.Errorf("FindMaxNoRepititionStrSlideWindow() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindMaxNoRepititionStrSlideWindow() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindMaxNoRepititionStrDynamicPlan(t *testing.T) {
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
			name:  "case1",
			args:  args{
				str: "arabcacfr",
			},
			want:  1,
			want1: 4,
		},
		{
			name:  "case2",
			args:  args{
				str: "abcdeefghij",
			},
			want:  5,
			want1: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindMaxNoRepititionStrDynamicPlan(tt.args.str)
			fmt.Println(got, got1)
			if got != tt.want {
				t.Errorf("FindMaxNoRepititionStrDynamicPlan() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindMaxNoRepititionStrDynamicPlan() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
