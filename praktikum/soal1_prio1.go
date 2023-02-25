package main

import "fmt"

func main() {

	var luas, a, b, t float32

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&t)
	fmt.Print("Masukkan panjang sisi a: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan panjang sisi b: ")
	fmt.Scanln(&b)

	luas = 0.5 * t * (a + b)

	fmt.Println("Luas Trapesium adalah", luas)

}
