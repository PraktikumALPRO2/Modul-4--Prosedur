package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    if a >= b {
        permutasi(a, b)
    } else {
        permutasi(b, a)
    }
}

// Prosedur faktorial untuk menghitung dan menampilkan faktorial
func faktorial(n int) {
    var hasil int = 1
    var i int
    for i = 1; i <= n; i++ {
        hasil = hasil * i
    }
    fmt.Println("Faktorial dari", n, "adalah:", hasil)
}

// Prosedur permutasi untuk menghitung dan mencetak hasil permutasi
func permutasi(n, r int) {
    faktorialN := 1
    faktorialNR := 1
    // Hitung faktorial n
    for i := 1; i <= n; i++ {
        faktorialN *= i
    }
    // Hitung faktorial (n-r)
    for i := 1; i <= (n - r); i++ {
        faktorialNR *= i
    }
    // Cetak hasil permutasi
    fmt.Println("Permutasi dari", n, "dan", r, "adalah:", faktorialN/faktorialNR)
}
