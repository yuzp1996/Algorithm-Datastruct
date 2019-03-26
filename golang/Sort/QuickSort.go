package Sort

func Quick_Sort(array []int)[]int{
	if len(array)<=1{
		return array
	}
	middle := array[0]
	bigger := GetBigger(array,middle)
	smaller := GetSmaller(array,middle)
	Middle := FindSame(array,middle)

	FinnalArray := append(Quick_Sort(smaller),Middle...)
	FinnalArray = append(FinnalArray,Quick_Sort(bigger)...)
	return FinnalArray
}

func GetBigger(Array []int,middle int )[]int{
	Finnal := []int{}
	for _,value := range Array{
		if value > middle{
			Finnal = append(Finnal,value)
		}
	}
	return Finnal
}

func GetSmaller(Array []int,middle int )[]int{
	Finnal := []int{}
	for _,value := range Array{
		if value < middle{
			Finnal = append(Finnal,value)
		}
	}
	return Finnal
}
func FindSame(Array []int,middle int )[]int{
	Finnal := []int{}
	for _,value := range Array{
		if value == middle{
			Finnal = append(Finnal,value)
		}
	}
	return Finnal

}

// 2 3 4 2 3 4 56 74 2
// 2 3 3 4 2	4   56 74
// 2 2  < 3 < 3 4    < 4 <   56 74
// 2 2 3 3 4    4   56  74