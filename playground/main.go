package main
import "fmt"

func try_vim() string {
	return "hello vim"
}

func chang_arr(arr *[5]int) {
	arr[4] = 200
}

func main(){
	// name := "ayoub"
	// age := 127
	// note := 16.5
	// active := true

	// fmt.Printf("%v | %T\n", name, name)
	// fmt.Printf("%v | %T\n", age, age)
	// fmt.Printf("%v | %T\n", note, note)
	// fmt.Printf("%v | %T\n", active, active)

	// var x int8
	// x = 127

	// z := x + 127
	// fmt.Printf("-->%v\n", z)


	// var str string

	// s := "界ello"
	// x := '界'
	// fmt.Printf("%v\n", s[:3])
	// fmt.Printf("%T\n", s[0])
	// fmt.Printf("%c\n", x)
	// fmt.Printf("%v\n", x)
	// fmt.Printf("%T\n", x)
	// fmt.Printf("%v\n", len(s))

	x := '界'

	fmt.Printf("%%v = %v\n", x)
	
	fmt.Printf("%%T = %T\n", x)
	fmt.Printf("%%c = %c\n", x)
	fmt.Printf("%%d = %d\n", x)
	fmt.Printf("%%x = %x\n", x)
	fmt.Printf("%%b = %b\n", x)
	fmt.Printf("%%q = %q\n", x)
	var gp_array = [...]int{12, 13, 14, 15, 16}
	fmt.Println(gp_array)
	chang_arr(&gp_array)
	fmt.Println(gp_array)
}
