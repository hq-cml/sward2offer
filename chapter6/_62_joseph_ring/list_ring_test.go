package _62_joseph_ring

import (
    "testing"
)

func TestCalclRing(t *testing.T) {
    type args struct {
        n int
        m int
    }
    tests := []struct {
        name    string
        args    args
        want    int
        wantErr bool
    }{
        {
            name:    "case1",
            args:    args{
                n: 5,
                m: 3,
            },
            want:    3,
            wantErr: false,
        },
        {
            name:    "case1",
            args:    args{
                n: 5,
                m: 1,
            },
            want:    4,
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := CalclRing(tt.args.n, tt.args.m)
            if (err != nil) != tt.wantErr {
                t.Errorf("CalclRing() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("CalclRing() got = %v, want %v", got, tt.want)
            }
        })
    }
}