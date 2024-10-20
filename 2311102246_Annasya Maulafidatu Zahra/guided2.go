package main

import "fmt"

func main() {
	var s int
	fmt.Print("Masukkan Sisi: ")
	fmt.Scan(&s)

	luasPersegi(s)
	kelilingPersegi(s)

}
func luasPersegi(s int) {
	hasil := s * s
	fmt.Println("Hasil Luas: ", hasil)
}

func kelilingPersegi(s int) {
	hasil := 4 * s
	fmt.Print("Hasil keliling: ", hasil)
}
