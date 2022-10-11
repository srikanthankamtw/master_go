package main

import "fmt"

func main() {

	// if condition_that_evaluates_to_boolean{
	//      perform action1
	// }else if condition_that_evaluates_to_boolean{
	//      perform action2
	// }else{
	//      perform action3
	// }

	price, inStock := 100, true

	if price >= 80 { // parenthesis are no required to enclose the testing condition
		fmt.Println("Too Expensive")
	}

	if price <= 100 && inStock { //the same with: if price <= 100 && inStock { }
		fmt.Println("Buy it!")
	}

	// In Go there is not such a thing like the Truthiness of a variable.
	// Error:
	// if price {
	//  fmt.Println("We have price!")
	// }

	// only one if branch will be executed
	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else { //executed only once if all the if branches are false (it's optional)
		fmt.Println("It's Expensive!")
	}

	language := "golang"

	switch language {
	case "Python": //values must be comparable (compare string to string)
		fmt.Println("You are learning Python! You don't use { } but indentation !! ")
		// an implicit break is added here
	case "Go", "golang": //compare language with "Go" OR "golang"
		fmt.Println("Good, Go for Go!. You are using {}!")
	default:
		// the default clause the equivalent of the else clause of an if statement
		// and gets executed if no testing condition is true.
		fmt.Println("Any other programming language is a good start!")
	}

	n := 5
	// comparing the result of an expression which is bool to another bool value
	switch true {
	case n%2 == 0:
		fmt.Println("Even!")
	case n%2 != 0:
		fmt.Println("Odd!")
	default:
		fmt.Println("Never here!")
	}

	//** Switch simple statement **//

	// Syntax: statement (n:=10), semicolon and a switch condition
	//(true in this case, we are comparing boolean expressions that return true)
	// we can remove the word "true" because it's the default
	switch n := 10; true {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
