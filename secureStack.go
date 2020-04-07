package goStack

import "container/list"

type secureStack struct {
	data *list.List
	size int
	lock chan int8
}

func NewSecureStack() *secureStack {
	q := new(secureStack)
	q.init()
	return q
}

func (q *secureStack) init() {
	q.data = list.New()
	q.lock = make(chan int8, 1)
}

func (q *secureStack) Size() int {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	return q.size
}

func (q *secureStack) Empty() bool {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	return q.size == 0
}

func (q *secureStack) Top() interface{} {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	return q.data.Back().Value
}

func (q *secureStack) Push(value interface{}) {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	q.data.PushBack(value)
	q.size++
}

func (q *secureStack) Pop() {
	q.lock <- 1
	defer func(lock chan int8) {<- lock}(q.lock)
	if q.size > 0 {
		tmp := q.data.Back()
		q.data.Remove(tmp)
		q.size--
	}
}