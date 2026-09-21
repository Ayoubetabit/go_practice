package main
import "fmt"

func main(){



	name := "ayoub"
	age := 127
	note := 16.5
	active := true

	fmt.Printf("%v | %T\n", name, name)
	fmt.Printf("%v | %T\n", age, age)
	fmt.Printf("%v | %T\n", note, note)
	fmt.Printf("%v | %T\n", active, active)

	var x int8
	x = 127

	z := x + 127
	fmt.Printf("-->%v\n", z)


	// var str string

	fmt.Printf("%v\n", int8(127))
}
