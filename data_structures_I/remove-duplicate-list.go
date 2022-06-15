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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	curr := head
	check := head.Next
	count := 0
	for {
		//print(curr.Val, check.Val, count, "\n")
		if check.Val == curr.Val {
			if check.Next == nil {
				curr.Next = nil
				return head
			}
			check = check.Next
			count++
		} else if check.Val != curr.Val && count > 0 {
			curr.Next = check
			count = 0
		} else if check.Val != curr.Val && count == 0 {
			if check.Next == nil {
				return head
			}
			curr = curr.Next
			check = check.Next
		}
	}

	return head
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
	input1 := []int{1, 1, 2, 3, 3}
	n1 := ListNode{Next: nil, Val: input1[0]}
	list1.head = &n1
	for i := 1; i < len(input1); i++ {
		list1.Insert(input1[i])
	}

	printFrom(list1.head)
	printFrom(deleteDuplicates(list1.head))
}
