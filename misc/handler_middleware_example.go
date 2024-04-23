package main

import "fmt"

// The Problem

// Part 1
// Imagine you have an object, that has a method that takes any type of function
// as its argument
// Notice we don't have control over this object since it is part of another
// package

type Operation struct {}

func (o *Operation) DoOperation(fn func(int) int){
	// We don't have control of the inner implementation, we just know that it will do whatever
	// the function we pass does
	result := fn(10)
	fmt.Println("Result of the operation: ", result)
}

func MultiplyByTwo(val int) int {
	fmt.Println("Multiplying by two")
	return val * 2
}

func DivideByTwo(val int) int {
	fmt.Println("Dividing by two")
	return val / 2
}

// Part 2
// I don't always want the function to do the same since in my application I have a point
// in which I cannot do more things.
// In this case, my application is a Math object with two element

type Math struct {
	operation *Operation
	canDoMore bool
}

// The idea will be to modify each operation function to return 0 if 
// canDoMore is set up to false. But we have three problems
// 
// 1. The functions we defined above are outside the scope to check
// the value of canDoMore
// 2. Even if we could, we would need to write the same condition to all functions
// 3. We cannot edit the DoOperation as it is part of a package.

// The Solution
// What if we create a method for the Math struct that can mimic the DoOperation 
// method from Operations, in fact, it will event use it
// This method now will be able to check canDoMore

func (m *Math) DoOperation(fn func(int) int){ // Notice we still take the same type of argument

	// We can now instead of passing the fn to the Operation.DoRequest
	// create a new function that implements the additional logic
	newFunc := func(val int) int {
		// New Logic
		// Checks if we can't do more
		if !m.canDoMore {
			return -1
		}

		// If it is true, then just do what the original function did
		return fn(val)
	}

	m.operation.DoOperation(newFunc) // DoRequest Now takes this function
}

// The code above kind of follows the concept of a middleware.
// We are not overwriting the function because we cannot modify it.
// We are just creating a function with the same signature as the 
// operation.DoOperation expects with additional logic.

func main() {

	// Part 1
	operation := &Operation{}

	operation.DoOperation(MultiplyByTwo) // 20
	operation.DoOperation(DivideByTwo) // 5

	math := &Math{
		operation: operation,
		canDoMore: false,
	}

	math.DoOperation(MultiplyByTwo) // -1
	math.DoOperation(DivideByTwo) // -1

	math.canDoMore = true

	math.DoOperation(MultiplyByTwo) // -20
	math.DoOperation(DivideByTwo) // 5

}