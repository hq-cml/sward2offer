package _05_replace_blank

// blank => %20
func ReplaceBlank(s []byte) ([]byte, error) {
    if s == nil {
        return s, nil
    }

    cnt := 0
    for _, v := range s {
        if string(v) == " " {
            cnt ++
        }
    }
    if cnt == 0 {
        return s, nil
    }

    add := make([]byte, cnt * 2)
    p1 := len(s) - 1
    s = append(s, add ...) // 续上空白的长度
    p2 := len(s) - 1

    for p1 < p2 {
        if string(s[p1]) != " " {
            s[p2] = s[p1]
            p2 --
        } else {
            s[p2-2], s[p2-1], s[p2] = byte('%'), byte('2'), byte('0')
            p2 -= 3
        }
        p1 --
    }
    return s, nil
}