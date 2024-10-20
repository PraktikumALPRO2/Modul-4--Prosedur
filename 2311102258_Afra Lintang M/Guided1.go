package main

import "fmt"

func main() {
    var x, y int
    fmt.Print("Masukkan dua angka: ")
    fmt.Scanln(&x, &y)
   
    if x < y {
        x, y = y, x
    }

    hitungPermutasi(x, y)
}

func hitungPermutasi(n, r int) {
    fmt.Printf("Menghitung P(%d, %d)\n", n, r)
    permutasi(n, r)
}

func faktorial(n int, hasil *int) {
    *hasil = 1
    for j := 2; j <= n; j++ {  
        *hasil *= j
    }
}

func permutasi(n, r int) {
    var factN, factNMinR int
    faktorial(n, &factN)
    faktorial(n-r, &factNMinR)
    fmt.Printf("Hasil permutasi(%d, %d) adalah: %d\n", n, r, factN/factNMinR)
}
