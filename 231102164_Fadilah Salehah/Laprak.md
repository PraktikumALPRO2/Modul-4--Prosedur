<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL IV</strong></h2>
<h2 align="center"><strong> PROSEDUR </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fadilah Salehah / 2311102164<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## I. Dasar Teori
### Pengertian Fungsi
Prosedur adalah blok kode yang tidak mengembalikan nilai tetapi dirancang untuk menjalankan tugas tertentu. Berbeda dengan fungsi yang mengembalikan hasil, prosedur hanya menjalankan sekumpulan instruksi tanpa memberikan nilai kembali ke pemanggil. Prosedur berguna untuk memecah program menjadi bagian yang lebih sederhana dan mudah dipahami, terutama ketika tugas tidak memerlukan pengembalian nilai [1].

Prosedur membantu meningkatkan keterbacaan dan organisasi program karena memungkinkan pembagian logika program menjadi beberapa bagian yang terpisah namun tetap koheren dalam menjalankan tugas tertentu [2].

### Deklarasi Prosedur
Dalam bahasa pemrograman, deklarasi prosedur biasanya sangat mirip dengan fungsi, namun prosedur tidak memiliki tipe pengembalian. Pada beberapa bahasa seperti Pascal atau Go (dimana prosedur disebut sebagai "fungsi tanpa pengembalian"), prosedur dideklarasikan menggunakan kata kunci yang sama dengan fungsi namun tanpa tipe pengembalian.

Contoh deklarasi prosedur dalam bahasa Go:

```go
func cetakNama(nama string) {
    fmt.Println("Nama:", nama)
}

```
#### Penjelasan:
- cetakNama: Nama prosedur yang digunakan untuk pemanggilan.
- nama: Parameter yang diterima oleh prosedur, dalam hal ini berupa string.
- Blok kode: Berisi instruksi yang akan dijalankan oleh prosedur.

### Pemanggilan Prosedur
Sama seperti fungsi, pemanggilan prosedur adalah cara untuk mengeksekusi instruksi yang berada di dalamnya. Prosedur dapat dipanggil kapan saja selama berada dalam lingkup yang benar. Untuk memanggil prosedur, kita cukup menuliskan nama prosedur diikuti oleh kurung buka dan tutup beserta argumen yang diperlukan (jika ada).

Contoh pemanggilan prosedur dari contoh sebelumnya:

```go
func main() {
    cetakNama("Andi") // Memanggil prosedur 'cetakNama' dengan argumen "Andi"
}
```

Dalam contoh di atas, prosedur cetakNama dipanggil dengan memberikan argumen "Andi". Prosedur tersebut tidak mengembalikan nilai, hanya mencetak nama ke layar.

### Prosedur dengan Banyak Parameter
Prosedur juga dapat menerima lebih dari satu parameter untuk melakukan tugas yang lebih kompleks. Berikut contoh prosedur dengan beberapa parameter:

```go
func tampilkanDetail(nama string, umur int) {
    fmt.Printf("Nama: %s, Umur: %d\n", nama, umur)
}

```
#### Pemanggilan prosedur:
```go
func main() {
    tampilkanDetail("Andi", 25) // Output: Nama: Andi, Umur: 25
}
```
Pada contoh di atas, prosedur tampilkanDetail menerima dua parameter: nama dan umur, lalu mencetak informasi detail tersebut.

## II. GUIDED
## 1. Program Perhitungan Permutasi Dua Bilangan Bulat

#### Source Code
```go
package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    if a >= b {
        permutasi(a, b)
    } else {
        permutasi(b, a)
    }
}

// Prosedur faktorial untuk menghitung dan menampilkan faktorial
func faktorial(n int) {
    var hasil int = 1
    var i int
    for i = 1; i <= n; i++ {
        hasil = hasil * i
    }
    fmt.Println("Faktorial dari", n, "adalah:", hasil)
}

// Prosedur permutasi untuk menghitung dan mencetak hasil permutasi
func permutasi(n, r int) {
    faktorialN := 1
    faktorialNR := 1
    // Hitung faktorial n
    for i := 1; i <= n; i++ {
        faktorialN *= i
    }
    // Hitung faktorial (n-r)
    for i := 1; i <= (n - r); i++ {
        faktorialNR *= i
    }
    // Cetak hasil permutasi
    fmt.Println("Permutasi dari", n, "dan", r, "adalah:", faktorialN/faktorialNR)
}
```

