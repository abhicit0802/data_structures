package linkedlist

type LinkedList interface {
	Add(int)
	AddAll([]int)
	Print()
	IsPresent(int) bool
	PrintReverse()
}