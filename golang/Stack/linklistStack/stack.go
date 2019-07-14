package linklistStack

import (
	"fmt"
)

type node struct {
	next  *node
	Value interface{}
}

// only need a top
type LinkStack struct {
	Top *node
}

func NewLinkStack() *LinkStack {
	return &LinkStack{
		Top: nil,
	}
}

func (this *LinkStack) IsEmpty() bool {
	if this.Top == nil {
		return true
	}
	return false
}

func (this *LinkStack) Push(v interface{}) {
	if this.IsEmpty() {
		this.Top = &node{
			next:  nil,
			Value: v,
		}
	} else {
		this.Top = &node{
			next:  this.Top,
			Value: v,
		}
	}
	return
}

func (this *LinkStack) Pop() {
	if this.IsEmpty() {
		fmt.Println("Empty Stack")
	} else {
		this.Top = this.Top.next
	}
}

func (this *LinkStack) Printstack() {
	Top := this.Top
	for {
		if this.Top != nil {
			fmt.Printf("%v -> ", this.Top.Value)
			this.Top = this.Top.next
		} else {
			break
		}
	}
	this.Top = Top
	fmt.Println()
}

func (this *LinkStack) Flush() {
	if this.IsEmpty() {
		return
	}
	this.Top = nil
}
