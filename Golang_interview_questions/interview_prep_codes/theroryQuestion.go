package main

import "fmt"

func main() {
	s := []int{1, 2}
	adding(s)
	fmt.Println("s from main", s) //s from main [1 5]
}

func adding(s []int) {
	s[1] = 5
	s = append(s, 3)
	fmt.Println("s is", s) //s is [1 5 3]

}

/*
Ques 1 : what will be value of s in main and in adding func and why?
A: In main s is [1 2] and in adding func s is [1 5 3] because in adding func we are not passing s by value but by reference.

Ques 2 : Why the length of S is different in adding func and main func?
A: Because in adding func we are appending 3 to s which is a reference to the original slice s in main func. So the length of s in adding func is 3 and in main
func is 2.

Ques 3 : What will be len and capacity of s in main and in adding func?
A: In main len(s) = 2 and cap(s) = 2 and in adding
func len(s) = 3 and cap(s) = 4
*/

// Another Question
/*

func main(){
a:= [3]string{"a","b","c"}
b:= a[0:1]
what will be values in b?
A: b will be ["a"] as when we do slicing first index is considered but last one is discarded
}
*/

//Another Question
/*
func main(){
something := make([]int,8) -> if we have declare like this then it means that, length and capacity both are 8 and also make keyword allocates
memory for the slice. So all the values will be initialised with 0
something := append(somenthing,1)
fmt.Println(something, len(something), cap(something)) -> predict the output
--> A: [0 0 0 0 0 0 0 0 1] 9 16
GetValue(something)
fmt.Println(something)
--> A: [10 0 0 0 0 0 0 0 1]
}
now lets say we have a function
GetValue(a []int){  --> slices are always passed by reference
a[0]=10
}
*/
