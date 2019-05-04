package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxLevel = 7

const Probability = 0.25

func randlevel() (level int) {
	rand.Seed(time.Now().UnixNano())
	for level = 1; rand.Float32() < Probability && level <= MaxLevel; level++ {

	}
	return
}

type node struct {
	key     int
	forward []*node
}

type skiplist struct {
	head  *node
	level int
}

func NewNode(key, level int) *node {
	return &node{key: key, forward: make([]*node, level)}
}

func NewSkipList() *skiplist {
	return &skiplist{head: NewNode(0, MaxLevel), level: 1}
}

func (list *skiplist) Insert(key int) {
	current := list.head
	update := make([]*node, MaxLevel)
	//先一层一层的确定这个update的 应该是被哪些节点来指向的
	for i := list.level - 1; i >= 0; i-- {
		if current.forward[i] == nil || current.forward[i].key > key{
			update[i] = current
		} else {
			for current.forward[i] != nil && current.forward[i].key < key {
				current = current.forward[i]
			}
			update[i] = current
		}
	}
	level := randlevel()
	if level > list.level {
		// 新的节点层数高于最高的层数的时候， 就要更新这个list的层数了
		for i := list.level; i < level; i++ {
			update[i] = list.head
		}
		list.level = level
	}

	node := NewNode(key, level)
	for i := 0; i < level; i++ {
		node.forward[i] = update[i].forward[i]
		update[i].forward[i] = node
	}
}

func (s *skiplist) Print() {
	fmt.Println()
	for i := s.level - 1; i >= 0; i-- {
		current := s.head
		for current.forward[i] != nil {
			fmt.Printf("%d ", current.forward[i].key)
			current = current.forward[i]
		}
		fmt.Printf("**************** level %d \n", i+1)
	}
}

func main() {
	list := NewSkipList()
	for i := 0; i < 200; i++ {
		list.Insert(rand.Intn(100))
	}
	list.Print()
}
