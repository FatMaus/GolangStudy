package algorithm

import (
	"reflect"
	"testing"
)

// TestBubble 冒泡排序测试
func TestBubble(t *testing.T) {
	type testSample []struct {
		input []int
		want  []int
	}
	tests := testSample{
		{input: []int{85, 6, 55, 54, 87, 32, 92, 17, 12, 19}, want: []int{6, 12, 17, 19, 32, 54, 55, 85, 87, 92}},
		{input: []int{83, 39, 56, 50, 22, 58, 53, 38, 88, 39}, want: []int{22, 38, 39, 39, 50, 53, 56, 58, 83, 88}},
		{input: []int{65, 75, 99, 11, 65, 33, 34, 24, 56, 29}, want: []int{11, 24, 29, 33, 34, 56, 65, 65, 75, 99}},
	}
	for index, ts := range tests {
		var helpArr = make([]int, len(ts.input), len(ts.input))
		copy(helpArr, ts.input)
		BubbleSort(helpArr)
		if !reflect.DeepEqual(helpArr, ts.want) {
			t.Errorf("No %v: want %v, but got %v", index, ts.want, helpArr)
		}
	}
}

// TestInsert 插入排序测试
func TestInsert(t *testing.T) {
	type testSample struct {
		input []int
		want  []int
	}
	tests := []testSample{
		testSample{input: []int{85, 6, 55, 54, 87, 32, 92, 17, 12, 19}, want: []int{6, 12, 17, 19, 32, 54, 55, 85, 87, 92}},
		testSample{input: []int{83, 39, 56, 50, 22, 58, 53, 38, 88, 39}, want: []int{22, 38, 39, 39, 50, 53, 56, 58, 83, 88}},
		testSample{input: []int{65, 75, 99, 11, 65, 33, 34, 24, 56, 29}, want: []int{11, 24, 29, 33, 34, 56, 65, 65, 75, 99}},
	}
	for index, ts := range tests {
		var helpArr = make([]int, len(ts.input), len(ts.input))
		copy(helpArr, ts.input)
		InsertSort(helpArr)
		if !reflect.DeepEqual(helpArr, ts.want) {
			t.Errorf("insert sort sample %v: want %v, but got %v", index, ts.want, helpArr)
		}
	}
}

// TestQuick 快排测试
func TestQuick(t *testing.T) {
	type testSample struct {
		input []int
		want  []int
	}
	tests := []testSample{
		testSample{input: []int{85, 6, 55, 54, 87, 32, 92, 17, 12, 19}, want: []int{6, 12, 17, 19, 32, 54, 55, 85, 87, 92}},
		testSample{input: []int{83, 39, 56, 50, 22, 58, 53, 38, 88, 39}, want: []int{22, 38, 39, 39, 50, 53, 56, 58, 83, 88}},
		testSample{input: []int{65, 75, 99, 11, 65, 33, 34, 24, 56, 29}, want: []int{11, 24, 29, 33, 34, 56, 65, 65, 75, 99}},
	}
	for index, ts := range tests {
		var helpArr = make([]int, len(ts.input), len(ts.input))
		copy(helpArr, ts.input)
		QuickSort(helpArr, 0, len(helpArr)-1)
		if !reflect.DeepEqual(helpArr, ts.want) {
			t.Errorf("quick sort sample %v: want %v, but got %v", index, ts.want, helpArr)
		}
	}
}

// TestMerge 并归排序测试
func TestMerge(t *testing.T) {
	type testSample struct {
		input []int
		want  []int
	}
	tests := map[string]testSample{
		"merge_sample_1": {input: []int{85, 6, 55, 54, 87, 32, 92, 17, 12, 19}, want: []int{6, 12, 17, 19, 32, 54, 55, 85, 87, 92}},
		"merge_sample_2": {input: []int{83, 39, 56, 50, 22, 58, 53, 38, 88, 39}, want: []int{22, 38, 39, 39, 50, 53, 56, 58, 83, 88}},
		"merge_sample_3": {input: []int{65, 75, 99, 11, 65, 33, 34, 24, 56, 29}, want: []int{11, 24, 29, 33, 34, 56, 65, 65, 75, 99}},
	}
	for name, ts := range tests {
		var helpArr = make([]int, len(ts.input), len(ts.input))
		copy(helpArr, ts.input)
		t.Run(name, func(t *testing.T) {
			MergeSort(helpArr, 0, len(helpArr)-1)
			if !reflect.DeepEqual(helpArr, ts.want) {
				t.Errorf("want %v, but got %v", ts.want, helpArr)
			}
		})
	}
}

// TestRotateArray 测试旋转数组最小值
func TestRotateArray(t *testing.T) {
	type testSample struct {
		input []int
		want  int
	}
	tests := map[string]testSample{
		"rotate1": {input: []int{5, 6, 7, 8, 9, 1, 2, 3, 4}, want: 1},
		"rotate2": {input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, want: 0},
		"rotate3": {input: []int{8, 9, 1, 2, 3, 4, 5, 6, 7}, want: 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) {
			get := RotateArray(ts.input)
			if !reflect.DeepEqual(get, ts.want) {
				t.Errorf("want %v, but got %v", ts.want, get)
			}
		})
	}
}
