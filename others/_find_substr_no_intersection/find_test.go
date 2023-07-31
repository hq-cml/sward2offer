package _find_substr_no_intersection

import (
    "reflect"
    "testing"
)

func TestFindSubStr1(t *testing.T) {
    type args struct {
        str string
    }
    tests := []struct {
        name string
        args args
        want []string
    }{
        {
            name: "case1",
            args: args{
                str: "abc",
            },
            want: []string{"abc"},
        },
        {
            name: "case2",
            args: args{
                str: "arrar",
            },
            want: []string{"arrar"},
        },
        {
            name: "case4",
            args: args{
                str: "arrabc",
            },
            want: []string{"arra", "bc"},
        },
        {
            name: "case3",
            args: args{
                str: "arrarbc",
            },
            want: []string{"arrar", "bc"},
        },
        {
            name: "case5",
            args: args{
                str: "arrarbcdkkdkef",
            },
            want: []string{"arrar", "bc", "dkkdk", "ef"},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := FindSubStr(tt.args.str); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("FindSubStr1() = %v, want %v", got, tt.want)
            }
        })
    }
}