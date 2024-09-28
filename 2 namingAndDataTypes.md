# Naming and Data Types
This file references the [Go programming language documentation](https://go.dev/ref/spec).

Additional sources are cited below.

## Naming Requirements and Conventions for Variables, Structure Instances, and Functions
- Names are case sensitive
- Names must begin with a letter or underscore and cannot begin with a digit or symbol
- camelCase is conventionally used rather than snake_case
- Internal identifiers should begin with a lowercase letter
- Exported identifiers should begin with a capital letter

```
// internal
func calcAge(birthYear int, currYear int) int { return currYear - birthYear }

// exported
func CalcAge(birthYear int, currYear int) int { return currYear - birthYear }
```

[W3Schools. (n.d.). Go Variable Naming Rules. Retrieved September 20, 2024.](https://www.w3schools.com/go/go_variable_naming_rules.php)

## Data Types
- bool - represents a boolean value and is either true or false
- string - represents a string value, which is any sequence of characters
- int - represents an integer 
- float32 - represents a 32 bit floating point value
- float64 - represents a 64 bit floating point value
- (*) - represents the set of all pointers of a specified base type
- [array](https://www.w3schools.com/go/go_arrays.php) - holds a list of values of the same type in a single variable
- [slice](https://www.w3schools.com/go/go_slices.php) - like an array, contains a set of values, but is uniquely determined by its ability to grow and shrink
- [map](https://go.dev/blog/maps) - represents a hash table and is the go equivalent of a dictionary
- [chan](https://go.dev/ref/spec#Channel_types) - represents a channel and permits functions in concurrent execution to communicate by transmitting values of a specified type
- [interface](https://go.dev/ref/spec#Interface_types) - represents a type set and contains a list of interface elements (a method or type element)

### A Note About Floats
Float point values with only trailing zeros after the decimal point will be outputted without any decimal places (2.0 → 2).

### Static and Strong
Go is a statically and strongly typed language (Grasberger, 2022). A language is static if data types are checked at compile time (Oracle, 2015). Similarly, a language is strong if data types are known at compile time (Gosling et al., 2000).

### Mutability
In Go, slices, arrays, maps, and channels are mutable whereas strings, pointers, interfaces, booleans, ints, and floats are immutable (Gor, 2022).

[Oracle. (2015). Dynamic typing vs. static typing. Retrieved September 21 , 2024.](https://docs.oracle.com/cd/E57471_01/bigData.100/extensions_bdd/src/cext_transform_typing.html#:~:text=First%2C%20dynamically%2Dtyped%20languages%20perform,type%20checking%20at%20compile%20time)

[Gor, R. (2022, June 21). Golang: Mutable and Immutable Data Types. Dev. Retrieved September 21 , 2024.](https://dev.to/mr_destructive/golang-mutable-and-immutable-data-types-4p6)

[Gosling, J. et al. (2000). The Java Language Specification. Addison-Wesley.](https://docs.oracle.com/javase/specs/jls/se6/html/typesValues.html#:~:text=The%20Java%20programming%20language%20is,is%20known%20at%20compile%20time)

[Grasberger, M. (2022, October 24). What Golang generics support means for code structure. Tech Target. Retrieved September 21 , 2024.](https://www.techtarget.com/searchitoperations/tip/What-Golang-generics-support-means-for-code-structure#:~:text=Go%2C%20also%20known%20as%20Golang,considered%20strongly%20and%20statically%20typed)

## Variable Declarations
Go is explicitly and implicitly typed. See [dataTypesExample](https://github.com/danielleWilliams4dx/Go-CS330/tree/main/dataTypesExample).

Variables can be declared with implicit type through the **:=** short assignment statement.

```a := 17``` ```b, c := true, false```

Variables can be declared explicitly with **var** and its type (with its name sandwiched between)

```var a int = 17``` ```var b, c bool = true, false```

The reflect package has a **TypeOf** method that can be used to check the type of a variable.

## Function Declarations
Functions are declared with the **func** type. Each parameter must have a designated type. All parameters must have a name, otherwise they must be solely identified by their type such that each is the only instance of its type in the function. Zero or more arguments can be passed into the final parameter of a _variadic_ function. In these functions, the last parameter is prefixed with a **...**. The return type(s) should be declared after the parameter list in the result list. Functions without a declared return type return **nil**.

```
func(a int, b int) int // takes in two ints and returns an int

func(int, float32) (int, float32) // takes in and returns an int and float32

func(bool, ...string) // takes in a bool and any number of strings and returns nil
```

## A Quick Note About Structs
In Go, **structs**, or structures, are user-defined types and are similar to classes in traditional object-oriented programming. Structs are declared with the **type** and **struct** keywords. Methods can be defined on types with a receiver argument, which specifies the type and is placed between the **func** keyword and the function name. Struct attributes and methods can be accessed using the "." dot operator.

```
package main

import "fmt"

type Cat struct {
	name string
	breed string
	age int
}

func (c Cat) Meow(){
	fmt.Println(c.name + " says meow!")
}

func main() {
	luna := Cat{"Oops!", "Siamese", 8}
	luna.name = "Luna" // changes the name attribute from Oops! to Luna
	luna.Meow() // Luna says meow!
}
```
[W3Schools. (n.d). Go Structures. Retrieved 28 September 20, 2024.](https://www.w3schools.com/go/go_struct.php)

[The Go Programming Language. (n.d). Methods. Retrieved 28 September 20, 2024.](https://go.dev/tour/methods/1)

## Type Conversions
Go does not support implicit type conversion except for mixed operations of variables with constants in which the lower data type is cast to a higher data type. Attempting to implicitly cast in the other direction, such as converting a floating point constant to an integer, throws an error.

The syntax for converting between types is as follows:

```variable = desiredType(variable)```

[The Go Programming Language. (n.d). Type Conversions. Retrieved 21 September 20, 2024.](https://go.dev/tour/basics/13)

## Operators
- The assignment operator = can be used with any data type.
- Other assignment operators (+=, -=, *=, /=, %=, &=, |=, ^=, >>=, <<=) can be used with numeric types.
- The arithmetic operators (+, -, *, /, %, ++, --) can be used with numeric types.
- The comparison operators (==, !=) can be used with any data type.
- Other comparison operators (>, <, >=, <=) can be used with numeric types and strings.
- Logical operators (&&, ||, !), also known as boolean operators, are used to solve the logic between variables or values.
- [Bitwise Operators](https://www.geeksforgeeks.org/go-operators/#Bitwise%20Operators) (&, |, ^, <<, >>) can be used with integers and have applications in binary computations.

## Reserved Keywords
Go has the following 25 reserved keywords:

| break | case | chan | const | continue |
| :---: | :---: | :---: | :---: | :---: |
| default | defer | else | fallthrough | for |
| func | go | goto | if | import | 
| interface | map | package | range | return |
| select | struct | switch | type | var |

## Binding
- Identifiers are bound to their type and address at compile time.
- Operators are also bound at compile time. Urinary operators bind first, then multiplication operators, addition operators, comparison operators, && and ||.

## Limitations
- Arrays can only store data of the same type.
- A compiler may flag the declaration of an unused variable in a function body as illegal.
- A compiler does not throw an error if the type of an operand belongs to an empty type set. A function with one of these types parameters cannot be declared and will cause an error.
- A compiler may incorrectly round when evaluating floating-point constant expressions in an integer context.
- The conversion between pointers and integers can only be implemented with the unsafe package and this is done in limited cases.



