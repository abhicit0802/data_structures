package linkedlist

import "fmt"

type circularlyNode struct {
	data int
	next *circularlyNode
}

type circularList struct {
	head *circularlyNode
}

func NewCircularList() *circularList {
	return &circularList{}
}

func (cl *circularList) Add(data int) {
	if cl.head == nil {
		cl.head = &circularlyNode{
			data: data,
		}
		cl.head.next = cl.head
		return
	}
	p := cl.head
	for {
		if p.next == cl.head {
			p.next = &circularlyNode{
				data: data,
				next: cl.head,
			}
			return
		}
		p = p.next
	}
}

func (cl *circularList) AddAll(vals []int) {
	p := cl.head
	for _, val := range vals {
		if cl.head == nil {
			cl.head = &circularlyNode{
				data: val,
			}
			cl.head.next = cl.head
		}
		for {
			if p.next == cl.head {
				break
			}
			p = p.next
		}
		p.next = &circularlyNode{
			data: val,
			next: cl.head,
		}
		p = p.next
	}
}

func (cl *circularList) Print() {
	if cl.head == nil {
		return
	}
	p := cl.head
	for {
		fmt.Print(p.data, " ")
		p = p.next
		if p == cl.head {
			return
		}
	}
}

func (cl *circularList) IsPresent(val int) bool {
	if cl.head == nil {
		return false
	}
	p := cl.head
	for p.next != cl.head {
		if p.data == val {
			return true
		}
		p = p.next
	}
	return false
}

func (cl *circularList) PrintReverse() {
	cl.reverseCircular(cl.head)
}

func (cl *circularList) reverseCircular(p *circularlyNode) {
	if p.next == cl.head {
		fmt.Print(p.data, " ")
		return
	}
	cl.reverseCircular(p.next)
	fmt.Print(p.data, " ")
}
