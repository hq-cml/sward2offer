package _devide_group

import (
	"testing"
)

func TestDevideGroup(t *testing.T) {
	type args struct {
		groupList [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		//{
		//	name: "case1--注定无结果，因为必须同时存在2人组和3人组",
		//	args: args{
		//		groupList: [][]string{
		//			{"1", "2"},
		//			{"A", "B"},
		//			{"张", "王"},
		//		},
		//	},
		//	want:    nil,
		//	wantErr: false,
		//},
		{
			name: "case2",
			args: args{
				groupList: [][]string{
					{"少华", "少平", "少军", "少安", "少康"},
					{"福军", "福堂", "福民", "福平", "福心"},
					{"小明", "小红", "小花", "小丽", "小强"},
					{"大壮", "大力", "大1", "大2", "大3"},
					{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
					{"A", "B", "C", "D", "E"},
					{"一", "二", "三", "四", "五"},
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DevideGroup(tt.args.groupList)
			_ = err
			t.Log(got)
		})
	}
}
