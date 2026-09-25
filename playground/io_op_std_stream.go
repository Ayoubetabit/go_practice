package main

import (
	"os"
	"fmt"
)

func main() {
	///// Print in the std streams
	fmt.Println("use os package to reach to std streams to write (1, 2)\n")
	err1, b1 := os.Stdin.Write([]byte("hello im printing from stdin stream by os package"))
	fmt.Println("\nthe error rsturns value: ", err1, "\nbytes printed: ", b1)
	err2, b2 := os.Stderr.Write([]byte("hello im printing from stderr stream by os package"))
	fmt.Println("\nthe error rsturns value: ", err2, "\nbytes printed: ", b2)

	fmt.Println("use fmt package to reach to std streams via Fprintf function to write (1, 2)\n")
	fmt.Fprintf(os.Stdin, "hello im printing from stdin stream by fmt/Fprintf package\n")
	fmt.Fprintf(os.Stderr, "hello im printing from stderr stream by fmt/Fprintf package\n")
	///// 
}


