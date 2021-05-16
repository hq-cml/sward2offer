package _16_pow

import "testing"

func TestPow1(t *testing.T) {
	type args struct {
		f    float64
		base int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{
				f: 3.3,
				base: 3,
			},
			want:    func(f float64, exp int) float64{
				r, _ := Pow2(f, exp)
				return r
			}(3.3, 3),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pow1(tt.args.f, tt.args.base)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pow1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pow1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_powRecurse(t *testing.T) {
	type args struct {
		f   float64
		exp int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case1",
			args: args{
				f:   4,
				exp: 2,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := powRecurse(tt.args.f, tt.args.exp); got != tt.want {
				t.Errorf("powRecurse() = %v, want %v", got, tt.want)
			}
		})
	}
}