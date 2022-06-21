package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testQueueInternal(t *testing.T, qi QueueInternal) {
	var err error

	qi.Append(1)

	item, err := qi.Top()
	assert.NoError(t, err)
	assert.Equal(t, item, 1)

	err = qi.Pop()
	assert.NoError(t, err)

	size := qi.Len()
	assert.Equal(t, size, 0)

	empty := qi.Empty()
	assert.Equal(t, empty, true)
}

func TestArrayQueueInternal(t *testing.T) {
	qi := NewArrayQueueInternal()
	testQueueInternal(t, qi)
}

func TestLinkedListQueueInternal(t *testing.T) {
	qi := NewLinkedListQueueInternal()
	testQueueInternal(t, qi)
}

func testQueue(t *testing.T, q Queue) {
	q.Put(1)
	item, err := q.Get()
	assert.NoError(t, err)
	assert.Equal(t, item, 1)
}

func TestBlockingQueue(t *testing.T) {
	qi := NewArrayQueueInternal()
	q := NewBlockingQueue(qi)
	testQueue(t, q)
}
