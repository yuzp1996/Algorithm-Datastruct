package Recursion


func Ways(n int)int{
	if n==1{
		return 1
	}
	if n==2{
		return 2
	}
	return Ways(n-1)+Ways(n-2)
}

func WaysIteration(n int)int{
	if n == 1{
		return 1
	}
	if n==2{
		return 2
	}
	per := 1
	prepare := 2
	res := 0
	for i:=3;i<=n;i++{
		res = per + prepare
		per = prepare
		prepare = res
	}
	return res
}


func Fibonacci(n int)int{
	if n<0{
		return 0
	}
	if n == 0{
		return 0
	}
	if n == 1{
		return 1
	}
	result := Fibonacci(n-1)+Fibonacci(n-2)
	return result
}

func Factorial(n int)int{
	if n == 2{
		return 2
	}
	result := Factorial(n-1)*n
	return result
}


func Arrangement(num int)int{
	if num == 1{
		return 1
	}
	if num == 2{
		return 2
	}
	return Arrangement(num-1)*(num+1)
}