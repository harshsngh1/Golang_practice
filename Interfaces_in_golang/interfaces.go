package main

import "fmt"

type value struct { // value is a struct that has two int fields
	first  int
	second int
}

type mathTest interface { // mathTest is an interface that has a method called add which takes two int parameters and returns an int
	add(a, b int) int
}

func (v value) add(a, b int) int {
	return a + b + v.first + v.second
}

func main() {
	var te mathTest = value{1, 2}
	fmt.Println(te.add(1, 2))
}

/*
Explanation :
var te mathTest = value{1,2}
Is line mein hum ek variable te declare kar rahe hain jo mathTest type ka hai (yaani ki ek interface).
Phir hum us variable ko value{1, 2} se assign kar rahe hain, jo ek struct instance hai jisme first = 1 aur second = 2.
Yeh kaise kaam karta hai?
- value{1,2} ek struct instance hai jisme first ki value 1 aur second ki value 2 hai.
- mathTest ek interface hai jisme ek method add(a, b int) int defined hai.
Kyunki value struct mein bhi ek add method hai jo isi signature ko follow karta hai,
toh value struct mathTest interface ko implement karta hai.
- Go language mein interface ka implementation implicit hota hai,
yaani agar ek struct ke paas wohi methods hain jo interface ko chahiye,
toh wo us interface ko implement kar raha hai bina explicitly bataye.
Iska matlab, ab te ek value type ka instance hai, lekin use mathTest interface ke taur pe treat kiya ja raha hai.
*/
