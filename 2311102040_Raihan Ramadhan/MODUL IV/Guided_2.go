package main

import "fmt"

var sisi int

func main() {
	fmt.Print("Masukan Panjang Sisi: ")
	fmt.Scan(&sisi)

	var luas, keliling int
	LuasPersegi(sisi, &luas)
	KelilingPersegi(sisi, &keliling)

	fmt.Printf("Luas Persegi adalah: %d\n", luas)
	fmt.Printf("Keliling Persegi adalah: %d\n", keliling)
}

// Prosedur LuasPersegi tanpa return
func LuasPersegi(sisi int, result *int) {
	*result = sisi * sisi
}

// Prosedur KelilingPersegi tanpa return
func KelilingPersegi(sisi int, result *int) {
	*result = 4 * sisi
}
