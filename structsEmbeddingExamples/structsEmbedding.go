package main

import "fmt"
/*
This example program covers:
structs
struct embedding
struct methods
*/

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

// getAge is defined on the Person type with the reciever argument p
func (p Person) getAge(currYear int) int{
	return currYear - p.birthYear
}

// calls getAge as a helper method
func (p Person) sayHello(){
	fmt.Println("Hello! My name is " + p.name + ". I am " + fmt.Sprint(p.getAge(2024)) + " years old.")
}

// replaces the student's major with newMajor
// uses pointers to modify the object
func (s *Student) changeMajor(newMajor string){
	s.major = newMajor // s --> *s implicitly
}

// calls person's sayHello method and adds to the output with Student unique information
// calls getNumAddOn as a helper method.
func (s Student) sayHello(){
	s.person.sayHello()
	fmt.Println("I am a " + getNumAddOn(s.year) + " year and a " + s.major + " Major.\n")
}

// returns the number with the proper add-on
// ex: 1 --> 1st
func getNumAddOn(num int) string{
	s := fmt.Sprint(num) //converts num to a string and stores it in s
	switch num{
		case 1:
			return s + "st"
		case 2:
			return s + "nd"
		case 3:
			return s + "rd"
		default:
			return s + "th"
	}
}

func main(){
	var naomi Person = Person{"Naomi", 1999}
	var danielle Student = Student{Person{"Danielle", 2004}, 3, "Undeclared"}
	naomi.sayHello()
	fmt.Println()
	danielle.sayHello()
	danielle.major = "Computer Science and Design" // change major directly through its attribute
	danielle.sayHello()
	danielle.changeMajor("Computer Science, Design, and Web Design and Development") // change major with class method
	danielle.sayHello()
}
