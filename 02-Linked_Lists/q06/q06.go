// Given a circular linked list, implement an algorithm which returns the node
// at the beginning of the loop. DEFINITION Circular linked list: A (corrupt)
// linked list in which a node's next pointer points to an earlier node, so as // to make a loop in the linked list. EXAMPLE Input:A ->B->C->D->E->C
// [the same C as earlier]
// Output:C

package q06

import (
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
	// "fmt"
)

// runtime O(n+l*a)
// n = length of loop
// a = length of list up to loop
// l = length of loop in list
func FindLoopUnknownLength(l *linkedlist.LinkedList) *linkedlist.Node {
	slowRunner := l.Root
	fastRunner := l.Root.Next

	nPlusK := 0
	for slowRunner != fastRunner {
		if fastRunner == nil || fastRunner.Next == nil {
			// return nil
			return nil
		}
		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next.Next
		nPlusK++
	}

	// we know this is inside the loop
	intersection := slowRunner

	// start a runner at the start of the chain
	listRunner := l.Root

	// move it forward until it meets the intersection
	for listRunner != intersection {

		// move a runner around the loop
		loopRunner := intersection.Next
		for loopRunner != intersection {
			// if the runner intersects our runner
			// from the start of the list, that
			// means that node is the start of
			// the loop
			if loopRunner == listRunner {
				return listRunner
			}
			loopRunner = loopRunner.Next
		}

		listRunner = listRunner.Next
	}

	// otherwise we can assume where our fast/slowRunner's
	// met was our loop start
	return intersection
}

// runtime O(n)
func FindLoopKnownLength(l *linkedlist.LinkedList) *linkedlist.Node {
	slowRunner := l.Root
	fastRunner := l.Root.Next

	nPlusK := 0
	for slowRunner != fastRunner {
		if fastRunner == nil || fastRunner.Next == nil {
			return nil
		}
		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next.Next
		nPlusK++
	}

	// Now we're inside the loop, and we know that the length of the
	// 'list' is equal to n+k+l, where n is the length until we reach the
	// loop start, k is the length we iterated into the loop until our
	// runners intersected, and l is the length of our loop (not as
	// important). Our intersection point is length n+k into the loop, so
	// it's length-(n+k) from the end of the list / the intersection.
	LoopMinusK := l.Length - nPlusK
	for i:=0;i<LoopMinusK;i++ {
		slowRunner = slowRunner.Next
	}

	return slowRunner
}
