package main 

import "fmt"

func main(){
    //// new code 
	p:= new(int)
	fmt.Println(*p)
    
	*p = 42
	fmt.Println(*p)

     ///make code 

	 arr:= make([]int, 3)
	 arr[0],arr[1],arr[2]= 21,18,24
	fmt.Println(arr)
}
