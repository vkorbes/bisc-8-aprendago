package main

func main() {
	memory := data{frameBuffer: [256]byte{0b00110011, 0b11000011, 0b11111111, 0b00000111}}
	drawDisplay(&memory)
}
