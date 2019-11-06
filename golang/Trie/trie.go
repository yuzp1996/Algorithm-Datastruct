package Trie

import (
	"log"
	"sync"
	"time"
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
	x := 2
	y := 4
	table := make([][]int,x)
	for i := range table{
		table[i] = make([]int,y)
	}
	table[0][0] = 1

	log.Printf("table is %v\n", table)


	h,w := 2,4
	raw := make([]int, h*w)
	for i := range raw {
		raw[i] = i
	}
	log.Printf("raw is %v &raw is %v \n", raw,&raw[4])
	table = make([][]int,h)
	for i := range table{
		table[i] =  raw[i*w:i*w+w]
	}
	log.Printf("table is %v and &table[1][0] is %v",table, &table[1][0])

}

func Test18(){
	TestString := "zpyutest"
	log.Printf("TestString[0] is %v", TestString[0])
	TestChinese := "于志鹏"
	log.Printf("TestChinese[0] is %v", TestChinese[0])

}

//WaitGroup make sure all worker will be done
func Test31(){
	var wg sync.WaitGroup

	workcount := 2
	for i:=0;i<workcount;i++{
		wg.Add(1)
		go test31DoIt(i, &wg)
	}

	time.Sleep(1 * time.Second)
	wg.Wait()
	log.Printf("all done")


}

func test31DoIt(workId int, wg *sync.WaitGroup){
	log.Printf("[%v]start running... \n", workId)
	time.Sleep(3*time.Second)
	log.Printf("[%v]running done \n",workId)
	wg.Done()
}

// WaitGroup make sure all worker will be done
// but chan done is for stop a goroutine when it is not done because we want to exit all the sub program
func Test311(){
	done := make(chan struct{})
	var wg sync.WaitGroup
	ch := make(chan interface{})


	workcount := 2
	for i:=0;i<workcount;i++{
		wg.Add(1)
		go test311DoIt(i, done,ch, &wg)
	}
	for i := 0;i<workcount;i++{
		ch <-i
	}

	close(done)
	wg.Wait()
	close(ch)

	log.Printf("all done")
}

func test311DoIt(workId int, done <-chan struct{},ch <-chan interface{}, wg *sync.WaitGroup){
	log.Printf("[%v]start running... \n", workId)
	defer wg.Done()
	for{
		select {
		case m:= <-ch:
			log.Printf("[%v] m is %v \n ", workId,m)
		case <-done:
			log.Printf("[%v] is done",workId)
			return
		}
	}
}

func Test33(){
	done := make(chan struct{})
	ch := make(chan int)
	for i:= 0;i < 3;i++{
		go func(index int) {
			select {
			case ch <- (index+1)*2:
				log.Println(index, "Send Result")
			case <-done:
				log.Println(index," Exits")
			}
		}(i)
	}

	log.Println("Result: ",<-ch)
	close(done)
	time.Sleep(time.Second *2)
	log.Println("done")


}

//在一个值为 nil 的 channel 上发送和接收数据将永久阻塞
//向已关闭的 channel 发送数据会造成 panic
//两者不同
func Test34(){
	inch := make(chan int)
	outch := make(chan int)

	go func() {
		var in <-chan int = inch
		var out chan<- int
		var val int

		for{
			select {
			case out <- val:
				log.Println("-----")
				out = nil
				in = inch
			case val = <-in:
				log.Println("+++++")
				out = outch
				in = nil
			}
		}
	}()
	go func() {
		for r:= range outch{
			log.Println("Result: ", r)
		}
	}()
	inch <- 1
	inch <- 2
	time.Sleep(3*time.Second)


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

func Test53(){


	var data *PointTester
	log.Println("data is nil ", data == nil)

	// it will be nil only if it's type and val be nil
	var inter interface{}
	log.Println("inter is nil ", inter == nil)

	inter = data
	log.Println("after inter=data inter is nil", inter == nil)

}