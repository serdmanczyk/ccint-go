// Singly linked list, because most questions in chapter call
// for a singly linked list, and golang library only provides
// doubly linked list.  Also, exports values for easy access
// in separate question packages.

package linkedlist

import (
	"encoding/json"
	"strings"
	"strconv"
)

type Node struct {
	Next *Node
	Value int
}

type LinkedList struct {
	Root,Head *Node
	Length int
}

func New() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (l *LinkedList) Add (val int) *Node {
	if l.Root == nil {
		l.Root = &Node{nil, val}
		l.Head = l.Root
		l.Length = 1
	} else {
		l.Head.Next = &Node{nil, val}
		l.Head = l.Head.Next
		l.Length++
	}

	return l.Head
}

func (l *LinkedList) Find (val int) *Node {
	cur := l.Root
	for cur != nil {
		if cur.Value == val {
			break
		}
		cur = cur.Next
	}

	return cur
}

func (l *LinkedList) NodeAtPos (pos int) *Node {
	if l.Root == nil || pos > l.Length-1 || pos < 0 {
		return nil
	}

	cur := l.Root
	for i:=0;i<pos;i++ {
		cur = cur.Next
	}

	return cur
}

func (l *LinkedList) Delete (e *Node) {
	if e == l.Root {
		if l.Root == l.Head {
			l.Root = nil
			l.Head = nil
			l.Length = 0
		} else {
			l.Root = l.Root.Next
			l.Length--
		}

		return
	}

	prev := l.Root
	cur := prev.Next
	for cur != nil {
		if cur == e {
			break
		}
		prev = cur
		cur = cur.Next
	}

	if cur != nil {
		prev.Next = cur.Next
		l.Length--
		if cur == l.Head {
			if cur.Next != nil {
				l.Head = cur.Next
			} else {
				l.Head = prev
			}
		}
	}
}

func (l *LinkedList) Reverse() {
	var prev,cur,next *Node

	cur = l.Root
	l.Head = cur
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Root = prev
}

func (l *LinkedList) JSON() string {
	if l.Root == nil {
		return "[]"
	}

	vals := make([]string, 0, l.Length)
	cur := l.Root
	for cur != nil {
		vals = append(vals, strconv.Itoa(cur.Value))
		cur = cur.Next
	}

	return "[" + strings.Join(vals, ",") + "]"
}

func FromJSON(s string) *LinkedList {
	var vals []int
	err := json.Unmarshal([]byte(s), &vals)
	if err != nil {
		panic(err)
	}

	l := &LinkedList{nil, nil, 0}
	for _,i := range vals {
		l.Add(i)
	}

	return l
}
