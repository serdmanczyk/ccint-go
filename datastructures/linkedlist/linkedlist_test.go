package linkedlist

import (
	"testing"
    . "gopkg.in/check.v1"
)

func Test (t *testing.T) { TestingT(t) }

type LListTests struct{}

var _ = Suite(&LListTests{})

func (t *LListTests) TestAdd(c *C) {
	l := New()
	
	l.Add(1)
	c.Assert(l.Root, NotNil)
	c.Assert(l.Head, NotNil)
	c.Assert(l.Root, Equals, l.Head)
	c.Assert(l.Root.Value, Equals, l.Head.Value)
	c.Assert(l.Root.Next,  IsNil)
	c.Assert(l.Length, Equals, 1)

	l.Add(2)
	c.Assert(l.Root, Not(Equals), l.Head)
	c.Assert(l.Root.Value, Not(Equals), l.Head.Value)
	c.Assert(l.Root.Next,  Equals, l.Head)
	c.Assert(l.Root.Next.Next,  IsNil)
	c.Assert(l.Root.Value,  Equals, 1)
	c.Assert(l.Root.Next.Value,  Equals, 2)
	c.Assert(l.Head.Value,  Equals, 2)
	c.Assert(l.Length, Equals, 2)

	l.Add(3)
	c.Assert(l.Root.Next.Next,  NotNil)
	c.Assert(l.Root.Next.Next,  Equals, l.Head)
	c.Assert(l.Root.Value,  Equals, 1)
	c.Assert(l.Root.Next.Value,  Equals, 2)
	c.Assert(l.Root.Next.Next.Value,  Equals, 3)
	c.Assert(l.Head.Value,  Equals, 3)
	c.Assert(l.Length, Equals, 3)
}

func (t *LListTests) TestDelete(c *C) {
	l := New()
	_, two, _ := l.Add(1), l.Add(2), l.Add(3)

	l.Delete(two)
	c.Assert(l.Length, Equals, 2)
	c.Assert(l.Root, NotNil)
	c.Assert(l.Root.Next, NotNil)
	c.Assert(l.Head, NotNil)
	c.Assert(l.Root.Value, Equals, 1)
	c.Assert(l.Root.Next.Value, Equals, 3)
	c.Assert(l.Head.Value, Equals, 3)

	l = New()
	one, _, _ := l.Add(1), l.Add(2), l.Add(3)

	l.Delete(one)
	c.Assert(l.Length, Equals, 2)
	c.Assert(l.Root, NotNil)
	c.Assert(l.Root.Next, NotNil)
	c.Assert(l.Head, NotNil)
	c.Assert(l.Root.Value, Equals, 2)
	c.Assert(l.Root.Next.Value, Equals, 3)
	c.Assert(l.Head.Value, Equals, 3)

	l = New()
	_, _, three := l.Add(1), l.Add(2), l.Add(3)

	l.Delete(three)
	c.Assert(l.Length, Equals, 2)
	c.Assert(l.Root, NotNil)
	c.Assert(l.Root.Next, NotNil)
	c.Assert(l.Head, NotNil)
	c.Assert(l.Root.Value, Equals, 1)
	c.Assert(l.Root.Next.Value, Equals, 2)
	c.Assert(l.Head.Value, Equals, 2)

	l = New()
	one, two, three = l.Add(1), l.Add(2), l.Add(3)

	l.Delete(two)
	l.Delete(three)
	c.Assert(l.Length, Equals, 1)
	c.Assert(l.Root, NotNil)
	c.Assert(l.Root.Next, IsNil)
	c.Assert(l.Root, Equals, l.Head)
	c.Assert(l.Head, NotNil)
	c.Assert(l.Root.Value, Equals, 1)
	c.Assert(l.Head.Value, Equals, 1)

	l.Delete(one)
	c.Assert(l.Length, Equals, 0)
	c.Assert(l.Root, IsNil)
	c.Assert(l.Head, IsNil)
}

