package Sort

func QuickSort(array []int)[]int{
	if len(array)<=1{
		return array
	}
	middle := array[0]
	bigger := GetBigger(array,middle)
	smaller := GetSmaller(array,middle)
	Middle := FindSame(array,middle)

	FinnalArray := append(QuickSort(smaller),Middle...)
	FinnalArray = append(FinnalArray,QuickSort(bigger)...)
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