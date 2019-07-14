package main

//
//import (
//	"fmt"
//	"rand"
//	"time"
//)
//
//const MaxLevel = 7
//
//const Probability = 0.25
//
//func randlevel() (level int) {
//	rand.Seed(time.Now().UnixNano())
//	for level = 1; rand.Float32() < Probability && level <= MaxLevel; level++ {
//
//	}
//	return
//}
//
//type node struct {
//	key     int
//	forward []*node
//}
//
//type skiplist struct {
//	head  *node
//	level int
//}
//
//func NewNode(key, level int) *node {
//	return &node{key: key, forward: make([]*node, level)}
//}
//
//func NewSkipList() *skiplist {
//	return &skiplist{head: NewNode(0, MaxLevel), level: 1}
//}
//
//func (list *skiplist) Insert(key int) {
//	current := list.head
//	update := make([]*node, MaxLevel)
//	//先一层一层的确定这个update的 应该是被哪些节点来指向的
//	for i := list.level - 1; i >= 0; i-- {
//		if current.forward[i] == nil || current.forward[i].key > key {
//			update[i] = current
//		} else {
//			for current.forward[i] != nil && current.forward[i].key < key {
//				current = current.forward[i]
//			}
//			update[i] = current
//		}
//	}
//	level := randlevel()
//	if level > list.level {
//		// 新的节点层数高于最高的层数的时候， 就要更新这个list的层数了
//		for i := list.level; i < level; i++ {
//			update[i] = list.head
//		}
//		list.level = level
//	}
//
//	//最后插入这个节点
//	node := NewNode(key, level)
//	for i := 0; i < level; i++ {
//		node.forward[i] = update[i].forward[i]
//		update[i].forward[i] = node
//	}
//}
//
//func (s *skiplist) Delete(key int) {
//	current := s.head
//	for i := s.level - 1; i >= 0; i-- {
//		for current.forward[i] != nil {
//			if current.forward[i].key == key {
//				tmp := current.forward[i]
//				current.forward[i] = tmp.forward[i]
//				tmp.forward[i] = nil
//			} else if current.forward[i].key > key {
//				break
//			} else {
//				current = current.forward[i]
//			}
//		}
//	}
//}
//
//func (s *skiplist) Search(key int) bool {
//	current := s.head
//	for i := s.level; i >= 0; i-- {
//		for current.forward[i] != nil {
//			if current.forward[i].key == key {
//				return true
//			} else if current.forward[i].key > key {
//				break
//			} else {
//				current = current.forward[i]
//			}
//		}
//	}
//	return false
//}
//
//func (s *skiplist) Print() {
//	fmt.Println()
//	for i := s.level - 1; i >= 0; i-- {
//		current := s.head
//		for current.forward[i] != nil {
//			fmt.Printf("%d ", current.forward[i].key)
//			current = current.forward[i]
//		}
//		fmt.Printf("**************** level %d \n", i+1)
//	}
//}
//
//func main() {
//	list := NewSkipList()
//	for i := 0; i < 20; i++ {
//		list.Insert(i)
//	}
//	list.Print()
//	list.Delete(2)
//	list.Print()
//	fmt.Println(list.Search(2))
//}
