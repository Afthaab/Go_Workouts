package main

import "fmt"

func main() {
	students := make(map[string]map[string]interface{})
	students["Afthab"] = map[string]interface{}{
		"Institue": "Brototype",
		"Age":      22,
		"Place":    "Bangalore",
	}
	students["Poorvi"] = map[string]interface{}{
		"Institue": "Pentagon",
		"Age":      22,
		"Place":    "Bangalore",
	}
	students["Jeevan"] = map[string]interface{}{
		"Institue": "J Spyder",
		"Age":      22,
		"Place":    "Bangalore",
	}
	fmt.Println("The Details are :")
	for details, v := range students {
		fmt.Printf("The Name is '%s' and the details is '%v' \n", details, v)
	}
}
