package main

import (
	"fmt"
	"learn-go/data_structures/linkedlist"
)

func main() {
	l := linkedlist.New()
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.AddAll([]int{8, 9, 10})
	l.Print()
	fmt.Println()
	fmt.Println(l.IsPresent(9))
	l.PrintReverse()
}
