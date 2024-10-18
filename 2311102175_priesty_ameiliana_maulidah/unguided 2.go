package main

import (
	"fmt"
	"math"
)

func hitungSkor(nama string, waktu []int) (int, int) {
	totalSoal := 0
	totalWaktu := 0

	for _, w := range waktu {
		if w <= 301 {
			totalSoal++
			totalWaktu += w
		}
	}

	return totalSoal, totalWaktu
}

func main() {
	var peserta string
	var maxSoal, minWaktu int
	var pemenang string

	maxSoal = -1
	minWaktu = math.MaxInt32

	for {
		_, err := fmt.Scan(&peserta)
		if err != nil {
			break
		}

		waktu := make([]int, 8)
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		totalSoal, totalWaktu := hitungSkor(peserta, waktu)

		if totalSoal > maxSoal || (totalSoal == maxSoal && totalWaktu < minWaktu) {
			maxSoal = totalSoal
			minWaktu = totalWaktu
			pemenang = peserta
		}
	}

	fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
}
