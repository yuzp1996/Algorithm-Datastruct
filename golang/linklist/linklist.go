package main

import "fmt"

type LinkNode struct {
	next *LinkNode
	value int
}

type LinkList struct {
	head *LinkNode
	length int
}

func NewLinkNode(value int)*LinkNode{
	return &LinkNode{nil,value}
}

func NewLinkList()*LinkList{
	return &LinkList{NewLinkNode(0),0}
}

func (this *LinkNode)Getvalue()int{
	return this.value
}

func (this *LinkNode)GetNext()*LinkNode{
	return this.next
}





func (this *LinkList)InsterAfter(p *LinkNode,value int){
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
			fmt.Printf("%d -> ",per.value)
			per = per.next
		}else{
			break
		}
	}
	fmt.Println()
}

func(this *LinkList)DeleteNode(deletenode *LinkNode){
	per := this.head
	//only one element
	if this.head == deletenode{
		this.head.next = nil
		this.head.value = 0
		this.length--
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
		this.length--
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

func (this *LinkList)InsterInhead(value int){
	newnode := NewLinkNode(value)
	pernext := this.head
	newnode.next = pernext
	this.head = newnode
	this.length++
}


func (this *LinkList)InsterInTail(value int){
	newnode := NewLinkNode(value)
	per := this.head
	for{
		if per.next != nil{
			per = per.next
		}else{
			per.next = newnode
			this.length++
			return
		}
	}
}

//在index的下面加上这个
func(this *LinkList)InsertAtIndex(index int, value int){
	if index > this.length || index<0 {
		fmt.Printf("wrong index when insert at index %d  length is %d\n",index, this.length)
		return
	}
	newnode := NewLinkNode(value)
	per := this.head
	for i:=0;i<index;i++{
		per = per.next
	}
	beforeindex := per
	afterindex := per.next
	newnode.next = afterindex
	beforeindex.next = newnode
	this.length++
}

func(this *LinkList)Deletehead(){
	newhead := this.head.next
	this.head = newhead
	this.length--
}

//    1 -> 0 -> 5
//head     1    2
//index         2


func(this *LinkList)findfianlNode()(*LinkNode){
	per := this.head
	for{
		if per.next != nil{
			per = per.next
		}else{
			return per
		}
	}
}

func main(){
	NewList := NewLinkList()
	fmt.Printf("len is %d\n",NewList.length)
	//NewList.InsterAfter(NewList.searchwithvalue(0),1)
	//NewList.InsterAfter(NewList.searchwithvalue(1),2)
	//NewList.InsterAfter(NewList.searchwithvalue(2),3)
	//NewList.InsterAfter(NewList.searchwithvalue(1),4)
	//NewList.DeleteNode(NewList.searchwithvalue(4))
	//NewList.DeleteNode(NewList.searchwithvalue(1))
	//NewList.InsterInTail(4)
	//NewList.InsterInTail(5)
	//NewList.InsterInhead(1)
	NewList.InsterInTail(1)
	NewList.InsterInTail(2)
	NewList.InsterInTail(3)


	NewList.PrintList()

	SecondList := NewLinkList()
	SecondList.InsterInTail(4)
	SecondList.InsterInTail(5)
	SecondList.InsterInTail(6)
	SecondList.PrintList()
	FinalList := MergeLinkList(NewList,SecondList)
	FinalList.PrintList()




}
func MergeLinkList(firstlink *LinkList, secondlink *LinkList)(finallink *LinkList) {
	finallink = NewLinkList()
	for {
		if firstlink.length != 0 && secondlink.length != 0 {
			if firstlink.head.Getvalue() <= secondlink.head.Getvalue() {
				finallink.InsertAtIndex(finallink.length, firstlink.head.Getvalue())
				firstlink.Deletehead()
			} else {
				finallink.InsertAtIndex(finallink.length, secondlink.head.Getvalue())
				secondlink.Deletehead()
			}
		} else {
			break
		}
	}
	if firstlink.length != 0 {
		finallink.findfianlNode().next = firstlink.head
	} else {
		finallink.findfianlNode().next = secondlink.head
	}
	return finallink
}




//正常的应该就是头部插入 尾部插入  按照索引插入


