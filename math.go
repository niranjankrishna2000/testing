package main

func Abs(a int)int{
	if a<0{
		return -a
	}
	return a
}
func Factorial(num int)int{
	fact:=1
	for num>0{
		fact=fact*num
		num--
	}
	return fact
}