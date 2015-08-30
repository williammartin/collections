package set

type Set struct {
	values []string
}

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

func (s *Set) Contains(value string) bool {
	for i := 0; i < s.Size(); i++ {
		if s.values[i] == value {
			return true
		}
	}
	return false
}

func (s *Set) Remove(value string) {
	for i := 0; i < s.Size(); i++ {
		if s.values[i] == value {
			s.values[i] = s.values[s.Size()-1]
			s.values = s.values[:s.Size()-1]
		}
	}
}
