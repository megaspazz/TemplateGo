/**
 * A simple queue.
 * Replace T with the actual type for the queue.
 */

type Queue struct {
	index int
	array []T
}

func NewQueue() *Queue {
	return &Queue{
		index: 0,
		array: make([]T, 0)
	}
}

func (q *Queue) Enqueue(x T) {
	q.array = append(q.array, x)
}

func (q *Queue) Dequeue() T {
	x := q.array[q.index]
	q.index++
	return x
}

func (q *Queue) Count() int {
	return len(q.array) - q.index
}

func (q *Queue) Empty() bool {
	return q.Count() <= 0
}
