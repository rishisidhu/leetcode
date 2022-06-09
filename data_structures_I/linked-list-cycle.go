package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
}

func hasCycle(head *ListNode) bool {
	count := 0
	if head == nil {
		return false
	}
	for {
		if count > 10000 {
			return true
		}
		if head.Next != nil {
			head = head.Next
			count++
			//fmt.Printf("\rOn %d", count)
		} else {
			return false
		}
	}
}
func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast_ptr := head
	for {

		if head.Next != nil {
			head = head.Next
			//fmt.Printf("curr slow: ")
			//fmt.Printf("%+v\n", *head)
		} else {
			return false
		}
		if fast_ptr.Next == nil {
			return false
		}
		if fast_ptr.Next.Next != nil {
			fast_ptr = fast_ptr.Next.Next
			//fmt.Printf("curr fast: ")
			//fmt.Printf("%+v\n******\n", *fast_ptr)
		} else {
			return false
		}
		if fast_ptr == head {
			return true
		}
	}
}

func (L *List) Insert(val int) {
	l := L.head
	n := &ListNode{
		Next: nil,
		Val:  val,
	}
	for l.Next != nil {
		l = l.Next
	}
	l.Next = n
}

func (l *List) addLoop(pos int) {
	specialNode := l.head
	currentNode := l.head
	for i := 0; i < pos; i++ {
		specialNode = specialNode.Next
	}
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = specialNode
}

func (l *List) showAllNodes(limit int) error {
	currentNode := l.head
	if currentNode == nil {
		//fmt.Println("List is empty.")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for i := 0; i < limit; i++ {
		currentNode = currentNode.Next
		//fmt.Printf("%+v\n", *currentNode)
	}

	return nil
}

func main() {
	input := []int{3, 2, 0, -4, 1, 6, 7, 8, 14, -12, 0, 23}
	pos := -1
	var mylist List
	n := ListNode{Next: nil, Val: input[0]}
	mylist.head = &n
	for i := 1; i < len(input); i++ {
		mylist.Insert(input[i])
	}
	if pos >= 0 {
		mylist.addLoop(pos)
	}
	mylist.showAllNodes(len(input))
	fmt.Print(hasCycle2(mylist.head))
}
