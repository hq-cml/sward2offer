package _39_find_majority_number

import "testing"

func TestFindMajority1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		//{
		//	name: "case1",
		//	args: args{
		//		arr: []int{3, 2, 5, 6, 7, 2, 2, 2, 2},
		//	},
		//	want:    2,
		//	wantErr: false,
		//},
		//{
		//	name: "case2",
		//	args: args{
		//		arr: []int{3, 2, 5, 6, 7, 2, 2, 2, 2, 2},
		//	},
		//	want:    2,
		//	wantErr: false,
		//},
		//{
		//	name: "case3",
		//	args: args{
		//		arr: []int{2, 2, 2, 2, 2, 2},
		//	},
		//	want:    2,
		//	wantErr: false,
		//},
		//{
		//	name: "case4",
		//	args: args{
		//		arr: []int{2},
		//	},
		//	want:    2,
		//	wantErr: false,
		//},
		//{
		//	name: "case5",
		//	args: args{
		//		arr: []int{2, 1, 1},
		//	},
		//	want:    1,
		//	wantErr: false,
		//},
		{
			name: "case6",
			args: args{
				arr: []int{1, 2, 3, 2, 2, 2, 5, 4, 2},
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMajority1(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMajority1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindMajority1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMajority2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "case0",
			args: args{
				arr: []int{1, 2, 3, 2, 2, 2, 5, 4, 2},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "case1",
			args: args{
				arr: []int{3, 2, 5, 6, 7, 2, 2, 2, 2},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				arr: []int{3, 2, 5, 6, 7, 2, 2, 2, 2, 2},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "case3",
			args: args{
				arr: []int{2, 2, 2, 2, 2, 2},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "case4",
			args: args{
				arr: []int{2},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "case5",
			args: args{
				arr: []int{2, 1, 1},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMajority2(tt.args.arr)
			if got != tt.want {
				t.Errorf("FindMajority2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
