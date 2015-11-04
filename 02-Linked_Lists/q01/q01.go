// Write code to remove duplicates from an unsorted linked list. FOLLOW UP How would
//  you solve this problem if a temporary buffer is not allowed?
package q01

import (
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
)

func RemoveDuplicates(l *linkedlist.LinkedList) {
	if l.Length == 0 {
		return
	}

	used := make(map[int]bool, l.Length)

	prev := l.Root
	cur := prev.Next
	used[prev.Value] = true
	for cur != nil {
		if _, dupe := used[cur.Value]; dupe {
			prev.Next = cur.Next
			cur = cur.Next
			l.Length--
		} else {
			used[cur.Value] = true
			prev = cur
			cur = cur.Next
		}
	}

	l.Head = prev
}

// O(n^2)
func RemoveDuplicatesNoStruct(l *linkedlist.LinkedList) {
	var prev *linkedlist.Node
	cur := l.Root
	for cur != nil {
		// for each node we need to iterate through the 
		// list and delete all duplicates ahead of it
		RemoveDuplicatesHelper(cur, l, cur.Value)
		prev = cur
		cur = cur.Next
	}

	if prev != nil {
		l.Head = prev
	}
}

func RemoveDuplicatesHelper(start *linkedlist.Node, parent *linkedlist.LinkedList, Value int) {
	prev := start
	cur := start.Next
	for cur != nil {
		if cur.Value == Value {
			prev.Next = cur.Next
			cur = prev.Next
			parent.Length--
		} else {
			prev = cur
			cur = cur.Next
		}
	}
}
