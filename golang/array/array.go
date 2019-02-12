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
	if index < 0 || index >= cap(this.data) {
		return true
	}
	return false
}

func (this *Array) Find(index int) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, fmt.Errorf("error happend %s", "out of range")
	}
	return this.data[index], nil
}
func (this *Array) Insert(index, value int)(error){
	if this.isIndexOutOfRange(index)||this.Len()==cap(this){
		return errors.New("index out of range")
	}
	for i:=this.length-1;i>=index-1;i--{
		this.data[i+1] = this.data[i]
	}
	this.data[index] = value
	this.length++
	return nil
}

func (this *Array) Delete(index int)(int, error){
	if this.isIndexOutOfRange(index){
		return 0,errors.New("index out the range")
	}
	v := this.data[index]
	for i:=index-1;i<this.length;i++{
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v,nil
}














