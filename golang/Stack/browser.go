package main

import (
	. "Algorithm-Datastruct/golang/Stack/linklistStack"
	"fmt"
)

func NewBrowser()*Browser{
	return &Browser{
		forword:NewLinkStack(),
		back:NewLinkStack(),
	}
}
type Browser struct {
	forword *LinkStack
	back *LinkStack
}

func (this *Browser)CanFord()bool{
	if this.forword.IsEmpty(){
		return false
	}
	return true
}

func(this *Browser)CanBack()bool{
	if this.back.IsEmpty(){
		return false
	}
	return true
}

func(this *Browser)Open(v interface{}){
	this.forword.Push(v)
	this.back.Flush()
}

func(this *Browser)Forward(){
	if this.CanFord(){
		this.back.Push(this.forword.Top.Value)
		this.forword.Pop()
	}
}

func (this *Browser)Back(){
	if this.CanBack(){
		this.forword.Push(this.back.Top.Value)
		this.back.Pop()
	}
}

func (this *Browser)PrintBrowser(){
	fmt.Printf("back is \n")
	this.back.Printstack()
	fmt.Printf("forword is \n")
	this.forword.Printstack()
}

func main(){
	Browser := NewBrowser()
	Browser.Open("www.baidu.com")
	Browser.Open("www.google.com")
	Browser.Open("www.JD.com")
	Browser.Open("www.yuzp1996.com")
	Browser.PrintBrowser()

	Browser.Forward()

	Browser.PrintBrowser()
	Browser.Forward()

	Browser.PrintBrowser()
	Browser.Open("www.yahoo.com")
	Browser.PrintBrowser()

	fmt.Printf("\nNow it is %v", Browser.forword.Top.Value)

}
