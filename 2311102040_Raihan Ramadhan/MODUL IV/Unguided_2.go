package main

import (
	"fmt"
)

// Fungsi hitungSkor menghitung jumlah soal yang diselesaikan dan menambahkan total waktu ke parameter skor dan waktu
func hitungSkor(soal *[8]int, skor *int, totalWaktu *int) {
	*skor = 0
	*totalWaktu = 0

	for i := 0; i < 8; i++ {
		if soal[i] < 301 {
			*skor++
			*totalWaktu += soal[i] // Tambahkan waktu soal yang valid
		}
	}
}

func main() {
	var jumlahPeserta int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	namaPeserta := make([]string, jumlahPeserta)
	skorPeserta := make([]int, jumlahPeserta)
	totalWaktuPeserta := make([]int, jumlahPeserta)

	for i := 0; i < jumlahPeserta; i++ {
		fmt.Printf("Masukkan nama peserta %d: ", i+1)
		fmt.Scan(&namaPeserta[i])

		var soal [8]int
		fmt.Printf("Masukkan waktu penyelesaian 8 soal (dalam menit) untuk %s: ", namaPeserta[i])
		for j := 0; j < 8; j++ {
			fmt.Scan(&soal[j])
		}

		// Panggil fungsi hitungSkor tanpa return, gunakan parameter input-output
		hitungSkor(&soal, &skorPeserta[i], &totalWaktuPeserta[i])
	}

	// Menentukan pemenang
	pemenang := 0
	for i := 1; i < jumlahPeserta; i++ {
		if skorPeserta[i] > skorPeserta[pemenang] || (skorPeserta[i] == skorPeserta[pemenang] && totalWaktuPeserta[i] < totalWaktuPeserta[pemenang]) {
			pemenang = i
		}
	}

	// Menampilkan hasil pemenang
	fmt.Printf("Pemenangnya adalah %s dengan %d soal diselesaikan dalam total waktu %d menit\n",
		namaPeserta[pemenang], skorPeserta[pemenang], totalWaktuPeserta[pemenang])
}
