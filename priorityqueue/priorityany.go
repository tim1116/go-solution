package priorityqueue

type ItemAny struct {
	value    interface{}
	priority int
}

type PriorityaAnyQueue []*ItemAny

func (pq PriorityaAnyQueue) Len() int { return len(pq) }
func (pq PriorityaAnyQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityaAnyQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityaAnyQueue) Push(x interface{}) {
	item := x.(*ItemAny)
	*pq = append(*pq, item)
}

func (pq *PriorityaAnyQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
