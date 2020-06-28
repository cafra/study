package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: struct{ values []int }{values: []int{4, 2, 3}},
		},
		{
			name: "test2",
			args: struct{ values []int }{values: []int{4, 2, 3, 7, 8, 1, 5, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.values)
			fmt.Print(tt.args.values)
		})
	}
}

func TestMyTest(t *testing.T) {
	a, b := 1, 2

	a, b = b, a

	t.Log(a, b)
}

// 参考文章 https://blog.csdn.net/adusts/article/details/80882649
func TestMyQSort(t *testing.T) {
	type args struct {
		data  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

		{
			name: "自己写的快排1",
			args: struct {
				data  []int
				left  int
				right int
			}{data: []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}, left: 0, right: 9},
		},
		{
			name: "自己写的快排临界值2",
			args: struct {
				data  []int
				left  int
				right int
			}{data: []int{6, 1}, left: 0, right: 1},
		},
		{
			name: "自己写的快排临界值3",
			args: struct {
				data  []int
				left  int
				right int
			}{data: []int{1, 6}, left: 0, right: 1},
		},
		{
			name: "自己写的快排全部一样",
			args: struct {
				data  []int
				left  int
				right int
			}{data: []int{1, 1, 1, 1, 1, 1}, left: 0, right: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MyQSort(tt.args.data, tt.args.left, tt.args.right)
			t.Log(tt.args.data)
		})
	}
}
