package Trie

type TrieLeaf struct {
	Value string
	Leafs []*TrieLeaf
	Endchar  bool
}

func NewTrieLeaf(val string, endchar bool)*TrieLeaf{
	return &TrieLeaf{
		Value:val,
		Endchar:endchar,
		Leafs:make([]*TrieLeaf,26),
	}
}

func (leaf TrieLeaf)IsEnding()bool{
	return leaf.Endchar
}


func NewTrieTree()(  root *TrieLeaf){
	return &TrieLeaf{
		"Root",
		make([]*TrieLeaf,26),
		true,
	}
}



func(TrieTree *TrieLeaf)Insert(keyword string)*TrieLeaf{
	Root := TrieTree
	for _, char := range keyword{
		if Root.Leafs[int(char)-97] == nil{
			Root.Leafs[int(char)-97] = NewTrieLeaf(string(char),false)
		}
		Root = Root.Leafs[int(char)-97]
	}
	Root.Endchar = true
	return TrieTree
}


func(TrieTree *TrieLeaf)Find(keyword string)bool{
	Root := TrieTree
	for _,char := range keyword{
		if Root.Leafs[int(char)-97] == nil{
			return false
		}
		if int32(Root.Leafs[int(char)-97].Value[0]) != char{
			return false
		}
		Root = Root.Leafs[int(char)-97]
	}
	return true
}