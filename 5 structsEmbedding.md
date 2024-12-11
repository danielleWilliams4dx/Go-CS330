# Structs and Embedding
This file references the [Go programming language documentation](https://go.dev/ref/spec).

Additional sources are cited below.

## Does Go have Objects?
Go does not support objects; however, Go has ```structs```, or structures, which are user-defined types that are similar to classes in traditional object-oriented programming. Structs are declared with the ```type``` and ```struct``` keywords. Methods can be defined on types with a receiver argument, which specifies the type and is placed between the ```func``` keyword and the function name. Struct attributes and methods can be accessed using the ```.``` dot operator.

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
See [structsEmbeddingExamples](https://github.com/danielleWilliams4dx/Go-CS330/tree/main/structsEmbeddingExamples) for more examples.

[W3Schools. (n.d). Go Structures. Retrieved September 28, 2024.](https://www.w3schools.com/go/go_struct.php)

[The Go Programming Language. (n.d). Methods. Retrieved September 28, 2024.](https://go.dev/tour/methods/1)

### Naming Conventions
Like everything else in Go, ```struct``` instances and functions should use the camelCase naming convention and begin with a capital letter if it is exported (accessible outside of the package) and begin with a lowercase letter otherwise.

## Standard Methods
Go does not have standard methods that serve a similar purpose across all structs such as ```toString()``` in Java or  ```__str__``` in Python.

## Inheritance
Go does not support traditional inheritance seen in languages like Java and Python that are built for object-oriented programming. There are no **IS-A** relationships in Go, only **HAS-A** relationships are viable and can be created through _struct embedding_ (vanigupta20024, 2020). This notion in which a ```struct``` is utilized to create another is called _composition_. A single ```struct``` can embed any number of other structures.

```
type Person struct {
	name string
	birthYear int
}

// Go doesn't have inheritance,
// but it does have embedding
type Student struct{
	person Person
	year int
	major string
}
```

[vanigupta20024. (June 22, 2020). Inheritance in GoLang. Geeks for Geeks. Retrieved November 9, 2024.](https://www.geeksforgeeks.org/inheritance-in-golang/)

### Simulating IS-A Relationships
In the example above, a ```Student``` instance can only call ```Person``` attributes and methods through its own ```person``` attribute.

```
s.person.name
```
**Getter** and **setter** functions can be defined on the embedding structure, in this case ```Student```, to treat the embedded attributes as its own.
```
func (s Student) getName() string{
	return s.person.name
}

// in another function:
fmt.Println("My name is " + s.getName())
```
