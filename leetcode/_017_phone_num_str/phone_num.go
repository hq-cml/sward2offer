/*
 * 将电话号码转化成为可能的字符串
 * 例如，输入“2,3"
 * 输出：[ad, ae, af, bd, be, bf, cd, ce, cf]
 */
package _017_phone_num_str

import (
	"errors"
	"log"
)

var Map = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	// ...
}

// 递归 + 回溯
func PhoneNumStr(phoneNum string) []string {
	var ret []string
	err := calc(phoneNum, "", &ret)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return ret
}

func calc(phoneNum string, curr string, ret *[]string) error {
	if phoneNum == "" {
		*ret = append(*ret, curr)
		return nil
	}
	c := phoneNum[0]
	list, ok := Map[c]
	if !ok {
		return errors.New("invalid char")
	}
	for _, c := range list {
		err := calc(phoneNum[1:], curr+string(c), ret)
		if err != nil {
			return err
		}
	}
	return nil
}
