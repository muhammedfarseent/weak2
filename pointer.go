package main 

import "fmt"

func changevalue(num *int){
	*num = *num + 10

}
func main(){
	num := 40
	fmt.Println("before:",num)

	changevalue(&num)
	fmt.Println("change value:",num)
}