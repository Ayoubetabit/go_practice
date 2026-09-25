package main

import ("fmt";"slices")

func chang_arr(arr []int) *[]int {
	arr[4] = 200
	n_arr := []int{12, 30, 40}
	up_arr := append(arr, n_arr...)
	return &up_arr
}


func main(){
	// 1 - manage array's by function
	// gp_array := []int{12, 13, 14, 15, 16}
	// fmt.Println(gp_array)
	// fmt.Printf("%v=> %p\n", gp_array, gp_array)
	// n_arr := chang_arr(gp_array)
	// fmt.Printf("%v=> %p\n", gp_array, gp_array)
	// fmt.Printf("%v=> %p\n", *n_arr, *n_arr)
	// n1_arr := chang_arr(*n_arr)
	// fmt.Printf("%v=> %p\n", *n1_arr, *n1_arr)

	// 2 - check addresses
	test_arr := []int{12, 13, 14, 15, 16}
	fmt.Printf("%v=> %p\n", test_arr, test_arr)
	test2_arr := slices.Clone(test_arr)
	chang_arr(test_arr)
	fmt.Printf("%v=> %p\n", test2_arr, test2_arr)
	fmt.Printf("t1 again %v=> %p\n", test_arr, test_arr)

	// 1 - practice append
	var new_arr [10]int
	for i:=0; i < 10; i++{
		new_arr[i] = i
	}
	fmt.Println(new_arr)
}


