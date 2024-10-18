<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL IV</strong></h2>
<h2 align="center"><strong> PROSEDUR </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Lutfiana Isnaeni L /2311102165<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh S.Kom., M.Kom. 
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## Daftar Isi

1. [Dasar Teori](#dasar-teori)
2. [Guided](#guided)
3. [Unguided](#unguided)

## Dasar Teori




## Guided
### 1. 

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		cetakPermutasi(a, b)
	} else {
		cetakPermutasi(b, a)
	}
}

func faktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Prosedur untuk mencetak permutasi
func cetakPermutasi(n, r int) {
	perm := permutasi(n, r)
	fmt.Println(perm)
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
```
#### Output
![image](https://github.com/user-attachments/assets/18b9b03d-dac8-47b0-af3d-576c80f6fb87)
### Deskripsi Program
Program ini adalah aplikasi sederhana yang digunakan untuk menghitung permutasi dari dua angka yang diberikan oleh pengguna. Permutasi menggambarkan jumlah cara untuk memilih dan mengatur `r` objek dari `n` objek yang tersedia. Program ini menggunakan beberapa fungsi untuk menghitung faktorial dan permutasi, serta menampilkan hasilnya di layar.
### Algoritma Program
Berikut adalah langkah-langkah algoritma untuk program ini:
1.	Mulai Program
   
2.	Inisialisasi Variabel:
   
    •	Deklarasikan dua variabel bertipe integer, yaitu `a` dan` b`.
3.	Input Angka:
   
    •	Minta pengguna untuk memasukkan dua angka (a dan b).
  
4.	Periksa dan Hitung Permutasi:
   
    •	Jika a lebih besar atau sama dengan `b`, panggil fungsi `cetakPermutasi(a, b)`.
  
    •	Jika tidak, panggil fungsi `cetakPermutasi(b, a)`.
  
5.	Fungsi `cetakPermutasi(n, r)`:

    •	Hitung permutasi dengan memanggil fungsi permutasi`(n, r)`.

    •	Tampilkan hasil permutasi di layar.

6.	Fungsi `permutasi(n, r)`:
    
    •	Hitung permutasi menggunakan rumus P(n,r)=n!(n−r) dengan memanggil fungsi `faktorial(n)` dan `faktorial(n-r)`.

7.	Fungsi `faktorial(n)`:
    
    •	Hitung faktorial dari `n` dengan cara mengalikan semua angka dari 1 hingga `n`.
  
8.	Akhiri Program

### Cara Kerja Program
1.	Input dari Pengguna: Pengguna diminta untuk memasukkan dua angka yang akan digunakan sebagai parameter untuk menghitung permutasi.
   
2.	Proses Input: Program membaca kedua angka tersebut dan menyimpannya dalam variabel `a` dan `b`.
   
3.	Pemanggilan Fungsi `cetakPermutasi`: Program membandingkan nilai `a` dan `b`. Fungsi `cetakPermutasi` dipanggil dengan argumen yang tepat (angka yang lebih besar sebagai parameter pertama).
   
4.	Penghitungan Permutasi:
   
    •	Fungsi `cetakPermutasi` memanggil fungsi permutasi, yang selanjutnya menggunakan fungsi faktorial untuk menghitung faktorial dari `n, dan `n-r`.
  	
    •	Permutasi dihitung dengan membagi hasil faktorial `n` dengan faktorial `n-r`.

6.	Menampilkan Hasil: Hasil perhitungan permutasi dicetak di konsol.
    
7.	Program Selesai: Setelah hasil ditampilkan, program selesai dan bisa ditutup.

 ### 2. 
 ```go
package main

import (
	"fmt"
)

// Prosedur untuk menghitung luas dan keliling persegi
func hitungPersegi(panjangSisi float64) (float64, float64) {
	// Menghitung luas dan keliling persegi
	LuasPersegi := panjangSisi * panjangSisi
	KelilingPersegi := 4 * panjangSisi

	return LuasPersegi, KelilingPersegi
}

func main() {
	// Meminta input panjang sisi persegi
	var panjangSisi float64
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&panjangSisi)

	// Memanggil prosedur untuk menghitung luas dan keliling
	Luas, Keliling := hitungPersegi(panjangSisi)

	// Menampilkan hasil
	fmt.Printf("Luas persegi: %g\n", Luas)
	fmt.Printf("Keliling persegi: %g\n", Keliling)
}
```
#### Output
![image](https://github.com/user-attachments/assets/ed5ddda6-859e-44a7-9722-bf888b13ccfb)

### Deskripsi Program
Program ini digunakan untuk menghitung luas dan keliling persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna. Program ini menggunakan fungsi untuk melakukan perhitungan dan kemudian menampilkan hasilnya di layar. Fungsi ini menerima panjang sisi sebagai argumen dan mengembalikan dua nilai: luas dan keliling persegi.
### Algoritma Program
1.	Mulai Program
   
2.	Inisialisasi Variabel:
   
    •	Buat variabel panjangSisi dengan tipe `float64`.

3.	Input Panjang Sisi:
   
    •	Minta pengguna untuk memasukkan panjang sisi dari persegi.

4.	Fungsi `hitungPersegi`:
   
    •	Terima panjangSisi sebagai parameter.

    •	Hitung luas dengan rumus Luas=panjangSisi × panjangSisi

    •	Hitung keliling dengan rumus Keliling=4 × panjangSisiKeliling

    •	Kembalikan nilai luas dan keliling.

5.	Panggil Fungsi `hitungPersegi`:
    
    •	Panggil fungsi dengan argumen `panjangSisi` dan simpan hasilnya ke dalam variabel `Luas` dan `Keliling`.

6.	Tampilkan Hasil:
    
    •	Tampilkan luas dan keliling yang telah dihitung di layar.

7.	Akhiri Program

### Cara Kerja Program
1.	Input dari Pengguna: Pengguna diminta untuk memasukkan panjang sisi dari persegi melalui input.
  
2.	Proses Input: Program membaca nilai yang dimasukkan dan menyimpannya dalam variabel `panjangSisi`.
   
3.	Panggilan Fungsi hitungPersegi: Fungsi `hitungPersegi` dipanggil dengan parameter `panjangSisi`.
	
4.	Perhitungan Luas dan Keliling:
   
    •	Di dalam fungsi `hitungPersegi`, luas dihitung dengan mengalikan panjang sisi dengan dirinya sendiri.

    •	Keliling dihitung dengan mengalikan panjang sisi dengan 4.\

    •	Fungsi kemudian mengembalikan nilai luas dan keliling ke fungsi utama.

5.	Menampilkan Hasil: Program menampilkan hasil luas dan keliling yang telah dihitung ke konsol menggunakan `fmt.Printf`.



