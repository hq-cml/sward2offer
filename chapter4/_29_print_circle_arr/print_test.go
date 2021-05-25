package _29_print_circle_arr

import "testing"

func TestPrint(t *testing.T) {
	type args struct {
		arr  []int
		rows int
		cols int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{
				arr:  []int{1,2,3,4,5,6,7,8,9},
				rows: 3,
				cols: 3,
			},
			wantErr: false,
		},
		{
			name:    "case2",
			args:    args{
				arr:  []int{1,2,3,4,5,6},
				rows: 2,
				cols: 3,
			},
			wantErr: false,
		},
		{
			name:    "case3",
			args:    args{
				arr:  []int{1},
				rows: 1,
				cols: 1,
			},
			wantErr: false,
		},
		{
			name:    "case4",
			args:    args{
				arr:  []int{1,2,3,4,5,6},
				rows: 3,
				cols: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Print(tt.args.arr, tt.args.rows, tt.args.cols); (err != nil) != tt.wantErr {
				t.Errorf("Print() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}