// Perulangan
/* for i := 5; i >= 0; i-- {
	for j := i; j >= 0; j-- {
		fmt.Print(j, " ")
	}
	fmt.Println()
} */

// Perulangan Array
/* var fruits = [4]string{"banana", "apple", "orange", "grape"}
for i, fruit := range fruits {
	fmt.Printf("%d : %s, ", i, fruit)
} */

// Slice
// films := []string{"GOTG3", "The Marvels", "Secret Invasion"}

/* for index, film := range films {
	fmt.Println(index+1, film)
} */

package main

import "fmt"

func main() {

	// Slice of Map
	films := []map[string]string{
		{"name": "GOTG3", "year": "2024"},
		{"name": "The Marvels", "year": "2023"},
		{"name": "Secret Invasion", "year": "2023"},
		{"name": "Antman and The Wasp : Quantumania", "year": "2023"},
	}

	for index, film := range films {
		fmt.Println(index+1, "Film Title :", film["name"], "release on", film["year"])
	}

	fmt.Println(len(films))
}
