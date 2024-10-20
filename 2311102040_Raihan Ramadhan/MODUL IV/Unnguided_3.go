package main

import "fmt"

// Fungsi untuk mencetak deret
func cetakDeret(n int) {
    for n != 1 {
        fmt.Printf("%d ", n)
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
    }
    fmt.Printf("1\n") // mencetak 1 sebagai elemen terakhir
}

func main() {
    // Contoh masukan
    var n int
    fmt.Print("Masukkan nilai awal: ")
    fmt.Scan(&n)
    if n > 0 && n < 1000000 {
        cetakDeret(n)
    } else {
        fmt.Println("Nilai harus antara 1 dan kurang dari 1.000.000")
    }
}
