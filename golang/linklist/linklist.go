package main

import (
	"fmt"
)

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

func NewLinkList(headvalue int)*LinkList{
	return &LinkList{NewLinkNode(headvalue),1}
}

func (this *LinkNode)Getvalue()int{
	return this.value
}

func (this *LinkNode)GetNext()*LinkNode{
	return this.next
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



func (this *LinkList)searchwithvalue(value interface{})(*LinkNode,int){
	per := this.head
	index := 1
	for{
		if per != nil {
			if per.value == value {
				return per,index
			} else {
				per = per.next
				index++
			}
		}else{
			return nil,0
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

func(this *LinkList)FindwithIndex(index int)*LinkNode{
	if index<1||index>this.length{
		fmt.Println("index unexpected")
		return nil
	}
	head := this.head
	for i:=1;i<index;i++{
		head = head.next
	}
	return head
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

func(this *LinkList)DeleteHead(){
	if this.length>=1{
		newhead := this.head.next
		this.head = newhead
		this.length--
	}
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
	Merge()
	LRU()

}

func Merge(){
	NewList := NewLinkList(1)
	NewList.InsterInTail(11)
	NewList.InsterInTail(23)
	NewList.InsterInTail(34)

	SecondList := NewLinkList(2)
	SecondList.InsterInTail(21)
	SecondList.InsterInTail(43)
	SecondList.InsterInTail(54)


	FinalList := MergeLinkList(NewList,SecondList)
	FinalList.PrintList()
}

func LRU(){
	fmt.Println("Begin LRU")
	linklistforLRU := NewLinkList(11)
	linklistforLRU.PrintList()
	for i:=0; i<10;i++{
		linklistforLRU.LRU(i)
		linklistforLRU.PrintList()
	}
	linklistforLRU.LRU(1)
	linklistforLRU.PrintList()
	linklistforLRU.LRU(6)
	linklistforLRU.PrintList()

	middlenode := linklistforLRU.FindMiddlenode()
	fmt.Println(middlenode.Getvalue())
}


func MergeLinkList(firstlink *LinkList, secondlink *LinkList)(finallink *LinkList) {
	finallink = NewLinkList(0)

	firstlink.PrintList()
	secondlink.PrintList()
	finallink.PrintList()

	for {
		if firstlink.length != 0 && secondlink.length != 0{
			if firstlink.head.Getvalue() <= secondlink.head.Getvalue() {
				finallink.InsertAtIndex(finallink.length-1, firstlink.head.Getvalue())
				firstlink.DeleteHead()
			} else {
				finallink.InsertAtIndex(finallink.length-1, secondlink.head.Getvalue())
				secondlink.DeleteHead()
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


func (this *LinkList)LRU(value int){
	//三种情况  有，删放到头部即可、 没有，没满直接放在头部、 没有，满了 去尾头增
	capoflru := 5
	//有
	if node,index := this.searchwithvalue(value);node!=nil&&index!=0{
		if index==1{
			return
		}else{
			pre:=this.FindwithIndex(index-1)
			nextnode := this.FindwithIndex(index+1)
			pre.next=nextnode
			this.InsterInhead(value)
		}
	}else{
		//没有
		if this.length<capoflru{
			this.InsterInhead(value)
		}else{
			finalnode := this.FindwithIndex(capoflru-1)
			finalnode.next=nil
			this.InsterInhead(value)
		}

	}
}


func(this *LinkList)FindMiddlenode()*LinkNode{
	if this.length<=2{
		return this.head
	}
	fast, slow := this.head, this.head
	//fast要首先保证不为nil  可以试想  只有在fast在最后一个元素或者最后一个元素的next(nil)的时候  才应该结束
	for fast!=nil && fast.next!=nil{
		fast = fast.next.next
		slow = slow.next
	}
	return slow

}


//正常的应该就是头部插入 尾部插入  按照索引插入
//https://blog.csdn.net/Charliewolf/article/details/82622014
//https://blog.csdn.net/Charliewolf/article/details/82687642
func (this *LinkList)reserve(){
	this.InsterInhead(1)
}