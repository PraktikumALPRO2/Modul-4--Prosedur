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

### 1. Buatlah program beserta prosedur yang digunakan untuk menghitung nilai faktorial dan permutasi

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


### 2. Buatlah program beserta prosedur untuk menghitung luas persegi dan keliling persegi

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
![Screenshot 2024-10-15 221705](https://github.com/user-attachments/assets/60dd4435-7f88-4f00-9979-4fe6677f6698)

![Screenshot 2024-10-15 221754](https://github.com/user-attachments/assets/89bc1e81-57f9-4403-bd95-e41650f0db7f)

### Full code Screenshot:
![Screenshot 2024-10-15 221813](https://github.com/user-attachments/assets/6ae2f635-89da-4f18-9a0f-9fb628cbf84f)

### Deskripsi Program : 
Program ini menerima empat bilangan asli (`a`, `b`, `c`, dan `d`) sebagai input dari pengguna, kemudian menghitung dan menampilkan hasil permutasi dan kombinasi dari `a` terhadap `c`, serta dari `b` terhadap `d`. Program ini menggunakan rumus permutasi `P(n, r) = n!/(n-r)!` dan kombinasi `C(n, r) = n!/r!(n-r)!`, dan juga memverifikasi bahwa nilai `a` lebih besar atau sama dengan `c`, serta `b` lebih besar atau sama dengan `d`.

### Algoritma Program
1. Input Bilangan: Program meminta pengguna memasukkan empat bilangan asli `a`, `b`, `c`, dan `d`.
2. Validasi Kondisi: Program mengecek apakah nilai `a` lebih besar atau sama dengan `c`, dan `b` lebih besar atau sama dengan `d`. Jika kondisi tidak terpenuhi, program menampilkan pesan error dan berhenti.
3. Hitung Permutasi dan Kombinasi:
   - Program menghitung permutasi dan kombinasi dari pasangan `a` terhadap `c` dengan memanggil prosedur `HitungPermutasi` dan `HitungKombinasi`.
   - Program juga menghitung permutasi dan kombinasi dari pasangan `b` terhadap `d`.
4. Output Hasil: Program menampilkan hasil perhitungan permutasi dan kombinasi untuk kedua pasangan bilangan.

### Cara Kerja Program:
1. Program meminta input empat bilangan dari pengguna.
2. Setelah memeriksa apakah syarat `a >= c` dan `b >= d` terpenuhi, program melanjutkan dengan menghitung permutasi dan kombinasi menggunakan dua prosedur:
   - Prosedur `HitungPermutasi`: Menghitung permutasi `P(n, r)` dengan membagi faktorial dari `n` dengan faktorial dari `(n - r)`.
   - Prosedur `HitungKombinasi`: Menghitung kombinasi `C(n, r)` dengan membagi faktorial dari `n` dengan hasil perkalian faktorial `r` dan faktorial `(n - r)`.
3. Hasil perhitungan permutasi dan kombinasi ditampilkan ke layar untuk kedua pasangan bilangan (`a`, `c` dan `b`, `d`).

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

```

### Output:
![Screenshot 2024-10-15 234325](https://github.com/user-attachments/assets/0f6adad2-3530-47bb-ba10-f4d874f73a55)

### Full code Screenshot:
![Screenshot 2024-10-15 234343](https://github.com/user-attachments/assets/bdad2794-85f7-4e68-a8d5-e5e8f25afbe4)

### Deskripsi Program : 
Program ini menerapkan algoritma Collatz, yang juga dikenal sebagai masalah 3n + 1. Program ini meminta pengguna untuk memasukkan sebuah bilangan bulat positif dan kemudian mencetak urutan angka yang dihasilkan berdasarkan aturan berikut:
- Jika bilangan tersebut genap, bagi dengan 2.
- Jika bilangan tersebut ganjil, kalikan dengan 3 dan tambahkan 1. Proses ini berlangsung hingga bilangan tersebut mencapai 1, dan program akan mencetak seluruh urutan tersebut.
  
### Algoritma Program
1. Input: Minta pengguna untuk memasukkan bilangan bulat positif.
2. Inisialisasi: Ambil nilai bilangan yang dimasukkan oleh pengguna.
3. Proses:
    - Selama bilangan tidak sama dengan 1:
    - Cetak nilai bilangan.
    - Jika bilangan genap, bagi bilangan tersebut dengan 2.
    - Jika bilangan ganjil, kalikan bilangan tersebut dengan 3 dan tambahkan 1.
4. Setelah bilangan mencapai 1, cetak 1 dan akhiri proses.
5. Output: Tampilkan urutan yang dihasilkan.

### Cara Kerja Program:
1. Program dimulai dengan mendefinisikan prosedur `cetakDeret` yang menerima satu parameter `bilangan`.
2. Di dalam prosedur, selama `bilangan` tidak sama dengan 1, program mencetak nilai `bilangan` saat ini.
3. Jika `bilangan` genap, program membagi `bilangan` dengan 2.
4. Jika `bilangan` ganjil, program mengalikan `bilangan` dengan 3 dan menambahkan 1.
5. Ketika `bilangan` akhirnya menjadi 1, program mencetak 1 sebagai bagian terakhir dari urutan.
6. Dalam fungsi `main`, program meminta input dari pengguna dan memanggil prosedur `cetakDeret` untuk memproses dan mencetak urutan tersebut.



## Kesimpulan 


## Daftar Pustaka
[1]
