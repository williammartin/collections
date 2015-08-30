package set

type Set struct {
	size int
}

func New() *Set {
	var s = new(Set)
	return s
}

func (s *Set) IsEmpty() bool {
	return s.size == 0
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) Add(value string) {
	s.size++
}
