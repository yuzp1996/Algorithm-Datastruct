package heap

func (heap *Heap) TopKInsert(value int) {
	heap.Count += 1

	if heap.Count >= heap.Capacity {
		heap.Count -= 1
		if value > heap.Data[1] {
			heap.Data[1] = value
		}
		//heapify
		for index := heap.Count / 2; index > 0; index-- {
			// index should token out of this function
			heapindex := index
			// make the top the biggest value
			heap.smallheapify(heapindex)
		}

		return
	} else {
		// build the smalltapheap
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

}

func (heap *Heap) GetTopKData() []int {
	return heap.Data[1:]
}
