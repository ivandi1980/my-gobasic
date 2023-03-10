package main

import (
	"fmt"
	"github.com/ivandi1980/my-basics/arrays"
	"github.com/ivandi1980/my-basics/datatypes"
	"github.com/ivandi1980/my-basics/functions"
	"github.com/ivandi1980/my-basics/variables"
)

func main() {

	/**
	 * Variable
	 * below is the example of the variable
	 */

	// Print the string "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Call the function "ExampleOfVariable" from the package "variables"
	variables.ExampleOfVariable()

	// Call the function "ExampleOfVariableWithParameter" from the package "variables"
	variables.ExampleOfVariableWithParameter("ivan")

	// Call the function "ExampleOfVariableWithMultipleParameters" from the package "variables"
	variables.ExampleOfVariableWithMultipleParameters("ivan", 43, "Jakarta")

	// Call the function "ExampleOfVariableWithMultipleParametersAndReturnValue" from the package "variables"
	name, age, address := variables.ExampleOfVariableWithMultipleParametersAndReturnValue(
		"ivan",
		43,
		"Jakarta",
	)
	fmt.Printf("Variable Example with parameter is : %s, %d, %s\n", name, age, address)

	// Call the function "PrintsAllVariables" from the package "variables"
	variables.PrintsAllVariables()

	/**
	 * Array
	 * below is the example of the Array
	 */

	// Call the function "ExampleOfArray" from the package "arrays"
	arrays.ExampleArray()

	// Call the function "getTheLengthOfArray" from the package "arrays"
	arrays.GetTheLengthOfArray()

	/**
	 * Datatype
	 * below is the example of the Datatype
	 */

	// Call the function "GetStringDataType" from the package "datatypes"
	datatypes.GetStringDataType()

	// Call the function "GetIntDataType" from the package "datatypes"
	datatypes.GetIntDataType()

	// Call the function "Add" from the package "functions"
	functions.Add(1, 2)

	// Call the function "AddWithReturnValue" from the package "functions"
	sum := functions.AddWithReturnValue(1, 2)
	fmt.Println("The sum of the two numbers is :", sum)

	// Call the function "AddWithMultipleReturnValues" from the package "functions"
	sum, sub := functions.AddWithMultipleReturnValues(1, 2)
	_ = sum
	fmt.Println("The sum of the two numbers is :", sum)
	fmt.Println("The sub of the two numbers is :", sub)

	// Call the function "Swap" from the package "functions"
	a, b := functions.Swap("hello", "world")
	fmt.Println(a, b)
}
