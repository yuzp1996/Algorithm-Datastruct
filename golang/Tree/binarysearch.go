package Linkedtree

import "fmt"

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

func (Tree *Leaf)Delete(value int)int{

	leaf := Tree
	parentleaf := NewTree(-1)
	// 找到节点
	for {
		if Tree == nil {
			fmt.Print("There is no such node value as %d", value)
			return -1
		}
		if Tree.Value == value {
			parentleaf = parentleaf
			leaf = Tree
			break
		}
		if value > Tree.Value {
			parentleaf = Tree
			Tree = Tree.Right
		} else {
			parentleaf = Tree
			Tree = Tree.Left
		}
	}
	if parentleaf.Value == -1{
		parentleaf = nil
	}



	// 处理子节点有两个节点的  最终变成parentleaf 和 leaf  使得这中情况的处理方法和其他的一个或者没有子节点的处理方法一致

	if leaf.Right != nil && leaf.Left != nil{
		DeletePTree := leaf
		DeleteTree := leaf.Right
		for {
			if(DeleteTree.Left != nil){
				DeletePTree = DeleteTree
				DeleteTree = DeleteTree.Left
			}
			break
		}
		leaf.Value = DeleteTree.Value
		parentleaf = DeletePTree
		leaf = DeleteTree
	}

	//上面已经处理成子节点了 也就是只有一个 或者没有节点的情况了

	// 把需要被删除节点的 子节点 都挂到child上  这里其实还剩只有一个节点需要处理的情况
	child := NewTree(-1)
	if leaf.Left != nil{
		child = leaf.Left
	}else if leaf.Right != nil{
		child= leaf.Right
	}
	if child.Value == -1{
		child = nil
	}

	// 把child放到正确的位置上
	if parentleaf == nil{  // 处理删除根节点
		Tree = child
		if Tree == nil{
			fmt.Print("\n Root can not be deleted ")
		}
		return value
	}else if parentleaf.Left == leaf{
		parentleaf.Left = child
		return value
	}else{
		parentleaf.Right = child
		return value
	}
}

//有两个节点的 其实就比有一个节点的多出了一步  那就是把值给赋过去  最终就是变成最小值的叶子的删除


































