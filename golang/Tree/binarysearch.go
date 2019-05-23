package Linkedtree

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
