package _19_simple_regexp

func Regexp(str, pattern string) bool {
    if len(str) == 0 || len(pattern) == 0 {
        return false
    }

    return recurseReg([]byte(str), []byte(pattern))
}

func recurseReg(str, pattern []byte) bool {
    if len(str) == 0 && len(pattern) == 0 {
        return true
    }

    if len(str) != 0 || len(pattern) != 0 {
        return false
    }

    if len(pattern) > 1 &&
}