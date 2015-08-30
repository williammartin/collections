package linkedlist

type LinkedList struct {
	size int
}

func New() *LinkedList {
	return new(LinkedList)
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Add(value string) {
	l.size++
}

func (l *LinkedList) Size() int {
	return l.size
}
