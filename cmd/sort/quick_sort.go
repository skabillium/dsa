package sort

func partition(arr []int, lo int, hi int) int {
	idx := lo - 1
	pivot := arr[hi]

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++

			tmp := arr[idx]
			arr[idx] = arr[i]
			arr[i] = tmp
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

func qs(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)
	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func QuickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}
