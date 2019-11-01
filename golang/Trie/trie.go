package Trie

import "log"

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

type PointTester struct {
	Name int
}


func TestArrayPointer(Array []*PointTester)bool{
	resultlist := make(map[int]*PointTester)

	for index, value := range Array{
		resultlist[index] = value
		log.Printf("index is %v, value is %v and address is %v\n", index,value,value)
	}
	log.Printf("now \n next")


	for index, value := range resultlist{
		log.Printf("index is %v, value is %v and address is %v\n", index,value,&value)
	}


	log.Printf("resultlist is %v", resultlist)

	return  true
}


func TestArray(Array []PointTester)bool{
	resultlist := make(map[int]*PointTester)

	for index, value := range Array{
		//value := _value
		resultlist[index] = &value
		log.Printf("index is %v, value is %v and address is %v\n", index,value,value)
	}
	log.Printf("now \n next")
	for index, value := range resultlist{
		log.Printf("index is %v, value is %v and address is %v\n", index,value,&value)
	}


	log.Printf("resultlist is %v", resultlist)

	return  true
}



func TestPointerArray(Array *[]PointTester)bool{
	resultlist := make(map[int]*PointTester)

	for index, value := range *Array{
		//value := _value
		resultlist[index] = &value
		log.Printf("index is %v, value is %v and address is %v\n", index,value,value)
	}
	log.Printf("now \n next")

	for index, value := range resultlist{
		log.Printf("index is %v, value is %v and address is %v\n", index,value,&value)
	}


	log.Printf("resultlist is %v", resultlist)

	return  true
}