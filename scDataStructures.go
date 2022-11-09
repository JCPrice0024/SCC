package SCC

import (
	"errors"
	"fmt"
	"hash/fnv"
)

// Stack of int
type Stack struct {
	slice []int
}

// Push adds the int provided to the top of the stack.
func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}

// Peek returns the top item from the stack, but DOES NOT
// remove it from the stack.
func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

// Size checks the size of the stack and returns whether
// it's empty (true) or not (false).
func (s *Stack) Size() bool {
	return len(s.slice) == 0
}

// Pop removes and returns the top item from the stack.
func (s *Stack) Pop() int {
	var ret int = s.Peek()
	s.slice = s.slice[:len(s.slice)-1]
	return ret
}

// Search Tree
type Node struct {
	Val   int
	Right *Node
	Left  *Node
}
type SearchTree struct {
	Root *Node
}

// Search trees have 3 operations; Insert, Find, and Delete
func (s *SearchTree) Insert(value int) error {
	if s == nil {
		return errors.New("cannot insert on an empty tree")
	}
	value = HashVal(value)
	if s.Root == nil {
		s.Root = &Node{Val: value}
		return nil
	}
	s.Root.insert(value)
	return nil
}

func (n *Node) insert(value int) {
	switch {
	case n.Val == value:
		return
	case value < n.Val:
		if n.Left == nil {
			n.Left = &Node{Val: value}
			return
		}
		n.Left.insert(value)
	case value > n.Val:
		if n.Right == nil {
			n.Right = &Node{Val: value}
			return
		}
		n.Right.insert(value)
	}
}

func (s *SearchTree) String() string {
	if s == nil {
		return ""
	}
	return s.Root.string()
}
func (n *Node) string() string {
	if n == nil {
		return ""
	}
	var tempStr string
	var tempStr2 string
	fmt.Println(n.Left)
	//fmt.Println(n)
	tempStr = n.Left.string()
	if n.Left != nil {
		//tempStr = n.Left.string()
		tempStr = fmt.Sprintf("Left: %s %d", tempStr, n.Val)
	}
	//tempStr = fmt.Sprintf("Left: %s %d", tempStr, n.Val)
	tempStr2 = n.Right.string()
	if n.Left != nil {
		//tempStr2 = n.Right.string()
		tempStr = fmt.Sprintf("Right: %s %s", tempStr, tempStr2)
	}
	//tempStr = fmt.Sprintf("Right: %s %s", tempStr, tempStr2)
	return tempStr
}

func (s *SearchTree) Find(value int) (int, bool) {
	if s == nil {
		return 0, false
	}
	value = HashVal(value)
	return s.Root.find(value)

}

func (n *Node) find(value int) (int, bool) {
	if n == nil {
		return 0, false
	}

	switch {
	case n.Val == value:
		return value, true
	case value < n.Val:
		return n.Left.find(value)
	case value > n.Val:
		return n.Right.find(value)

	}
	return 0, false
}

// shaVal := sha256.New()
func HashVal(val int) int {
	h := fnv.New64()
	h.Write([]byte(fmt.Sprintf("%d", val)))
	return int(h.Sum64())
}

func (n *Node) Delete() {}
