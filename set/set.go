package set

type Set struct {
	values []string
}

const INDEX_NOT_FOUND = -1

func New() *Set {
	return new(Set)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) Size() int {
	return len(s.values)
}

func (s *Set) Add(value string) {
	if s.Contains(value) {
		return
	}
	s.values = append(s.values, value)
}

func (s *Set) indexOf(value string) int {
	for i := 0; i < s.Size(); i++ {
		if s.values[i] == value {
			return i
		}
	}
	return INDEX_NOT_FOUND
}

func (s *Set) Contains(value string) bool {
	return s.indexOf(value) != INDEX_NOT_FOUND
}

func (s *Set) Remove(value string) {
	if !s.Contains(value) {
		return
	}

	index := s.indexOf(value)
	s.values[index] = s.values[s.Size()-1]
	s.values = s.values[:s.Size()-1]
}
