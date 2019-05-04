package main

import (
	"math/rand"
	"time"
)

const MaxLevel  = 7

const Probability = 0.25

func randlevel(level int){
	rand.Seed(time.Now().UnixNano())
	for level = 1; rand.Float32()<Probability && level <= MaxLevel; level++{

	}
	return
}



type node struct {
	key int
	forward []*node
}

type skiplist struct {
	head *node
	level int
}

func NewNode(key, level int)*node{
	return &node{key:key,forward:make([]*node, level)}
}

func NewSkipList()*skiplist{
	return &skiplist{head:NewNode(0, MaxLevel),level:1}
}


