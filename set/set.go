package set

type Set struct {
	values []string
	size   int
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
	if s.Contains(value) {
		return
	}
	s.values = append(s.values, value)
	s.size++
}

func (s *Set) Contains(value string) bool {
	for i := 0; i < s.size; i++ {
		if s.values[i] == value {
			return true
		}
	}
	return false
}

func (s *Set) Remove(value string) {
	for i := 0; i < s.size; i++ {
		if s.values[i] == value {
			s.values[i] = s.values[len(s.values)-1]
			s.values = s.values[:len(s.values)-1]
			s.size--
		}
	}
}
