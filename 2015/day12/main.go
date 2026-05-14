package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getSum(data interface{}) int {
	switch v := data.(type) {
	case float64:
		return int(v)
	case map[string]interface{}:
		sum := 0
		for _, value := range v {
			sum += getSum(value)
		}
		return sum
	case []interface{}:
		sum := 0
		for _, item := range v {
			sum += getSum(item)
		}
		return sum
	default:
		return 0
	}
}

func main() {
	file, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var data interface{}

	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Printf("TOTAL SUM:%d\n", getSum(data))
}
