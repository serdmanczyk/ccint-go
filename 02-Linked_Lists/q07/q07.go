// Implement a function to check if a linked list is a palindrome.

package q07

import (
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
)

func IsPalindrome(l *linkedlist.LinkedList) bool {
	slowRunner := l.Root
	fastRunner := l.Root

	// use fast runner and slow runnner to find middle
	midpoint := 0
	for fastRunner != nil && fastRunner.Next != nil {
		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next.Next
		midpoint++
	}

	// Odd length list, ignore middle
	if fastRunner != nil {
		slowRunner = slowRunner.Next 
	}

	// reverse first half of list
	var prev *linkedlist.Node
	cur := l.Root
	for i:=0; i < midpoint;i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// Now compare values that should be opposite each other
	opposite := prev
	for slowRunner != nil {
		if slowRunner.Value != opposite.Value {
			return false
		}

		slowRunner = slowRunner.Next
		opposite = opposite.Next
	}

	return true
}
