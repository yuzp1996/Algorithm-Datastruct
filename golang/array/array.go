package main

import (
"fmt"
"errors"
)

type Array struct {
	data   []int
	length int
}

func NewArray(capcity int) *Array {
	if capcity <= 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capcity, capcity),
		length: 0,
	}
}

func (this *Array) Len() int {
	return this.length
}

func (this *Array) isIndexOutOfRange(index int) bool {
	if index <= 0 || index > cap(this.data) {
		return true
	}
	return false
}

func (this *Array) Find(index int) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, fmt.Errorf("error happend %s", "out of range")
	}
	return this.data[index-1], nil
}
func (this *Array) Insert(index, value int)(error){
	if this.Len() == cap(this.data){
		fmt.Println("full array")
		return errors.New("full array")
	}
	if this.isIndexOutOfRange(index){
		fmt.Println("index out of range")
		return errors.New("index out of range")
	}
	if this.data[index-1] == 0{
		this.data[index-1] = value
		this.length++
		return nil
	}
	for i:=this.Len();i>=index;i--{
		this.data[i] = this.data[i-1]
	}
	this.data[index-1] = value
	this.length++
	return nil
}

func (this *Array) Delete(index int)(int, error){
	if this.isIndexOutOfRange(index){
		return 0,errors.New("index out the range")
	}
	v := this.data[index-1]
	for i:=index-1;i<this.length-1;i++{
		this.data[i] = this.data[i+1]
	}
	this.data[this.length-1]=0
	this.length--
	return v,nil
}



func main(){
    newArray := NewArray(6)
    for i:=1;i<=6;i++{
    	 if err:= newArray.Insert(i,i); err != nil{
    	 	 fmt.Println("err is ", err)
			 return
		 }
	}
    bignewarray := NewArray(6)
    fmt.Println(bignewarray)
	for i:=1;i<=6;i++{
		if err:= bignewarray.Insert(i,i-11); err != nil{
			fmt.Println("err is ", err)
			return
		}
	}

	fianalArray := NewArray(20)

	if newArray.Len()==0{
		fianalArray.data=bignewarray.data
	}else if bignewarray.Len()==0{
		fianalArray.data=newArray.data
	}

	fmt.Println(newArray)
	fmt.Println(bignewarray)
	newindex := 0
	bigindex := 0
	for{
		if newArray.Len()!=0 && bignewarray.Len()!=0{
			if newArray.data[newindex]<=bignewarray.data[bigindex]{
				_ = fianalArray.Insert(fianalArray.length+1,newArray.data[0])
				_,_ = newArray.Delete(1)
			}else{
				_ = fianalArray.Insert(fianalArray.length+1,bignewarray.data[0])
				_,_ = bignewarray.Delete(1)
			}
		}else{
			if newArray.Len()==0{
				for i:=0;i<bignewarray.length;i++{
					_=fianalArray.Insert(fianalArray.length+1,bignewarray.data[i])
				}
			}else if bignewarray.Len()==0{
				for i:=0;i<newArray.length;i++{
					_=fianalArray.Insert(fianalArray.length+1,newArray.data[i])
				}
			}
			break
		}
	}
	fmt.Printf("finalArray is %v",fianalArray)
}





//你假设用户输入的 和 计算机里面的索引是不一样的
//所以0的输入应该是不对的 用户不应该有0的输入 最小是1
//最好是用数据手写一下 试试看
//如果insert加一了  那就应该进行一次length++
//自己手写 也是很重要的

//insert的时候 应该是用户输入 没有0  但是用length的话 有0 所以应该是 length+1




