package search

func BinarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack)

	for low < high {
		mfloat := float64(low + (high-low)/2)
		m := int(mfloat)
		val := haystack[m]

		if val == needle {
			return true
		} else if val < needle {
			low = m + 1
		} else {
			high = m
		}
	}
	return false
}
