package main

import "fmt"

func drawDisplay(m *data) {
	// 256 bytes
	// 2048 bits
	// 64 horizontal
	// 32 vertical
	// cada linha horizontal tem 64 bits
	// cada linha horizontal tem 8 bytes
	// * 32 linhas

	black := '█'
	white := ' '
	lineBuffer := []rune{}
	var mask byte = 0b10000000
	for line := 0; line < 32; line++ {
		for column := 0; column < 8; column++ {
			for bit := 0; bit < 8; bit++ {
				// se o bit "numero bit" do byte for 1
				// desenha o pixel
				// 0b00110011
				//    ↑ zero ou um?
				if m.frameBuffer[(line*8)+column]&(mask>>bit) != 0 {
					lineBuffer = append(lineBuffer, black)
					lineBuffer = append(lineBuffer, black)
				} else {
					lineBuffer = append(lineBuffer, white)
					lineBuffer = append(lineBuffer, white)
				}
			}
		}
		fmt.Println(string(lineBuffer[:]))
		lineBuffer = []rune{}
	}

}

// explicação operações AND e SHIFT RIGHT
// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
//
// 	byte1 := 0b00111100
// 	mask1 := 0b10000000
// 	// operação AND (&)
//
//
// 	fmt.Printf("mask: %08b\n", mask1)
// 	// operação RIGHT SHIFT
// 	// pega os bits e "todo mundo pula pra direita"
// 	//   10000000 >> 2
// 	// = 00100000
// 	fmt.Printf("mask: %08b\n\n", mask1>>1)
//
// for oito := 0; oito < 8; oito++ {
// // operacao mágica????
// 	fmt.Printf("%08b\n", byte1)
// 	fmt.Printf("%08b\n", mask1>>oito)
// 	fmt.Printf("%08b\n\n", byte1&(mask1>>oito))
// 	}
// }
//