#### Screenshoot Output
![image](https://github.com/user-attachments/assets/1832add6-b678-4a6f-a825-926144e08129)

#### Deskripsi Program
Program ini berfungsi untuk menghitung dan mencetak nilai permutasi dari dua bilangan bulat yang diinputkan oleh pengguna.

## 2. Program Penghitungan Luas dan Keliling Persegi

#### Source Code
```go
package main

import (
	"fmt"
)

// Prosedur untuk menghitung dan menampilkan luas persegi
func hitungLuas(sisi float64) {
	luas := sisi * sisi
	fmt.Printf("Luas persegi: %.2f\n", luas)
}

// Prosedur untuk menghitung dan menampilkan keliling persegi
func hitungKeliling(sisi float64) {
	keliling := 4 * sisi
	fmt.Printf("Keliling persegi: %.2f\n", keliling)
}

func main() {
	var sisi float64

	fmt.Print("Sisi Persegi: ")
	fmt.Scan(&sisi)

	// Panggil prosedur untuk menghitung luas dan keliling
	hitungLuas(sisi)
	hitungKeliling(sisi)
}

```

#### Screenshoot Output
![image](https://github.com/user-attachments/assets/001f8bc4-0fcf-4cff-b650-6fa0e6342254)

#### Deskripsi Program
Program Go ini berfungsi untuk menghitung dan menampilkan luas dan keliling dari sebuah persegi berdasarkan panjang sisinya yang dimasukkan oleh pengguna. 

## III. UNGUIDED
## 1. Program Penghitung Permutasi dan Kombinasi

#### Source Code
```go
package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial dari sebuah bilangan
func hitungFaktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi P(n, r)
func hitungPermutasi(n, r int) int {
	faktorialN := hitungFaktorial(n)
	faktorialNR := hitungFaktorial(n - r)
	return faktorialN / faktorialNR
}

// Fungsi untuk menghitung kombinasi C(n, r)
func hitungKombinasi(n, r int) int {
	faktorialN := hitungFaktorial(n)
	faktorialR := hitungFaktorial(r)
	faktorialNR := hitungFaktorial(n - r)
	return faktorialN / (faktorialR * faktorialNR)
}

func main() {
	// Input
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan a, b, c, d: ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// Tampilkan hasil permutasi dan kombinasi untuk a dan c
	fmt.Printf("Permutasi P(%d, %d) adalah: %d\n", a, c, hitungPermutasi(a, c))
	fmt.Printf("Kombinasi C(%d, %d) adalah: %d\n", a, c, hitungKombinasi(a, c))

	// Tampilkan hasil permutasi dan kombinasi untuk b dan d
	fmt.Printf("Permutasi P(%d, %d) adalah: %d\n", b, d, hitungPermutasi(b, d))
	fmt.Printf("Kombinasi C(%d, %d) adalah: %d\n", b, d, hitungKombinasi(b, d))
}

```

#### Screenshoot Output
![image](https://github.com/user-attachments/assets/67630283-08c4-4d86-aebf-5786c36d0ab3)

#### Deskripsi Program
Program ini berfungsi untuk menghitung dan menampilkan permutasi dan kombinasi dari dua pasangan bilangan yang dimasukkan oleh pengguna. Program ini mengimplementasikan tiga fungsi utama: menghitung faktorial, menghitung permutasi, dan menghitung kombinasi.

## 2. Program Penentuan Pemenang Kompetisi Pemrograman Berdasarkan Skor dan Waktu Penyelesaian

#### Source Code
```go

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Prosedur untuk menghitung skor peserta
func hitungSkor(soal [8]int, totalSoal *int, totalWaktu *int) {
	*totalSoal = 0
	*totalWaktu = 0

	// Hitung jumlah soal yang berhasil diselesaikan dan total waktu yang dibutuhkan
	for _, waktu := range soal {
		if waktu < 301 {
			*totalSoal++
			*totalWaktu += waktu
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pemenangNama string
	var maxSoalDiselesaikan int
	var minWaktu int = math.MaxInt32

	for {
		fmt.Printf("Masukkan nama peserta diikuti dengan waktu untuk setiap soal atau ketik 'Selesai' untuk berhenti:")

		// Membaca input peserta
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		// Memisahkan input menjadi nama peserta dan waktu pengerjaan soal
		parts := strings.Fields(input)
		if len(parts) < 9 {
			fmt.Println("Input tidak valid. Pastikan memasukkan nama peserta dan 8 waktu pengerjaan soal.")
			continue
		}

		pesertaNama := parts[0]

		// Simpan waktu pengerjaan soal-soal ke array
		var soal [8]int
		for i := 1; i <= 8; i++ {
			soal[i-1], _ = strconv.Atoi(parts[i])
		}

		// Variabel untuk menyimpan hasil hitung skor
		var totalSoal, totalWaktu int

		// Hitung skor peserta menggunakan prosedur
		hitungSkor(soal, &totalSoal, &totalWaktu)

		// Tentukan pemenang berdasarkan jumlah soal yang diselesaikan dan waktu total
		if totalSoal > maxSoalDiselesaikan || (totalSoal == maxSoalDiselesaikan && totalWaktu < minWaktu) {
			pemenangNama = pesertaNama
			maxSoalDiselesaikan = totalSoal
			minWaktu = totalWaktu
		}
	}

	// Tampilkan hasil
	if pemenangNama != "" {
		fmt.Printf("Pemenang: %s\n", pemenangNama)
		fmt.Printf("Jumlah soal yang diselesaikan: %d\n", maxSoalDiselesaikan)
		fmt.Printf("Total waktu: %d menit\n", minWaktu)
	} else {
		fmt.Println("Tidak ada data peserta.")
	}
}

```

#### Screenshoot Output
![image](https://github.com/user-attachments/assets/7f0d0c37-8b45-4476-bcfb-8653f767748a)

#### Deskripsi Program
Program ini dirancang untuk menghitung dan menentukan pemenang dari suatu kompetisi berdasarkan waktu penyelesaian soal oleh peserta. Program ini mengumpulkan data waktu penyelesaian soal dari peserta dan menentukan siapa yang menyelesaikan jumlah soal terbanyak dengan waktu total tercepat.

## 3. Program untuk Mencetak Deret Bilangan Berdasarkan Algoritma 3n+1

#### Source Code
```go
package main

import (
	"fmt"
)

// Fungsi untuk mencetak deret berdasarkan nilai n
func cetakDeret(n int) {
	// Teruskan pencetakan deret selama n tidak sama dengan 1
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n /= 2 // Jika n adalah genap
		} else {
			n = 3*n + 1 // Jika n adalah ganjil
		}
	}
	// Cetak 1 sebagai elemen terakhir dari deret
	fmt.Print(1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai awal: ")
	fmt.Scan(&n)

	// Memanggil fungsi untuk mencetak deret
	cetakDeret(n)
}


```

#### Screenshoot Output
![image](https://github.com/user-attachments/assets/7255bbeb-70f5-4edf-8ac2-b279328eae88)

#### Deskripsi Program
Program ini berfungsi untuk mencetak deret yang dihasilkan dari sebuah bilangan bulat 
𝑛
n menggunakan algoritma yang dikenal sebagai "konjectur Collatz" atau "masalah 3n + 1". Deret ini dihasilkan dengan aturan berikut:

Jika 
𝑛
n adalah genap, maka 
𝑛
n dibagi dua.
Jika 
𝑛
n adalah ganjil, maka 
𝑛
n dikalikan tiga dan ditambahkan satu.
Proses ini diulang sampai 
𝑛
n mencapai 1.

### Kesimpulan
Penggunaan prosedur dalam pemrograman adalah alat yang sangat berguna untuk menyusun kode secara terstruktur dan efisien. Dengan memahami cara mendeklarasikan dan menggunakan prosedur, programmer dapat meningkatkan keterbacaan dan keorganisasian kode mereka, serta memudahkan proses debugging dan pengembangan di masa mendatang. Prosedur memberikan cara untuk membagi program menjadi bagian yang lebih kecil dan lebih mudah dikelola, menjadikannya komponen penting dalam pemrograman modern.

## Referensi 
[1] Go.dev. (n.d.). Introduction to Go functions. Go Documentation.

[2] Zakhour, M., et al. (2010). The Java programming language. Addison-Wesley.
