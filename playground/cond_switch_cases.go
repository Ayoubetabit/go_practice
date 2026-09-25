package main
import "fmt"

var (
	x int = 120
	y int = 120
)


func main(){
	if (x > y) {
		fmt.Println(x, "is greater then ", y)
	} else if (x < y) {
		fmt.Println(y, "is greater then ", x)
	} else {
		fmt.Println(x, " and ", y," are equale")
	}

	switch x {
	case x > y:
		fmt.Println(x, "is greater then ", y)
	case x < y:
		fmt.Println(y, "is greater then ", x)
	default:
		fmt.Println(x, " and ", y," are equale")
	}
}