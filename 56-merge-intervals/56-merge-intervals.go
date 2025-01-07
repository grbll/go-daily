package mergeintervals

func rise(heap [][]int) {
	var index int = len(heap) - 1
	for index > 0 && heap[(index-1)/2][0] > heap[index][0] {
		heap[index], heap[(index-1)/2] = heap[(index-1)/2], heap[index]
		index = (index - 1) / 2
	}
}

func sink(heap [][]int) {
	var n, index int = len(heap), 0
	for index*2+2 < n && (heap[index*2+1][0] < heap[index][0] || heap[index*2+2][0] < heap[index][0]) {
		if heap[index*2+1][0] <= heap[index*2+2][0] {
			heap[index], heap[index*2+1] = heap[index*2+1], heap[index]
			index = index*2 + 1
		} else {
			heap[index], heap[index*2+2] = heap[index*2+2], heap[index]
			index = index*2 + 2
		}
	}
	if index*2+1 < n && heap[index][0] > heap[index*2+1][0] {
		heap[index], heap[index*2+1] = heap[index*2+1], heap[index]
	}
}
func pop(heap *[][]int) []int {
	var pair []int = (*heap)[0]
	(*heap)[0], (*heap)[len((*heap))-1] = (*heap)[len((*heap))-1], (*heap)[0]
	*heap = (*heap)[:len(*heap)-1]
	sink(*heap)
	return pair
}

func merge(intervals [][]int) [][]int {
	var n int = len(intervals)
	for i := 0; i < n; i++ {
		intervals = intervals[:i+1]
		rise(intervals)
	}

	var output [][]int = make([][]int, 0, n)
	var pair []int = pop(&intervals)
	output = append(output, pair)
	for len(intervals) > 0 {
		pair = pop(&intervals)
		if pair[0] > output[len(output)-1][1] {
			output = append(output, pair)
		} else {
			output[len(output)-1][1] = max(pair[1], output[len(output)-1][1])
		}
	}
	return output
}
