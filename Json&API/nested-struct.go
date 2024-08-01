// Use nested structs and print the number of message that are of type string and the number of  
// message that are of type int.
// {
//   "alerts": [
//     {
//       "message": "hello world"
//     },
//     {
//       "message": 911
//     },
//     {
//       "message": "hello india"
//     }
//   ]
// }
// Expected output: string messages-2, int messages-1

package main

import (
    "encoding/json"
    "fmt"
)

type Alert struct {
    Message interface{} `json:"message"`
}

type Alerts struct {
    Alerts []Alert `json:"alerts"`
}

func main() {
    data := []byte(`
    {
        "alerts": [
            {
                "message": "hello world"
            },
            {
                "message": 911
            },
            {
                "message": "hello india"
            }
        ]
    }`)

    var alerts Alerts
    if err := json.Unmarshal(data, &alerts); err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }

    stringCount := 0
    intCount := 0

    for _, alert := range alerts.Alerts {
        switch alert.Message.(type) {
        case string:
            stringCount++
        case float64: // JSON numbers are unmarshalled as float64
            intCount++
        }
    }

    fmt.Printf("string messages-%d, int messages-%d\n", stringCount, intCount)
}

// Follow-up question on this, what if the stuct is not given to you then how will you find alerts?

// If you receive a JSON payload and do not have a predefined struct to unmarshal it into, 
// you can use Go's map or interface{} types to handle the JSON data dynamically
// If the JSON data is a JSON object (a collection of key-value pairs), you can unmarshal it into a 
// map[string]interface{}. This allows you to work with the data dynamically, where the keys are strings and 
// the values are of type interface{}

package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    data := []byte(`
    {
        "alerts": [
            {
                "message": "hello world"
            },
            {
                "message": 911
            },
            {
                "message": "hello india"
            }
        ]
    }`)

    var result map[string]interface{}
    if err := json.Unmarshal(data, &result); err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }

    // Accessing the "alerts" field
    alerts, ok := result["alerts"].([]interface{})
    if !ok {
        fmt.Println("Error: 'alerts' field is not an array")
        return
    }

    stringCount := 0
    intCount := 0

    for _, alert := range alerts {
        alertMap, ok := alert.(map[string]interface{})
        if !ok {
            fmt.Println("Error: alert is not a map")
            continue
        }

        message, ok := alertMap["message"]
        if !ok {
            fmt.Println("Error: 'message' field is missing")
            continue
        }

        switch message.(type) {
        case string:
            stringCount++
        case float64: // JSON numbers are unmarshalled as float64
            intCount++
        }
    }

    fmt.Printf("string messages-%d, int messages-%d\n", stringCount, intCount)
}

// If the JSON data is more complex or its structure is unknown, you can unmarshal it into 
// interface{} and then use type assertions to work with the data.

package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    data := []byte(`
    {
        "alerts": [
            {
                "message": "hello world"
            },
            {
                "message": 911
            },
            {
                "message": "hello india"
            }
        ]
    }`)

    var result interface{}
    if err := json.Unmarshal(data, &result); err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }

    // Assert the result as a map
    resultMap, ok := result.(map[string]interface{})
    if !ok {
        fmt.Println("Error: JSON is not a map")
        return
    }

    // Access the "alerts" field
    alerts, ok := resultMap["alerts"].([]interface{})
    if !ok {
        fmt.Println("Error: 'alerts' field is not an array")
        return
    }

    stringCount := 0
    intCount := 0

    for _, alert := range alerts {
        alertMap, ok := alert.(map[string]interface{})
        if !ok {
            fmt.Println("Error: alert is not a map")
            continue
        }

        message, ok := alertMap["message"]
        if !ok {
            fmt.Println("Error: 'message' field is missing")
            continue
        }

        switch message.(type) {
        case string:
            stringCount++
        case float64: // JSON numbers are unmarshalled as float64
            intCount++
        }
    }

    fmt.Printf("string messages-%d, int messages-%d\n", stringCount, intCount)
}
