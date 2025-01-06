package slices

func Sum(number []int) int {
	sum := 0
	for _, num := range number {
		sum += num
	}
	return sum
}
