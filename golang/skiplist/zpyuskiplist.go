package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxListLevel = 3
	Probability  = 0.75
)

type node struct {
	value   int
	forword []*node
}

type skiplist struct {
	head  *node
	level int
}

func NewNode(value, level int) *node {
	return &node{value, make([]*node, level)}
}

func NewSkipList() *skiplist {
	head := NewNode(0, MaxListLevel)
	return &skiplist{head, 1}
}

func RandLevel() (level int) {
	rand.Seed(time.Now().UnixNano())
	for level = 1; level < MaxListLevel && rand.Float32() > Probability; level++ {

	}
	return level
}

func (skiplist *skiplist) Insert(insertvalue int) {
	updateforward := make([]*node, MaxListLevel)
	current := skiplist.head
	// 先把updateforward找好
	for level := skiplist.level - 1; level >= 0; level-- { //这里的条件处理的有问题
		if current.forword[level] == nil || current.forword[level].value > insertvalue { // something wrong
			updateforward[level] = current
		} else {
			for current.forword[level] != nil && current.forword[level].value < insertvalue {
				current = current.forword[level]
			}
			updateforward[level] = current
		}
	}
	// 之前确定的位置 现在要确定层数了
	newlevel := RandLevel()

	if newlevel > skiplist.level {
		for level := skiplist.level; level < newlevel; level++ {
			updateforward[level] = skiplist.head
		}
		skiplist.level = newlevel
	}

	// 插入该节点
	newnode := NewNode(insertvalue, newlevel)
	for level := 0; level < newlevel; level++ {
		newnode.forword[level] = updateforward[level].forword[level]
		updateforward[level].forword[level] = newnode
	}

}

func (skiplist *skiplist) Print() {
	for level := skiplist.level - 1; level >= 0; level-- { //条件处理的有问题
		current := skiplist.head
		for {
			if current != nil {
				fmt.Printf("%d ", current.value)
				current = current.forword[level]
			} else {
				break
			}
		}
		fmt.Printf("***************   level %d \n", level+1)
	}
}

func (skiplist *skiplist) Search(key int) bool {
	current := skiplist.head
	for level := skiplist.level - 1; level >= 0; level-- {
		for current.forword[level] != nil {
			if current.forword[level].value == key {
				return true
			} else if current.forword[level].value > key {
				break
			} else {
				current = current.forword[level]
			}
		}
	}
	return false
}

func (skiplist *skiplist) Delete(key int) {
	current := skiplist.head
	for level := skiplist.level - 1; level >= 0; level-- {
		for current.forword[level] != nil {
			if current.forword[level].value == key {
				current.forword[level] = current.forword[level].forword[level]
			} else if current.forword[level].value > key {
				break
			} else {
				current = current.forword[level]
			}
		}
	}
}

func main() {
	skiplist := NewSkipList()
	for i := 1; i <= 20; i++ {
		skiplist.Insert(i)
	}

	skiplist.Print()
	fmt.Printf("find the key Y/N?   %t\n", skiplist.Search(12))
	skiplist.Delete(13)
	skiplist.Print()
	fmt.Printf("find the key Y/N?   %t", skiplist.Search(13))
}
