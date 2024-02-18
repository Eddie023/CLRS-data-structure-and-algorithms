package sort

func InsertionSort(inputs []int) []int {
	for i := 1; i < len(inputs); i++ {
		key := inputs[i]
		j := i - 1
		for j >= 0 && inputs[j] > key {
			inputs[j+1] = inputs[j]
			j--
		}
		inputs[j+1] = key
	}
	return inputs
}
