package Linkedtree

import (
	"errors"
	"fmt"
)

type Leaf struct {
	Value int
	Right *Leaf
	Left  *Leaf
}

func NewTree(value int) *Leaf {
	return &Leaf{value, nil, nil}
}

func (Root *Leaf) TestMakeitnil() {
	//Root=nil
	fmt.Print("after nil")
	Root.Value = 111
	fmt.Print(Root)
	return
}

func (Root *Leaf) GetLeaf() int {
	if Root == nil || Root.Left == nil {
		return -1
	}
	return Root.Left.Value
}

func (Root *Leaf) GetRight() int {
	if Root == nil || Root.Right == nil {
		return -1
	}
	return Root.Right.Value
}

func (Root *Leaf) LeafAdd(value int) error {
	if Root.Left == nil {
		newleaf := Leaf{value, nil, nil}
		Root.Left = &newleaf
		return nil
	} else {
		return errors.New("Can not add element to this node's left leaf")
	}
}

func (Root *Leaf) RightAdd(value int) error {
	if Root.Right == nil {
		newleaf := Leaf{value, nil, nil}
		Root.Right = &newleaf
		return nil
	} else {
		return errors.New("Can not add element to this node's right leaf")
	}
}

func (Root *Leaf) Add(value int) {

}

func (Root *Leaf) Preorder(result []int) []int {
	if Root == nil {
		return result

	}
	result = append(result, Root.Value)
	result = Root.Left.Preorder(result)
	result = Root.Right.Preorder(result)
	return result
}

func (Root *Leaf) Middleorder(result []int) []int {
	if Root == nil {
		return result
	}
	result = Root.Left.Middleorder(result)
	result = append(result, Root.Value)
	result = Root.Right.Middleorder(result)
	return result
}

func (Root *Leaf) BehindOrder(result []int) []int {
	if Root == nil {
		return result
	}
	result = Root.Left.BehindOrder(result)
	result = Root.Right.BehindOrder(result)
	result = append(result, Root.Value)
	return result
}

func (Root *Leaf) LevelOrder(result []int) []int {
	if Root == nil {
		return result
	}
	Queue := []Leaf{}
	Queue = append(Queue, *Root)
	for len(Queue) > 0 {
		root := Queue[0]
		Queue = Queue[1:]
		result = append(result, root.Value)
		if root.Left != nil {
			Queue = append(Queue, *root.Left)
		}
		if root.Right != nil {
			Queue = append(Queue, *root.Right)
		}
	}
	return result
}

func (Root *Leaf) Level() []int {
	result := []int{}
	if Root == nil {
		return result
	}
	Queue := []*Leaf{}
	// 添加数据比较靠谱的方式
	Queue = append(Queue, Root)
	for len(Queue) > 0 {
		Root = Queue[0]
		Queue = Queue[1:]
		result = append(result, Root.Value)
		if Root.Left != nil {
			Queue = append(Queue, Root.Left)
		}
		if Root.Right != nil {
			Queue = append(Queue, Root.Right)
		}
	}
	return result
}
