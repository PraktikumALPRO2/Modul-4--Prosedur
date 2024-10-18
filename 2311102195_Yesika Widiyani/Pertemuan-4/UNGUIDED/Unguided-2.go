package main

import "fmt"

// Prosedur untuk menghitung skor peserta
func hitungSkor(waktu []int, totalSoal *int, totalSkor *int) {
	*totalSoal = 0
	*totalSkor = 0
	for _, t := range waktu {
		if t <= 300 { // Jika soal diselesaikan dalam waktu <= 300 menit
			*totalSoal++    // Soal yang berhasil diselesaikan
			*totalSkor += t // Tambahkan waktu pengerjaan ke total skor
		}
	}
}

// Fungsi utama program untuk mencari pemenang
func main() {
	var namaPemenang string
	var soalPemenang, skorPemenang int
	var selesai bool = false

	for !selesai {
		var nama string
		var waktu [8]int
		fmt.Print("Masukkan nama peserta dan waktu pengerjaan (atau ketik 'Selesai' untuk mengakhiri): ")
		fmt.Scan(&nama)

		if nama == "Selesai" { // Pengguna memasukkan kata "Selesai" untuk menghentikan input
			selesai = true
			break
		}

		// Baca waktu pengerjaan peserta
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var totalSoal, totalSkor int
		hitungSkor(waktu[:], &totalSoal, &totalSkor)

		// Tentukan pemenang sementara
		if totalSoal > soalPemenang || (totalSoal == soalPemenang && totalSkor < skorPemenang) {
			namaPemenang = nama
			soalPemenang = totalSoal
			skorPemenang = totalSkor
		}
	}

	// Menampilkan hasil pemenang
	fmt.Printf("Pemenangnya adalah %s dengan %d soal yang diselesaikan dan total waktu %d menit.\n", namaPemenang, soalPemenang, skorPemenang)
}
