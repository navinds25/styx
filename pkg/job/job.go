package main

import (
	"container/list"
	"log"
)

/*
Channel + Linked List
Channel for queue for starting tasks
Linked List for letting the user know the status of the tasks

funcs:
Add to Queue, Add to Channel + Lock
Check Queue, List all items in the queue + Lock
Remove from Queue, Remove item from the top of the queue once it is complete + Lock
Initialize Queue
Clear all Jobs -Reinitialize queue
*/

// Job is the struct for implementing a job
type Job struct {
	ID string
}

func main() {
	j := Job{ID: "blah"}
	e := Job{ID: "so blah"}
	queue := list.New()
	queue.PushBack(j)
	queue.PushBack(e)
	log.Println(queue.Len())
	elem := queue.Front()
	for i := 0; i < queue.Len(); i++ {
		log.Println(elem.Value)
		elem = elem.Next()
	}
}
