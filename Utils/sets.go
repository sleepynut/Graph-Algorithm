package Utils

type Sets[k comparable, v any] map[k]v

func (s *Sets[K, V]) Has(k K) bool {
	_, ok := (*s)[k]
	return ok
}

func (s *Sets[K, V]) Add(k K, v V) {
	(*s)[k] = v
}

func (s *Sets[K, V]) Remove(k K) {
	if s.Has(k) {
		delete(*s, k)
	}
}

func NewSets[k comparable, v any]() Sets[k, v] {
	return make(map[k]v)
}
