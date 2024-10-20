package main

import "fmt"

func luas_persegi(s int) {
	luas := s * s
	fmt.Println("Jadi luas persegi adalah", luas)
}
func keliling_persegi(s int) {
	keliling := 4 * s
	fmt.Println("Jadi keliling persegi adalah", keliling)
}
func main() {
	var s int
	fmt.Print("Masukkan sisi persegi : ")
	fmt.Scan(&s)
	luas_persegi(s)
	keliling_persegi(s)
}
