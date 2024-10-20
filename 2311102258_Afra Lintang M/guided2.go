package main

import "fmt"

func main() {
    var sisi int
    fmt.Print("Masukkan panjang sisi persegi: ")
    fmt.Scanln(&sisi)

    tampilkanHasil(sisi)
}

func tampilkanHasil(s int) {
    fmt.Println("Menghitung luas dan keliling persegi dengan sisi:", s)
    fmt.Println("Luas Persegi:", hitungLuas(s))
    fmt.Println("Keliling Persegi:", hitungKeliling(s))
}

func hitungLuas(s int) int {
    return s * s
}

func hitungKeliling(s int) int {
  return 4*s
}
