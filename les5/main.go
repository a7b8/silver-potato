package main

import "fmt"

func main() {
	appleCost := 5.99
	pearCost := 7
	totalOfMoney := 23

	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? ", 9*appleCost+float64(8*pearCost))

	fmt.Println("Скільки груш ми можемо купити? ", totalOfMoney/pearCost)
	fmt.Println("Скільки яблук ми можемо купити? ", int(float64(totalOfMoney)/appleCost))
	fmt.Printf("Чи ми можемо купити 2 груші та 2 яблука? %t\n", (float64(2*pearCost)+2*appleCost) <= float64(totalOfMoney))
}
