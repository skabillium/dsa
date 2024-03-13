package structures

type ArrayList struct {
	Length   int
	Capacity int
	arr      []any
}

func NewArrayList(capacity int) *ArrayList {
	return &ArrayList{Length: 0, Capacity: capacity, arr: make([]any, capacity)}
}

func (list *ArrayList) grow() {
	capacity := 2 * list.Capacity
	arr := make([]any, capacity)

	copy(arr[:list.Capacity], list.arr[:])
	list.arr = arr
}

func (list *ArrayList) Prepend(item any) {
	list.InsertAt(item, 0)
}

func (list *ArrayList) InsertAt(item any, idx int) {
	if list.Length+1 == list.Capacity {
		list.grow()
	}

	if idx < 0 || idx > list.Capacity {
		return
	}

	for i := list.Length; i > idx; i-- {
		list.arr[i] = list.arr[i-1]
	}

	list.arr[idx] = item
	list.Length++
}

func (list *ArrayList) Append(item any) {
	if list.Length+1 == list.Capacity {
		list.grow()
	}

	list.arr[list.Length] = item
	list.Length++
}

func (list *ArrayList) Remove(item any) any {
	for i := 0; i < list.Length; i++ {
		if list.arr[i] == item {
			it := list.arr[i]
			for j := i + 1; j < list.Length; j++ {
				list.arr[j-1] = list.arr[j]
			}
			list.Length--
			return it
		}
	}

	return nil
}

func (list *ArrayList) RemoveAt(idx int) any {
	if idx < 0 || idx > list.Length-1 {
		return nil
	}

	item := list.arr[idx]

	for i := idx + 1; i < list.Length; i++ {
		list.arr[i-1] = list.arr[i]
	}

	list.Length--
	return item
}

func (list *ArrayList) Get(idx int) any {
	if idx > list.Length-1 {
		return nil
	}

	return list.arr[idx]
}
