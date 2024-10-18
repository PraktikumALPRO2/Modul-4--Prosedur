package main

import (
    "fmt"
    "math"
)

// Function to calculate score
func hitungSkor(waktuPengerjaan []int) (int, int) {
    soalSelesai := 0
    totalSkor := 0
    for _, waktu := range waktuPengerjaan {
        if waktu < 301 {
            soalSelesai++
            totalSkor += waktu
        }
    }
    return soalSelesai, totalSkor
}

// Function to determine the winner
func cariPemenang(peserta map[string][]int) (string, int, int) {
    pemenang := ""
    maxSoal := -1
    minSkor := math.MaxInt64
    
    for nama, waktuPengerjaan := range peserta {
        soalSelesai, totalSkor := hitungSkor(waktuPengerjaan)
        if soalSelesai > maxSoal || (soalSelesai == maxSoal && totalSkor < minSkor) {
            pemenang = nama
            maxSoal = soalSelesai
            minSkor = totalSkor
        }
    }
    return pemenang, maxSoal, minSkor
}

func main() {
    var n int
    fmt.Println("Masukkan jumlah peserta:")
    fmt.Scan(&n)
    
    peserta := make(map[string][]int)
    
    for i := 0; i < n; i++ {
        var nama string
        fmt.Printf("Masukkan nama peserta %d:\n", i+1)
        fmt.Scan(&nama)
        
        waktuPengerjaan := make([]int, 8)
        fmt.Printf("Masukkan waktu pengerjaan soal untuk %s:\n", nama)
        for j := 0; j < 8; j++ {
            fmt.Scan(&waktuPengerjaan[j])
        }
        
        peserta[nama] = waktuPengerjaan
    }
    
    // Determine the winner
    pemenang, maxSoal, minSkor := cariPemenang(peserta)
    fmt.Printf("%s %d %d\n", pemenang, maxSoal, minSkor)
}
