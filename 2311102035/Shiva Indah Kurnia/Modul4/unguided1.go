package main 
 import (     "fmt" 
) 
// Prosedur untuk menghitung faktorial 
func factorial(n int, result *int) { 
    *result = 1     
    if n == 0 {         
        *result = 1         
        return 
    } 
    for i := 1; i <= n; i++ { 
        *result *= i 
    } 
} 
// Prosedur untuk menghitung permutasi 
func permutation(n, r int, result *int) {     
    var fn, fnr int     
    factorial(n, &fn)     
    factorial(n-r, &fnr) 
    *result = fn / fnr 
} 
// Prosedur untuk menghitung kombinasi 
func combination(n, r int, result *int) {     
    var fn, fr, fnr int     
    factorial(n, &fn)     
    factorial(r, &fr)     
    factorial(n-r, &fnr)     
    *result = fn / (fr * fnr) 
} 
func main() {     
    var a, b, c, d int 
    fmt.Print("Masukkan 4 bilangan: ")     
    fmt.Scan(&a, &b, &c, &d) 
	var p1, c1, p2, c2 int 
    // Baris pertama permutasi dan kombinasi a terhadap c 
    permutation(a, c, &p1)     
    combination(a, c, &c1)     
    fmt.Printf("%d %d\n", p1, c1) 
    // Baris kedua permutasi dan kombinasi b terhadap d 
    permutation(b, d, &p2)     
    combination(b, d, &c2)     
    fmt.Printf("%d %d\n", p2, c2) } 
