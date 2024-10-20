package main

import (
	"fmt"
)

func skor(soal [8]int, skor *int, waktu *int) {
	*skor = 0
	*waktu = 0
	for i := 0; i < 8; i++ {
		if soal[i] <= 301 {
			*skor++
			*waktu += soal[i]
		}
	}
}

func main() {
	var peserta1, peserta2 string
	var soalselesai1, soalselesai2 [8]int
	var skor1, skor2, total1, total2 int

	fmt.Print("PESERTA: ")
	fmt.Scan(&peserta1)
	fmt.Println("WAKTU PENGERJAAN:")
	for i := 0; i < 8; i++ {
		fmt.Scan(&soalselesai1[i])
	}

	fmt.Print("PESERTA: ")
	fmt.Scan(&peserta2)
	fmt.Println("WAKTU PENGERJAAN:")
	for i := 0; i < 8; i++ {
		fmt.Scan(&soalselesai2[i])
	}

	skor(soalselesai1, &skor1, &total1)
	skor(soalselesai2, &skor2, &total2)

	if skor1 > skor2 || (skor1 == skor2 && total1 < total2) {
		fmt.Printf("peserta %s menang dengan skor %d dalam total waktu %d menit\n", peserta1, skor1, total1)
	} else if skor1 == skor2 {
        fmt.Printf("kedua peserta seri\n")
    } else {
		fmt.Printf("peserta %s menang dengan skor %d dalam total waktu %d menit\n", peserta2, skor2, total2)
	}
}
