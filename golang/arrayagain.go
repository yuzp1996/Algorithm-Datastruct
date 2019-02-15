package main

import (
	"fmt"
	"github.com/go-errors/errors"
)

type ArrayList struct {
	data []int
	length int
}


func NewArray(cap int)*ArrayList{
	if cap <=0 {
		return nil
	}
	return &ArrayList{
		make([]int,cap,cap),
		0,
	}
}

func main() {
	newarray := NewArray(15)
	for i := 1; i <= 15; i++ {
		if err := newarray.Insert(i, i); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(newarray)

	newarray1 := NewArray(7)
	for i := 1; i <= 6; i++ {
		if err := newarray1.Insert(i, i+3); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(newarray,newarray1)
	CombineSortedSlice(newarray,newarray1)
}

func (this *ArrayList)isOutofIndex(index int)bool{
	if index<=0||index>cap(this.data){
		return true
	}
	return false
}

func (this *ArrayList)Insert(index,value int)error{
	if this.length==len(this.data){
		this.data = append(this.data,make([]int,1,1)...)
	}
	if this.isOutofIndex(index){
		return errors.New("index out of")
	}
	if this.data[index-1] == 0{
		this.data[index-1] = value
		this.length++
		return nil
	}
	for i:= this.length-1;i>=index-1;i--{
		this.data[i+1] = this.data[i]
	}
	this.data[index-1] = value
	this.length++
	return nil
}

func (this *ArrayList)Delete(index int)error{
	if this.isOutofIndex(index){
		return errors.New("delete out of index")
	}
	if this.data[index-1]==0{
		return errors.New("now val in the index")
	}
	for i:=index-1;i<this.length-1;i++{
		this.data[i] = this.data[i+1]
	}
	this.data[this.length-1] = 0
	this.length--
	return nil
}


func CombineSortedSlice(firarray, secarray *ArrayList){
	finalarray := NewArray(5)
	for{
		if firarray.length>0 && secarray.length>0{
			if firarray.data[0]>=secarray.data[0]{
				if err := finalarray.Insert(finalarray.length+1,secarray.data[0]); err!= nil{
					fmt.Printf("Insert err err is %v", err)
				}
				if err := secarray.Delete(1);err!=nil{
					fmt.Printf("Delete err err is %v", err)
				}
			}else{
				if err:= finalarray.Insert(finalarray.length+1,firarray.data[0]);err!=nil{
					fmt.Printf("Insert firarray err")
				}
				if err:=firarray.Delete(1);err!=nil{
					fmt.Printf("Delete firarry err")
				}
			}
		}else{
			if firarray.length==0{
				for{
					if secarray.length > 0{
						finalarray.Insert(finalarray.length+1,secarray.data[0])
						secarray.Delete(1)
					}else{
						break
					}
				}
			}else if secarray.length==0{
				for{
					if firarray.length > 0{
						finalarray.Insert(finalarray.length+1,firarray.data[0])
						firarray.Delete(1)
					}else{
						break
					}
				}
				}
			break
		}
	}
	fmt.Println(finalarray)
	good := false

	for i,v := range finalarray.data{
		if firarray.data[i+1] == 0{
			break
		}else if v <= finalarray.data[i+1]{
			continue
		}else{
			good = true
			break
		}
	}
	if !good{
		fmt.Printf("good algorithm")
	}else{
		fmt.Printf("not good algotrithm")
	}


}
