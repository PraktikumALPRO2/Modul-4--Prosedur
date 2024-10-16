# <h2 align="center">LAPORAN PRAKTIKUM</h2>
# <h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>
# <h2 align="center">MODUL 4</h2>
# <h2 align="center">PROSEDUR</h2>
<p align="center">
    <img src="https://github.com/user-attachments/assets/3ccfed0b-72d1-4349-ac08-c4dce379c827" alt="Gambar">
</p>
 <h3  align="center" >Disusun Oleh : </h3>
<p align="center">Amanda Windhu Gustyas - 2311102121</p>
<p align="center">IF-11-05</p>
 <h3 <p align="center" >Dosen Pengampu : </h3> </p>
 <p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>

# <h3 align="center"> PROGRAM STUDI S1 TEKNIK INFORMATIKA </h3>
# <h3 align="center"> FAKULTAS INFORMATIKA </h3>
# <h3 align="center"> TELKOM UNIVERSITY PURWOKERTO </h3>
# <h3 align="center"> 2024 </h3>

## I. DASAR TEORI



## II. GUIDED

### 1. Contoh Program dengan Function

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b) // Membaca input dari pengguna untuk dua nilai integer a dan b
	if a >= b { 
		// Jika a lebih besar atau sama dengan b, panggil prosedur permutasi dengan parameter (a, b)
		permutasi(a, b)
	} else { 
		// Jika b lebih besar dari a, panggil prosedur permutasi dengan parameter (b, a)
		permutasi(b, a)
	}
}

func faktorial(n int) int {
	var hasil int = 1
	// Loop untuk menghitung faktorial dari n
	for i := 1; i <= n; i++ {
		hasil *= i // Mengalikan hasil dengan nilai i pada setiap iterasi
	}
	return hasil // Mengembalikan hasil faktorial
}

func permutasi(n, r int) {
	// Menghitung permutasi nPr dan langsung mencetak hasilnya
	hasil := faktorial(n) / faktorial(n-r)
	fmt.Println(hasil) // Mencetak hasil permutasi
}
```
## Output: ![image](https://github.com/user-attachments/assets/6ed035eb-545b-4835-a4f0-ec0d64af7862)

Kode di atas menghitung dan mencetak hasil permutasi dari dua bilangan integer yang dimasukkan oleh pengguna. Program dimulai dengan membaca dua nilai, `a` dan `b`. Jika `a` lebih besar atau sama dengan `b`, program akan memanggil prosedur `permutasi` dengan parameter `(a, b)`. Sebaliknya, jika `b` lebih besar dari `a`, prosedur permutasi dipanggil dengan parameter `(b, a)`.<br/>
Prosedur `faktorial` bertugas menghitung faktorial dari sebuah bilangan dengan menggunakan loop untuk mengalikan semua angka dari 1 hingga `n`. Prosedur `permutasi` kemudian menghitung permutasi 𝑛𝑃𝑟, dengan menggunakan hasil faktorial dari `n` dan `n-r`, kemudian mencetak hasilnya.<br/>
Secara keseluruhan, program ini mengimplementasikan konsep permutasi dengan cara modular, memisahkan perhitungan faktorial dan permutasi dalam dua prosedur terpisah, serta mencetak hasil perhitungan langsung ke layar.

### 2. Menghitung Luas dan Keliling Persegi

```go
package main

import "fmt"

// Prosedur untuk menghitung dan mencetak luas persegi
func hitungLuas(sisi float64) {
    luas := sisi * sisi
    fmt.Printf("Luas persegi: %.2f\n", luas) // Mencetak hasil luas
}

// Prosedur untuk menghitung dan mencetak keliling persegi
func hitungKeliling(sisi float64) {
    keliling := 4 * sisi
    fmt.Printf("Keliling persegi: %.2f\n", keliling) // Mencetak hasil keliling
}

func main() {
    var sisi float64

    fmt.Print("Masukkan panjang sisi persegi: ")
    fmt.Scan(&sisi)

    hitungLuas(sisi)      // Memanggil prosedur hitungLuas
    hitungKeliling(sisi)  // Memanggil prosedur hitungKeliling
}
```
## Output: ![image](https://github.com/user-attachments/assets/d33e1c01-7b2e-48ab-926c-58807500733d)

Kode di atas

## III. UNGUIDED

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)<br/> Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. <br/> Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. <br/> Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!<br/> ![image](https://github.com/user-attachments/assets/5b90c7e3-9f76-45eb-bf14-8f1bca637918) <br/> Selesaikan program tersebut dengan memanfaatkan procedure yang diberikan berikut ini!<br/>![image](https://github.com/user-attachments/assets/d7a28bc4-25bd-4c1d-9091-e058e26a1407)<br/>

```go
package main

