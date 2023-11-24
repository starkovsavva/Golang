package main

import (
	"fmt"
	"LinkedList"
)

func main() {

	l := LinkedList{}

	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.DeleteFrom(2)
	l.printData()

}
