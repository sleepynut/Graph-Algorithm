package Utils

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
