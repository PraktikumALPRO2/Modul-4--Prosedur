package main

import (
	"fmt"
)

// Prosedur untuk menghitung jumlah soal yang diselesaikan dan total waktu (skor)
func hitungSkor(waktu []int, soal *int, skor *int) { // Mendefinisikan fungsi bernama hitungSkor yang menerima 3 parameter
	*soal = 0 // Menginisialsiasi nilai soal melalui pointer
	*skor = 0 // Menginisialisasi nilai skor melalui pointer
	for  _, waktuSoal := range waktu { // Loop untuk setiap elemen dalam waktu. Elemen disimpan dalam variabel waktuSoal
		if waktuSoal <= 300 { // Memeriksa apakah waktusoal diselesaikan dalam waktu kurang dari atau sama dengan 300 menit 
			*soal++
			*skor += waktuSoal
		}
	}
}

// Prosedur untuk mencari pemenang dari daftar peserta
func cariPemenang(namaPeserta []string, waktuTotal [][]int, pemenang *string, hasilSoal *int, hasilSkor *int) { // Mendefinisikan fungsi cariPemenang yang menerima 5 parameter
	*pemenang = "" // inisialisasi variabel  pemenang dengan string 
	*hasilSoal = -1 // Inisialisasi hasilSoal dengan -1 sebagai nilai awal
	*hasilSkor = 9999999 // Insialisasi hasilSkor dengan nilai yang sabgat besar untuk memudahkan pencarian skor yang lebih rendah nantinya

	for i, waktuSoal := range waktuTotal {
		var soal int
		var skor int

		// Memanggil prosedur hitungSkor untuk menghitung soal dan skor peserta
		hitungSkor(waktuSoal, &soal, &skor)

		// Mengecek apakah peserta ini menyelesaikan lebih banyak soal atau memiliki skor lebih kecil
		if soal > *hasilSoal || (soal == *hasilSoal && skor < *hasilSkor) {
			*pemenang = namaPeserta[i]
			*hasilSoal = soal
			*hasilSkor = skor
		}
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n) // Input jumlah peserta

	namaPeserta := make([]string, n)
	waktuTotal := make([][]int, n)

	for i := 0; i < n; i++ {
		// Input nama peserta dan waktu pengerjaan soal
		var nama string
		waktu := make([]int, 8)
		fmt.Printf("Masukkan: ")
		fmt.Scan(&nama)  // Membaca nama peserta
		for j := 0; j < 8; j++ {
			fmt.Scan(&waktu[j]) // Membaca waktu pengerjaan soal
		}
		namaPeserta[i] = nama
		waktuTotal[i] = waktu
	}

	fmt.Println("Selesai")

	var pemenang string
	var hasilSoal int
	var hasilSkor int

	// Memanggil prosedur untuk mencari pemenang
	cariPemenang(namaPeserta, waktuTotal, &pemenang, &hasilSoal, &hasilSkor)

	// Menampilkan hasil
	fmt.Printf("Pemenang adalah %s yang menyelesaikan %d soal dengan waktu %d menit", pemenang, hasilSoal, hasilSkor)
}
