# Selection, Loops and Conditionals
This file references the [Go programming language documentation](https://go.dev/ref/spec).

Additional sources are cited below.

## Boolean Values
In Go, boolean values are predetermined constants, ```true``` and ```false```. If an initial value is not specified, a boolean variable is false by default.

[W3Schools. (n.d.). Go Boolean Data Type. Retrieved October 13, 2024.](https://www.w3schools.com/go/go_boolean_data_type.php)

## Conditional Statements
Go supports if, if-else, if-else if-else, and switch statements. 
### if, if-else, if-else if-else
For these selection control statements, Go delimits code blocks under each condition using curly braces ```{ }```, which avoids the dangling else problem. 

**Note:** ```else if``` and ```else``` must be on the same line as the closing brace ```}``` that delimits the end of the above code block.

```
if(len(mySlice) >= 3){
		fmt.Println("That's enough fruits for a fruit salad!")
	}else if(len(mySlice) == 0){ // 'else if' must be on the same line as the closing '}' to prevent a syntax error
		fmt.Println("What are you doing? Are you making an imaginary fruit salad?")
	}else{ // 'else' must be on the same line as the closing '}' to prevent a syntax error
		fmt.Println("That would be a very sad fruit salad. Let's get some more fruit.")
	} 
``` 
See [conditionalsExamples](https://github.com/danielleWilliams4dx/Go-CS330/tree/main/conditionalsExamples) for more examples.

### Switch Statements
Breaks are implicit and do not need to be written explicitly inside of individual cases. However, if you want your case to fall through into the next one, use the ```fallthrough``` keyword. ```continue``` cannot be used to evaluate all of the conditions. For switch statements, Go delimits code blocks under each case using a colon ```:```.

```
func numberCompliments(num int){
	switch num{
		case 1:
			fmt.Print("There’s only one of you!")
		case 2:
			fmt.Print("You’re too amazing!")
		case 3:
			fmt.Print("You’re magic! ")
			fallthrough
		case 4:
			fmt.Print("You’re forever my favorite!")
		case 5:
			fmt.Print("You deserve a high five!")
		default:
			fmt.Print("I’m sure you’re great!")
	}
}
numberCompliments(3) // You’re magic! You’re forever my favorite!
```
[The Go Programming Language. (n.d.). Go Wiki: Switch. Retrieved October 13, 2024.](https://go.dev/wiki/Switch) 

## Short-Circuit Evaluation
Go supports short-circuit evaluation, which means that conditional statements can be evaluated as ```true``` or ```false``` without evaluating all conditions. When conditional statements short-circuit, the code in subsequent conditions are never reached.

```
x := true
y := false
if(!x && y){ // the compiler never checks y (false && ____ → false)
	fmt.Println("All false here!")
}else{
	fmt.Println("All true or mixed.")
}
``` 

## Loops
Unlike other programming languages like Java and C++, Go does not have while loops. Go does have for loops, which can be implemented like while loops.

```
 var repeat bool = true
// Go has no while loops
// This for loop is implemented like a while loop
// It will loop until the loop condition isn’t met
// When repeat is false
for repeat == true{
	// some statements
}
```

The ```continue``` statement can be used to skip over for loop iterations.

```
// Print 0 - 9, and skips 5
for i := 0; i < 10; i++{
	if(i == 5){
		fmt.Println("Skip!")
		continue
	}
	fmt.Println(i)
}
```

The ```break``` statement can be used to break out of a loop, but it is poor practice to use them unless truly necessary.

```
// Print 0 - 4, breaks out of the loop when i == 5
for i := 0; i < 10; i++{
	if(i == 5){
		fmt.Println("I’m breaking out!")
		break
	}
	fmt.Println(i)
}
```

The ```range``` keyword can be used to simply traverse arrays, slices, and maps. It returns both the index and value of each element.

```
// prints every fruit in fruitSlice
for i,fruit := range fruitSlice{
		fmt.Println("index: " + fmt.Sprint(i) + " item: " + fruit)
}
```
[W3Schools. (n.d.). For Loops. Retrieved October 13, 2024.](https://www.w3schools.com/go/go_loops.php)
