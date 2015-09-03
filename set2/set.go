package set

//Set is a collection without any particular order and no repeated values
type Set struct {
	Values map[string]bool
}

func New() *Set {
	return &Set{Values: make(map[string]bool)}
}

func (s *Set) IsEmpty() bool {
	return len(s.Values) != 0
}

func (s *Set) Size() int {
	return len(s.Values)
}

func (s *Set) Add(value string) {
	s.Values[value] = false
}

func (s *Set) Contains(value string) bool {
	_, isPresent := s.Values[value]
	return isPresent
}

func (s *Set) Remove(value string) {
	delete(s.Values, value)
}
