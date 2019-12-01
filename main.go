package main

import "fmt"

func main() {
	// long syntax
	// variable i is pointing @ integer 42 in memory
	var i int
	i = 42
	fmt.Println(i)

	// implicit variable / type declaration
	firstName := "Arthur"
	fmt.Println(firstName)

	/* POINTER DATA TYPE
	Instead of holding the value directly in the variable,
	we hold an address of the location in memory that holds
	the variable for us.

	Our variable is going to point to another location that's going to hold the information that we're interested in.
	*/

	// firstName1 is going to be a pointer to a string value.
	var firstName1 *string = new(string)
	// NOTE: a pointer that doesn't point to anything has value of <nil>; it is an uninitialized pointer.

	// Dereference Pointer (When you want to access the data/value in the memory that the pointer pointes to)
	*firstName1 = "Arthur"
	// to dereference a pointer, you're basically saying, I've got this pointer that's pointing @ some data. I want you to reach through the pointer, grab that data, and give that back to me. That's called dereferencing

	fmt.Println(firstName1)  // address in memory
	fmt.Println(*firstName1) // "Arthur"

	// I specify pointers by preceding with an asterisk *,
	// and I dereference pointers by preceding the declared variable with an asterisk *

	// =================

	// & - Address Of - Operator

	firstName2 := "Stevie"
	fmt.Println(firstName2)

	// create a pointer thats pointing to the firstName2 variable
	ptr := &firstName2
	fmt.Println(ptr, *ptr) // memory address = 1 - dereference value = "Stevie"

	// change the value stored @ firstName2
	firstName2 = "Tricia"
	fmt.Println(ptr, *ptr) // memory address = 1 (ptr address does not chance) = updated dereference value = "Tricia"

	/*
		NOTE:
		Why is this?
		Because the firstName2 variable is still in the same location.
		We just changed the data that is stored there.
		The pointer is pointing at firstName2,
		so it's storing the same memory address.
		But since we're dereferencing the variable after it has been changed,
		the pointer dereferences to that new value as well.
		The value of the pointer changes, but the memory address does not.
	*/

}
