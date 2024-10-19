//Haifa Zahra Azzimmi
//2311102163

package main

import (
    "fmt"
)

// Fungsi untuk menghitung faktorial dari sebuah angka
func hitungFaktorial(x int) int {
    if x == 0 || x == 1 {
        return 1
    }
    hasil := 1
    for i := 2; i <= x; i++ {
        hasil *= i
    }
    return hasil
}

// Fungsi untuk menghitung permutasi dari n dan r
func hitungPermutasi(n, r int) int {
    return hitungFaktorial(n) / hitungFaktorial(n-r)
}

// Fungsi untuk menghitung kombinasi dari n dan r
func hitungKombinasi(n, r int) int {
    return hitungFaktorial(n) / (hitungFaktorial(r) * hitungFaktorial(n-r))
}

func main() {
    var angka1, angka2, angka3, angka4 int

    // Meminta pengguna memasukkan 4 angka
    fmt.Println("Masukkan empat angka dengan syarat angka1 >= angka3 dan angka2 >= angka4:")
    fmt.Scan(&angka1, &angka2, &angka3, &angka4)

    // Mengecek apakah syarat angka1 >= angka3 dan angka2 >= angka4 terpenuhi
    if angka1 < angka3 || angka2 < angka4 {
        fmt.Println("Syarat angka1 >= angka3 dan angka2 >= angka4 tidak terpenuhi!")
        return
    }

    // Menghitung permutasi dan kombinasi untuk dua pasangan angka
    permutasiA := hitungPermutasi(angka1, angka3)
    kombinasiA := hitungKombinasi(angka1, angka3)
    permutasiB := hitungPermutasi(angka2, angka4)
    kombinasiB := hitungKombinasi(angka2, angka4)

    // Menampilkan hasil perhitungan
    fmt.Printf("Hasil permutasi dari %d terhadap %d: %d, kombinasi: %d\n", angka1, angka3, permutasiA, kombinasiA)
    fmt.Printf("Hasil permutasi dari %d terhadap %d: %d, kombinasi: %d\n", angka2, angka4, permutasiB, kombinasiB)
}
