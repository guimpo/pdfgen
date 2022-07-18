package data

import (
	"fmt"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

type Fruit struct {
	Name        string  `fake:"{fruit}"`
	Description string  `fake:"{loremipsumsentence:10}"`
	Price       float64 `fake:"{price:1,10}"`
}

func generateFruit() []string {
	var f Fruit
	gofakeit.Struct(&f)

	froot := []string{}
	froot = append(froot, f.Name)
	froot = append(froot, f.Description)
	froot = append(froot, fmt.Sprintf("%.2f", f.Price))
	return froot
}

func FruitList(length int) ([][]string, float64) {
	var fruits [][]string
	var total = 0.0
	for i := 0; i < length; i++ {
		onefruit := generateFruit()
		var priceStr = onefruit[2]
		if f, err := strconv.ParseFloat(priceStr, 64); err == nil {
			total += f
		}

		fruits = append(fruits, onefruit)
	}

	return fruits, total
}
