package main

import (
	"fmt"	
	"math"
)


func main(){
	//十进制
	i:=10
	fmt.Printf("%d\n",i)
	//二进制
	fmt.Printf("%b\n",i)
	//八进制
	m:=077
	fmt.Printf("%d\n",m)
	fmt.Printf("%o\n",m)
	//十六进制
	n:=0xff
	fmt.Printf("%d\n",n)
	fmt.Printf("%o\n",n)
	fmt.Printf("%x\n",n)

	var age uint8
	fmt.Println(age)

	//溢出会编译报错
	//age=256

	//浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	

	//复数
	var c1 complex64
	c1=1+2i
	var c2 complex128
	c2=3+4i
	fmt.Println(c1)
	fmt.Println(c2)

	//不可以给布尔赋值整数或者浮点，不会发生强转
	//布尔无法参与数值转化，也不能与其他类型数值转换
	//var isTrue bool=15
	//fmt.Println(isTrue)
	fmt.Println(false)


	//字符串
	s1:="string UTF=8"
	fmt.Println(s1)

	//多行字符串用反引号
	s2:=`
呃呃呃
	多行
	hhhhhhhhhhh`
	fmt.Println(s2)
}