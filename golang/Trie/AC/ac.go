package AC

import (
	. "Algorithm-Datastruct/golang/Queue/arrayqueue"
	"log"
)

type ACNode struct {
	Value   rune
	Leafs   []*ACNode
	FailurePointer *ACNode
	Endchar bool
	Length int
}

func NewACNode(val rune, endchar bool) *ACNode {
	return &ACNode{
		Value:   val,
		Endchar: endchar,
		Leafs:   make([]*ACNode, 26),
		FailurePointer:nil,
	}
}

func (leaf ACNode) IsEnding() bool {
	return leaf.Endchar
}

func NewACNodeTree() (root *ACNode) {
	return &ACNode{
		Value:'&',
		Leafs:make([]*ACNode, 26),
		FailurePointer: nil,
		Endchar:true,
	}
}

func (ACNode *ACNode) Insert(keyword string) *ACNode {
	Root := ACNode
	Root.Endchar = false
	var length int
	for _, char := range keyword {
		if Root.Leafs[int(char)-97] == nil {
			Root.Leafs[int(char)-97] = NewACNode(char, false)
		}
		Root = Root.Leafs[int(char)-97]
		length++
	}
	Root.Length = length
	Root.Endchar = true
	return ACNode
}

func(Root *ACNode)BuildFailurePointer(){

	queue := NewArrayQueue(100)
	queue.Enqueue(Root)

	for(queue.Candequeue()){
		p := queue.Dequeue().(*ACNode)
		for i:=0; i<26; i++{
			pc := p.Leafs[i]
			if pc == nil{
				continue
			}
			q := p.FailurePointer
			for q != nil{
				num := pc.Value
				if q.Leafs[num - 97]!=nil{
					pc.FailurePointer = q.Leafs[num-97]
					break
				}
				q = q.FailurePointer
			}
			if q == nil{
				pc.FailurePointer = Root
			}
			queue.Enqueue(pc)
		}
	}

}


func(ACTree *ACNode)FindString(text string)[]string{
	result := []string{}
	head := ACTree
	root := ACTree
	for index,value := range text{
		if root.Leafs[value-97] != nil{
			root = root.Leafs[value-97]
			if root.IsEnding(){
				log.Printf("the matched is %v",text[index-root.Length+1:index+1])
				result = append(result,text[index-root.Length+1:index+1])
			}
		}else{
			root = root.FailurePointer
		}
		if root == nil{
			root = head
		}
	}
	return result
}



//for(queue.Candequeue()){
//	p := queue.Dequeue().(ACNode)
//	//每次都是下一层的值 需要放到queue中
//	for i:=0;i<26;i++{
//		//我们是一直在找pc的失败指针qc
//		pc := p.Leafs[i]
//		if pc == nil{
//			continue
//		}
//		// 这里是在找q
//		q := p.FailurePointer
//		for q != nil{
//			num, _:=strconv.Atoi(pc.Value)
//			if q.Leafs[num-97] != nil{
//				// find it, end this, start next
//				pc.FailurePointer = q.Leafs[num-97]
//				break
//			}
//			q = q.FailurePointer
//		}
//		if q == nil{
//			pc.FailurePointer = Root
//		}
//
//		queue.Enqueue(pc)
//	}
//}

