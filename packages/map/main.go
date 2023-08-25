package func_map

import "fmt"

func Map() {
	scores := map[string]int{
        "Alice": 90,
        "Bob":   85,
        "Eve":   95,
    }
    fmt.Println("Scores:", scores)
}