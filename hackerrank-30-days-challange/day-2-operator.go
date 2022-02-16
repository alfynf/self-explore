// Day 2: Operators

// Task: Given the meal price (base cost of a meal), tip percent (the percentage of the meal price being added as tip), and tax percent (the percentage of the meal price being added as tax) for a meal.
// find and print the meal's total cost.
// Round the result to the nearest integer.

package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'solve' function below.
 *
 * The function accepts following parameters:
 *  1. DOUBLE meal_cost
 *  2. INTEGER tip_percent
 *  3. INTEGER tax_percent
 */

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
	// Write your code here
	cost := meal_cost + (float64(tip_percent) / 100 * meal_cost) + (float64(tax_percent) / 100 * meal_cost)
	fmt.Println(int(math.Round(cost)))
}

func main() {
	solve(12.00, 20, 8)
}
