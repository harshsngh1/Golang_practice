package main

import (
	"encoding/json"
	"fmt"
)

type Alert struct {
	Message interface{} `json:"message"`
}

// nested struct
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

//ANOTHER WAY TO DO SO

// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// type Alert struct {
//     Message interface{} `json:"message"`
// }

// type Alerts struct {
//     Alerts []Alert `json:"alerts"`
// }

// func main() {
//     // JSON data as a string
//     jsonString := `{
//         "alerts": [
//             {"message": "hello world"},
//             {"message": 911},
//             {"message": "hello india"}
//         ]
//     }`

//     var alerts Alerts
//     err := json.Unmarshal([]byte(jsonString), &alerts)
//     if err != nil {
//         fmt.Println("Error unmarshalling JSON:", err)
//         return
//     }

//     stringCount := 0
//     intCount := 0

//     for _, alert := range alerts.Alerts {
//         switch alert.Message.(type) {
//         case string:
//             stringCount++
//         case float64: // JSON numbers are unmarshalled as float64
//             intCount++
//         }
//     }

//     fmt.Printf("string messages-%d, int messages-%d\n", stringCount, intCount)
// }
