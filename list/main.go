package main

import "fmt"

func main() {
	// arraySlices()
	// dinamycSlices()
	practice()
}

func arraySlices() {
	prices := [4]float64{10.99, 10.99, 10.99, 10.99}
	var productNames [4]string = [4]string{"Book"}

	// NOTE: Slices keep pointing to selected items to the original arrat item
	featuredPrices := prices[1:3]
	featuredPrices[0] = 100.00

	fmt.Println(featuredPrices)
	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(len(featuredPrices), cap(featuredPrices))
}

func dynamicSlices() {
	prices := []float64{10.99, 8.99}
	prices[1] = 9.99
	updatedPrices := append(prices, 8.99)
	fmt.Println(prices, updatedPrices)

	// NOTE: We use append to add new item to and create a brand new array under the hood
}

func practice() {
	hobbies := [3]string{"Hobbie 1", "Hobbie 2", "Hobbie 3"}
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	hobbies2 := hobbies[:2]
	hobbies3 := hobbies[0:2]
	fmt.Println(hobbies2, hobbies3)

	hobbies3 = hobbies3[1:3]

	fmt.Println(hobbies2, hobbies3)

	goals := []string{"Goal 1.0", "Goal 2"}

	goals[1] = "Goal 2.0"
	goals = append(goals, "Goal 3.0")

	fmt.Println(goals)

	type Product struct {
		id    string
		title string
		price float64
	}

	products := []Product{{"1", "1", 100.0}, {"2", "2", 200.0}}

	products = append(products, Product{"3", "3", 300.0})

	fmt.Println(products)
}
