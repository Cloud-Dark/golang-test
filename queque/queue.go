package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type queueBox struct {
	maxItem int
	items   []interface{}
}

func New(size int) queueBox {
	q := queueBox{}
	q.maxItem = size
	return q
}

func (q *queueBox) Push(key interface{}) {
	if len(q.items) < q.maxItem {
		q.items = append(q.items, key)
	} else {
		q.items = append(q.items[1:], key)
	}

}

func (q *queueBox) Pop() interface{} {
	itemPop := q.items[0]
	q.items = q.items[1:]
	return itemPop
}

func (q *queueBox) Contains(key interface{}) bool {
	var has bool
	for _, val := range q.items {
		if val == key {
			has = true
			break
		} else {
			has = false
		}
	}
	return has
}

func (q *queueBox) Len() int {
	return len(q.items)
}

func (q *queueBox) Keys() []interface{} {
	return q.items
}
