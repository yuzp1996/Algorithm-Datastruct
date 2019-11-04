package Trie

import (
	"log"
)

type TrieLeaf struct {
	Value   string
	Leafs   []*TrieLeaf
	Endchar bool
}

func NewTrieLeaf(val string, endchar bool) *TrieLeaf {
	return &TrieLeaf{
		Value:   val,
		Endchar: endchar,
		Leafs:   make([]*TrieLeaf, 26),
	}
}

func (leaf TrieLeaf) IsEnding() bool {
	return leaf.Endchar
}

func NewTrieTree() (root *TrieLeaf) {
	return &TrieLeaf{
		"Root",
		make([]*TrieLeaf, 26),
		true,
	}
}

func (TrieTree *TrieLeaf) Insert(keyword string) *TrieLeaf {
	Root := TrieTree
	for _, char := range keyword {
		if Root.Leafs[int(char)-97] == nil {
			Root.Leafs[int(char)-97] = NewTrieLeaf(string(char), false)
		}
		Root = Root.Leafs[int(char)-97]
	}
	Root.Endchar = true
	return TrieTree
}

func (TrieTree *TrieLeaf) Find(keyword string) bool {
	Root := TrieTree
	for _, char := range keyword {
		if Root.Leafs[int(char)-97] == nil {
			return false
		}
		if int32(Root.Leafs[int(char)-97].Value[0]) != char {
			return false
		}
		Root = Root.Leafs[int(char)-97]
	}
	return true
}

type PointTester struct {
	Name int
}

func TestArray(Array []PointTester) bool {
	resultlist := make(map[int]*PointTester)

	for index, value := range Array {
		//value := _value
		//最主要的是要意识到 这个value这个的指针 一直都是一个
		//这里的value指的是这个循环中value的地址 一直都是一个
		//但是如果加上一个value := _value  相当于把这个地址给换成另一个指针了 以后每次循环时 这个指针都会不同了 不会再使用循环中使用的那个指针了
		resultlist[index] = &value
		log.Printf("index is %v, value is %v and address is %v\n", index, value, value)
	}
	log.Printf("now \n next")
	for index, value := range resultlist {
		log.Printf("index is %v, value is %v and address is %v\n", index, value, &value)
	}

	log.Printf("resultlist is %v", resultlist)

	return true
}

func TestArrayPointer(Array []*PointTester) bool {
	resultlist := make(map[int]*PointTester)

	for index, value := range Array {
		//这里的value指向的是Array中某个数据的指针地址
		resultlist[index] = value
		log.Printf("index is %v, value is %v and address is %v\n", index, value, &value)
	}
	log.Printf("now \n next")

	for index, value := range resultlist {
		log.Printf("index is %v, value is %v and address is %v\n", index, value, &value)
	}

	log.Printf("resultlist is %v", resultlist)

	return true
}

//func TestPointerArray(Array *[]PointTester)bool{
//	resultlist := make(map[int]*PointTester)
//
//	for index, value := range *Array{
//		//value := _value
//		resultlist[index] = &value
//		log.Printf("index is %v, value is %v and address is %v\n", index,value,&value)
//	}
//	log.Printf("now \n next")
//
//	for index, value := range resultlist{
//		log.Printf("index is %v, value is %v and address is %v\n", index,value,&value)
//	}
//
//
//	log.Printf("resultlist is %v", resultlist)
//
//	return  true
//}

//https://segmentfault.com/a/1190000013739000#articleHeader5

func Test12(){


	// this is array this will not work
	TestArray := [3]int{1,2,3}
	log.Printf("before func TestArray it is %v",TestArray)

	func(arr [3]int){
		arr[0] = 0
		log.Printf("TestArray in func it is %v",arr)
	}(TestArray)

	log.Printf("out of ht func it is %v", TestArray)


	//this is array but use pointer it willl work
	log.Printf("before func TestArray it is %v",TestArray)

	func(arr *[3]int){
		arr[0] = 0
		log.Printf("TestArray in func it is %v",arr)
	}(&TestArray)

	log.Printf("out of ht func it is %v", TestArray)



	// this is slice it will work
	TestSlice := []int{1,2,3}
	log.Printf("before func TestArray it is %v",TestSlice)

	func(arr []int){
		arr[0] = 0
		log.Printf("TestArray in func it is %v",arr)
	}(TestSlice)

	log.Printf("out of ht func it is %v", TestSlice)


}

func Test14(){

}

func Test18(){
	TestString := "zpyutest"
	log.Printf("TestString[0] is %v", TestString[0])
	TestChinese := "于志鹏"
	log.Printf("TestChinese[0] is %v", TestChinese[0])

}


func Test31(){

}

func Test34(){

}

func Test45(){
	loop:
		for{
			switch {
			case true:
				log.Printf("breaking out ...")
				break loop
			}
		}

	loop1: for{
		switch {
		case true:
			log.Printf("breaking1 out ...")
			break loop1
		}
	}
	log.Printf("out loop...")

}

func Test47(){
	var i = 1
	i++
	defer log.Printf("result is %v", func()int {return i*2}())
	i++
}

func Test50(){

}

func Test51(){}

func Test53(){}