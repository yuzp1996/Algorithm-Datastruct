package BF

import "fmt"

func BFfindstring(Main, Pattern string)bool{
	patternlength := len(Pattern)
	mainlength := len(Main)
	for i:=0;i<=mainlength-patternlength;i++{
		if Main[i:i+patternlength] == Pattern{
			return true
		}
	}
	fmt.Println(constructdict())
	return false
}

//func RKfindstring(Main, Pattern string)bool{
//	patternlength := len(Pattern)
//	mainlength := len(Main)
//}

func constructdict()map[string]int{
	letterarray := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	letterdict := make(map[string]int)
	for index,value := range letterarray{
		letterdict[value] = index
	}
	fmt.Printf("dict is %v\n",letterdict)
	return letterdict

}