package main

import "fmt"

func hitungSkor(waktu []int) (int, int) {
	totalSoal := 0
	totalSkor := 0

	for _, w := range waktu {
		if w < 301 {
			totalSoal++
			totalSkor += w
		}
	}
	return totalSoal, totalSkor
}

func main() {
	var peserta int
	fmt.Print("Input Jumlah Peserta: ")
	fmt.Scan(&peserta)

	pemenang := ""
	maxSoal := -1
	minSkor := 99999

	for i := 0; i < peserta; i++ {
		var nama string
		var waktu [8]int
		fmt.Print("Input Nama dan Waktu: ")
		fmt.Scan(&nama, &waktu[0], &waktu[1], &waktu[2], &waktu[3], &waktu[4], &waktu[5], &waktu[6], &waktu[7])

		totalSoal, totalSkor := hitungSkor(waktu[:])

		if totalSoal > maxSoal || (totalSoal == maxSoal && totalSkor < minSkor) {
			pemenang, maxSoal, minSkor = nama, totalSoal, totalSkor
		}
	}

	fmt.Printf("Pemenang: %s\nJumlah soal: %d\nTotal waktu: %d\n", pemenang, maxSoal, minSkor)
}