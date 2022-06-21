package collection

import (
	"errors"
	"sync"
)

var (
	QueueEmptyError = errors.New("queue is empty")
)

type QueueInternal interface {
	Append(item interface{})
	Top() (interface{}, error)
	Pop() error
	Len() int
	Empty() bool
}

type ArrayQueueInternal struct {
	queue []interface{}
}

func NewArrayQueueInternal() *ArrayQueueInternal {
	qi := ArrayQueueInternal{
		queue: make([]interface{}, 0),
	}
	return &qi
}

func (qi *ArrayQueueInternal) Append(item interface{}) {
	qi.queue = append(qi.queue, item)
}

func (qi *ArrayQueueInternal) Top() (interface{}, error) {
	if len(qi.queue) == 0 {
		return nil, QueueEmptyError
	}
	return qi.queue[0], nil
}

func (qi *ArrayQueueInternal) Pop() error {
	if len(qi.queue) == 0 {
		return QueueEmptyError
	}
	qi.queue = qi.queue[1:]
	return nil
}

func (qi *ArrayQueueInternal) Len() int {
	return len(qi.queue)
}

func (qi *ArrayQueueInternal) Empty() bool {
	return len(qi.queue) == 0
}

type LinkedListQueueInternal struct {
	head *node
	tail *node
	size int
}

func NewLinkedListQueueInternal() *LinkedListQueueInternal {
	qi := LinkedListQueueInternal{
		head: nil,
		tail: nil,
		size: 0,
	}
	return &qi
}

func (qi *LinkedListQueueInternal) Append(item interface{}) {
	qi.size++
	if qi.head == nil {
		n := newNode(item)
		qi.head = n
		qi.tail = n
		return
	}
	qi.tail.next = newNode(item)
}

func (qi *LinkedListQueueInternal) Top() (interface{}, error) {
	if qi.head == nil {
		return nil, QueueEmptyError
	}
	return qi.head.item, nil
}

func (qi *LinkedListQueueInternal) Pop() error {
	if qi.head == nil {
		return QueueEmptyError
	}
	qi.size--
	qi.head = qi.head.next
	if qi.head == nil {
		qi.tail = nil
	}
	return nil
}

func (qi *LinkedListQueueInternal) Len() int {
	return qi.size
}

func (qi *LinkedListQueueInternal) Empty() bool {
	return qi.size == 0
}

type Queue interface {
	Put(item interface{})
	Get() (interface{}, error)
}

type BlockingQueue struct {
	qi    QueueInternal
	cond  *sync.Cond
}

func NewBlockingQueue(qi QueueInternal) *BlockingQueue {
	q := BlockingQueue{
		qi:   qi,
		cond: sync.NewCond(&sync.Mutex{}),
	}
	return &q
}

func (q *BlockingQueue) Put(item interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.qi.Append(item)
	q.cond.Signal()
}

func (q *BlockingQueue) Get() (interface{}, error) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for q.qi.Empty() {
		q.cond.Wait()
	}
	item, _ := q.qi.Top()
	_ = q.qi.Pop()
	return item, nil
}
