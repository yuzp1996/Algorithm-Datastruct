package main

import (
	"fmt"
)

type node struct {
	next *node
	value interface{}
}

type LinkStack struct {
	top *node
}


func NewLinkStack()*LinkStack{
	return &LinkStack{
		top:nil,
	}
}


func(this *LinkStack)IsEmpty()bool{
	if this.top == nil{
		return true
	}
	return false
}


func(this *LinkStack)push(v interface{}){
	if this.IsEmpty(){
		this.top = &node{
			next:nil,
			value:v,
		}
	}else{
		this.top = &node{
			next:this.top,
			value:v,
		}
	}
	return
}

func(this *LinkStack)pop(){
	if this.IsEmpty(){
		fmt.Println("Empty Stack")
	}else{
		this.top = this.top.next
	}
}

func(this *LinkStack)Printstack(){
	top := this.top
	for{
		if this.top != nil{
			fmt.Printf("%v -> ", this.top.value)
			this.top = this.top.next
		}else{
			break
		}
	}
	this.top = top
	return
}



func(this *LinkStack)Flush(){
	if this.IsEmpty(){
		return
	}
	this.top = nil
}


func main(){
	stack := NewLinkStack()
	fmt.Printf("1 in, ")
	stack.push(1)
	fmt.Printf("second in, ")

	stack.push("second")
	fmt.Printf("third in, \n")

	stack.push("third")
	stack.Printstack()
}