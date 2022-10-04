package main

import "fmt"

// 前缀和作用 用来快速计算某段区间的和
// 技巧 数组下标从1到n 数组和也从1到n 方便处理边界值  s(l,r) = s(r) -s(l-1), 用在二维数组的时候，求某个范围内数组和，比如S[(2,2)->(4,4)],
// 就会等于S[4,4] - S[4,1] - S[1,4] + s[1,1] (数组从1，1起)  所以公式为S[x2,y2]- S[x2, (y1)-1] - S[(x1)-1,y2] + S[x1-1, y1-1]

// 题目 给定一个长度为n的整数序列，然后输入m的询问，每个询问输入一对l,r,求每一对l,r的整数区间和



func main() {
	var input []int
	for {
		var v int
		_, err := fmt.Scanf("%d", &v)
		if err != nil {
			break
		}
		input = append(input, v)
		fmt.Println(v)
	}

}
