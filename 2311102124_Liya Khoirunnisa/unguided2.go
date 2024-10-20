/* Liya Khoirunnisa - 2311102124 */

package main

import (
	"fmt"
	"strings"
)

// Waktu maksimum (5 jam 1 menit)
const maksWaktu = 301

func hitungSkor(jumlahSoal *int, totalSkor *int, dikerjakan int) {
	var waktu int

	// Perulangan waktu soal
	for i := 0; i < dikerjakan; i++ {
		_, error := fmt.Scan(&waktu)
		if error != nil {
			fmt.Println("Kesalahan input:", error)
			return
		}

		// Hanya menghitung soal yang dikerjakan dalam waktu kurang dari batas maksimum
		if waktu < maksWaktu {
			*totalSkor += waktu
			*jumlahSoal++
		}
	}
}

func main() {
	// Insialisai
	pemenang := ""
	soalTerbanyak := 0
	totalWaktu := maksWaktu

	// Deklarasi variabel
	var namaPeserta string

	for {
		// Meminta input pengguna
		fmt.Print("Masukkan nama peserta: ")
		fmt.Scan(&namaPeserta)

		// Menghentikan untuk menginputkan nama
		if strings.ToLower(namaPeserta) == "selesai" {
			break // Hentikan input jika pengguna mengetik 'selesai'
		}

		// Reset nilai jumlah soal dan total skor
		jumlahSoal := 0
		totalSkor := 0

		// Input soal yang dikerjakan
		fmt.Print("Berapa banyak soal yang dikerjakan dari 8 soal: ")
		var dikerjakan int
		fmt.Scan(&dikerjakan)

		// Cek jumlah soal
		if dikerjakan < 1 || dikerjakan > 8 {
			fmt.Println("Jumlah soal harus antara 1 dan 8.")
			continue
		}

		// Input waktu pengerjaan
		fmt.Print("Masukkan waktu untuk ", dikerjakan, " soal: ")
		hitungSkor(&jumlahSoal, &totalSkor, dikerjakan) // Mengoper jumlah soal yang dikerjakan

		// Menentukan pemenang
		if jumlahSoal > soalTerbanyak || (jumlahSoal == soalTerbanyak && totalSkor < totalWaktu) {
			pemenang = namaPeserta
			soalTerbanyak = jumlahSoal
			totalWaktu = totalSkor
		}
	}

	// Menampilkan hasil akhir
	if pemenang != "" {
		fmt.Printf("Pemenang: %s\n", pemenang)
		fmt.Printf("Jumlah Soal: %d\n", soalTerbanyak)
		fmt.Printf("Total Waktu: %d\n", totalWaktu)
	} else {
		fmt.Println("Tidak ada peserta yang dimasukkan.")
	}
}
