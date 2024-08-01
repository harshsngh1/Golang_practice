# Questions

## what is marshal and unmarshal in golang ?
Marshalling is the process of converting a Go data structure (such as a struct, slice, or map) into a JSON (or other format) string. This is useful for sending data over the network, saving to a file, or otherwise serializing data.
- Function: json.Marshal
- Purpose: Converts Go objects into JSON data
- Return Value: A []byte containing the JSON data and an error (if any).
```
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    p := Person{Name: "Alice", Age: 30}

    // Marshal the Go struct into JSON
    jsonData, err := json.Marshal(p)
    if err != nil {
        fmt.Println("Error marshalling JSON:", err)
        return
    }

    fmt.Println(string(jsonData)) // Output: {"name":"Alice","age":30}
}
```
Unmarshalling is the process of converting JSON data (or other formats) into Go data structures. This is useful for reading data from files, APIs, or other sources where data is in JSON format.
- Function: json.Unmarshal
- Purpose: Converts JSON data into Go objects
- Parameters: JSON data as []byte and a pointer to the Go data structure where the unmarshalled data will be stored.
- Return Value: An error (if any).
```
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    jsonData := []byte(`{"name":"Alice","age":30}`)

    var p Person

    // Unmarshal JSON data into the Go struct
    err := json.Unmarshal(jsonData, &p)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }

    fmt.Println(p) // Output: {Alice 30}
}
```