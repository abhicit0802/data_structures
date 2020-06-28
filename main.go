package main

import (
	"fmt"
	"learn-go/data_structures/linkedlist"
)

func main() {
	//singly()
	//doubly()
	circularly()

}

func circularly() {
	l := linkedlist.NewCircularList()
	l.Add(1)
	l.AddAll([]int{2, 3, 4, 5})
	l.Add(10)
	l.AddAll([]int{6, 7, 8, 9})
	l.Print()
	fmt.Println()
	fmt.Println(l.IsPresent(9))
	l.PrintReverse()
}

func doubly() {
	l := linkedlist.NewDoublyList()
	l.AddAll([]int{1, 2, 3, 4, 5, 6})
	l.Add(7)
	l.AddAll([]int{11, 12, 13, 14, 15, 16})
	l.Add(17)
	l.Print()
	l.Add(18)
	fmt.Println()
	l.Add(19)
	l.PrintReverse()
	fmt.Println()
}

func singly() {
	l := linkedlist.NewSinglyList()
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.AddAll([]int{8, 9, 10})
	l.AddAll([]int{18, 19, 20})
	l.Print()
	fmt.Println()
	fmt.Println(l.IsPresent(9))
	l.PrintReverse()
	fmt.Println()
}
