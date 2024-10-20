package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cetakDeret(n)
}

// Prosedur cetakDeret menerima n sebagai input (parameter in)
func cetakDeret(n int) {
	for n > 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2 // Jika genap
		} else {
			n = 3*n + 1 // Jika ganjil
		}
	}
	fmt.Print(n) // Mencetak 1 di akhir deret
}
