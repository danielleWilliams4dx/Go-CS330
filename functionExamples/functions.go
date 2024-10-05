package main

import "fmt"

func main(){
	a := 5
	b := 7
	banana := "banana"

	// multiply test
	fmt.Println(a, "*", b, "=", multiply(a,b))

	// factorial test
	fmt.Println(fmt.Sprint(a) + "! = " + fmt.Sprint(factorial(a)))

	// min test
	c := 3.14
	fmt.Println("Minimum of " + fmt.Sprint(a) + " and " + fmt.Sprint(c) + " is " + fmt.Sprint(min(float64(a),c)))

	// string split test
	bananaSplit := split(banana)
	fmt.Println("\n" + banana + " --Split--> " + bananaSplit[0] + " " + bananaSplit[1])


	// swapping test for pass by reference or pass by value
	var oldA, oldB int = a, b
	var aAddr, bAddr *int = &a, &b

	fmt.Println("\nInside main before swapping: a = " + fmt.Sprint(a) + ", b = " + fmt.Sprint(b))
	swap1(a,b)
	fmt.Println("Inside main after swapping: a = " + fmt.Sprint(a) + ", b = " + fmt.Sprint(b))

	if(oldA == a && oldB == b){
		fmt.Println("Go is pass by value by default")
	}else{
		fmt.Println("Go is pass by reference by default")
	}

	fmt.Println("\nUsing Pointers:\nInside main before swapping: a = " + fmt.Sprint(a) + ", b = " + fmt.Sprint(b))
	swap2(aAddr,bAddr)
	fmt.Println("Inside main after swapping: a = " + fmt.Sprint(a) + ", b = " + fmt.Sprint(b))
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

// returns the minimum of two numbers
func min[T ~int|~float64](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

// takes in a string, splits it into two strings, then returns both strings
func split(str string) [2]string{
	var mid int = len(str)/2
	var firstHalf string = str[:mid]
	var secondHalf string = str[mid:]
	var arr = [2]string{firstHalf, secondHalf}
	return arr
}

// tests whether Go is pass-by-reference or pass-by-value by default
func swap1(a int, b int){
	var temp int = a
	a = b
	b = temp
	fmt.Println("Inside the function after swapping: a = " + fmt.Sprint(a) + ", b = " + fmt.Sprint(b))
}

func swap2(a *int, b *int){
	var temp int = *a
	*a = *b
	*b = temp
	fmt.Println("Inside the function after swapping: a = " + fmt.Sprint(*a) + ", b = " + fmt.Sprint(*b))
}
