package main

import "fmt"

func HitungSKOR(waktu []int, soal *int, skor *int) {
	for _, w := range waktu {
		if w < 301 {
			*soal++
			*skor += w
		}
	}
}

func main() {
	var peserta int
	fmt.Print("Input jumlah peserta: ")
	fmt.Scan(&peserta)

	pemenang, maxSoal, minSkor := "", -1, 99999

	for i := 0; i < peserta; i++ {
		var nama string
		var waktu [8]int
		fmt.Print("Input nama dan waktu: ")
		fmt.Scan(&nama, &waktu[0], &waktu[1], &waktu[2], &waktu[3], &waktu[4], &waktu[5], &waktu[6], &waktu[7])

		var totalSoal, totalSkor int
		HitungSKOR(waktu[:], &totalSoal, &totalSkor)

		if totalSoal > maxSoal || (totalSoal == maxSoal && totalSkor < minSkor) {
			pemenang, maxSoal, minSkor = nama, totalSoal, totalSkor
		}
	}

	fmt.Printf("Pemenang: %s\nJumlah soal: %d\nTotal waktu: %d\n", pemenang, maxSoal, minSkor)
}
