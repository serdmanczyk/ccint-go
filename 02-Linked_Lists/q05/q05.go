// You have two numbers represented by a linked list, where each node contains a single
//  digit. The digits are stored in reverse order, such that the Ts digit is at the
//  head of the list. Write a function that adds the two numbers and returns the sum
//  as a linked list. EXAMPLE Input:(7-> 1 -> 6) + (5 -> 9 -> 2).Thatis,617 + 295.
//  Output: 2 -> 1 -> 9.That is, 912. FOLLOW UP Suppose the digits are stored in forward
//  order. Repeat the above problem. EXAMPLE Input:(6 -> 1 -> 7) + (2 -> 9 -> 5).Thatis,617
//  + 295. Output: 9 -> 1 -> 2.That is, 912

package q05

import (
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
)

func SumLinkedLists(l1, l2 *linkedlist.LinkedList) *linkedlist.LinkedList {
	sum := linkedlist.New()

	carry := 0
	cur1, cur2 := l1.Root, l2.Root
	for cur1 != nil && cur2 != nil {
		total := cur1.Value + cur2.Value + carry
		val := total % 10
		carry = total / 10
		sum.Add(val)

		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	if cur1 != nil {
		sum.Add(cur1.Value + carry)
	} else if (cur2 != nil) {
		sum.Add(cur2.Value + carry)
	} else if (carry > 0) {	
		sum.Add(carry)
	}

	return sum
}

func SumLinkedListsReverse(l1, l2 *linkedlist.LinkedList) *linkedlist.LinkedList {
	l1.Reverse()
	l2.Reverse()
	result := SumLinkedLists(l1, l2)
	result.Reverse()
	return result
}