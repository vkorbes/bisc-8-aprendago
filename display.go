package main

import "fmt"

func drawSprite(m *data, x, y, height int, address uint16) bool {
	colisão := false
	if x > 63 {
		return false
	}
	if y > 31 {
		return false
	}
	//  0         1        2...
	//	|xxxxxxxxx|xxxxxxxx|
	//              x <- aqui
	byteNaLinha := x / 8
	offsetProBit := x % 8
	// quero escrever 5 bytes começando em...
	//                  x <- aqui
	//	|xxxxxxxxx|xxxxxxxx|xx...
	//  0         1        2...
	// escrever denovo, com offset "ao contrário"

	// vamos ler x bytes, onde cada byte é uma unidade de altura
	for byte := 0; byte < height; byte++ {
		byteAddress := address + uint16(byte)
		line := (y + byte) % 32
		fb1 := line*8 + byteNaLinha
		fb2 := line*8 + ((byteNaLinha + 1) % 8)
		spr1 := m.ram[byteAddress] >> offsetProBit
		spr2 := m.ram[byteAddress] << (8 - offsetProBit)
		if m.frameBuffer[fb1]&spr1 != 0 {
			colisão = true
		}
		if m.frameBuffer[fb2]&spr2 != 0 {
			colisão = true
		}
		m.frameBuffer[fb1] ^= m.ram[byteAddress] >> offsetProBit
		m.frameBuffer[fb2] ^= m.ram[byteAddress] << (8 - offsetProBit)
	}
	return colisão
}

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
		fmt.Println(line, "\t", string(lineBuffer[:]))
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
