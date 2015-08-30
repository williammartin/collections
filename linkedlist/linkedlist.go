package linkedlist

type LinkedList struct {
	isEmpty bool
}

func New() *LinkedList {
	l := new(LinkedList)
	l.isEmpty = true
	return l
}

func (l *LinkedList) IsEmpty() bool {
	return l.isEmpty
}

func (l *LinkedList) Add(value string) {
	l.isEmpty = false
}
