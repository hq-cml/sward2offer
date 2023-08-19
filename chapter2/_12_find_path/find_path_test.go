package _12_find_path

import "testing"

func TestFindPath(t *testing.T) {
	type args struct {
		arr  [][]byte
		rows int
		cols int
		find string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				arr: [][]byte{
					[]byte("abtg"),
					[]byte("cfcs"),
					[]byte("jdeh"),
				},
				rows: 3,
				cols: 4,
				find: "bfce",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				arr: [][]byte{
					[]byte("abtg"),
					[]byte("cfcs"),
					[]byte("jdeh"),
				},
				rows: 3,
				cols: 4,
				find: "abfb",
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				arr: [][]byte{
					[]byte("abtg"),
					[]byte("cfcs"),
					[]byte("jdeh"),
				},
				rows: 3,
				cols: 4,
				find: "abtgshedjc",
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				arr: [][]byte{
					[]byte("abtg"),
					[]byte("cfcs"),
					[]byte("jdeh"),
				},
				rows: 3,
				cols: 4,
				find: "abtgshedjca",
			},
			want: false,
		},
		{
			name: "case5",
			args: args{
				arr: [][]byte{
					[]byte("abtg"),
					[]byte("cfcs"),
					[]byte("jdeh"),
				},
				rows: 3,
				cols: 4,
				find: "q",
			},
			want: false,
		},
		{
			name: "case6",
			args: args{
				arr: [][]byte{
					[]byte("abtg"),
					[]byte("cfcs"),
					[]byte("jdeh"),
				},
				rows: 3,
				cols: 4,
				find: "g",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(tt.args.arr, tt.args.find); got != tt.want {
				t.Errorf("FindPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
