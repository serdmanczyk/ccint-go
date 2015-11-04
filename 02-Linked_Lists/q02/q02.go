// Implement an algorithm to find the kth to last element of a singly linked list.

package q02

import (
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
)

// O(N)
// Requires N iterations to end, then N-K iterations to final element
func KthToLastUnknownLength(l *linkedlist.LinkedList, k int) *linkedlist.Node {
	if k > l.Length || k < 1 {
		return nil
	}

	// pretend we don't have l.Length, use runners to count
	slowRunner := l.Root
	fastRunner := l.Root
	midpoint := 0
	for fastRunner != nil && fastRunner.Next != nil {
		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next.Next
		midpoint++
	}

	var runner *linkedlist.Node
	var length, pos int
	if fastRunner == nil {
		length = midpoint * 2
	} else {
		length = (midpoint * 2) + 1
	}

	// determine if endpoint is before
	// or after midpoint, select
	// runner based on calculation
	endpoint := length - k
	if endpoint > (length / 2) {
		runner = slowRunner
		pos = midpoint
	} else {
		runner = l.Root
		pos = 0
	}

	// iterate runner to endpoint
	for i := pos; i != endpoint && runner != nil; i++ {
		runner = runner.Next
	}

	return runner
}

// O(N) memory O(N)
func KthToLastUnknownLengthAlternative(l *linkedlist.LinkedList, k int) *linkedlist.Node {
	if k > l.Length || k < 1 {
		return nil
	}

	storage := make([]*linkedlist.Node, 0, l.Length)
	runner := l.Root
	length := 0
	for runner != nil {
		storage = append(storage, runner)
		runner = runner.Next
		length++
	}

	return storage[length-k]
}

// O(N-K)
func KthToLastKnownLength(l *linkedlist.LinkedList, k int) *linkedlist.Node {
	if k > l.Length || k < 1 {
		return nil
	}

	endpoint := l.Length - k
	runner := l.Root
	for i := 0; i != endpoint; i++ {
		runner = runner.Next
	}

	return runner
}
