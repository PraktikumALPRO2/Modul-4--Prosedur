package main

import (
	"fmt"
)

func cetakDeret_145(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1) 
}

func main() {
	var n int
	

	for {
		fmt.Print("Masukkan satu bilangan integer positif yang lebih kecil dari 1000000: ")
		fmt.Scan(&n)
		
		if n > 0 && n < 1000000 {
			break 
		} else {
			fmt.Println("Bilangan harus positif dan lebih kecil dari 1000000")
		}
	}
	
	
	cetakDeret_145(n)
}
