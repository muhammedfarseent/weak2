package main 

import "fmt"

func findmax(arr []int)int{
	max:= arr[0]
    for i:=1; i<len(arr);i++{
		if arr[i]> max{
			max = arr[i]
		}
	}
  return max
}

func main(){
	numbers:=[]int{1,6,8,25,23,45}
	fmt.Println(findmax(numbers))

}