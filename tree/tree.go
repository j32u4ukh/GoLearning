package tree

import "fmt"

type Tree struct {
	Left  *Tree
	Value int32
	Right *Tree
}

func NewTree(value int32) *Tree {
	tree := &Tree{
		Left:  nil,
		Value: value,
		Right: nil,
	}

	return tree
}

func (t_ptr *Tree) AddLeft(value int32) *Tree {
	(*t_ptr).Left = NewTree(value)

	return (*t_ptr).Left
}

func (t_ptr *Tree) AddRight(value int32) *Tree {
	(*t_ptr).Right = NewTree(value)

	return (*t_ptr).Right
}

func (t Tree) String() string {
	var result string

	if t.Left != nil {
		result += fmt.Sprintf("%v,", t.Left.String())
	}

	if t.Left == nil && t.Right == nil {
		result += fmt.Sprintf("%d", t.Value)
	} else {
		result += fmt.Sprintf("%d,", t.Value)
	}

	if t.Right != nil {
		result += fmt.Sprintf("%v,", t.Right.String())
	}

	return result
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	defer close(ch)

	var walker func(t *Tree)
	walker = func(t *Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- int(t.Value)
		walker(t.Right)
	}
	walker(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 && !ok2 {
			break
		}
	}

	return true
}
