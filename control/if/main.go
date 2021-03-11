package main

import "fmt"

func main(){

	//基本写法
	var score =65
	if score>90{
		fmt.Println("A")
	}else if score>75{
		fmt.Println("B")
	}else{
		fmt.Println("C")
	}

	//特殊写法
	//score只在代码块有效
	if score:=65;score>90{
		fmt.Println("A")
	}else if score>75{
		fmt.Println("B")
	}else{
		fmt.Println("C")
	}
}