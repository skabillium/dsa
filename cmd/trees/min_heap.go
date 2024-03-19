package trees

type MinHeap struct {
	Length int
	data   []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{Length: 0, data: []int{}}
}

func (h *MinHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(h.Length)
	h.Length++
}

func (h *MinHeap) Delete() int {
	if h.Length == 0 {
		return -1
	}

	out := h.data[0]
	h.Length--
	if h.Length == 0 {
		h.data = []int{}
		return out
	}

	h.data[0] = h.data[h.Length]
	h.heapifyDown(0)

	return out
}

func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := h.parent(idx)
	parentValue := h.data[p]
	value := h.data[idx]
	if parentValue > value {
		h.data[p] = value
		h.data[idx] = parentValue
		h.heapifyUp(p)
	}
}

func (h *MinHeap) heapifyDown(idx int) {
	if idx >= h.Length {
		return
	}

	lIdx := h.leftChild(idx)
	rIdx := h.rightChild(idx)
	if idx >= h.Length || lIdx >= h.Length {
		return
	}

	value := h.data[idx]
	leftValue := h.data[lIdx]
	rightValue := h.data[rIdx]

	if leftValue > rightValue && value > rightValue {
		h.data[idx] = rightValue
		h.data[rIdx] = value
		h.heapifyDown(rIdx)
	} else if rightValue > leftValue && value > leftValue {
		h.data[idx] = leftValue
		h.data[lIdx] = value
		h.heapifyDown(lIdx)
	}
}

func (h *MinHeap) parent(idx int) int     { return (idx - 1) / 2 }
func (h *MinHeap) leftChild(idx int) int  { return (idx * 2) + 1 }
func (h *MinHeap) rightChild(idx int) int { return (idx * 2) + 2 }
