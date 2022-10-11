package main

import (
	"fmt"
	"math"
)

func foo() {
	fmt.Println("This is foo()!")
}

func bar() {
	fmt.Println("This is bar()!")
}

func foobar() {
	fmt.Println("This is foobar()!")
}

// defining a function with no parameters
func f1() {
	fmt.Println("This is f1() function")
}

// defining a function with 2 parameters, a and b
func f2(a int, b int) {
	//a and b are local to the function
	fmt.Println("Sum:", a+b)
}

// defining a function using shorthand parameters notation
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// defining a function that have one parameter of type float64 and returns a value of type float64
func f4(a float64) float64 {
	return math.Pow(a, a)
	//any statements below the return statement are never executed

}

// defining a function that have two parameters of type int and returns two values of type int
func f5(a, b int) (int, int) {
	return a * b, a + b
}

// defining a function that have one parameter of type int and returns a "named parameter"
func sum(a, b int) (s int) {
	fmt.Println("s:", s) // -> s is a variable with the zero value inside the function
	s = a + b

	// it automatically return s
	return // This is known as a "naked" return.
}

func main() {

	// calling functions
	f1() // => This is f1() function

	f2(10, 20)

	f3(3, 4, 5, 4., 5.5, "ss") // => 3 4 5 4 5.5 ss

	fmt.Println(f4(5.1))

	a, b := f5(6, 7)
	fmt.Printf("a:%d, b:%d\n", a, b) // => a:42, b:13

	ss := sum(4, 5)
	fmt.Println(ss) // -> 9

	// There are no default arguments in Go //

	// a defer statement defers or postpones the execution of a function until the surrounding function returns.

	// by deferring foo() it will execute it just before exiting the surrounding function which is main()
	defer foo()
	bar()

	fmt.Println("Just a string after deferring foo() and calling bar()")

	// if there are more functions deferred, Go will execute them in the reverse order they were deferred
	defer foobar()

	/*
	   When executing the program the fallowing output is printed out:

	   This is bar()!
	   Just a string after deferring foo() and calling bar()
	   This is foobar()!
	   This is foo()!
	*/

}
