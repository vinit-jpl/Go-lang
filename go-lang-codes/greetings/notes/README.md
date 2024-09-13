## function.

func Hello(name string) string
    * Hello => name of function.
    * (name string) => parameter and it's type.
    * string => retrun type of the function.

In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name. 

## variable declaration.

In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type).

## sprintf

fmt.Sprintf("Hi, %v. Welcome!", name)
Sprintf formats according to a format specifier and returns the resulting string.