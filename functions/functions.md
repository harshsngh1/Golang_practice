##Theory about functions:

Variadic Functions :
Variadic function are the functions that can take any number of arguments of a define type
Declaration :  func <func_name> (param-1 type, param-2 type,para-3 ...type)<return_type>
Example : func sumNumbers (numbers ...int) 
    - Now this function can take any number of arguments of type "int"
Example : func sumNumber (str string, number ...int)


Anonymous functions : 
- An anonymous function is a function that is declared without any named identifier to refer to it.
- They can accept inputs and return outputs, just as standard functions do.
- They can be used for containing functionality that need not be named and possibly for short-term use.

High order functions : 
- Composition
    - creating smaller functions that take care of certain piece of logic.
    - composing complex function by using different smaller functions.
- Reduces bugs
- Code gets easier to read and understand
- The concept of higher-order functions is powerful and commonly used in functional programming paradigms.

Methods : 
- In Go, a method is a function that is associated with a specific type. It allows you to define behaviors and actions that can be     performed by values of that type. Methods are similar to functions but have a special receiver parameter that binds the method to the type.
-Syntax : 
func (receiver Type) methodName(parameters) returnType {
    // Method implementation
}

Receiver : The receiver is a parameter in the method declaration that specifies the type to which the method is bound. It associates the method with that type. The receiver can be of any named type (including struct types, pointer types, and even user-defined types). You can access the receiver's fields and call other methods on it inside the method's implementation.

Type: The type that the method is bound to. It can be a built-in type (e.g., int, string) or a user-defined type (e.g., a struct that you've defined).


Inferfaces : 
In Go, an interface is a collection of method signatures. It defines a set of methods that a type must implement to be considered as implementing that interface. In other words, an interface specifies a contract that a type must fulfill to be used in a particular context.
Interfaces enable polymorphism in Go, allowing you to write more flexible and reusable code. They help decouple the implementation details from the usage, making it easier to swap different types that satisfy the same interface without changing the code that uses them.