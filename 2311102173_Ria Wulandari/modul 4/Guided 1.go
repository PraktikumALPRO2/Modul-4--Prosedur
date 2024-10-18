package main

import "fmt"

// Prosedur untuk membaca input dan mencetak hasil permutasi
func bacaDanCetakPermutasi() {
    var a, b int
    fmt.Scan(&a, &b)
    if a >= b {
        fmt.Println(permutasi(a, b))
    } else {
        fmt.Println(permutasi(b, a))
    }
}

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
    var hasil int = 1
    for i := 1; i <= n; i++ {
        hasil *= i
    }
    return hasil
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int) int {
    return faktorial(n) / faktorial(n-r)
}

func main() {
    bacaDanCetakPermutasi()
}
