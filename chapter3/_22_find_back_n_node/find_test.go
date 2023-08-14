package _22_find_back_n_node

import (
	"github.com/hq-cml/sward2offer/common"
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantV   int
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				n: 1,
			},
			wantV:   8,
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				n: 4,
			},
			wantV:   2,
			wantErr: false,
		},
		{
			name: "case3",
			args: args{
				n: 5,
			},
			wantV:   0,
			wantErr: true,
		},
		{
			name: "case4",
			args: args{
				n: 6,
			},
			wantV:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l *common.ListNode
			l = l.PushNode(2)
			l = l.PushNode(4)
			l = l.PushNode(5)
			l = l.PushNode(8)

			gotV, err := FindBackN(l, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotV != tt.wantV {
				t.Errorf("Find() gotV = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}