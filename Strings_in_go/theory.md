# Strings
- Strings are immutable in golang and they are readonly
- String defined using backticks [``] are called raw literals and are different from the ones defined usinging double quotes, as in strings with back ticks the escape charecters like \n or \t are not considered and are treated as part of string only
    ### Example : s := `sdasdsadsadasd \n adsaasasas`
    here above s will have the whole word with it.
- strings package have the following inbuilt methods to trim
```
strings.Trim(<string_name>,<charecter_from_hich_to_trim>)
strings.TrimLeft(<string_name>,<charecter_from_hich_to_trim>)
strings.TrimRight(<string_name>,<charecter_from_hich_to_trim>)
strings.TrimSpace(<string_name>)
```
- strings package have the following inbuilt methods to split
```
strings.Split(<string_name>,<from_where_to_split>)
example : sgtrings.Split(str,",")
strings.SplitN(<string_name>,<from_where_to_split>,<number_of_splits>)
example : strings.SplitN(str,",",2)
```
- we also have strings.Compare method to compare two strings
- strings package have the following inbuilt methods to join
```
strings.Join(<string_name>,<from_where_to_split>)
example : strings.Join(str,",")
strings.JoinN(<string_name>,<from_where_to_split>,<number_of_splits>)
example : strings.JoinN(str,",",2)
```
- The strings.Fields method in Go is used to split a given string into a slice of strings, using one or more consecutive whitespace characters as the delimiter. It returns a slice containing all the fields of the input string, with any leading and trailing whitespace removed.
```
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "  Hello, World!   This is a test string.  "
	fields := strings.Fields(str)
	fmt.Println(fields)
}
```
Output : 
```
[Hello World This is a test string]
```
- strings have compare method as well

    The return value of strings.Compare can be one of the following:

    0 if the strings are equal
    a negative value if the first string is less than the second string
    a positive value if the first string is greater than the second string

    Note that strings.Compare is case-sensitive, so uppercase letters are considered less than lowercase letters.