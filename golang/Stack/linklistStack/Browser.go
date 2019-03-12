package main


func NewBrowser()*Browser{
	return &Browser{
		forword:NewLinkStack(),
		back:NewLinkStack(),
	}
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
	this.forword.push(v)
	this.back.Flush()
}

func(this *Browser)Forward(){
	if this.CanFord(){
		this.back.push(this.forword.top.value)
		this.forword.pop()
	}
}

func (this *Browser)Back(){
	if this.CanBack(){
		this.forword.push(this.back.top.value)
		this.back.pop()
	}
}

func (this *Browser)PrintBrowser(){
	this.back.Printstack()
	this.forword.Printstack()
}


type Browser struct {
	forword *LinkStack
	back *LinkStack
}


func main(){
	Browser := NewBrowser()
	Browser.Open("www.baidu.com")
	Browser.Open("www.google.com")
	Browser.Open("www.JD.com")
	Browser.Back()

}
