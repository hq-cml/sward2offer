package basic

// ¶ş·Ö²éÕÒ
func BinarySearch(arr []int, need int) int {
    if len(arr) == 0 {
        return -1
    }

    beg := 0
    end := len(arr)-1

    for beg <= end {
        mid := (end-beg)/2 + beg
        if arr[mid] == need {
            return mid
        } else if arr[mid] > need {
            end = mid - 1
        } else {
            beg = mid + 1
        }
    }

    return -1
}