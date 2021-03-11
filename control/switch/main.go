package main
import "fmt"
func main(){
	finger :=3
	//swtich 没有break
	switch finger{
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无")
	}

	//case 里可以有多个值
	num:=9
	switch num{
	case 1,3,5,7,9:
		fmt.Println("奇数")
	default:
		fmt.Println("偶数")
	}

	//case中使用表达式
	age:=25
	switch{
	case age>30:
		fmt.Println("三十而立")
	case age<30:
		fmt.Println("哈哈哈")
	}
}