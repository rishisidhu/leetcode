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

func removeElements(head *ListNode, val int) *ListNode {
	head_l := head
	prev := head
	if ((head.Next == nil) && head.Val == val) || head == nil {
		return &ListNode{}
	} else if head.Next == nil && head.Val != val {
		return head
	}
	currentNode := head.Next
	flag := true
	for {
		//fmt.Print(prev.Val, currentNode.Val, "\n")
		if prev.Val == val && flag {
			if prev.Next == nil {
				return &ListNode{}
			}
			prev = prev.Next
			head_l = prev
			if currentNode.Next == nil {
				continue
			}
			currentNode = currentNode.Next

			continue
		} else {
			flag = false
			if currentNode.Val == val {
				if currentNode.Next == nil {
					prev.Next = nil
					break
				}
				prev.Next = currentNode.Next
				currentNode = currentNode.Next
			} else {
				if currentNode.Next == nil {
					break
				}
				prev = prev.Next
				currentNode = currentNode.Next
			}
		}

	}
	return head_l
}

func main() {
	var list1 List
	input1 := []int{1, 2}
	n1 := ListNode{Next: nil, Val: input1[0]}
	list1.head = &n1
	for i := 1; i < len(input1); i++ {
		list1.Insert(input1[i])
	}
	printFrom(removeElements(list1.head, 1))

}
