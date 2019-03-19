package arrayqueue

import "fmt"

type ArrayQueue struct {
	Cap int
	Head int
	Tail int
	Data []interface{}
}

func NewArrayQueue(Cap int)*ArrayQueue{
	return &ArrayQueue{
		Head:0,
		Tail:0,
		Cap:Cap,
		Data:make([]interface{},Cap,Cap),
	}
}

func (this *ArrayQueue)Canenqueue()bool{
	if this.Tail>=this.Cap{
		fmt.Println("Can Not Enqueue Now")
		return false
	}
	return true
}

func (this *ArrayQueue)Candequeue()bool{
	if this.Tail>this.Head{
		return true
	}
	fmt.Println("Can Not Dequeue Now")
	return false
}

func (this *ArrayQueue)Enqueue(value interface{}){
	if this.Canenqueue(){
		this.Data[this.Tail]=value
		this.Tail++
	}else{
		if this.Head > 0{
			// move the data now
			for index:=this.Head;index<this.Tail;index++ {
				this.Data[index-this.Head] = this.Data[index]
			}
			this.Tail -= this.Head
			this.Head = 0
			//后面这两步都是自己丢掉的
			this.Data[this.Tail] = value
			this.Tail++
		}else{
			fmt.Println("It is a Really Full Data Queue")
			return
		}
	}
}

func (this *ArrayQueue)Dequeue(){
	if this.Candequeue(){
		this.Data[this.Head] = nil
		this.Head++
	}
}

func (this *ArrayQueue)PrintQueue()[]interface{}{
	if this.Head < this.Tail{
		index := this.Head
		fmt.Printf("Equeue is ")
		for(index <= this.Tail){
			fmt.Printf(" %v",this.Data[index])
		}
		return this.Data
	}else{
		fmt.Println("Empty Queue, Can not print")
		return  nil
	}
}