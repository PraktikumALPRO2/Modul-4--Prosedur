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
### Definisi Prosedur
Prosedur adalah blok kode yang dirancang untuk melakukan tugas tertentu, yang dapat dipanggil berulang kali dari berbagai tempat dalam program. Di Golang, prosedur diimplementasikan sebagai fungsi (function), meskipun istilah "prosedur" merujuk lebih spesifik pada fungsi yang tidak mengembalikan nilai. Prosedur biasanya digunakan untuk mengorganisir kode agar lebih modular dan mudah dipelihara.

Contoh sederhana prosedur dalam Golang:
```go
func tampilkanSelamat() {
    fmt.Println("Selamat datang!")
}
```
Prosedur ini mencetak pesan "Selamat datang!" ketika dipanggil.

### Cara Pemanggilan Prosedur
Setelah prosedur atau fungsi didefinisikan, pemanggilan prosedur dilakukan dengan menggunakan nama fungsi diikuti oleh tanda kurung. Jika prosedur memiliki parameter, nilai yang sesuai harus diberikan pada saat pemanggilan.

Contoh pemanggilan prosedur tanpa parameter:
```go
func main() {
    tampilkanSelamat() // Pemanggilan prosedur
}
```

Jika prosedur memiliki parameter, nilai aktual diberikan dalam tanda kurung saat memanggilnya. Misalnya:
```go


### Keuntungan Menggunakan Prosedur:
1. Modularitas: Dengan membagi program besar ke dalam beberapa fungsi kecil, kode menjadi lebih mudah dipahami dan dikelola. Setiap fungsi bertanggung jawab atas 	satu tugas spesifik.func tampilkanPesan(nama string) {
    fmt.Println("Halo", nama)
}

func main() {
    tampilkanPesan("Andi") // Pemanggilan dengan parameter "Andi"
}
```
### Pengertian Parameter
Parameter adalah nilai masukan yang diterima oleh prosedur atau fungsi untuk digunakan dalam proses perhitungan atau tugas lainnya di dalam fungsi tersebut. Parameter berperan penting dalam membuat prosedur atau fungsi menjadi lebih dinamis dan fleksibel.

Ada dua jenis parameter dalam konteks pemanggilan fungsi:

1. Parameter Formal: Parameter yang dideklarasikan dalam definisi fungsi. Ini adalah variabel placeholder yang menampung nilai saat fungsi dipanggil.
   
2. Parameter Aktual: Nilai nyata yang diberikan kepada fungsi saat dipanggil. Nilai ini dipetakan ke parameter formal dalam fungsi.
   
 ### Parameter Formal
Parameter formal adalah variabel yang disebutkan dalam definisi fungsi. Mereka berfungsi sebagai placeholder untuk menerima nilai saat fungsi dipanggil. Ketika fungsi dipanggil, parameter formal ini akan mengambil nilai dari parameter aktual yang diberikan dalam pemanggilan fungsi.

Contoh parameter formal:
```go
func hitungKeliling(panjangSisi float64) float64 { // 'panjangSisi' adalah parameter formal
    return 4 * panjangSisi
}
```
Pada contoh di atas, panjangSisi adalah parameter formal. Variabel ini akan menerima nilai aktual saat fungsi dipanggil.

### Parameter Aktual
Parameter aktual adalah nilai nyata yang diberikan ketika fungsi dipanggil. Nilai ini dipetakan ke parameter formal dalam fungsi. Parameter aktual bisa berupa nilai langsung atau variabel yang menyimpan nilai.

contoh paramenter aktual
```go
func main() {
    hasil := hitungKeliling(5.0) // '5.0' adalah parameter aktual
    fmt.Println("Keliling persegi:", hasil)
}
```
Pada contoh ini, `5.0` adalah parameter aktual yang diberikan saat fungsi `hitungKeliling` dipanggil. Nilai ini dipetakan ke `panjangSisi`, yang merupakan parameter formal di definisi fungsi `hitungKeliling`.



## Guided
### 1. Buatlah program beserta prosedur yang digunakan untuk menghitung nilai faktorial dan permutasi

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

 ### 2. Buatlah program beserta prosedur untuk menghitung luas persegi dan keliling persegi
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

