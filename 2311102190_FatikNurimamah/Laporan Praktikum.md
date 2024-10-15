<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL 4</strong></h2>
<h2 align="center"><strong> PROSEDUR </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fatik Nurimamah / 2311102190<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh,S.Kom.,M.Kom.
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
4. [Kesimpulan](#kesimpulan)
5. [Daftar Pustaka](#daftar-pustaka)


## Dasar Teori





## Guided 

### 1. Buatlah program 

### Source Code :
```go
package main

import "fmt"

func main() {
	// Mendeklarasikan variabel untuk menampung dua input bilangan bulat
	var a, b int
	fmt.Scan(&a, &b)

	// Mendeklarasikan variabel untuk menampung hasil permutasi
	var hasilPermutasi int
	// Memanggil prosedur hitungPermutasi berdasarkan nilai a dan b
	if a >= b {
		hitungPermutasi(a, b, &hasilPermutasi)
	} else {
		hitungPermutasi(b, a, &hasilPermutasi)
	}

	// Menampilkan hasil permutasi
	fmt.Println(hasilPermutasi)
}

// Prosedur untuk menghitung faktorial dari bilangan n
func hitungFaktorial(n int, hasil *int) {
	*hasil = 1
	// Melakukan perulangan untuk menghitung faktorial
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur untuk menghitung permutasi dari n dan r
func hitungPermutasi(n, r int, hasil *int) {
	var fn, fnr int
	hitungFaktorial(n, &fn)
	hitungFaktorial(n-r, &fnr)
	// Menghitung permutasi dengan rumus n! / (n-r)! dan menyimpannya di hasil
	*hasil = fn / fnr
}

```

### Output:
![Screenshot 2024-10-15 212936](https://github.com/user-attachments/assets/9a60908c-fb4f-4a7c-a326-daae4c20e744)

### Full code Screenshot:
![Screenshot 2024-10-15 213026](https://github.com/user-attachments/assets/d637d8d6-5bc4-4b01-8a20-f266e2c1b614)

### Deskripsi Program : 
Program ini digunakan untuk menghitung permutasi dari dua bilangan bulat yang dimasukkan oleh pengguna. Permutasi dihitung menggunakan rumus `P(n, r) = n!/(n-r)!`, di mana `n` adalah bilangan yang lebih besar atau sama dengan `r`. Program akan membaca dua input bilangan, menghitung permutasi, dan menampilkan hasilnya.

### Algoritma Program
1. Input Bilangan: Program meminta dua bilangan bulat `a` dan `b` dari pengguna.
2. Pengecekan Nilai: Program menentukan mana yang lebih besar antara `a` dan `b`, kemudian menjadikan bilangan yang lebih besar sebagai `n` dan yang lebih kecil sebagai `r`.
3. Hitung Faktorial: Program menghitung faktorial `n` dan `n-r` dengan menggunakan prosedur `hitungFaktorial`.
4. Hitung Permutasi: Program kemudian menghitung permutasi dengan rumus `P(n, r) = n!/(n-r)!` di dalam prosedur `hitungPermutasi`.
5. Output Hasil: Program menampilkan hasil permutasi yang telah dihitung ke layar.

### Cara Kerja Program:
1. Program dimulai dengan menerima dua input bilangan dari pengguna dan menyimpannya dalam variabel `a` dan `b`.
2. Program membandingkan nilai `a` dan `b` untuk memastikan bilangan yang lebih besar diperlakukan sebagai `n` dan yang lebih kecil sebagai `r`, kemudian memanggil prosedur `hitungPermutasi`.
3. Dalam prosedur `hitungPermutasi`, dilakukan perhitungan faktorial untuk `n` dan `n-r` menggunakan prosedur `hitungFaktorial`.
4. Setelah perhitungan faktorial selesai, permutasi dihitung dengan membagi faktorial `n` dengan faktorial `n-r`.
5. Hasil akhirnya ditampilkan di layar sebagai hasil permutasi dari dua bilangan yang dimasukkan.


### 2. Buatlah program 

### Source Code :
```go
package main

import (
	"fmt"
)

func main() {
	// Meminta input panjang sisi persegi
	var panjangSisi float64
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&panjangSisi)

	var luasPersegi, kelilingPersegi float64
	// Memanggil prosedur untuk menghitung luas dan keliling persegi
	hitungLuasPersegi(panjangSisi, &luasPersegi)
	hitungKelilingPersegi(panjangSisi, &kelilingPersegi)

	// Menampilkan hasil
	fmt.Printf("Luas persegi: %g\n", luasPersegi)
	fmt.Printf("Keliling persegi: %g\n", kelilingPersegi)
}

// Prosedur untuk menghitung luas persegi
func hitungLuasPersegi(sisi float64, hasil *float64) {
	*hasil = sisi * sisi
}

// Prosedur untuk menghitung keliling persegi
func hitungKelilingPersegi(sisi float64, hasil *float64) {
	*hasil = 4 * sisi
}
```

### Output:
![Screenshot 2024-10-15 214702](https://github.com/user-attachments/assets/10fa92f0-747d-446a-b880-f92c97706b92)

### Full code Screenshot:
![Screenshot 2024-10-15 214723](https://github.com/user-attachments/assets/89e6486b-a56b-4def-8f18-31754c0685b2)

### Deskripsi Program : 
Program ini digunakan untuk menghitung luas dan keliling sebuah persegi berdasarkan panjang sisi yang diberikan oleh pengguna. Program menggunakan dua prosedur terpisah untuk menghitung luas dan keliling persegi, kemudian menampilkan hasilnya.

### Algoritma Program
1. Input Panjang Sisi: Pengguna diminta untuk memasukkan panjang sisi persegi.
2. Hitung Luas: Program menghitung luas persegi menggunakan rumus `Luas = sisi*sisi` melalui prosedur `hitungLuasPersegi`.
3. Hitung Keliling: Program menghitung keliling persegi menggunakan rumus `Keliling = 4*sisi` melalui prosedur `hitungKelilingPersegi`.
4. Tampilkan Hasil: Program menampilkan hasil perhitungan luas dan keliling persegi kepada pengguna.

### Cara Kerja Program:
1. Program meminta pengguna memasukkan panjang sisi persegi, yang kemudian disimpan dalam variabel `panjangSisi`.
2. Program memanggil prosedur `hitungLuasPersegi`, yang menghitung luas dengan mengalikan panjang sisi dengan dirinya sendiri.
3. Setelah itu, program memanggil prosedur `hitungKelilingPersegi` untuk menghitung keliling persegi dengan mengalikan panjang sisi dengan 4.
4. Hasil dari perhitungan luas dan keliling persegi kemudian ditampilkan kepada pengguna dalam format tertentu.


## Unguided 

### 1.
### Source Code :
```go
package main

import (
	"fmt"
)

// Prosedur untuk menghitung faktorial
func HitungFaktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur untuk menghitung permutasi P(n, r)
func HitungPermutasi(n, r int, hasil *int) {
	var fn, fnr int
	HitungFaktorial(n, &fn)
	HitungFaktorial(n-r, &fnr)
	*hasil = fn / fnr
}

// Prosedur untuk menghitung kombinasi C(n, r)
func HitungKombinasi(n, r int, hasil *int) {
	var fn, fnr, fr int
	HitungFaktorial(n, &fn)
	HitungFaktorial(n-r, &fnr)
	HitungFaktorial(r, &fr)
	*hasil = fn / (fnr * fr)
}

func main() {
	// Input empat bilangan asli a, b, c, d
	var a, b, c, d int
	fmt.Print("Masukkan empat bilangan asli (a, b, c, d): ")
	fmt.Scan(&a, &b, &c, &d)

	// Syarat: a >= c dan b >= d
	if a < c || b < d {
		fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus >= d.")
		return
	}

	// Variabel untuk hasil permutasi dan kombinasi
	var hasilPermutasiA, hasilKombinasiA, hasilPermutasiB, hasilKombinasiB int

	// Hitung permutasi dan kombinasi untuk a terhadap c
	HitungPermutasi(a, c, &hasilPermutasiA)
	HitungKombinasi(a, c, &hasilKombinasiA)

	// Hitung permutasi dan kombinasi untuk b terhadap d
	HitungPermutasi(b, d, &hasilPermutasiB)
	HitungKombinasi(b, d, &hasilKombinasiB)

	// Output hasil
	fmt.Printf("\nPermutasi dan Kombinasi untuk %d dan %d: %d %d\n", a, c, hasilPermutasiA, hasilKombinasiA)
	fmt.Printf("Permutasi dan Kombinasi untuk %d dan %d: %d %d\n", b, d, hasilPermutasiB, hasilKombinasiB)
}

```

### Output:

### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program


### Cara Kerja Program:


### 2.
### Source Code :
```go

```

### Output:

### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program


### Cara Kerja Program:


### 3.
### Source Code :
```go

```

### Output:

### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program


### Cara Kerja Program:




## Kesimpulan 


## Daftar Pustaka
[1]
