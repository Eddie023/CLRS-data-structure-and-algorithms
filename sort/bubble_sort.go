package sort

func BubbleSort(inputs []int) []int {
	// out := make([]int, 0, len(inputs))

	for i := 0; i < len(inputs)-1; i++ {
		for j := i + 1; j < len(inputs); j++ {
			if inputs[j] < inputs[i] {
				tmp := inputs[j]
				inputs[j] = inputs[i]
				inputs[i] = tmp
			}
		}

	}

	return inputs
}
