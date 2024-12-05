package set

type Set[E comparable] map[E]bool

func New[E comparable](vals ...E) Set[E] {
	s := Set[E]{}
	for _, v := range vals {
		s[v] = true
	}
	return s
}

func (s Set[E]) Add(vals ...E) {
	for _, v := range vals {
		s[v] = true
	}
}

func (s Set[E]) Remove(v E) {
	delete(s, v)
}

func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

func (s Set[E]) Members() []E {
	result := make([]E, 0, len(s))
	for v := range s {
		result = append(result, v)
	}
	return result
}