import (
	"fmt"
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
	// Input empat bilangan: a, b, c, d
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan: ")
	fmt.Scan(&a, &b, &c, &d)

	// Variabel untuk menyimpan hasil
	var p1, c1, p2, c2 int

	// Baris pertama: permutasi dan kombinasi a terhadap c
	permutation(a, c, &p1)
	combination(a, c, &c1)
	fmt.Printf("%d %d\n", p1, c1)

	// Baris kedua: permutasi dan kombinasi b terhadap d
	permutation(b, d, &p2)
	combination(b, d, &c2)
	fmt.Printf("%d %d\n", p2, c2)
}
```
## Output: ![image](https://github.com/user-attachments/assets/39e62db0-e5d2-4312-accd-9d73df59666d)

Kode di atas

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. <br/> Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.<br/> ![image](https://github.com/user-attachments/assets/b8f89dce-7575-4a33-ad7a-61ece8188938)<br/> Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban, maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). <br/> Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.<br/>![image](https://github.com/user-attachments/assets/3117bbb6-fc8f-45e5-9027-890dea0cb6dc)<br/>Keterangan:<br/> Astuti menyelesaikan 6 soal dalam waktu 287 menit, sedangkan Bertha 7 soal dalam waktu 294 menit. Karena Bertha menyelesaikan lebih banyak, maka Bertha menang. Jika keduanya menyelesaikan sama banyak, maka pemenang adalah yang menyelesaikan dengan total waktu paling kecil.<br/>

```go
package main

import (
	"fmt"
	"strings"
)

const maxTime = 301

// Prosedur untuk menghitung skor dan total waktu
func hitungSkor(soal [8]int, totalSoal *int, totalWaktu *int) {
	*totalSoal = 0
	*totalWaktu = 0
	for _, waktu := range soal {
		if waktu < maxTime {
			*totalSoal++
			*totalWaktu += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal [8]int
	var totalSoal, totalWaktu int
	maxSoal := -1
	minWaktu := maxTime * 8 // Inisialisasi dengan nilai maksimal

	for {
		// Input nama peserta dan soal-soal yang dikerjakan
		fmt.Print("Masukkan nama peserta: ")
		fmt.Scan(&nama)
		if strings.ToLower(nama) == "selesai" {
			break
		}

		fmt.Println("Masukkan waktu pengerjaan: ")
		for i := 0; i < 8; i++ {
			fmt.Scan(&soal[i])
		}

		// Hitung total soal yang diselesaikan dan total waktu yang dibutuhkan
		hitungSkor(soal, &totalSoal, &totalWaktu)

		// Tentukan pemenang sementara
		if totalSoal > maxSoal || (totalSoal == maxSoal && totalWaktu < minWaktu) {
			maxSoal = totalSoal
			minWaktu = totalWaktu
			pemenang = nama
		}
	}

	// Output pemenang
	fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
}
```
## Output: ![image](https://github.com/user-attachments/assets/58671054-6ef9-40e6-8188-0ab45d7a86d8)

Kode di atas

### 3. Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah ½n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh, jika dimulai dengan n = 22, maka deret bilangan yang diperoleh adalah:<br/> ![image](https://github.com/user-attachments/assets/eb8656b7-5bd8-4bc3-8e9e-ac2e315c7cff)<br/>Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1.<br/>Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.<br/> 
### ![image](https://github.com/user-attachments/assets/3abb1b23-463e-4877-918c-5e045d2a246e) <br/>Masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000.<br/> Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang sama dan dipisahkan oleh sebuah spasi.<br/>

```go
package main

import "fmt"

// Prosedur untuk mencetak deret
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	// Cetak angka terakhir (1)
	fmt.Printf("%d\n", n)
}

func main() {
	var n int

	// Input bilangan awal
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	// Cetak deret
	cetakDeret(n)
}
```
## Output: ![image](https://github.com/user-attachments/assets/2ccb1a4c-8909-4acc-8947-c58a108a890a)

Kode di atas

### KESIMPULAN
### REFERENSI
