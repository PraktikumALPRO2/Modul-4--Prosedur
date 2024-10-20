package main

import "fmt"

// Prosedur hitungSkor menghitung berapa banyak soal yang diselesaikan dan berapa total waktu yang dibutuhkan
func hitungSkor(waktu [8]int, soal *int, skor *int) {
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		if waktu[i] <= 300 { // Waktu lebih dari 300 dianggap gagal
			*soal++
			*skor += waktu[i]
		}
	}
}

func main() {
	var nama, pemenang string
	var waktu [8]int
	var soal, skor, soalMax, skorMin int

	soalMax = 0
	skorMin = 301 * 8 // Nilai skor maksimal yang mungkin (301 menit per soal)

	for {
		// Baca nama peserta
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		// Baca waktu penyelesaian untuk 8 soal
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		// Hitung jumlah soal yang diselesaikan dan total skor
		hitungSkor(waktu, &soal, &skor)

		// Tentukan pemenang berdasarkan soal yang diselesaikan dan skor
		if soal > soalMax || (soal == soalMax && skor < skorMin) {
			pemenang = nama
			soalMax = soal
			skorMin = skor
		}
	}

	// Cetak pemenang, jumlah soal yang diselesaikan, dan total skor
	fmt.Println(pemenang, soalMax, skorMin)
}
