package main

/*
Cara multiimport menggunakan ()
*/
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
}
