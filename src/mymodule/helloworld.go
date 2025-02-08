package main

import (
	"fmt"
	"mymodule/datatype"
	"mymodule/mathutil"
)

func main() {
	fmt.Println("Hello world!!...")
	datatype.Datatype1()
	calc := mathutil.Calculator{Name: "Basic Calculator"}

	fmt.Println("Addition: ", calc.Add(10, 30))
	fmt.Println("Subtraction: ", calc.Subtract(20, 10))
	fmt.Printf("SquareRoot: %.2f", calc.SquareRoot(25))
	//calc.PrintInfo()
}
