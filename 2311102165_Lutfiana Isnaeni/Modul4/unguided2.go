// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jumlah soal selesai dan total waktu
func kalkulasiSkor(waktu []int) (int, int) {
	jumlahSelesai := 0
	totalWaktu := 0
	for _, t := range waktu {
		if t < 301 {
			jumlahSelesai++
			totalWaktu += t
		}
	}
	return jumlahSelesai, totalWaktu
}

// Fungsi untuk mencari peserta terbaik
func temukanPemenang(dataPeserta map[string][]int) (string, int, int) {
	namaPemenang := ""
	soalMaksimum := -1
	skorMinimum := math.MaxInt64

	for nama, waktu := range dataPeserta {
		soalSelesai, totalWaktu := kalkulasiSkor(waktu)
		if soalSelesai > soalMaksimum || (soalSelesai == soalMaksimum && totalWaktu < skorMinimum) {
			namaPemenang = nama
			soalMaksimum = soalSelesai
			skorMinimum = totalWaktu
		}
	}
	return namaPemenang, soalMaksimum, skorMinimum
}

func main() {
	var jumlahPeserta int
	fmt.Println("Masukkan jumlah peserta (minimal 1):")
	fmt.Scan(&jumlahPeserta)

	if jumlahPeserta < 1 {
		fmt.Println("Jumlah peserta harus minimal 1.")
		return
	}

	dataPeserta := make(map[string][]int)

	for i := 0; i < jumlahPeserta; i++ {
		var namaPeserta string
		fmt.Printf("Masukkan nama peserta ke-%d:\n", i+1)
		fmt.Scan(&namaPeserta)

		waktuPengerjaan := make([]int, 8)
		fmt.Printf("Masukkan waktu pengerjaan untuk %s (8 soal):\n", namaPeserta)
		for j := 0; j < 8; j++ {
			fmt.Scan(&waktuPengerjaan[j])
		}

		dataPeserta[namaPeserta] = waktuPengerjaan
	}

	// Menentukan pemenang
	pemenang, soalTerbanyak, skorTerendah := temukanPemenang(dataPeserta)
	fmt.Printf("Pemenang: %s, Soal Selesai: %d, Total Waktu: %d\n", pemenang, soalTerbanyak, skorTerendah)
}
