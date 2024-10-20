package main

import "fmt"

//mencetak deret bilangan
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n) //cetak nilai n
		if n%2 == 0 {        //jika n genap
			n = n / 2
		} else { //jika n ganjil
			n = 3*n + 1
		}
	}
	fmt.Println(1) //cetak nilai akhir 1
}
func main() {
	var n int
	fmt.Print("Masukkan nilai awal deret : ")
	fmt.Scan(&n)  //input dari pengguna
	cetakDeret(n) //panggil prosedur cetakDeret
}
