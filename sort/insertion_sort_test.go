package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {

	tests := []struct {
		name  string
		given []int
		want  []int
	}{
		{
			name:  "ok",
			given: []int{10, 5, 2, 3, 7},
			want:  []int{2, 3, 5, 7, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.given); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	inputSize := []int{10, 100, 1000, 10000, 100000}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			testList := GetRandomInputs(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				InsertionSort(testList)
			}
		})
	}
}
