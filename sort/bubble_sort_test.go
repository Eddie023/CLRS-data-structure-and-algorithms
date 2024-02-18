package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name   string
		inputs []int
		want   []int
	}{
		{
			name:   "ok",
			inputs: []int{10, 2, 3, 1, 5, 7, 6},
			want:   []int{1, 2, 3, 5, 6, 7, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func GetRandomInputs(size int) []int {
	testList := make([]int, size)
	for i := 0; i < size; i++ {
		testList[i] = rand.Intn(size)
	}

	return testList
}

func BenchmarkBubbleSort(b *testing.B) {
	inputSize := []int{10, 100, 1000, 10000, 100000}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			testList := GetRandomInputs(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				BubbleSort(testList)
			}
		})
	}
}
