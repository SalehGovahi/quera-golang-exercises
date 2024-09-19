package main

import "fmt"

func main() {
	var coins int

	fmt.Scanf("%d",&coins)
	
	var tax float64
	tax = 0.00

	if coins <= 100 {
		tax += float64(coins) / 100 * 5
	} else if coins > 100 && coins <= 500 {
		coins = coins - 100
		tax += float64(coins) / 100 * 10 + 5
	} else if coins > 500 && coins <= 1000 {
		coins = coins - 500
		tax += float64(coins) / 100 * 15 + 5 + 40
	} else if coins > 1000 {
		coins = coins - 1000
		tax += float64(coins) / 100 * 20 + 5 + 40 + 75
	}

	fmt.Printf("%d", int(tax))
}
