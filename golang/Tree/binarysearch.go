package Linkedtree

import (
	"fmt"
)

func (Tree *Leaf)Insert(value int){
	for{
		if Tree != nil {
			if value > Tree.Value{
				if Tree.Right == nil{
					Tree.Right = NewTree(value)
					return
				}
				Tree = Tree.Right
			}else if value < Tree.Value{
				if Tree.Left == nil{
					Tree.Left = NewTree(value)
					return
				}
				Tree = Tree.Left
			}else if value == Tree.Value{
				return
			}
		}
	}
}


func (Tree *Leaf)Search(value int)bool{
	for{
		if Tree == nil{
			return false
		}
		if Tree.Value == value{
			return true
		}
		if value > Tree.Value{
			Tree = Tree.Right
		}else {
			Tree = Tree.Left
		}
	}
}



func (Tree *Leaf)DeleteLeaf(value int)int{
	Root := Tree
	ParentRoot := NewTree(-1)
	// find the leaf and it's parent leaf
	for{
		if Root == nil{
			fmt.Printf("no such leaf")
			return -1
		}else if value > Root.Value{
			ParentRoot = Root
			Root = Root.Right
		}else if value < Root.Value{
			ParentRoot = Root
			Root = Root.Left
		}else if value == Root.Value{
			break
		}
	}

    //经过这个的处理 之后只有没有或者单个leaf的情况了
	if Root.Right != nil && Root.Left != nil{
		pleaf := Root
		leaf := Root.Right
		for{
			if leaf.Left != nil{
				pleaf = leaf
				leaf =leaf.Left
			}else{
				break
			}
		}
		Root.Value, leaf.Value = leaf.Value, Root.Value
		Root = leaf
		ParentRoot = pleaf
	}

	// 经过上面的处理 现在如果paroot还是-1的话 那肯定是只有
	// 1、删除根节点
	// 2、根节点只有一个或者没有子树
	if ParentRoot.Value == -1{
		value := Root.Value
		if Root.Left != nil{
			*Tree = *(Root.Left)
		}else if Root.Right != nil{
			*Tree = *(Root.Right)
		}else {
			Tree = nil
		}
		return value
	}
	// just one leaf or no leaf
	// no leaf
	// 现在肯定是 最多一个leaf了
	child := NewTree(-1)
	if Root.Left != nil{
		child = Root.Left
	}else if Root.Right != nil{
		child = Root.Right
	}else{
		child = nil
	}

	if ParentRoot.Right == Root{
		ParentRoot.Right = child
	}else if ParentRoot.Left == Root{
		ParentRoot.Left = child
	}
	return Root.Value
}

































