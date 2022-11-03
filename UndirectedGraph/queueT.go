package main

type Passed struct {
	Node string
	Dist int
}

func (p *Passed) move() {
	p.Dist += 1
}

type Moveable interface {
	move()
}

type QueueT[T Moveable] []T

func (q *QueueT[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *QueueT[T]) Push(item T) {
	*q = append(*q, item)
}

func (q *QueueT[T]) Pop() (T, bool) {
	if (*q).IsEmpty() {
		var empty T
		return empty, false
	}

	el := (*q)[0]
	*q = (*q)[1:]
	return el, true
}
