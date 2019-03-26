package BinarySearch

// if exist,return the first index
// if not, return 0
func Binarysearch(searcharray []int,num int)(bool){
	last := len(searcharray)-1
	per := 0
	if num < searcharray[0]||num>searcharray[last]{
		return false
	}
	for last-per>1{
		if searcharray[(per+last)/2]>num{
			last = (per+last)/2
		}else if searcharray[(per+last)/2]<num{
			per = (per+last)/2
		}else{
			return true
		}
	}
	return false
}
