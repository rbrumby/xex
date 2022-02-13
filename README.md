# xex
e**X**tensible **EX**pression library for Go (golang)

## Overview
xex allows you to write, parse & evaluate expressions dynamically at runtime.

For example, to get the average number of books pers author in a library data structure:
```
count(lib.Books()) / lib.NumAuthors
```

## Supports
- Defining string & numeric literals
- Accessing object properties (including sub-properties and properites of objects returned from method calls)
- Calling methods on objects returned by accessing properties
- Executing ANY function registered with xex

## Usage
To build xex into your application (skipping error checking for brevity):
```
ex, _ := xex.NewStr("concat(myvar.someProp, "_is_xexy)")
//You can also use xex.New to use a *bufio.Reader instead of a string
r, _ := ex.Evaluate(Values{"myvar": anAppVar})
fmt.Println(r)
```
The expression references a top-level value "myvar" which we assign to application variable **anAppVar** when we evaluate the expression.

## Extensibility
xex includes numerous [built-in functions](builtins.md) but is fully extensible - you can add your own functions or any functions from any library.

### Register a function:
```
xex.RegisterFunction(
	NewFunction(
		"intersperse",
        func intersperse(in1, in2 string) string {
	        var b strings.Builder
        	for i, r := range []rune(in1) {
		        b.WriteRune(r)
        		if len([]rune(in2)) > i {
		        	b.WriteRune([]rune(in2)[i])
		        }
	        }
	        return b.String()
        }
	),
)
```
The above  example intersperses the characters from 2 strings & can now be used in an expression:
```
intersperse("Hello","World") //HWeolrllod
```
## Expression Syntax 
-  Literals may be expressed as numbers (with or without decimal points) or strings (enclosed in double quotes).
    - Numbers without decimal points will be parsed as int's
    - Numbers with decimal points will be parsed as float64's
    - It may be necessary to use builtin number conversion functions to convert to the types needed for calling functions (either directly or via binary operators)
    ```
    float64(5) + multiply(float64(7), 3.25) //returns 27.75
    ```
- Standard dot-notation is used to reference variables and their child properties & methods, starting with the top level variable names which are added into the Values provided to the *Expression.Evaluate call
    - xex can only access public properties & methods of an object
    ```
    myvar.SomeProperty.SomeSubProperty
    ```
- Methods & functions are identified by an open parenthesis at teh end of their identifier (with NO whitespace between the name & the open parenthesis)
    - Methods & function arguments are comma-separated within parentheses as they are in most languages 
    ```
    add( myvar.SomeMethod(), 5 )
    ```
- If a method or function returns multiple return values, by default the first one  returned is used. If you want to override this, you can use curly braces to use a different value
    ```
    someFunction(){1} //use the 2nd value returned from the function
    someFunction() //equivalent to writing someFunction{0}
    ```
- Methods & functions which return an error as the last argument have that argument checked during executionand if not nil on any call in the expression, evaluation is terminated & *Expression.Evaluate returns the error

- Array/slice & map indices can be accessed with square brackets
    ```
    concat(mySlice[3], myMap["mykey"])
    ```
- There are lots of example expressions in the unit tests which should give a good flavour of what's possible & how

## Binary Operators
The following binary operators & the built-in functions that they map to are as follows:

| operator |function          | description
| -------  | ---------------- | ----------- 
| +        | addOrConcat      | If both operands are numeric, adds them, else they are concatenated as strings
| -        | subtract         | Subtracts the 2nd operand from the 1st
| *        | multiply         | Multiplies the operands
| /        | divide           | Divides the 1st operand by the 2nd
| ^        | pow              | Raise the 1st operand to the power of the 2nd
| %        | mod              | Calculates the modulus of the 1st operand divided by the 2nd
| ==       | equals           | Returns a boolean indicating if the operands are equal
| !=       | notEquals        | Returns a boolean indicating if the operands are inequal
| >        | greaterThan      | Returns a boolean indicating if the 1st operandis greater than the 2nd
| >=       | greaterThanEqual | Returns a boolean indicating if the 1st operandis greater than or equal to the 2nd
| <        | lessThan         | Returns a boolean indicating if the 1st operandis less than the 2nd
| <=       | lessThanEqual    | Returns a boolean indicating if the 1st operandis less than or equal to the 2nd
| &&       | and              | Performs a logical AND on boolean operands
| \|\|     | or               | Performs a logical OR on boolean operands

## Unary Operators

| operator |function          | description
| -------  | ---------------- | ----------- 
| !        | not              | Inverts  boolean operand