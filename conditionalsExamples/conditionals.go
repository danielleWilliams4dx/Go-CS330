package main

import "fmt"

/*
This example program covers:
one-condition if statements
multi-condition if/else statements
if/elif/else statements
short-circuit logic
switch-case statements
for-loops
*/

// struct represents a calendar date
type Date struct{
	day int
	month int
	year int
}

func main() {
	fruit := "banana"
	fruitSlice := []string{} // create an empty slice
	fmt.Println(fruitSlice)
	enoughFruits(fruitSlice)
	fruitSlice = append(fruitSlice, "apple","pineapple") //append "apple" and "pinapple" to fruitSlice
	fmt.Println(fruitSlice)
	enoughFruits(fruitSlice)
	fruitSlice = append(fruitSlice, fruit) // appending fruit to fruitSlice
	fmt.Println(fruitSlice)
	enoughFruits(fruitSlice)
	// // prints every fruit in fruitSlice with its index
	// for i,fruit := range fruitSlice{
	// 	fmt.Println("index: " + fmt.Sprint(i) + " item: " + fruit)
	// }
	fmt.Println()

	today := Date{12,10,2024};
	bestBy := Date{17,12,2024};
	safeToEat(fruitSlice, "apple", bestBy, today)
	fruitSlice = fruitSlice[1:] // remove the "apple" from fruitSlice
	fmt.Println(fruitSlice)
	safeToEat(fruitSlice, "apple", bestBy, today)
	fmt.Println()

	numberCompliments(3) // try a different number!
	fmt.Println()
	favFruit()
}

// outputs whether there are enough fruits for a fruit salad
func enoughFruits(mySlice []string){
	// if/elif/else statement
	if(len(mySlice) >= 3){
		fmt.Println("That's enough fruits for a fruit salad!")
	}else if(len(mySlice) == 0){ // 'else if' must be on the same line as the closing '}' to prevent a syntax error
		fmt.Println("What are you doing? Are you making an imaginary fruit salad?")
	}else{ // 'else' must be on the same line as the closing '}' to prevent a syntax error
		fmt.Println("That would be a very sad fruit salad. Let's get some more fruit.")
	} 
}

// outputs whether a food is safe to eat by its ingredients and best by date
// showcases short-ciruit logic
func safeToEat(ingredients []string, allergen string, bestBy Date, today Date){
	// short-circuit example
	// if there is an allergen, doesn't check the best by date
	if(containsAllergen(ingredients, allergen) || !compareDate(today,bestBy)){
		fmt.Println("CAUTION! This is not safe to eat.")
	}else{ 
		fmt.Println("Enjoy! This is safe to eat.")
	}
}

// compares two dates, returns true if the first date is before or the same as the second
func compareDate(day1 Date, day2 Date) bool{
	fmt.Println("[Checking best by date]")
	// one-condition if statement
	if(day1.year <= day2.year){
		if(day1.month <= day2.month){
			if(day1.day <= day2.day){
				return true
			}
		}
	}
	return false
}

// returns true if a list of ingredients contains the given allergen, false otherwise
func containsAllergen(ingredients []string, allergen string) bool{
	fmt.Println("[Checking for allergen]")
	for i := 0; i < len(ingredients); i++{
		// fmt.Println("Ingredient: " + ingredients[i] + " Allergen: " + allergen)
		if(ingredients[i] == allergen){
			// fmt.Println("true")
			return true
		}
	}
	return false
}

// critiques your favorite fruit, prompts for user-input
// switch-case statement example
func favFruit(){
	var yourFav string = ""
	var myFav string = "watermelon"

	fmt.Print("What's your favorite fruit? ")
	fmt.Scanln(&yourFav) // pass in yourFav's address to modify with user input

	switch yourFav{
		case myFav:
			fmt.Println("The only correct answer :)")
		case "banana":
			fmt.Println("Let's make banana splits!")
		case "apple":
			fmt.Println("Macintosh or bust!")
		case "pineapple":
			fmt.Println("How tropical.")
		case "lemon":
			fmt.Println("Weird, though I admire it.")
		default:
			fmt.Println("You chose...Poorly.")
	}
}

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
