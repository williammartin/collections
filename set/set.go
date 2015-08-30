package set

type Set struct {
	isEmpty bool
}

func New() *Set {
	var s = new(Set)
	s.isEmpty = true

	return s
}

func (s *Set) IsEmpty() bool {
	return s.isEmpty
}

func (s *Set) Add(value string) {
	s.isEmpty = false
}
