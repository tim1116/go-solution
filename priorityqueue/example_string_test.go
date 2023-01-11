package priorityqueue

import (
	"container/heap"
	"fmt"
)

var queueString *PriorityStringQueue

func init() {
	queueString = new(PriorityStringQueue)
	heap.Init(queueString)

	heap.Push(queueString, &ItemString{value: "A", priority: 3})
	heap.Push(queueString, &ItemString{value: "B", priority: 6})
	heap.Push(queueString, &ItemString{value: "C", priority: 1})
	heap.Push(queueString, &ItemString{value: "D", priority: 2})

}

func ExampleF_stringQueue_len() {
	len := queueString.Len()
	fmt.Printf("%d", len)
	// Output:
	// 4
}

func ExampleF_stringQueue_top() {
	item := (*queueString)[0]
	fmt.Printf("%s:%d", item.value, item.priority)
	// Output:
	// B:6
}

func ExampleF_stringQueue_pop() {
	for queueString.Len() > 0 {
		fmt.Printf("%s ", heap.Pop(queueString).(*ItemString).value)
	}
	// Output:
	// B A D C
}
