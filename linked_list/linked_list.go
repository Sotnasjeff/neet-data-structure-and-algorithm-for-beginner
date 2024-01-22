package linkedlist

type node struct {
	Value int
	Next  *node
}

type LinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: &node{
			Value: 0,
			Next:  nil,
		},
		Tail: &node{
			Value: 0,
			Next:  nil,
		},
		Length: 0,
	}
}

func (l *LinkedList) Get(index int) int {
	currNode := l.Head.Next
	var i int
	for currNode != nil {
		if i == index {
			return currNode.Value
		}
		i += 1
		currNode = currNode.Next
	}
	return -1
}

func (l *LinkedList) InsertHead(value int) {
	newNode := &node{
		Value: value,
		Next:  nil,
	}
	newNode.Next = l.Head.Next
	l.Head.Next = newNode

	if newNode.Next == nil {
		l.Tail = newNode
	}
	l.Length++
}

func (l *LinkedList) InsertTail(value int) {
	if l.Length == 0 {
		l.InsertHead(value)
	}

	newNode := &node{
		Value: value,
		Next:  nil,
	}

	l.Tail.Next = newNode
	l.Tail = l.Tail.Next
	l.Length++
}

func (l *LinkedList) Remove(index int) bool {
	if index < 0 || index > l.Length {
		return false
	}
	var i int
	currNode := l.Head
	for i < index && currNode != nil {
		i += 1
		currNode = currNode.Next
	}

	if currNode != nil && currNode.Next != nil {
		if currNode.Next == l.Tail {
			l.Tail = currNode
		}
		currNode.Next = currNode.Next.Next
		return true
	}

	return false
}

func (l *LinkedList) GetValues() []int {
	currNode := l.Head.Next
	var i []int

	for currNode != nil {
		i = append(i, currNode.Value)
		currNode = currNode.Next
	}

	return i
}
