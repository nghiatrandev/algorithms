package bubble_sort

func BubboleSort(in []int) []int {
	count := 0
	for i := 0; i < len(in)-1; i++ {
		for j := 1; j < len(in); j++ {
			if in[j] < in[j-1] {
				in[j], in[j-1] = in[j-1], in[j]
				count++
			}
		}
	}

	return in
}