func (t *LListTests) TestNodeAtPOS(c *C) {
	l := New()

	c.Assert(l.NodeAtPos(-1), IsNil)
	c.Assert(l.NodeAtPos(0), IsNil)
	c.Assert(l.NodeAtPos(1), IsNil)

	l.Add(1)

	c.Assert(l.NodeAtPos(-1), IsNil)
	c.Assert(l.NodeAtPos(0), NotNil)
	c.Assert(l.NodeAtPos(0).Value, Equals, 1)
	c.Assert(l.NodeAtPos(1), IsNil)

	l.Add(2)

	c.Assert(l.NodeAtPos(-1), IsNil)
	c.Assert(l.NodeAtPos(0), NotNil)
	c.Assert(l.NodeAtPos(0).Value, Equals, 1)
	c.Assert(l.NodeAtPos(1), NotNil)
	c.Assert(l.NodeAtPos(1).Value, Equals, 2)
	c.Assert(l.NodeAtPos(2), IsNil)
	c.Assert(l.NodeAtPos(3), IsNil)

	l.Add(3)

	c.Assert(l.NodeAtPos(-1), IsNil)
	c.Assert(l.NodeAtPos(0), NotNil)
	c.Assert(l.NodeAtPos(0).Value, Equals, 1)
	c.Assert(l.NodeAtPos(1), NotNil)
	c.Assert(l.NodeAtPos(1).Value, Equals, 2)
	c.Assert(l.NodeAtPos(2), NotNil)
	c.Assert(l.NodeAtPos(2).Value, Equals, 3)
	c.Assert(l.NodeAtPos(3), IsNil)
	c.Assert(l.NodeAtPos(4), IsNil)
}

func (t *LListTests) TesttoJSON(c *C) {
	l := New()
	c.Assert(l.JSON(), Equals, "[]")

	l.Add(1)
	c.Assert(l.JSON(), Equals, "[1]")

	l.Add(2)
	c.Assert(l.JSON(), Equals, "[1,2]")

	l.Add(3)
	c.Assert(l.JSON(), Equals, "[1,2,3]")

	l.Add(4)
	l.Add(5)
	l.Add(6)
	c.Assert(l.JSON(), Equals, "[1,2,3,4,5,6]")
}

func (t *LListTests) TestFromJSON(c *C) {
	l := FromJSON("[1]")
	c.Assert(l.Root, NotNil)
	c.Assert(l.Head, NotNil)
	c.Assert(l.Root, Equals, l.Head)
	c.Assert(l.Root.Value, Equals, l.Head.Value)
	c.Assert(l.Root.Next,  IsNil)
	c.Assert(l.Length, Equals, 1)

	l = FromJSON("[1,2]")
	c.Assert(l.Root, Not(Equals), l.Head)
	c.Assert(l.Root.Value, Not(Equals), l.Head.Value)
	c.Assert(l.Root.Next,  Equals, l.Head)
	c.Assert(l.Root.Next.Next,  IsNil)
	c.Assert(l.Root.Value,  Equals, 1)
	c.Assert(l.Root.Next.Value,  Equals, 2)
	c.Assert(l.Head.Value,  Equals, 2)
	c.Assert(l.Length, Equals, 2)

	l = FromJSON("[1,2,3]")
	c.Assert(l.Root.Next.Next,  NotNil)
	c.Assert(l.Root.Next.Next,  Equals, l.Head)
	c.Assert(l.Root.Value,  Equals, 1)
	c.Assert(l.Root.Next.Value,  Equals, 2)
	c.Assert(l.Root.Next.Next.Value,  Equals, 3)
	c.Assert(l.Head.Value,  Equals, 3)
	c.Assert(l.Length, Equals, 3)

	l = FromJSON("[1,2,3,4,5,6]")
	c.Assert(l.Root.Value, Equals, 1)
	c.Assert(l.Root.Next.Value, Equals, 2)
	c.Assert(l.Root.Next.Next.Value, Equals, 3)
	c.Assert(l.Root.Next.Next.Next.Value, Equals, 4)
	c.Assert(l.Root.Next.Next.Next.Next.Value, Equals, 5)
	c.Assert(l.Root.Next.Next.Next.Next.Next.Value, Equals, 6)
}
func (t *LListTests) TestReverse(c *C) {
	l := FromJSON("[1,2,3,4]")
	l.Reverse()
	c.Assert(l.JSON(), Equals, "[4,3,2,1]")
	c.Assert(l.Root.Value, Equals, 4)
	c.Assert(l.Head.Value, Equals, 1)
	c.Assert(l.Length, Equals, 4)

	l = FromJSON("[]")
	l.Reverse()
	c.Assert(l.JSON(), Equals, "[]")

	l = FromJSON("[5,6,7,8]")
	l.Reverse()
	c.Assert(l.JSON(), Equals, "[8,7,6,5]")
	c.Assert(l.Root.Value, Equals, 8)
	c.Assert(l.Head.Value, Equals, 5)
	c.Assert(l.Length, Equals, 4)
}