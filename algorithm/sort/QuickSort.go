package sort

import "fmt"

// 分治和递归

// 第一种写法
func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		//
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	quickSort(values, 0, len(values)-1)
}

// 第二种写法
func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		fmt.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	Quick2Sort(values[:head])
	Quick2Sort(values[head+1:])
}

// 第三种写法
func Quick3Sort(a []int, left int, right int) {

	if left >= right {
		return
	}

	explodeIndex := left

	for i := left + 1; i <= right; i++ {

		if a[left] >= a[i] {

			//分割位定位++
			explodeIndex++
			a[i], a[explodeIndex] = a[explodeIndex], a[i]

		}

	}

	//起始位和分割位
	a[left], a[explodeIndex] = a[explodeIndex], a[left]

	Quick3Sort(a, left, explodeIndex-1)
	Quick3Sort(a, explodeIndex+1, right)

}

func MyQSort(data []int, left, right int) {
	if left >= right {
		return
	}
	tmp := data[left] // 基准数
	i, j := left, right

	for i < j {
		// 找到右侧小值
		for data[j] >= tmp && j > i {
			j--
		}
		// 找到左侧大值
		for data[i] <= tmp && j > i {
			i++
		}
		// 交换
		if i < j {
			//fmt.Println("交换", data[i], data[j], i, j)
			data[i], data[j] = data[j], data[i]
		}
	}
	// 移动基准数
	data[left] = data[i]
	data[i] = tmp
	// 递归左侧和右侧
	MyQSort(data, left, i-1)
	MyQSort(data, i+1, right)
}
