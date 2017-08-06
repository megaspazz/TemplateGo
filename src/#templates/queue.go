/**
 * A simple queue backed by a slice that resizes whenever it gets full.
 * The queue is defined by a number of elements starting at a given index.
 * The elements in the queue can wrap around from the end of the slice to the beginning.
 * Replace T with the actual type for the queue (6 occurrences).
 */

type Queue struct {
	index int
	count int
	array []T
}

func NewQueue() *Queue {
	return &Queue{
		index: 0,
		count: 0,
		array: make([]T, 4),
	}
}

func (q *Queue) Enqueue(x T) {
	if q.count == len(q.array) {
		q.Resize(len(q.array) << 1)
	}
	idx := (q.index + q.count) % len(q.array)
	q.array[idx] = x
	q.count++
}

func (q *Queue) Dequeue() (T, bool) {
	var x T
	if q.count == 0 {
		return x, false
	}
	x = q.array[q.index]
	q.index = (q.index + 1) % len(q.array)
	q.count--
	return x, true
}

func (q *Queue) Count() int {
	return q.count
}

func (q *Queue) Empty() bool {
	return q.Count() <= 0
}

func (q *Queue) Resize(newCapacity int) bool {
	if newCapacity < len(q.array) {
		return false
	}
	newArray := make([]T, newCapacity)
	for i := 0; i < q.count; i++ {
		idx := (i + q.index) % len(q.array)
		newArray[i] = q.array[idx]
	}
	q.index = 0
	q.array = newArray
	return true
}
