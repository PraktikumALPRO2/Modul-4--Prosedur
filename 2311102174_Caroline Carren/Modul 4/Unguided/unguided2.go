// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Prosedur untuk menghitung total soal yang dikerjakan dan total waktu
func hitungSkor(waktu []int) (int, int) {
	soal := 0
	totalWaktu := 0
	for _, w := range waktu {
		if w <= 300 { // jika waktu pengerjaan kurang dari 300 menit, soal selesai
			soal++
			totalWaktu += w // hanya tambahkan waktu soal yang selesai
		}
	}
	return soal, totalWaktu
}

func main() {
	var n int
	fmt.Print("Masukkan Jumlah Peserta: ")
	fmt.Scan(&n)

	var pemenang string
	maximalSoal := 0
	minimalWaktu := 1000 // nilai awal
	var peserta []string
	var waktu [][]int

	for i := 0; i < n; i++ {
		var namaPeserta string
		fmt.Printf("\nNama Peserta %d: ", i+1)
		fmt.Scan(&namaPeserta)
		peserta = append(peserta, namaPeserta)

		// Input waktu pengerjaan soal
		fmt.Print("Masukkan Waktu Pengerjaan Soal (8 soal, dalam menit): ")
		var waktuPeserta []int
		for j := 0; j < 8; j++ {
			var w int
			fmt.Printf("Soal %d: ", j+1)
			fmt.Scan(&w)
			waktuPeserta = append(waktuPeserta, w)
		}
		waktu = append(waktu, waktuPeserta)

		// Hitung soal yang selesai dan total waktu yang digunakan
		soal, totalWaktu := hitungSkor(waktuPeserta)

		// pemenang berdasarkan jumlah soal yang selesai dan waktu tersingkat
		if soal > maximalSoal || (soal == maximalSoal && totalWaktu < minimalWaktu) {
			pemenang = namaPeserta
			maximalSoal = soal
			minimalWaktu = totalWaktu
		}
	}

	// Output hasil akhir dalam format tabel
	fmt.Println("\n==============================")
	fmt.Println("Hasil Akhir:")
	fmt.Println("==============================")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Nama\tJumlah Soal\tTotal Waktu")
	fmt.Fprintln(w, "-----------------------------")
	for i, nama := range peserta {
		soal, totalWaktu := hitungSkor(waktu[i])
		fmt.Fprintf(w, "%s\t%d\t%d\n", nama, soal, totalWaktu)
	}
	w.Flush()

	fmt.Println("==============================")
	fmt.Printf("Nama Pemenang: %s\n", pemenang)
	fmt.Printf("Jumlah Soal yang Selesai: %d\n", maximalSoal)
	fmt.Printf("Total Waktu yang Dihabiskan: %d menit\n", minimalWaktu)
	fmt.Println("==============================")
}
