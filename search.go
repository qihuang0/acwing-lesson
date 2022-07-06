package main

// 二分 有整数二分和小数二分

// 整数二分 两个模板 一个是区间[l,r]被划分为[l, mid]和[mid+1, r]时使用, 另一个是被划分为[l, mid-1]和[mid, r]时使用
func search_1(l int, r int) {
	for l < r {
		mid := l + r>>1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
}

func check(v int) bool {
	// check用于判断mid是否满足某种性质
	return false
}

func search_2(l int, r int) {
	for l < r {
		mid := l + r + 1>>1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
}
