package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
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

func reverseList(head *ListNode) *ListNode {
	first := true
	curr := head
	prev := curr
	if head == nil {
		return head
	}
	for {
		//fmt.Print(curr.Val, prev.Val, "\n")
		if first {
			if curr.Next == nil {
				return curr
			}
			next := curr.Next
			curr.Next = nil
			curr = next
			first = false
			//fmt.Print("post first- ", curr.Val, prev.Val, "\n")
		} else {
			if curr.Next != nil {
				next := curr.Next
				curr.Next = prev
				prev = curr
				curr = next
			} else {
				curr.Next = prev
				//fmt.Print("before return- ", curr.Val, prev.Val, "\n")
				return curr
			}
			//fmt.Print("post else- ", curr.Val, prev.Val, "\n")
		}
	}
}

func printFrom(n *ListNode) {
	currentNode := n
	if currentNode == nil {
		fmt.Println("List is empty.")
	}
	for currentNode.Next != nil {
		fmt.Printf("%+v\n", *currentNode)
		currentNode = currentNode.Next
	}
	fmt.Printf("%+v\n", *currentNode)

}

func main() {
	var list1 List
	input1 := []int{1}
	n1 := ListNode{Next: nil, Val: input1[0]}
	list1.head = &n1
	for i := 1; i < len(input1); i++ {
		list1.Insert(input1[i])
	}

	printFrom(list1.head)
	printFrom(reverseList(list1.head))
}
