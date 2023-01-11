package priorityqueue

import (
	"container/heap"
	"fmt"
)

func ExampleF_anyQueue_pop() {
	queueAny := new(PriorityaAnyQueue)
	heap.Init(queueAny)

	heap.Push(queueAny, &ItemAny{value: map[string]int{"aaa": 1}, priority: 3})
	heap.Push(queueAny, &ItemAny{value: "B", priority: 6})
	heap.Push(queueAny, &ItemAny{value: []string{"a", "b", "c"}, priority: 1})
	heap.Push(queueAny, &ItemAny{value: 12, priority: 2})

	for queueAny.Len() > 0 {
		fmt.Printf("%v ", heap.Pop(queueAny).(*ItemAny).value)
	}
	// Output:
	// B map[aaa:1] 12 [a b c]
}
