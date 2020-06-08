package algorithm

import (
	"reflect"
	"testing"
)

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

func TestMerge(t *testing.T) {
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
		MergeSort(helpArr, 0, len(helpArr)-1)
		if !reflect.DeepEqual(helpArr, ts.want) {
			t.Errorf("merge sort sample %v: want %v, but got %v", index, ts.want, helpArr)
		}
	}
}
