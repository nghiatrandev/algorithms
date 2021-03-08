package selection_sort

func SelectionSort(ints []int) []int {
	for i := 0; i < len(ints); i++ {
		minI, minV := i, ints[i]
		for j := i + 1; j < len(ints); j++ {
			if minV > ints[j] {
				minI = j
				minV = ints[j]
			}
		}

		if minV != ints[i] {
			ints[minI], ints[i] = ints[i], ints[minI]
		}
	}
	return ints
}
