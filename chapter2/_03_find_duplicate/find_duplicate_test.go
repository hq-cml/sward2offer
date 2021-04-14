package _03_find_duplicate

import "testing"

func TestFindDuplicate(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name      string
		args      args
		wantD     int
		wantExist bool
	}{
		{
			name:      "case1",
			args:      args{
				s: []int{0,1,2,3,4,5,6,7},
			},
			wantD:     -1,
			wantExist: false,
		},
		{
			name:      "case2",
			args:      args{
				s: []int{1,2,0,3,4,2,6,3},
			},
			wantD:     2,
			wantExist: true,
		},
		{
			name:      "case3",
			args:      args{
				s: []int{2,3,1,0,2,5,3},
			},
			wantD:     2,
			wantExist: true,
		},
		{
			name:      "case4",
			args:      args{
				s: []int{2,1,1,3,4},
			},
			wantD:     1,
			wantExist: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, gotExist := FindDuplicate1(tt.args.s)
			if gotD != tt.wantD {
				t.Errorf("FindDuplicate() gotD = %v, want %v", gotD, tt.wantD)
			}
			if gotExist != tt.wantExist {
				t.Errorf("FindDuplicate() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}