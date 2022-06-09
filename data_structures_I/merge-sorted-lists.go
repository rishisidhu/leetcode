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

func (l *List) showAllNodes() error {
	currentNode := l.head
	if currentNode == nil {
		fmt.Println("List is empty.")
		return nil
	}
	for currentNode.Next != nil {
		fmt.Printf("%+v\n", *currentNode)
		currentNode = currentNode.Next
	}
	fmt.Printf("%+v\n", *currentNode)

	return nil
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list2 == nil {
		return list1
	}
	if list1 == nil {
		return list2
	}
	//l2_head := list2
	prev1 := &ListNode{Next: nil, Val: -100000}
	//prev2 := list2
	if list1.Val > list2.Val {
		//fmt.Print(l1_head, ",", list1, ",", list2, "\n")
		//fmt.Print(l1_head.Next, ",", list1.Next, ",", list2.Next, "\n")
		temp1 := &ListNode{Next: list1, Val: -100000}
		temp2 := &ListNode{Next: list2, Val: -100000}
		list1 = temp2.Next
		list2 = temp1.Next

	}
	l1_head := list1

	for {

		if list1.Val <= list2.Val {
			fmt.Print("l1: ", list1.Val, " l2: ", list2.Val, " prev1:", prev1.Val, "\n")
			prev1 = list1
			if list1.Next == nil {
				for {
					fmt.Print("l1: ", list1.Val, " l2: ", list2.Val, " prev1:", prev1.Val, "\n")
					list1.Next = list2
					if list2.Next == nil {
						return l1_head
					}
					list2 = list2.Next
					list1 = list1.Next
				}
			}
			list1 = list1.Next
		} else {
			fmt.Print("else l1: ", list1.Val, " l2: ", list2.Val, " prev1:", prev1.Val, "\n")
			if list2.Next == nil {
				prev1.Next = list2
				list2.Next = list1
				return l1_head
			}
			fmt.Print("else next2: ", list2.Next.Val, "prev1: ", prev1.Val, " prev1.Next: ", list2.Val, " list2.Next: ", list1.Val, "\n")

			next2 := list2.Next
			prev1.Next = list2
			list2.Next = list1
			list2 = next2
			prev1 = prev1.Next
		}
	}
}

func main() {
	var list1, list2 List
	input1 := []int{0, 3, 5, 6}
	n1 := ListNode{Next: nil, Val: input1[0]}
	list1.head = &n1
	for i := 1; i < len(input1); i++ {
		list1.Insert(input1[i])
	}
	input2 := []int{6, 7}
	n2 := ListNode{Next: nil, Val: input2[0]}
	list2.head = &n2
	for i := 1; i < len(input2); i++ {
		list2.Insert(input2[i])
	}
	fmt.Print(input1, "\n", input2, "\n")
	//list1.showAllNodes()
	//list2.showAllNodes()
	printFrom(mergeTwoLists(list1.head, list2.head))
	//list1.showAllNodes()
	//list2.showAllNodes()
}
