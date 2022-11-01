package Utils

type Queue []string

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Push(str string) {
	*q = append(*q, str)
}

func (q *Queue) Pop() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}

	el := (*q)[0]
	*q = (*q)[1:]

	return el, true
}
