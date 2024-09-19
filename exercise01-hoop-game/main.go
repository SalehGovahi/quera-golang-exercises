package main

import "fmt"

func main() {
	var number1, number2 int

	fmt.Scanf("%d %d",&number1,&number2)
	
	start := 1

	for start <= number2 {
		if start % number1 == 0 {
			hoopCount := start / number1
			for i := 0; i < hoopCount; i++ {
				fmt.Print("Hope ")
			}
			fmt.Printf("\n")
		} else {
			fmt.Printf("%d\n", start)
		}
		start++
	}
}
