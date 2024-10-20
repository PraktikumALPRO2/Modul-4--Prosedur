package main

import "fmt"


func MencetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2 
		} else {
			n = 3*n + 1 
		}
	}
	fmt.Print(n) 
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif (n < 1000000): ")
	fmt.Scan(&n)

	if n > 0 && n < 1000000 {
		MencetakDeret(n) 
	} else {
		fmt.Println("Input tidak valid. Pastikan n adalah bilangan positif dan kurang dari 1000000.")
	}
}
