package main

import "fmt"

func main() {
	// long syntax
	// variable i is pointing @ integer 42 in memory
	var i int
	i = 42
	fmt.Println(i)

	// implicit
	firstName := "Arthur"
	fmt.Println(firstName)

	/* POINTER DATA TYPE
	Instead of holding the value directly in the variable,
	we hold an address of the location in memory that holds
	the variable for us.

	Our variable is going to point to another location that's going to hold the information that we're interested in.
	*/

	// firstName1 is going to be a pointer to a string value.
	var firstName1 *string
	// NOTE: a pointer that doesn't point to anything has value of <nil>; it is an uninitialized pointer.

	// Dereference Pointer
	*firstName1 = "Arthur"
	// to dereference a pointer, you're basically saying, I've got this pointer that's pointing @ some data. I want you to reach through the pointer, grab that data, and give that back to me. That's called dereferencing

}
