package main
import "fmt"
func main(){
	//byte:uint8的别名 ASCII码
	c1:='a'
	c2:='b'
	fmt.Println(c1,c2)
	fmt.Printf("c1:%T c2:%T\n",c1,c2)

	//中文占了好几个字节，用rune来便利
	s:="hello杨超越"
	for i:=0;i<len(s);i++{
		fmt.Printf("%c\t%T\n",s[i],s[i])
	}
	for _,r:=range s{
		fmt.Printf("%c\t%T\n",r,r)

	}

	//rune:int32的别名
}