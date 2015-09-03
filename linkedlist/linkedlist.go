package linkedlist

type LinkedList struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	Value string
	Next *Node
}

func New() *LinkedList {
	return new(LinkedList)
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Add(value string) {
	
	newNode := new(Node)
	newNode.Value = value

	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *LinkedList) Contains(value string) bool {

	for node := l.head; node != nil; node = node.Next {
		if node.Value == value {
			return true
		}
	}
	return false
}

func (l *LinkedList) Head() *Node {
	return l.head
}
