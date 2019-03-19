package cyclequeue

import "fmt"

type CycleQueue struct {
	Cap int
	Head int
	Tail int
	Data []interface{}
}

func NewCycleQueue(cap int)*CycleQueue{
	return &CycleQueue{
		Cap:cap,
		Head:0,
		Tail:0,
		Data:make([]interface{},cap,cap),
	}
}
func(this *CycleQueue)Canenqueue()bool{
	if (this.Tail+1)%this.Cap==this.Head{
		return false
	}
	return true
}

func (this *CycleQueue)Candequeue()bool{
	if this.Tail == this.Head{
		return false
	}
	return true
}

func(this *CycleQueue)Enqueue(value interface{}){
	if this.Canenqueue(){
		this.Data[this.Tail] = value
		if this.Tail == this.Cap-1{
			this.Tail=0
			return
		}
		this.Tail++
	}
}

func(this *CycleQueue)Dequeue(){
	if this.Candequeue(){
		this.Data[this.Head] = nil
		//better way is this.head = (this.head+1)%this.cap
		if this.Head == this.Cap-1{
			this.Head=0
			return
		}
		this.Head++
	}
}

func (this *CycleQueue)PrintQueue()[]interface{} {
	for index := (this.Head + 1) % this.Cap; index < (1+this.Tail)%this.Cap; index = this.Head % this.Cap {
		fmt.Printf("%v ", this.Data[index])
	}
	return this.Data
}