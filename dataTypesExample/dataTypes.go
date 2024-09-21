package main

import "fmt"
import "reflect"

func main() {
	//implicit declaration
    // num1 := 17
	// ans1, ans2 := true, false //multiple variables can be declared at once
	// fruit := map[string]string{ //Go equivalent of dictionaries
	// 	"apple" : "red",
	// 	"pommegranate" : "pink",
	// 	"banana" : "yellow",
	// 	"grapes" : "purple",
	// 	"mango" : "orange",
	// 	"plums" : "purple",
	// 	"coconut" : "white",
	// 	"clementine" : "orange",
	// 	"blueberries" : "blue",
	// 	"blackberries" : "black",
	// }
	// lang := [...]string{"Java","JavaScript","Go","Python","C++"} //defining with inferred lengths
	// ages := [10]int{19,20,20,20,20,21,21,21,21,25} //defining with set lengths

	//explicit declaration
	var num1 int = 17
	var ans1, ans2 bool = true, false //multiple variables can be declared at once
	var num2 float32 = 2.0
	var hello string = "hello world"
	var fruit = map[string]string{ //Go equivalent of dictionaries
		"apple" : "red",
		"pommegranate" : "pink",
		"banana" : "yellow",
		"grapes" : "purple",
		"mango" : "orange",
		"plums" : "purple",
		"coconut" : "white",
		"clementine" : "orange",
		"blueberries" : "blue",
		"blackberries" : "black",
	}
	var lang = [...]string{"Java","JavaScript","Go","Python","C++"} //defining with inferred lengths
	var ages = [10]int{19,20,20,20,20,21,21,21,21,25} //defining with set lengths
	
	fmt.Println("\nVariables and Types\n")
	fmt.Println(num1, "Type:", reflect.TypeOf(num1))
	fmt.Println(ans1, "Type:", reflect.TypeOf(ans1))
	fmt.Println(ans2, "Type:", reflect.TypeOf(ans2))
	fmt.Println(num2, "Type:", reflect.TypeOf(num2))
	fmt.Println(hello, "Type:", reflect.TypeOf(hello))
	fmt.Println(fruit, "Type:", reflect.TypeOf(fruit))
	fmt.Println(lang, "Type:", reflect.TypeOf(lang)) //specifies an array length of 5
	fmt.Println(ages, "Type:", reflect.TypeOf(ages))

	//basic mathematic operations with the same type
	fmt.Println("\nMathematical Operations - Same Type\n")
	fmt.Println(num1, "+", 17, "=", num1+17) //addition
	fmt.Println(num1, "-", 63, "=", num1-63) //subtraction
	fmt.Println(num1, "*", 52, "=", num1*52) //multiplication
	fmt.Println(num1, "/", -5, "=", num1/-5) //integer division
	fmt.Println(num2, "/", 7.0, "=", num2/7.0) //float division
	fmt.Println(2, "/", 5, "=", 2/5) //integer division, yields 0 (truncated)

	//basic mathematic operations with different types succeed when one or both is a constant
	fmt.Println("\nMathematical Operations - Different Type\n")
	fmt.Println(3, "+", 8.2, "=", 2+8.2) //addition, yields a float
	// fmt.Println(12.1, "-", num1, "=", 12.1-num1) //subtraction, yields an int (truncated)
	fmt.Println(5, "*", -13.1, "=", 5*-13.1) //multiplication, yields a float
	fmt.Println(num2, "/", 10, "=", num2/10) //mixed division, yields a float
	fmt.Println(41, "/", 20.1, "=", 41/20.1) //mixed division, yields a float
	

	//basic mathematic operations with different types fail otherwise
	// fmt.Println("\nMathematical Operations Different Type\n")
	// fmt.Println(num1, "+", num2, "=", num1+num2) //addition
	// fmt.Println(num1, "-", num2, "=", num1-num2) //subtraction
	// fmt.Println(num1, "*", num2, "=", num1*num2) //multiplication
	// fmt.Println(num1, "/", num2, "=", num1/num2) //mixed division

	//basic mathematic operations with casting
	fmt.Println("\nMathematical Operations - Type Casting\n")
	fmt.Println(num1, "+", num2, "=", num1+int(num2)) //addition, yields an int
	fmt.Println(num1, "-", num2, "=", float32(num1)-num2) //subtraction, yields a float
	fmt.Println(num1, "*", num2, "=", num1*int(num2)) //multiplication, yields an int
	fmt.Println(num1, "/", num2, "=", float32(num1)/num2) //mixed division, yields a float
}