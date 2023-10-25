package main

import (
	"fmt"
	"go-learn/generic/genericFunc"
	"go-learn/generic/genericInterface"
)

func main() {
	fmt.Println("genericFunc.Run()")
	genericFunc.Run()
	fmt.Println("genericInterface.Run()")
	genericInterface.Run()
}
