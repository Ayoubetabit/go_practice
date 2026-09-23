package main

import "fmt"

func chang_arr(arr *[5]int) {
	arr[4] = 200
}

func main(){
	var gp_array = [...]int{12, 13, 14, 15, 16}
	fmt.Println(gp_array)
	chang_arr(&gp_array)
	fmt.Println(gp_array)
}


