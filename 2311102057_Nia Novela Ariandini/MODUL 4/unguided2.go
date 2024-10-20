package main

import (
	"fmt"
	"strings"
)

const maxTime = 301 //watu max kalo soal tidak selesai

// untuk menghitung skor (jumlah soal yang diselesaikan dan total waktu)
func hitungSkor(waktu [8]int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for _, waktuSoal := range waktu {
		if waktuSoal < maxTime { //kalo soal diselesaikan
			*soal++
			*skor += waktuSoal
		}
	}
}

func main() {
	var pemenang string
	var soalPemenang, skorPemenang int
	soalPemenang = 0
	skorPemenang = maxTime * 8 //nilai awal skor pemenang maksimal
	for {
		var nama string
		var waktu [8]int
		//membaca input nama peserta
		fmt.Print("Masukkan nama peserta (atau 'Selesai' untuk mengakhiri) : ")
		fmt.Scan(&nama)
		nama = strings.TrimSpace(nama)
		//kalo input adalah 'Selesai', hentikan loop
		if nama == "Selesai" {
			break
		}
		//membaca waktu penyelesaian 8 soal
		fmt.Print("Masukkan waktu penyelesaian 8 soal (dalam menit) : ")
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}
		//hitung soal yang diselesaikan dan total waktu
		var soal, skor int
		hitungSkor(waktu, &soal, &skor)
		//tentukan pemenang berdasarkan jumlah soal dan total waktu
		if soal > soalPemenang || (soal == soalPemenang && skor < skorPemenang) {
			pemenang = nama
			soalPemenang = soal
			skorPemenang = skor
		}
	}
	//cetak pemenang
	fmt.Printf("Pemenang : %s, Soal yang diselesaikan : %d, Total waktu : %d menit\n", pemenang, soalPemenang, skorPemenang)
}
