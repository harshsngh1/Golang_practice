## Theory about slices :

- Slice is reference to an underlying array. So when me make change in slice it will affect the underlying array also.

- Slices can be declared in this format as well using "make" keyword
var := make([]<data_type>, length, capacity)
example : make([]int, 5, 10)

- To add new element in slice we use "append function"
  working of append function :
  - it take two parameters (slice and then values) and returns a slice with the original + appended values
    func append (s []T, vs.....T) []T (Note : here T is data type and vs is value of same datatype)

  - example :- slice = append(slice, 10,20,30) 

- If original array cannot contain the appended elements then a new array of double capacity is allocated to slice.

- we can append one slice into another using append method with this syntax
  slice = append(slice1, slice2...)

- we can copy one slice into another slice using "copy" method
  example :- num := copy(destination_slice,source_slice)
  - here destination_slice will have elements from source slice and "num" will have number of elements copied.
  - slices should have same datatype to be copied


# Point to note : 
- when we use "make" keyword to initialise a slice then it helps in memory allocation which helps to improve performance
- This creates a slice with the specified length and capacity.
- The capacity allows you to pre-allocate memory for the slice, which can improve performance if you know the expected size of the slice in advance.
- The values of the elements in the slice will be set to their zero values based on the type (e.g., 0 for integers, empty string for strings, nil for pointers).
- In summary, if you already know the expected size of the slice and want to pre-allocate memory for better performance, use make. 
  If you are not sure about the size and want to start with an empty slice that you'll later populate dynamically, 
  use the empty slice syntax slice := []int{}.
