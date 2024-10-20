package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Procedure untuk menghitung skor
func hitungSkor_145(soal_145 *int, skor_145 *int, waktu_145 int) {
	if waktu_145 <= 300 {
		*skor_145 += 1
	}
	*soal_145 += 1
}

func main() {
	var pemenang_145 string
	var soalMax_145, skorMax_145, waktuTotalMax_145 int

	reader := bufio.NewReader(os.Stdin) // Menggunakan reader untuk input nama

	for {
		fmt.Print("\nMasukkan nama peserta (atau 'Selesai' untuk mengakhiri): ")
		nama_145, _ := reader.ReadString('\n') // Membaca hingga newline
		nama_145 = strings.TrimSpace(nama_145) // Trim spasi dan newline

		// Stop the input if the user types "Selesai"
		if nama_145 == "Selesai" {
			break
		}

		var totalSoal_145, skor_145, waktuTotal_145 int

		for j := 0; j < 8; j++ { // Asumsi ada 8 soal
			var waktu_145 int
			fmt.Printf("Masukkan waktu pengerjaan soal ke-%d (dalam menit): ", j+1)
			fmt.Scan(&waktu_145)
			fmt.Scanln() // Tambahkan ini untuk flush input buffer setelah setiap Scan

			// Hitung skor berdasarkan waktu pengerjaan
			hitungSkor_145(&totalSoal_145, &skor_145, waktu_145)

			// Hitung total waktu
			waktuTotal_145 += waktu_145
		}

		// Tentukan pemenang berdasarkan skor tertinggi dan total waktu terendah
		if skor_145 > skorMax_145 || (skor_145 == skorMax_145 && waktuTotal_145 < waktuTotalMax_145) {
			pemenang_145 = nama_145
			skorMax_145 = skor_145
			soalMax_145 = totalSoal_145
			waktuTotalMax_145 = waktuTotal_145
		}
	}

	// Output hasil
	fmt.Printf("\n%s %d %d\n", pemenang_145, soalMax_145, waktuTotalMax_145)
}
