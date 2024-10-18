package main

import "fmt"

// Prosedur untuk mencetak deret
func cetakDeret(bilangan int) {
    for bilangan != 1 {
        fmt.Printf("%d ", bilangan)
        if bilangan%2 == 0 {
            bilangan /= 2 // Jika n genap, bagi dua
        } else {
            bilangan = 3*bilangan + 1 // Jika n ganjil, kalikan tiga dan tambahkan satu
        }
    }
    fmt.Printf("1\n") 
}

func main() {
    var bilangan int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&bilangan)
    cetakDeret(bilangan)
}
