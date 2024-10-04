package main

import "fmt"

/*
Call your functions and save the results of the function calls in variables.
Write a function that tests whether your language is pass-by-reference or pass-by-value.
*/
func main(){
	a := 5
	b := 7
	banana := "banana"

	// multiply test
	fmt.Println(a, "*", b, "=", multiply(a,b))

	// factorial test
	fmt.Println(fmt.Sprint(a) + "! = " + fmt.Sprint(factorial(a)))

	// string split test
	bananaSplit := split(banana)
	fmt.Println("\n" + banana + " --Split--> " + bananaSplit[0] + " " + bananaSplit[1])


	// swapping test for pass by reference or pass by value
	var oldA, oldB int = a, b
	fmt.Println("\nInside main before swapping: a = " + fmt.Sprint(a) + ", b = " + fmt.Sprint(b))
	swap(a,b)
	fmt.Println("Inside main after swapping: a = " + fmt.Sprint(a) + ", b = " + fmt.Sprint(b))

	if(oldA == a && oldB == b){
		fmt.Println("Go is pass by value")
	}else{
		fmt.Println("Go is pass by reference")
	}
}

// takes in two numbers, multiplies them, and returns the output
func multiply(num1 int, num2 int) int{
	return num1 * num2
}

// recursively calculates a factorial
func factorial(num int) int{
	if(num == 1 || num == 0){ return 1 }
	return num*factorial(num-1)
}

// takes in a string, splits it into two strings, then returns both strings
func split(str string) [2]string{
	var mid int = len(str)/2
	var firstHalf string = str[:mid]
	var secondHalf string = str[mid:]
	var arr = [2]string{firstHalf, secondHalf}
	return arr
}

// tests whether Go is pass-by-reference or pass-by-value
func swap(a int, b int){
	var temp int = a
	a = b
	b = temp
	fmt.Println("Inside the function after swapping: a = " + fmt.Sprint(a) + ", b = " + fmt.Sprint(b))
}