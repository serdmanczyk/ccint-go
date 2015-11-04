// Implement an algorithm to delete a node in the middle of a singly linked list,
//  given only access to that node. EXAMPLE Input: the node c from the linked list
//  a->b->c->d->e Result: nothing isreturned, 
//  but the new linked list looks like a->b->d->e

package q03

import (
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
)

func DeleteNode(n *linkedlist.Node) {
	if n.Next != nil {
		n.Value = n.Next.Value
		n.Next = n.Next.Next
	}
	// This doesn't 'delete' node but
	// rather modifies its value to that of the
	// node in front of it, and deletes it.
	// Using this method we cannot delete
	// the node at the front of a singly linked list.
}
