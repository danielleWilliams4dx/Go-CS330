# Functions

This file references the [Go programming language documentation](https://go.dev/ref/spec).

Additional sources are cited below.

## Function Declarations
Functions are declared with the **func** type. Uninitialized functions have a value of **nil**. Functions can have zero or more parameters and can take in arguments of different data types. Functions that specify a list of type parameters are __generic__ and must be instantiated before a function call. Functions without type parameters do not need a body and will be implemented outside of Go. All parameters must have a name, otherwise they represent one instance of its type. Zero or more arguments can be passed into the final parameter of a _variadic_ function. In these functions, the last parameter is prefixed with a **...**. Multiple return values are supported and the return type(s) should be declared after the parameter list in the result list. Functions should terminate with a return statement containing data that matches the type(s) in the result list.

```
func f1(a int, b int) int{...} // takes in two ints and returns an int

func f2(int, float32) (int, float32){...} // takes in and returns an int and float32

func f3(bool, ...string){...} // takes in a bool and any number of strings and has no return type

func f4(a, b) // implemented externally; does not specify type parameters and omits a body

func f5[T ~int|~float64](a T, b T) T{...} // takes in two ints or float64s and returns an int or float64
```

## Function Features
Since Go is a compiled language, functions can be placed anywhere in the code file and it will still run (declarations can be placed above/below function calls). In addition to being a compiled language, Go is pass-by-value by default but supports pass-by-reference through the use of pointers (Agrawal, 2023). Go also supports recursive functions.

```
// without pointers, invokes pass-by-value substitution by default
func swap1(a int, b int){
	var temp int = a
	a = b
	b = temp
}

//with pointers, pass-by-reference
func swap2(a *int, b *int){
	var temp int = *a
	*a = *b
	*b = temp
}
```
See [functionExamples](https://github.com/danielleWilliams4dx/Go-CS330/tree/main/functionExamples) for more examples.

[Agrawal, M. (2023, March 15). Golang Call by Reference and Call by Value. Scaler. Retrieved October 5, 2024.](https://www.scaler.com/topics/golang/golang-call-by-reference-and-call-by-value/) 

## Side Effects

As seen in the ```swap2``` function, side-effects are possible in Go through aliasing. However, passing in an array as an argument and modifying it inside a function will not impact the original array. This is one protection that is not featured in other languages like C++ and Java. It is good practice to pass copies as arguments in order to prevent side-effects.

```
package main

import "fmt"

func main(){
	var a = [...]int{1,2,3}
	fmt.Println(a[0],a[1],a[2]) // 1 2 3
	changeArr(a)
	fmt.Println(a[0],a[1],a[2]) // 1 2 3
}

func changeArr(a [3]int){
	a[1] = 47
	fmt.Println(a[0],a[1],a[2]) // 1 47 3
}
```

## Scope
In Go, scope is determined by code blocks and is summarized in the table below.

→ means "to the end"

| Scope | Identifiers |
| --- | --- |
| Universe Block | [Predeclared identifiers](https://go.dev/ref/spec#Predeclared_identifiers) |
| Package Block | Constants, types, variables, or functions (but not type-defined methods) that are declared outside of any function |
| File Block | Imported package names |
| Function Body | Method receivers, function parameters, or result variables |
| After Function Name → Function Body | Function type parameters or method receiver type parameters|
| After Type Name → TypeSpec | Type parameters |
| After ConstSpec/VarSpec/ShortVarDecl → Innermost Containing Block | Local constants or variables in a function |
| TypeSpec Identifier → Innermost Containing Block | Local type identifiers in a function |

**Note:** an identifier may be redeclared in an inner code block, but all references within the scope of the redeclared identifier will refer to the redeclared version (it is a separate entity).

## Stack V.S. Heap

The stack contains data that is readily accessible whereas the heap contains data with much slower access. The stack acts as an intermediary for the heap and may store the addresses (pointers) of data stored on the heap. Go is garbage collected, which means that memory allocated on the heap becomes deallocated if the data does not have a pointer reference on the stack.

Arguments and parameters are stored on the stack (Kumar, 2021). Local variables are typically stored on the stack; however, if a pointer to a local variable is returned, the compiler generally places the local variable on the heap, which is called “escaping the heap.”

[Kumar, K. (2021, September 1). Memory allocations in Go. Dev. Retrieved October 5, 2024.](https://dev.to/karankumarshreds/memory-allocations-in-go-1bpa)
