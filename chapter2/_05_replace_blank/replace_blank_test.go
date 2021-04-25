package _05_replace_blank

import (
    "reflect"
    "testing"
)

func TestReplaceBlank(t *testing.T) {
    type args struct {
        s []byte
    }
    tests := []struct {
        name    string
        args    args
        want    []byte
        wantErr bool
    }{
        {
            name:    "case1",
            args:    args{
                s: []byte("we are happy."),
            },
            want:    []byte("we%20are%20happy."),
            wantErr: false,
        },
        {
            name:    "case2",
            args:    args{
                s: []byte("thereisnoblank."),
            },
            want:    []byte("thereisnoblank."),
            wantErr: false,
        },
        {
            name:    "case3",
            args:    args{
                s: []byte(" "),
            },
            want:    []byte("%20"),
            wantErr: false,
        },
        {
            name:    "case4",
            args:    args{
                s: []byte(" abc"),
            },
            want:    []byte("%20abc"),
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ReplaceBlank(tt.args.s)
            if (err != nil) != tt.wantErr {
                t.Errorf("ReplaceBlank() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ReplaceBlank() got = %v, want %v", got, tt.want)
            }
        })
    }
}