package main

import "fmt"

// Prosedur untuk menghitung faktorial
func faktorial(n_145 int, result *int) {
    *result = 1
    for i := 1; i <= n_145; i++ {
        *result *= i
    }
}

// Prosedur untuk menghitung permutasi
func permutasi(n_145, r int, result *int) {
    if n_145 < r {
        *result = 0
        
    }
    var fact_n_145, fact_n_r int
    faktorial(n_145, &fact_n_145)
    faktorial(n_145-r, &fact_n_r)
    *result = fact_n_145 / fact_n_r
}

// Prosedur untuk menghitung kombinasi
func kombinasi(n_145, r int, result *int) {
    if n_145 < r {
        *result = 0
        
    }
    var fact_n_145, fact_r, fact_n_r int
    faktorial(n_145, &fact_n_145)
    faktorial(r, &fact_r)
    faktorial(n_145-r, &fact_n_r)
    *result = fact_n_145 / (fact_r * fact_n_r)
}

func main() {
    var a, b, c, d int
    fmt.Println("Masukkan empat bilangan asli a, b, c, dan d, dengan syarat a >= c dan b >= d:")
    fmt.Scanln(&a, &b, &c, &d)

    var p1, c1, p2, c2 int

    // Menghitung permutasi dan kombinasi
    permutasi(a, c, &p1)
    kombinasi(a, c, &c1)
    permutasi(b, d, &p2)
    kombinasi(b, d, &c2)

    // Menampilkan hasil
    fmt.Println(p1, c1)
    fmt.Println(p2, c2)
}
