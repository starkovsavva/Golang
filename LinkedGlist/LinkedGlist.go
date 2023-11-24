package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) printData() {

	if l.Head != nil {
		iterator := l.Head
		for ; iterator.Next != nil; iterator = iterator.Next {
			fmt.Println(iterator.Value)
		}
		fmt.Println(iterator.Value)
	} else {
		fmt.Println("empty!")
	}

}

func New(q int) *LinkedList {
	l := &LinkedList{}
	if q <= 0 {
		return l
	}

	l.Head = &Node{Value: 0}
	iterator := l.Head
	for counter := 1; counter < q; counter++ {
		newNode := &Node{Value: 0}
		iterator.Next = newNode
		iterator = iterator.Next
	}

	return l
}

func (l LinkedList) Size() int {
	size := 0
	for iterator := l.Head; iterator.Next != nil; iterator = iterator.Next {
		size++
	}
	size++
	return size
}

func (l *LinkedList) Pop() {
	prevtoDelete := l.Head
	if prevtoDelete.Next == nil {
		return

	}
	for ; prevtoDelete.Next.Next != nil; prevtoDelete = prevtoDelete.Next {
	}
	prevtoDelete.Next = nil

}

func (l *LinkedList) Add(value int) {
	newNode := &Node{Value: value, Next: nil}
	if l.Head == nil {
		l.Head = newNode
	} else {
		iterator := l.Head
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = newNode
	}
}

func (l *LinkedList) At(pos int) int {
	counter := 0
	iterator := l.Head
	if l.Size() <= pos || pos < 0 {
		return -1
	}
	for ; counter != pos; counter++ {
		iterator = iterator.Next
	}

	return iterator.Value
}

func (l *LinkedList) DeleteFrom(pos int) {

	counter := pos - 1
	iteratorCurr := l.Head
	iteratorPrev := iteratorCurr
	// Удалить элемент из позиции pos.
	if pos < 0 || pos >= l.Size() { // недопустимая позиция
		return
	}
	for ; counter != pos; counter++ {
		iteratorCurr = iteratorCurr.Next
	}
	iteratorPrev.Next = iteratorCurr.Next

}

func (l *LinkedList) UpdateAt(pos, val int) {
	//Поменять значение на позиции
	counter := 0
	iterator := l.Head
	if pos < 0 || pos >= l.Size() { // недопустимая позиция
		return
	}
	for ; counter != pos; counter++ {
		iterator = iterator.Next
	}
	iterator.Value = val

}

func NewFromSlice(s []int) *LinkedList {
	l := &LinkedList{}
	size := len(s) - 1
	counter := 0
	for ; counter < size; counter++ {
		l.Add(s[counter])
	}
	return l
}

func main() {

	l := New(3)

	l.printData()

}
