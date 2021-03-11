package main 
import "fmt"
func main(){
	for i:=0;i<10;i++{
		fmt.Printf("%d\t",i)
	}
	fmt.Println()

	var  i=0
	for ;i<9;i++{
		fmt.Printf("%d\t",i)
	}
	fmt.Println()

	var m=10
	for m>0{
		fmt.Printf("%d\t",m)
		m--
	}
	fmt.Println()

	for i:=0;i<50;i++{
		if(i%2==0){
			fmt.Printf("%d\t",i)
		}
		if(i==25){
			fmt.Println()
			break
		}
	}

	for i:=0;i<50;i++{
		if(i%2==0){
			continue
		}
		fmt.Printf("%d\t",i)
	}
	fmt.Println()
}