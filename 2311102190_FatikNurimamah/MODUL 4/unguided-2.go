package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Prosedur menghitung soal yang selesai (<301 menit) dan total waktu
func HitungNilai(waktuSoal []int, SoalyangSelesai *int, totalWaktu *int) {
	*SoalyangSelesai = 0
	*totalWaktu = 0

	for _, waktu := range waktuSoal {
		if waktu < 301 {  // Hanya soal < 301 menit yang dihitung
			*SoalyangSelesai++
			*totalWaktu += waktu
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var namaPemenang string
	var soalTerbanyak, waktuTerkecil int
	waktuTerkecil = 9999

	fmt.Println("Masukkan data peserta (Nama dan 8 waktu), ketik 'Selesai' untuk berhenti:")

	for scanner.Scan() {
		input := scanner.Text()
		if input == "Selesai" {
			break
		}

		parts := strings.Fields(input)
		namaPeserta := parts[0]
		waktuSoal := make([]int, 8)

		for i := 0; i < 8; i++ {
			waktu, _ := strconv.Atoi(parts[i+1])
			waktuSoal[i] = waktu
		}

		var jumlahSoal, totalWaktu int
		HitungNilai(waktuSoal, &jumlahSoal, &totalWaktu)

// Tentukan pemenang berdasarkan soal terbanyak, jika sama lihat waktu terkecil
		if jumlahSoal > soalTerbanyak || (jumlahSoal == soalTerbanyak && totalWaktu < waktuTerkecil) {
			namaPemenang = namaPeserta
			soalTerbanyak = jumlahSoal
			waktuTerkecil = totalWaktu
		}
	}

// Tampilkan hasil akhir
	fmt.Printf("\n%s %d %d ", namaPemenang, soalTerbanyak, waktuTerkecil)
}
