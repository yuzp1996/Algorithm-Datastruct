package main

import "fmt"

type StackArray struct {
	data []interface{}
	top int
}

func NewStackArray()*StackArray{
	return &StackArray{
		data:make([]interface{},32,32),//这个是切片偶
		top:-1,
	}
}


func(this *StackArray)Top()interface{}{
	if this.IsEmpty(){
		fmt.Printf("Empty Stack")
		return nil
	}
	return this.data[this.top]
}

func(this *StackArray)Flush(){
	this.top=-1
}


func (this *StackArray)IsEmpty()(bool){
	if this.top<0{
		return true
	}else{
		return false
	}
}

func(this *StackArray)PrintStack(){
	if this.IsEmpty(){
		fmt.Println("empty stackarray")
	}else{
		for i := this.top;i>=0;i--{
			fmt.Printf("%v,",this.data[i])
		}
	}
}



func (this *StackArray)Push(v interface{}){
	this.top += 1
	this.data[this.top] = v
}

func (this *StackArray)Pop() (v interface{}){
	if this.IsEmpty(){
		fmt.Println("Empty Stack ....")
		return nil
	}
	v = this.data[this.top]
	this.top -= 1
	return v

}

func main(){
	newstack := NewStackArray()
	newstack.Push(1)
	fmt.Printf("pop is %v\n",newstack.Pop())
	newstack.PrintStack()
	newstack.Push(2)
	newstack.Push(3)
	newstack.Push(4)
	newstack.Push(5)
	newstack.Push(6)
	newstack.PrintStack()

	fmt.Printf("pop is %v\n",newstack.Pop())
	newstack.PrintStack()

	fmt.Printf("pop is %v\n",newstack.Pop())
	newstack.PrintStack()

	fmt.Printf("pop is %v\n",newstack.Pop())
	newstack.PrintStack()

	fmt.Printf("pop is %v\n",newstack.Pop())
	newstack.PrintStack()

	fmt.Printf("pop is %v\n",newstack.Pop())
	newstack.PrintStack()


}