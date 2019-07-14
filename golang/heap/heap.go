package heap

type Heap struct {
	Data     []int
	Count    int
	Capacity int
}

func NewHeap(capacity int) *Heap {
	data := make([]int, capacity+1, capacity+1)
	return &Heap{
		Data: data,
		// care place 1 the data struct
		Count:    0,
		Capacity: capacity + 1,
	}
}

func (heap *Heap) BigTopInsert(value int) {
	// care place 2
	// count should add 1 first  count is lower 1 than capacity
	heap.Count += 1
	// if this heap is full
	if heap.Count >= heap.Capacity {
		return
	}

	index := heap.Count
	heap.Data[index] = value

	for {
		// now can do the campare
		if (index)/2 > 0 && heap.Data[index] > heap.Data[index/2] {
			heap.Data[index], heap.Data[index/2] = heap.Data[index/2], heap.Data[index]
			index = index / 2
		} else {
			break
		}
	}
}

func (heap *Heap) SmallTopInsert(value int) {
	heap.Count += 1
	if heap.Count >= heap.Capacity {
		return
	}
	heap.Data[heap.Count] = value
	index := heap.Count
	for {
		if index/2 > 0 && heap.Data[index/2] > heap.Data[index] {
			heap.Data[index], heap.Data[index/2] = heap.Data[index/2], heap.Data[index]
			index = index / 2
		} else {
			break
		}
	}
}

func (heap *Heap) RemoveBigTop() {
	if heap.Count <= 1 {
		return
	}
	heap.Data[1] = heap.Data[heap.Count]
	heap.Data[heap.Count] = 0
	heap.Count -= 1

	index := 1
	maxposi := 1
	for {
		if index*2 <= heap.Count && heap.Data[index*2] > heap.Data[index] {
			maxposi = index * 2
		}
		if index*2+1 <= heap.Count && heap.Data[index*2+1] > heap.Data[maxposi] {
			maxposi = index*2 + 1
		}
		heap.Data[index], heap.Data[maxposi] = heap.Data[maxposi], heap.Data[index]
		index = maxposi
		if index*2 > heap.Count {
			break
		}
	}
	return
}

func (heap *Heap) Sort() {
	// Build the heap to make sure the top is the biggest element
	for index := heap.Count / 2; index > 0; index-- {
		// index should token out of this function
		heapindex := index
		// make the top the biggest value
		heap.heapify(heapindex)
	}
	for i := 1; i < heap.Count; {
		heap.heapsort()
	}
}

func (heap *Heap) heapify(heapindex int) {
	for {
		maxposi := 0
		// 这里这个条件很精妙
		if heapindex*2 <= heap.Count && heap.Data[2*heapindex] >= heap.Data[heapindex] {
			maxposi = heapindex * 2
		}
		// I lose the condition 2*index+1 should be lower than count
		if 2*heapindex+1 <= heap.Count && heap.Data[2*heapindex+1] > heap.Data[maxposi] {
			maxposi = 2*heapindex + 1
		}

		if maxposi == 0 {
			break
		}
		heap.Data[maxposi], heap.Data[heapindex] = heap.Data[heapindex], heap.Data[maxposi]
		heapindex = maxposi
	}

}

func (heap *Heap) heapsort() {
	if heap.Count <= 1 {
		return
	}
	heap.Data[1], heap.Data[heap.Count] = heap.Data[heap.Count], heap.Data[1]
	heap.Count -= 1

	index := 1
	maxposi := 1
	for {
		if index*2 <= heap.Count && heap.Data[index*2] > heap.Data[index] {
			maxposi = index * 2
		}
		if index*2+1 <= heap.Count && heap.Data[index*2+1] > heap.Data[maxposi] {
			maxposi = index*2 + 1
		}
		heap.Data[index], heap.Data[maxposi] = heap.Data[maxposi], heap.Data[index]
		index = maxposi
		if index*2 > heap.Count {
			break
		}
	}
	return

}
