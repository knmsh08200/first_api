package main

import (
	"fmt"
)

// type node struct {
// 	data int
// 	next *node
// }

// type linkedlist struct {
// 	head   *node
// 	length int
// }

// func (l *linkedlist) prepend(n *node) {
// 	second := l.head
// 	l.head = n
// 	l.head.next = second
// 	l.length++
// }
// func (l *linkedlist) printlist(n *node) {
// 	toPrint := l.head
// 	for l.length != 0 {
// 		fmt.Println(toPrint.data)
// 		toPrint = toPrint.next
// 		l.length--
// 	}

// }

// func (l *linkedlist) deleteListValue(val int) {
// 	if l.length == 0 {
// 		return
// 	}

// 	for l.head.data == val {
// 		l.head = l.head.next
// 		l.length--
// 		return
// 	}

// 	previous := l.head
// 	for previous.next.data != val {
// 		if previous.next.next == nil {
// 			return
// 		}

// 		previous = previous.next
// 	}

// 	previous.next = previous.next.next

// }

func main() {
	// mylist := linkedlist{}
	// node1 := &node{data: 48}
	// node2 := &node{data: 56}
	// mylist.prepend(node1)
	// mylist.prepend(node2)
	// fmt.Println(mylist)

	// mylist.deleteListValue(50)

	// mylist.printlist(mylist.head)
	list1 := ListNode{Val: 1}
	list2 := ListNode{Val: 1}
	list3 := ListNode{Val: 1}
	list4 := ListNode{Val: 3}
	list5 := ListNode{Val: 4}
	list1.Next = &list2
	list2.Next = &list3
	list3.Next = &list4
	list4.Next = &list5

	j := deleteDuplicates(&list1)

	for j != nil {
		fmt.Println(j.Val)
		j = j.Next
	}

}