## Unguided
### 1. Program untuk menghitung Faktorial, Permutasi dan Kombinasi

![image](https://github.com/user-attachments/assets/172b74a1-b5c8-4c0c-9adc-f8f98ea239d7)

```go
// Lutfiana Isnaeni Lathifah
// 2311102165
package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi P(n, r) = n! / (n - r)!
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r) = n! / (r! * (n - r)!)
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	// Input empat buah bilangan a, b, c, d
	var a, b, c, d int
	fmt.Print("Masukkan empat bilangan (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)

	// Asumsikan syarat a >= c dan b >= d sudah terpenuhi

	// Hitung permutasi dan kombinasi untuk (a, c) dan (b, d)
	permA := permutasi(a, c)
	kombA := kombinasi(a, c)
	permB := permutasi(b, d)
	kombB := kombinasi(b, d)

	// Cetak hasil
	fmt.Printf("Permutasi(%d, %d): %d\n", a, c, permA)
	fmt.Printf("Kombinasi(%d, %d): %d\n", a, c, kombA)
	fmt.Printf("Permutasi(%d, %d): %d\n", b, d, permB)
	fmt.Printf("Kombinasi(%d, %d): %d\n", b, d, kombB)
}
```
#### Output
![image](https://github.com/user-attachments/assets/a311b1d3-255d-4717-91dc-5cce0e02d62d)
### Deskripsi Program
Program di atas digunakan untuk menghitung permutasi dan kombinasi dari dua pasang bilangan yang diinput oleh pengguna. Dalam matematika, permutasi adalah cara menghitung jumlah kemungkinan pengaturan elemen dalam sebuah himpunan, sedangkan kombinasi menghitung jumlah cara memilih elemen tanpa memperhatikan urutannya. Program ini memanfaatkan tiga fungsi utama: `faktorial`, `permutasi`, dan `kombinasi`. Fungsi faktorial menghitung nilai `faktorial` dari suatu bilangan, yaitu hasil perkalian bilangan tersebut dengan bilangan-bilangan positif yang lebih kecil darinya. Fungsi permutasi menggunakan faktorial untuk menghitung jumlah cara mengatur `r` elemen dari`n` elemen, dengan rumus P(n, r) = n! / (n - r)!. Sedangkan fungsi `kombinasi` menghitung cara memilih `r` elemen dari `n` elemen, menggunakan rumus C(n, r) = n! / (r! * (n - r)!). Program ini meminta input dari pengguna berupa empat bilangan a, b, c, dan d. Setelah itu, program menghitung `permutasi` dan kombinasi untuk pasangan bilangan (a, c) dan (b, d) menggunakan fungsi-fungsi yang telah dibuat.

### Algoritma Program

•  Program dimulai dengan meminta input empat bilangan: `a`, `b`, `c`, dan `d`.

•  Fungsi `permutasi(a, c)` dipanggil untuk menghitung permutasi dari `a` dan `c`.

•  Fungsi `kombinasi(a, c)` dipanggil untuk menghitung kombinasi dari `a` dan `c`.

•  Fungsi `permutasi(b, d)` dipanggil untuk menghitung permutasi dari `b` dan `d`.

•  Fungsi `kombinasi(b, d)` dipanggil untuk menghitung kombinasi dari `b` dan `d`.

•  Program kemudian mencetak hasil perhitungan permutasi dan kombinasi.

### Cara kerja Program

1. Program meminta pengguna memasukkan empat bilangan integer sebagai input.

2. Bilangan-bilangan ini kemudian diproses oleh fungsi permutasi dan kombinasi.
   
3. Hasil dari perhitungan permutasi dan kombinasi untuk pasangan `(a, c)` dan `(b, d)` ditampilkan di layar.

### 2. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan
Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"strings"
)

func hitungSkor(soal []int) (int, int) {
	skor := 0
	totalWaktu := 0

	// Loop untuk menghitung soal yang berhasil diselesaikan dan total waktu
	for i := 0; i < len(soal); i++ {
		if soal[i] <= 301 {
			skor++
			totalWaktu += soal[i]
		}
	}

	return skor, totalWaktu
}

func main() {
	var peserta string
	var pesertaTerbaik string
	maxSkor := -1
	waktuTerbaik := 999999
	skorPeserta := 0
	totalWaktuPeserta := 0

	for {
		// Membaca input peserta
		fmt.Scanln(&peserta)
		if peserta == "Selesai" {
			break
		}

		// Memecah input menjadi bagian-bagian
		data := strings.Split(peserta, " ")
		nama := data[0]
		soal := make([]int, 8)

		for i := 0; i < 8; i++ {
			fmt.Sscan(data[i+1], &soal[i])
		}

		// Menghitung skor dan total waktu
		skorPeserta, totalWaktuPeserta = hitungSkor(soal)

		// Memilih pemenang berdasarkan skor dan waktu penyelesaian
		if skorPeserta > maxSkor || (skorPeserta == maxSkor && totalWaktuPeserta < waktuTerbaik) {
			maxSkor = skorPeserta
			waktuTerbaik = totalWaktuPeserta
			pesertaTerbaik = nama
		}
	}

	// Menampilkan hasil pemenang
	fmt.Printf("%s %d %d\n", pesertaTerbaik, maxSkor, waktuTerbaik)
}
```
#### Output

### Deskripsi Program 
### Algoritma Program
### Cara Kerja Program

### 3. Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.
![image](https://github.com/user-attachments/assets/9540f0a7-ac70-4a88-8bed-fa641069493e)
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi untuk mencetak deret bilangan Skiena
func cetakDeret(n int) {
	// Looping hingga n mencapai 1
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2 // Jika n genap, bagi 2
		} else {
			n = 3*n + 1 // Jika n ganjil, kalikan 3 lalu tambah 1
		}
	}
	// Cetak nilai akhir (1)
	fmt.Println(n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai awal: ")
	fmt.Scan(&n)

	// Panggil prosedur cetakDeret untuk menampilkan hasilnya
	cetakDeret(n)
}
```


#### Output
![image](https://github.com/user-attachments/assets/e02ea9e6-1fcc-4230-855f-b1bc3b19b4b3)

### Deskripsi Program
digunakan untuk menghasilkan deret bilangan berdasarkan aturan sederhana. Dimulai dengan bilangan `n`, jika bilangan tersebut genap, maka dibagi dua, sedangkan jika ganjil, dikalikan tiga lalu ditambah satu. Proses ini berlanjut hingga bilangan mencapai angka 1. Program ini meminta input nilai awal dari pengguna dan kemudian mencetak deret bilangan berdasarkan aturan tersebut.

Program ini menerima input integer n dari pengguna dan mencetak deret bilangan berdasarkan aturan:

1.Jika `n` adalah bilangan genap, maka dibagi dua.

2. Jika `n` adalah bilangan ganjil, maka dikalikan tiga dan ditambah satu. Proses ini berulang hingga `n` menjadi 1, dan pada akhirnya program mencetak 1 sebagai nilai akhir dari deret.
   
### Algoritma Program
1. Program dimulai dengan menerima input integer `n` dari pengguna.
   
2. Dilakukan looping yang terus berlanjut selama nilai `n` tidak sama dengan 1.
   
3. Jika `n` genap, maka `n` dibagi dua.
   
4. Jika `n` ganjil, maka` n` dikalikan tiga dan ditambah satu.
   
5. Setelah setiap perubahan nilai, `n` dicetak ke layar.
   
6. Ketika nilai `n` sudah sama dengan 1, looping berhenti dan program mencetak angka 1 sebagai penutup deret.
   
### Cara Kerja Program
1. Program meminta pengguna untuk memasukkan bilangan bulat n.
 
2. Program kemudian memanggil fungsi cetakDeret yang akan menjalankan loop untuk memproses nilai n hingga mencapai 1.

3. Selama looping, nilai n akan diperiksa apakah genap atau ganjil. Jika genap, n dibagi dua. Jika ganjil, n dikalikan tiga dan ditambah satu.

4. Setiap perubahan pada n akan dicetak di layar sampai akhirnya n menjadi 1.

5. Setelah itu, program mencetak 1 dan selesai.




