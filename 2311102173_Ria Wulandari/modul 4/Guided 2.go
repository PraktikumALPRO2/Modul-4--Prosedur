package main

import "fmt"

// Prosedur untuk menghitung luas persegi
func hitungLuasPersegi(sisi float64) float64 {
    return sisi * sisi
}

// Prosedur untuk menghitung keliling persegi
func hitungKelilingPersegi(sisi float64) float64 {
    return 4 * sisi
}

func main() {
    // Meminta input panjang sisi persegi
    var panjangSisi float64
    fmt.Print("Masukkan panjang sisi persegi: ")
    fmt.Scan(&panjangSisi)

    // Menghitung luas dan keliling persegi menggunakan prosedur
    luasPersegi := hitungLuasPersegi(panjangSisi)
    kelilingPersegi := hitungKelilingPersegi(panjangSisi)

    // Menampilkan hasil
    fmt.Printf("Luas persegi: %g\n", luasPersegi)
    fmt.Printf("Keliling persegi: %g\n", kelilingPersegi)
}
