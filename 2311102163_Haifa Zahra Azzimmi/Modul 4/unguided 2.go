package main

import (
    "fmt"
    "math"
)

// Fungsi untuk menghitung skor
func hitungSkor(waktuPengerjaan []int) (int, int) {
    jumlahSoalSelesai := 0
    totalWaktu := 0
    for _, waktu := range waktuPengerjaan {
        if waktu < 301 {
            jumlahSoalSelesai++
            totalWaktu += waktu
        }
    }
    return jumlahSoalSelesai, totalWaktu
}

// Fungsi untuk mencari pemenang
func cariPemenang(peserta map[string][]int) (string, int, int) {
    namaPemenang := ""
    jumlahSoalMax := -1
    skorMin := math.MaxInt64

    for nama, waktuPengerjaan := range peserta {
        jumlahSoalSelesai, totalWaktu := hitungSkor(waktuPengerjaan)
        if jumlahSoalSelesai > jumlahSoalMax || (jumlahSoalSelesai == jumlahSoalMax && totalWaktu < skorMin) {
            namaPemenang = nama
            jumlahSoalMax = jumlahSoalSelesai
            skorMin = totalWaktu
        }
    }
    return namaPemenang, jumlahSoalMax, skorMin
}

func main() {
    var jumlahPeserta int
    fmt.Println("Silakan masukkan jumlah peserta:")
    fmt.Scan(&jumlahPeserta)
    
    peserta := make(map[string][]int)
    
    for i := 0; i < jumlahPeserta; i++ {
        var nama string
        fmt.Printf("Masukkan nama untuk peserta ke-%d:\n", i+1)
        fmt.Scan(&nama)
        
        waktuPengerjaan := make([]int, 8)
        fmt.Printf("Masukkan waktu pengerjaan untuk setiap soal oleh %s:\n", nama)
        for j := 0; j < 8; j++ {
            fmt.Scan(&waktuPengerjaan[j])
        }
        
        peserta[nama] = waktuPengerjaan
    }
    
    // Mencari pemenang
    pemenang, maxSoal, minSkor := cariPemenang(peserta)
    fmt.Printf("Pemenangnya adalah %s dengan jumlah soal %d dan total waktu %d\n", pemenang, maxSoal, minSkor)
}
