package main

/*
1、确定分界点，q[l],q[(l+r)/2],q[r],随机
2、调整范围，所有<= x的在左半边，所有>= 的在右半边
3、递归处理左右两段
*/
func quickSort(q []int, l int, r int) {
	if l >= r {
		return
	}
	i := l
	j := r
	x := q[(i+j)/2]
	for i < j {
		for q[i] < x {
			i++
		}
		for q[j] > x {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

var temp []int

func mergeSort(q []int, l int, r int) {
	if l >= r {
		return
	}
	mid := l + r>>2
	mergeSort(q, l, mid)
	mergeSort(q, mid+1, r)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] < q[j] {
			temp[k] = q[i]
			k++
			i++
		} else {
			temp[k] = q[j]
			k++
			j++
		}
	}
	for i <= mid {
		temp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		temp[k] = q[j]
		k++
		j++
	}
	for i, j = 0, 0; i <= r; {
		q[i] = temp[j]
		i++
		j++
	}

}
