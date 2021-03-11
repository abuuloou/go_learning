package main

import (
	"fmt"	
	"strings"
)

func main(){
	//字符串
	s11:="string UTF=8"
	fmt.Println(s11)


	//多行字符串用反引号
	s22:=`
呃呃呃
	多行
	hhhhhhhhhhh`
	fmt.Println(s22)


	//求字符串长度
	s:="hello"
	fmt.Println(len(s))

	s2 :="杨超越"
	fmt.Println(len(s2))

	//拼接字符串
	fmt.Println(s+s2)
	s3 := fmt.Sprintf("%s+=%s",s,s2)
	fmt.Println(s3)

	//切割
	s4 := "how do you do"
	fmt.Println(strings.Split(s4," "))
	//字符串切片
	fmt.Printf("%T\n",strings.Split(s4," "))

	//是否包含
	fmt.Println(strings.Contains(s4,"do"))

	fmt.Println(strings.HasPrefix(s4,"h"))

	fmt.Println(strings.HasSuffix(s4,"d"))

	fmt.Println(len(s4))

	fmt.Println(strings.Index(s4,"o"))

	fmt.Println(strings.LastIndex(s4,"o"))


	//join操作，把字符串你切片用什么字符连接起来
	s5 := []string{"how","do","you","do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5,"#"))

}