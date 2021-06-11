package _46_number_translate_str

import "testing"

func Test_translate(t *testing.T) {
	type args struct {
		org string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				org: "12258",
			},
			want: 5,
		},
		{
			name: "case2",
			args: args{
				org: "33333",
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				org: "111",
			},
			want: 3,
		},
		{
			name: "case4",
			args: args{
				org: "222222", //斐波那契
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateRecurse(tt.args.org); got != tt.want {
				t.Errorf("translateRecurse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateLoop(t *testing.T) {
	type args struct {
		org string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				org: "12258",
			},
			want: translateRecurse("12258"),
		},
		{
			name: "case2",
			args: args{
				org: "33333",
			},
			want: translateRecurse("33333"),
		},
		{
			name: "case3",
			args: args{
				org: "111",
			},
			want: translateRecurse("111"),
		},
		{
			name: "case4",
			args: args{
				org: "222222", //斐波那契
			},
			want: translateRecurse("222222"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateLoop(tt.args.org); got != tt.want {
				t.Errorf("translateLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}