package genericFunc

import (
	"fmt"
)

func Run() {
	numsFloat := []float64{1.1, 1.2, 1.3, 1.4, 1.5}
	numsInt := []int{1, 2, 3, 4, 5}
	// use function parameter is specified type
	fmt.Println("sum with function parameter is specified type")
	fmt.Printf("sumFloat specified type: %#v\n", sumFloat(numsFloat))
	fmt.Printf("sumInt specified type: %#v\n", sumInt(numsInt))
	// Problem: function sumFLoat and sumInt is a same process
	// Sol: using generic function sum
	// use function parameter is generic type
	fmt.Println("sum with function parameter is generic type")
	fmt.Printf("sumFloat generic type: %#v\n", sumWithGeneric(numsFloat))
	fmt.Printf("sumInt generic type: %#v\n", sumWithGeneric(numsInt))

}

func sumFloat(nums []float64) float64 {
	var result float64
	for _, v := range nums {
		result += v
	}
	return result
}

func sumInt(nums []int) int {
	var result int
	for _, v := range nums {
		result += v
	}
	return result
}

// Problem: function sumFLoat and sumInt is a same process
// Sol: using generic function sum
func sumWithGeneric[T int | float64](nums []T) T {
	var result T
	for _, v := range nums {
		result += v
	}
	return result
}

type Movie struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Price    int    `json:"price"`
}

type Game struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
