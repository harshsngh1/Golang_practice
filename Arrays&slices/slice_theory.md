# Theory about Arrays and Slices :

- Arrays are fixed size, slices are not.
- Arrays are values, slices are references.

1-D Arrays declaration and initialisation : 
```
var arr [5]int
// Declares an array of 5 integers, initialized to zero values

//Declaration and initialisation
var myIntArray = [5]int{1, 2, 3, 4, 5}

//shorthand
myIntArray := [5]int{1, 2, 3, 4, 5}

//Using ... to Infer Length:
arr := [...]int{1, 2, 3, 4, 5} // Length is inferred to be 5
```
2-D Arrays declaration and initialisation :
```
var arr [5][5]int

//Declares a 2-D array of 5x5 integers, initialized to zero values

//Declaration and initialisation


//we can do this via shorthand method as well
    var matrix = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
```

- Slice is reference to an underlying array. So when me make change in slice it will affect the underlying array also.
```
var myIntSlice []int

var := make([]<data_type>, length, capacity)
example : make([]int, 5, 10)//first is length and second is capacity

my2DArray := [][]int{
    {1, 2, 3, 4},
    {5, 6, 7, 8},
    {9, 10, 11, 12},
}
```
- To add new element in slice we use "append function"
  working of append function :
- it take two parameters (slice and then values) and returns a slice with the original + appended values
- it will create a new slice and copy the original slice into it
- it will return the new slice
- it will not modify the original slice
```
slice = append(slice, 10,20,30) 
```
- If original array cannot contain the appended elements then a new array of double capacity is allocated to slice.

- we can append one slice into another using append method with this syntax
```
slice = append(slice1, slice2...)
```
- we can copy one slice into another slice using "copy" method
- The copy function takes two arguments: the destination slice and the source slice. It returns the number
of elements copied, which can be used to determine if the copy was successful.
```
num := copy(destination_slice,source_slice)
```
- here destination_slice will have elements from source slice and "num" will have number of elements copied.
- slices should have same datatype to be copied
- when we use "make" keyword to initialise a slice then it helps in memory allocation which helps to improve performance
- This creates a slice with the specified length and capacity.
- The capacity allows you to pre-allocate memory for the slice, which can improve performance if you know the expected size of the slice in advance.
- The values of the elements in the slice will be set to their zero values based on the type (e.g., 0 for integers, empty string for strings, nil for pointers).
- In summary, if you already know the expected size of the slice and want to pre-allocate memory for better performance, use make. 
- If you don't know the expected size of the slice or want to allocate memory dynamically, use append.
## Operations on slices
- To compare two slices, you need to manually compare their lengths and each corresponding element. Alternatively, you can use the reflect.DeepEqual function from the reflect package for deep comparison.
- Slices are also assignable, which means that you can assign one slice to another slice of the
same type. However, it's important to note that assigning one slice to another slice does not
copy the elements of the slice, it only copies the reference to the underlying array. This means
that if you modify the elements of the slice after the assignment, the original slice will also
be modified.
### Example
```
func main() {
    slice1 := []int{1, 2, 3}
    slice2 := slice1 // Assign slice1 to slice2

    // Modify slice2
    slice2[0] = 10

    // Both slices reflect the change
    fmt.Println("slice1:", slice1) // Output: [10 2 3]
    fmt.Println("slice2:", slice2) // Output: [10 2 3]
}
```
- To sort slices we can use "sort" method from sort package
### Example
```
package main

import (
    "fmt"
    "sort"
)

func main() {
    ints := []int{4, 2, 3, 1, 5}
    floats := []float64{4.2, 2.3, 3.1, 1.5, 5.0}
    strings := []string{"banana", "apple", "cherry"}

    // Sorting slices
    sort.Ints(ints)
    sort.Float64s(floats)
    sort.Strings(strings)

    fmt.Println("Sorted ints:", ints)       // Output: [1 2 3 4 5]
    fmt.Println("Sorted floats:", floats)   // Output: [1.5 2.3 3.1 4.2 5]
    fmt.Println("Sorted strings:", strings) // Output: [apple banana cherry]
}
```
- To reverse a slice we can use "reverse" method from "sort" package
### Example
```
sort.Sort(sort.Reverse(sort.IntSlice(slice)))
```



