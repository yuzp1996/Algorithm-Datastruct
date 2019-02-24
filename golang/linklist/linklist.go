package main

import "fmt"

type LinkNode struct {
	next *LinkNode
	value interface{}
}

type LinkList struct {
	head *LinkNode
	length int
}

func NewLinkNode(value interface{})*LinkNode{
	return &LinkNode{nil,value}
}

func NewLinkList()*LinkList{
	return &LinkList{NewLinkNode(0),0}
}

func (this *LinkNode)Getvalue()interface{}{
	return this.value
}

func (this *LinkNode)GetNext()*LinkNode{
	return this.next
}

func (this *LinkList)InsterAfter(p *LinkNode,value interface{}){
	newnode := NewLinkNode(value)
	if p.next == nil{
		p.next = newnode
	}else{
		pernext := p.next
		newnode.next = pernext
		p.next = newnode
	}
	this.length++
}


func(this *LinkList)PrintList(){
	per := this.head
	for{
		if per != nil{
			fmt.Println(per.value)
			per = per.next
		}else{
			return
		}
	}
}

func(this *LinkList)DeleteNode(deletenode *LinkNode){
	per := this.head
	//only one element
	if this.head == deletenode{
		this.head.next = nil
		this.head.value = 0
	}
	pernode := &LinkNode{}
	for{
		if per != nil{
			if per.next == deletenode{
				pernode = per
				break
			}else{
				per = per.next
			}
		}
	}
	if pernode.next!= nil{
		pernode.next = deletenode.next
	}
	return
}

func (this *LinkList)searchwithvalue(value interface{})*LinkNode{
	per := this.head
	for{
		if per != nil {
			if per.value == value {
				return per
			} else {
				per = per.next
			}
		}else{
			return nil
		}
	}
}


func main(){
	NewList := NewLinkList()
	NewList.InsterAfter(NewList.searchwithvalue(0),1)
	NewList.InsterAfter(NewList.searchwithvalue(1),2)
	NewList.InsterAfter(NewList.searchwithvalue(2),3)
	NewList.InsterAfter(NewList.searchwithvalue(1),4)
	NewList.DeleteNode(NewList.searchwithvalue(4))
	//NewList.DeleteNode(NewList.searchwithvalue(1))

	NewList.PrintList()
}






