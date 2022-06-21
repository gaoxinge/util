package collection

type node struct {
	item interface{}
	next *node
}

func newNode(item interface{}) *node {
	n := node{
		item: item,
		next: nil,
	}
	return &n
}
