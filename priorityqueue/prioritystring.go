package priorityqueue

type ItemString struct {
	value    string
	priority int
}

type PriorityStringQueue []*ItemString

func (pq PriorityStringQueue) Len() int { return len(pq) }
func (pq PriorityStringQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityStringQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityStringQueue) Push(x interface{}) {
	item := x.(*ItemString)
	*pq = append(*pq, item)
}

func (pq *PriorityStringQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
