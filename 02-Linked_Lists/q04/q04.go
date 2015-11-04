// Write code to partition a linked list around a value x, such that all nodes less
//  than x come before all nodes greater than or equal to x.

package q04

import (
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
)

func PartitionLinkedList(l *linkedlist.LinkedList, x int) {
	var ltroot,lessthan,gtroot,greaterthan *linkedlist.Node

	current := l.Root
	for current != nil {
		if current.Value < x {
			if ltroot == nil || lessthan == nil {
				ltroot = current
				lessthan = current
			} else {
				lessthan.Next = current
				lessthan = lessthan.Next
			}
		} else {
			if gtroot == nil || greaterthan == nil {
				gtroot = current
				greaterthan = current
			} else {
				greaterthan.Next = current
				greaterthan = greaterthan.Next
			}
		}

		current = current.Next
	}

	if l.Root != nil {
		if ltroot != nil && gtroot != nil {
			l.Root = ltroot
			l.Head = greaterthan
			lessthan.Next = gtroot
			greaterthan.Next = nil
		}
	}
}
