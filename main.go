package main

import "fmt"

func main() {
	memory := data{ram: characters}
	fmt.Println()
	fmt.Println()
	for l := 0; l < 6; l++ {
		drawSprite(&memory, l*5, l*5, 5, uint16(l*5))
	}
	colisão := drawSprite(&memory, 0, 15, 10, uint16(0x0*5))
	drawDisplay(&memory)
	if colisão {
		fmt.Println("deu colisão")
	}
}
