package main

func addHundred(x int) int {
	return x + 100
}

// high order func and variadic func
// partialsum has return type as func(), and this func() is anonymous func as it does not takes any arguments
// and the return type of this anonymous function is int
func partialSum(x ...int) func() int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return func() int {
		//fmt.Println(addHundred(sum))
		return addHundred(sum)
	}
}
func main() {
	partial := partialSum(1, 2, 3)
	partial()
	// there will be no output
}
