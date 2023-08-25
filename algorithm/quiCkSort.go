package algorithm

func QuickSorts(src []int, first, last int) {
	flag := first
	left := first
	right := last
	if first >= last {
		return
	}
	for first < last {
		for first < last {
			if src[last] >= src[flag] {
				last -= 1
				continue
			} else {
				tmp := src[last]
				src[last] = src[flag]
				src[flag] = tmp
				flag = last
				break
			}
		}
		for first < last {
			if src[first] <= src[flag] {
				first += 1
				continue
			} else {
				tmp := src[first]
				src[first] = src[flag]
				src[flag] = tmp
				flag = first
				break
			}
		}
	}
	QuickSorts(src, left, flag-1)
	QuickSorts(src, flag+1, right)
}

//func quickSort(arr []int,left,right int) {
//	key := arr[left]
//	p := key
//	i,j := left,right
//	for i <= j {
//
//	}
//}
