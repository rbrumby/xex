| function | args | description
| -------- | ---- | -----------
| add |> num2: The second number to add.<br/>> num1: The first number to add.<br/>| adds two numbers returning a single numerical result|
| addOrConcat |> val1: The first value to add / concat.<br/>> val2: The second value to add / concat.<br/>| Chooses to call add or concat depending if args are numeric or not.|
| and |> val1: The first bool value<br/>> val2: The second bool value<br/>| Returns true (bool) if both inputs are true, else false.|
| concat |> strs: variadic - the strings to concatentate.<br/>| concatenates any number of strings returning a single string result|
| count |> in: The number of elements in the collection.<br/>| Returns the number of elements in the passed in slice / array or map.|
| divide |> divisor: The number to divide by.<br/>> dividend: The number to be divided.<br/>| divides two numbers returning a single numerical result|
| entry |> key: The map entry key.<br/>> value: The map entry value.<br/>| Creates a map entry with the passed in key & value.|
| equals |> val1: The first value to compare<br/>> val2: The second value to compare<br/>| compares 2 inputs returning a bool|
| float32 |> number: The number to convert.<br/>| float32 converts the passed in value to an float32 or returns a error if conversion isn't possible|
| float64 |> number: The number to convert.<br/>| float64 converts the passed in value to an float64 or returns a error if conversion isn't possible|
| greaterThan |> val1: The first value.<br/>> val2: The second value.<br/>| Returns the result of val1 > val2. Values must be numeric or string.|
| greaterThanEqual |> val2: The second value.<br/>> val1: The first value.<br/>| Returns the result of val1 >= val2. Values must be numeric or string.|
| indexOf |> coll: The collection (array, slice or map) from which to extract a value.<br/>> index: The index / key to extract from coll<br/>| Returns the entry from the passed collection at the requested index.|
| instring |> input: The string to search.<br/>> search: The string to find in the input.<br/>| returns the start position in the input string of the search string or -1 if the search string is not found|
| int || int converts the passed in value to an int or returns a error if conversion isn't possible|
| int16 |> number: The number to convert.<br/>| int16 converts the passed in value to an int16 or returns a error if conversion isn't possible|
| int32 |> number: The number to convert.<br/>| int32 converts the passed in value to an int32 or returns a error if conversion isn't possible|
| int64 |> number: The number to convert.<br/>| int64 converts the passed in value to an int64 or returns a error if conversion isn't possible|
| int8 |> number: The number to convert.<br/>| int8 converts the passed in value to an int8 or returns a error if conversion isn't possible|
| len |> in: The string to measure.<br/>| returns the length of a string|
| lessThan |> val1: The first value.<br/>> val2: The second value.<br/>| Returns the result of val1 < val2. Values must be numeric or string.|
| lessThanEqual |> val1: The first value.<br/>> val2: The second value.<br/>| Returns the result of val1 <= val2. Values must be numeric or string.|
| map |> values: variadic - any number of MapEntry's can be passed to be built into a Map. Types must be compatible with the first value passed.<br/>| Makes a new map containing the passed in mapEntry values. 				The type of the map (key / value) created is determined by the types passed in the first element of values.|
| mod |> dividend: The number to be divided.<br/>> divisor: The number to divide by.<br/>| mod returns the remainder of dividend divided by divisor.|
| multiply |> multiplicand: The number to be multiplied.<br/>> multiplier: The number to multiply by.<br/>| multiplies two numbers returning a single numerical result|
| nil |> value: The value which will be returned as this function does nothing!<br/>| Returns what is passed - used to implement parenthesis grouping|
| not |> value: The value to invert.<br/>| Accepts a boolean & returns its inverse|
| notEquals |> val1: The first value to compare.<br/>> val2: The second value to compare.<br/>| Compares 2 inputs returning a bool.|
| or |> val1: The first bool value<br/>> val2: The second bool value<br/>| Returns true (bool) if either or both inouts are true, else false.|
| pow |> x: The base number.<br/>> y: The exponent (number of times x is multiplied by itself).<br/>| pow returns x to the power of y (x**y).|
| select |> coll: The collection (array, slice or map) to select from.<br/>> forEach: The name by which we will refer to each entry in coll<br/>> expr: An expression (Node) to apply using to each value in coll. MUST return a bool (true or false).<br/>> refs: An optional list values () which can be referenced as $0, $1, etc within the expression.<br/>| Returns the elements in the passed in collection (slice / array or map) for which expression evaluates to true. 				If an array is passed in, it is returned as a slice. 				If coll refers to a map, expression is evaluated on the map value, not the key. 				Example: 				//BookList is a collection. For each "book" in the list, we want to evaluate the equals Expression. 				//We also pass enother evaluated value SelectedAuthor which will be accessible as $0 in our expression. 				select(root.BookList, "book", "equals(book.Author, $0)", root.SelectedAuthor)|
| slice |> values: variadic - any number of values can be passed to be built into a slice. Types must be compatible with the first value passed.<br/>| Makes a new slice containing the passed in values. The type of slice created is determined by the type passed in the first element of values. 				slice can be used to create a list of values to test against - is myproperty x, y or z?: select(slice("x", "y", "z"), .myproperty) > 0|
| string |> in: The value to convert to a string.<br/>| Converts an input into a string using fmt.Sprint|
| substring |> input: The string take take a substring from.<br/>> start: The start index (counting from 0).<br/>> end: The end index. If this is less than 1, defaults to the end of the string.<br/>| returns the substring of the input string from index1 to index2 -1. If index2 is zero, everything to the end of the string is returned|
| subtract |> minuend: The initial number to subtract from.<br/>> subtrahend: The value to subreact from minuend.<br/>| subtracts two numbers returning a single numerical result|
| switch |> values: variadic - the value to test then alternate if/else pairs and finally an optional else value<br/>| Switches on the first value. 				The following values are equivalent to "case : result" pairs. 				If a final value is provided (an even number of arguments is passed in total), the final value is used as the default. 				If value1 equals value2, value3 is returned. Else if value1 equals value4, value5 is returned. And so on. 				If there is no default and no values matched, switch returns nil.|
| uint |> number: The number to convert.<br/>| uint converts the passed in value to an uint or returns a error if conversion isn't possible|
| uint16 |> number: The number to convert.<br/>| uint16 converts the passed in value to an uint16 or returns a error if conversion isn't possible|
| uint32 |> number: The number to convert.<br/>| uint32 converts the passed in value to an uint32 or returns a error if conversion isn't possible|
| uint64 |> number: The number to convert.<br/>| uint64 converts the passed in value to an uint64 or returns a error if conversion isn't possible|
| uint8 |> number: The number to convert.<br/>| uint8 converts the passed in value to an uint8 or returns a error if conversion isn't possible|
