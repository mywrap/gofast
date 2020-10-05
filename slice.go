package gofast

// IndexInts returns the index of the first instance of sub in slice,
// or -1 if sub is not present in slice
func IndexInts(slice []int, sub int) int {
	for i, v := range slice {
		if v == sub {
			return i
		}
	}
	return -1
}

// IndexStrings returns the index of the first instance of sub in slice,
// or -1 if sub is not present in slice
func IndexStrings(slice []string, sub string) int {
	for i, v := range slice {
		if v == sub {
			return i
		}
	}
	return -1
}

// SplitToBatches splits a long slice into smaller slices with a limited size
func SplitToBatches(slice []string, batchSize int) [][]string {
	ret := make([][]string, 0)
	batch := make([]string, 0)
	n := len(slice)
	for i := 0; i < n; i++ {
		batch = append(batch, slice[i])
		if len(batch) >= batchSize || i == n-1 {
			ret = append(ret, batch)
			batch = make([]string, 0)
		}
	}
	return ret
}
