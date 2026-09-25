package main

import "fmt"

////////////////////// 1- function as a parameter ////////////////////
// func loading() string{
// 	return "loading ...\n"
// }

// func concat(s1, s2 string, age int, load func() string) string {
// 	fmt.Println(load())
// 	str := fmt.Sprintf("%s! Im %s im %d years old", s1, s2, age) 
// 	return str
// }

// func start_process(x1, x2 string,
// 			con func(s1, s2 string, age int, f func() string) string,
// 			a int,
// 			load func() string){
// 	fmt.Println("hello to the program x_go !")
// 	fmt.Println("start processing the log ...")
// 	fmt.Println(con(x1, x2, a, load))
// }

///////////////////// 2- function with multiple return values ////////////////
func get_names(n1, n2 string) (string, string){
	return n1, n2
}

func greet_people(n1, n2 string) (g1, g2 string){
	g1 = fmt.Sprintf("hello %s", n1)
	g2 = fmt.Sprintf("hello %s", n2)
	return
}


func overwrite_data(n1, n2 *int) (x *int, y *int){
	x = n1
	y = n2
	*x = 150
	*y = 250
	var k, v int = 10, 20
	return &k, &v
}

func calculator(a, b int, operation rune) int {
	add := func(n1, n2 int) int {return n1 + n2}
	sub := func(n1, n2 int) int {return n1 - n2}
	mult := func(n1, n2 int) int {return n1 * n2}
	div := func(n1, n2 int) int {return n1 / n2}

	if operation == '+' {
		return add(a, b)
	} else if operation == '-'{
		return sub(a, b)
	}else if operation == '*'{
		return mult(a, b)
	}else if operation == '/'{
		return div(a, b)
	}else {
		return 0
	}
} 

func main() {
	///// 1 - function as a param :

	//start_process("hello ", "ayoub", concat, 23, loading)

	///// 2 - function multi returnes :

	_, res2 := get_names("ayoub", "ali")
	fmt.Println(res2)

	res1, _ := greet_people("ayoub", "ali")
	fmt.Println(res1)

	e1 := 100
	e2 := 200
	fmt.Println(e1, e2)
	var x, y *int = overwrite_data(&e1, &e2)
	fmt.Println(*x, *y)
	fmt.Println(e1, e2)

	///// 3 nested functions implementation 
	fmt.Printf("calculator ready !\n")
	fmt.Println("10 + 20", calculator(10, 20, '+'))
	fmt.Println("10 - 20", calculator(10, 20, '-'))
	fmt.Println("10 * 20", calculator(10, 20, '*'))
	fmt.Println("10 / 20", calculator(10, 20, '/'))
	fmt.Println("10 $ 20", calculator(10, 20, '$'))
}