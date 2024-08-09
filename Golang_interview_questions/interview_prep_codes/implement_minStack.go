package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	mainStack []int
	minStack  []int
}

// Constructor initializes a new MinStack
func Constructor() MinStack {
	return MinStack{
		mainStack: []int{},
		minStack:  []int{math.MaxInt64},
	}
}

// Push adds an element onto the stack
func (s *MinStack) Push(x int) {
	s.mainStack = append(s.mainStack, x)
	if x <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, x)
	}
}

// Pop removes the top element from the stack
func (s *MinStack) Pop() {
	if s.Top() == s.minStack[len(s.minStack)-1] {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
	s.mainStack = s.mainStack[:len(s.mainStack)-1]
}

// Top returns the top element of the stack
func (s *MinStack) Top() int {
	return s.mainStack[len(s.mainStack)-1]
}

// MinElement returns the minimum element in the stack
func (s *MinStack) MinElement() int {
	return s.minStack[len(s.minStack)-1]
}

func main() {
	minStack := Constructor()

	minStack.Push(3)
	minStack.Push(5)
	fmt.Println("Minimum element:", minStack.MinElement()) // Output: 3

	minStack.Push(2)
	minStack.Push(1)
	fmt.Println("Minimum element:", minStack.MinElement()) // Output: 1

	minStack.Pop()
	fmt.Println("Minimum element:", minStack.MinElement()) // Output: 2

	fmt.Println("Top element:", minStack.Top()) // Output: 2

	minStack.Pop()
	fmt.Println("Minimum element:", minStack.MinElement()) // Output: 3
}

// equivalent cpp code
// #include <bits/stdc++.h>
// using namespace std;

// class MinStack {
// private:
//     stack<int> mainStack;
//     stack<int> minStack;

// public:
//     MinStack() {
//         minStack.push(INT_MAX);
//     }

//     void push(int x) {
//         mainStack.push(x);
//         if (x <= minStack.top()) {
//             minStack.push(x);
//         }
//     }

//     void pop() {
//         if (mainStack.top() == minStack.top()) {
//             minStack.pop();
//         }
//         mainStack.pop();
//     }

//     int top() {
//         return mainStack.top();
//     }

//     int minelement() {
//         return minStack.top();
//     }
// };

// int main() {
//     MinStack minStack;

//     minStack.push(3);
//     minStack.push(5);
//     cout << "Minimum element: " << minStack.minelement() << endl; // Output: 3

//     minStack.push(2);
//     minStack.push(1);
//     cout << "Minimum element: " << minStack.minelement() << endl; // Output: 1

//     minStack.pop();
//     cout << "Minimum element: " << minStack.minelement() << endl; // Output: 2

//     cout << "Top element: " << minStack.top() << endl; // Output: 2

//     minStack.pop();
//     cout << "Minimum element: " << minStack.minelement() << endl; // Output: 3

//     return 0;
// }
