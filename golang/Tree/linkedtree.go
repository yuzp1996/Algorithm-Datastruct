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

func(Root *Leaf)Preorder(){
	if Root == nil{
		fmt.Print("nil")
	}else{
		fmt.Print(Root.Value)
	}
	Root.Left.Preorder()
	Root.Right.Preorder()
}