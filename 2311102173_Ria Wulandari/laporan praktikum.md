
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
  Ria Wulandari / 2311102173<br>
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

## Dasar Teori
#### 1. Definisi Prosedur
Dalam konteks pemrograman, prosedur atau fungsi adalah blok kode yang dirancang untuk melakukan tugas tertentu. Prosedur dapat menerima input (parameter), melakukan operasi, dan dapat mengembalikan nilai. Fungsi membantu dalam organisasi kode, membuatnya lebih terstruktur, dan memudahkan pemeliharaan serta penggunaan kembali.
#### 2. Deklarasi Prosedur
Deklarasi fungsi di Go dilakukan menggunakan kata kunci `func`, diikuti oleh nama fungsi, parameter (jika ada), tipe pengembalian (jika ada), dan blok kode yang berisi logika fungsi tersebut. Berikut adalah sintaks umum untuk mendeklarasikan fungsi:
```go
func namaFungsi(parameter1 tipe1, parameter2 tipe2, ...) returnType {
    // blok kode
    return nilai
}
```
##### Contoh deklarasi fungsi
Berikut adalah contoh deklarasi fungsi sederhana yang menjumlahkan dua bilangan:
```go
func tambah(a int, b int) int {
    return a + b
}
```
#### 3. Cara Pemanggilan Prosedur
Untuk memanggil fungsi yang telah dideklarasikan, Anda cukup menyebutkan nama fungsi dan memberikan argumen yang sesuai dengan parameter yang telah ditentukan. Berikut adalah contoh cara memanggil fungsi:
```go
func main() {
    hasil := tambah(3, 5) // Memanggil fungsi tambah dengan argumen 3 dan 5
    fmt.Println("Hasil penjumlahan:", hasil)
}
```
#### 4. Parameter Fungsi
Parameter dalam fungsi Go dapat dibagi menjadi dua kategori: parameter formal dan parameter aktual.
##### a.Parameter Formal
Parameter formal adalah parameter yang didefinisikan dalam deklarasi fungsi. Parameter ini bertindak sebagai variabel yang akan menerima nilai saat fungsi dipanggil. Dalam contoh fungsi `tambah`, `a` dan `b` adalah parameter formal.
```go
func tambah(a int, b int) int {
    return a + b
}
```
##### b.  Parameter Aktual
Parameter aktual adalah nilai yang sebenarnya diberikan kepada parameter formal saat fungsi dipanggil. Dalam pemanggilan fungsi `tambah(3, 5)`, nilai `3` dan `5` adalah parameter aktual yang diberikan kepada fungsi.
```go
func main() {
    hasil := tambah(3, 5) // 3 dan 5 adalah parameter aktual
}
```

## Guided
### 1. Buatlah program beserta prosedur yang digunakan untuk menghitung nilai faktorial dan permutasi
#### Source Code :
```go
package main

import "fmt"

// Prosedur untuk membaca input dan mencetak hasil permutasi
func bacaDanCetakPermutasi() {
    var a, b int
    fmt.Scan(&a, &b)
    if a >= b {
        fmt.Println(permutasi(a, b))
    } else {
        fmt.Println(permutasi(b, a))
    }
}

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
    var hasil int = 1
    for i := 1; i <= n; i++ {
        hasil *= i
    }
    return hasil
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int) int {
    return faktorial(n) / faktorial(n-r)
}

func main() {
    bacaDanCetakPermutasi()
}
```
#### Output
![image](https://github.com/user-attachments/assets/ac23afb4-ca41-4fae-81dc-a9463d7140f0)

#### Deskripsi program
Program ini membaca dua angka dari pengguna, lalu menghitung permutasi dengan menentukan angka yang lebih besar sebagai `n` dan yang lebih kecil sebagai `r`. Permutasi dihitung menggunakan rumus `n! / (n - r)!`, di mana faktorial dihitung dengan mengalikan semua bilangan dari 1 hingga angka tersebut. Program kemudian mencetak hasil perhitungan permutasi ke layar, membantu pengguna mengetahui berapa banyak cara memilih dan mengatur objek.
#### Algoritma program
1. Mulai.

2. Baca dua angka `a` dan `b` dari input pengguna.

3. Bandingkan `a` dan `b`:
- Jika `a >= b`, maka `n = a` dan `r = b`.

- Jika `b > a`, maka `n = b` dan `r = a`.

4. Hitung faktorial `n` dan faktorial `(n - r)`:
- Faktorial `n` dihitung sebagai hasil perkalian semua bilangan dari 1 hingga `n`.

- Faktorial `(n - r)` dihitung sebagai hasil perkalian semua bilangan dari 1 hingga `(n - r)`.

5. Hitung permutasi dengan rumus: `n! / (n - r)!`.

6. Cetak hasil perhitungan permutasi.

7. Selesai.
#### Cara kerja program
1. Membaca Input:

- Program mulai dengan meminta pengguna untuk memasukkan dua angka `(a dan b)`.
- Pengguna akan mengetikkan dua angka, dan program menyimpannya sebagai `a` dan `b`.

2. Mengecek Nilai Angka:

- Program mengecek apakah `a` lebih besar atau sama dengan `b`. Jika ya, `a` digunakan sebagai `n` dan `b` sebagai `r`.
- Jika `b` lebih besar, nilai `b` digunakan sebagai `n` dan `a` sebagai `r`.

3. Menghitung Faktorial:

- Program menghitung faktorial `n` menggunakan fungsi `faktorial`. Faktorial adalah hasil perkalian semua angka dari 1 hingga `n`.
- Program juga menghitung faktorial `(n - r)` dengan cara yang sama.

4. Menghitung Permutasi:
Permutasi dihitung dengan rumus `n! / (n - r)!`. Ini dilakukan dengan membagi faktorial `n` dengan faktorial `(n - r)`.

5. Menampilkan Hasil:
Hasil dari perhitungan permutasi ditampilkan kepada pengguna dalam bentuk angka yang menunjukkan jumlah cara memilih dan mengatur `r` objek dari `n` objek.

6. Program Selesai:

Setelah hasil ditampilkan, program selesai menjalankan tugasnya.

### 2. Buatlah program beserta prosedur untuk menghitung luas persegi dan keliling persegi
#### Source Code :
``` go
package main

import "fmt"

// Prosedur untuk menghitung luas persegi
func hitungLuasPersegi(sisi float64) float64 {
    return sisi * sisi
}

// Prosedur untuk menghitung keliling persegi
func hitungKelilingPersegi(sisi float64) float64 {
    return 4 * sisi
}

func main() {
    // Meminta input panjang sisi persegi
    var panjangSisi float64
    fmt.Print("Masukkan panjang sisi persegi: ")
    fmt.Scan(&panjangSisi)

    // Menghitung luas dan keliling persegi menggunakan prosedur
    luasPersegi := hitungLuasPersegi(panjangSisi)
    kelilingPersegi := hitungKelilingPersegi(panjangSisi)

    // Menampilkan hasil
    fmt.Printf("Luas persegi: %g\n", luasPersegi)
    fmt.Printf("Keliling persegi: %g\n", kelilingPersegi)
}
```
#### Output
![image](https://github.com/user-attachments/assets/6cc6c8b6-c947-4697-bb83-305c53c1b935)
#### Deskripsi program
Program ini menghitung luas dan keliling persegi berdasarkan input panjang sisi dari pengguna. Pertama, pengguna diminta untuk memasukkan panjang sisi persegi. Program kemudian menggunakan dua fungsi: `hitungLuasPersegi` untuk menghitung luas dengan rumus `sisi * sisi`, dan `hitungKelilingPersegi` untuk menghitung keliling dengan rumus `4 * sisi`. Setelah kedua nilai tersebut dihitung, program menampilkan hasil luas dan keliling persegi di layar.
#### Algoritma program
1. Mulai.

2. Minta Input: Tanyakan kepada pengguna untuk memasukkan panjang sisi persegi.

3. Baca Input: Simpan panjang sisi yang dimasukkan pengguna ke dalam variabel `panjangSisi`.

4. Hitung Luas:
- Panggil fungsi `hitungLuasPersegi` dengan `panjangSisi` sebagai argumen.
- Hitung luas dengan rumus: `luas = sisi * sisi`.

5. Hitung Keliling:
- Panggil fungsi `hitungKelilingPersegi` dengan `panjangSisi` sebagai argumen.
- Hitung keliling dengan rumus: `keliling = 4 * sisi`.

6. Tampilkan Hasil: Tampilkan hasil luas dan keliling ke layar.

7. Selesai.
#### Cara kerja program
1. Pengambilan Input: Program meminta pengguna untuk memasukkan panjang sisi persegi. Nilai yang dimasukkan disimpan dalam variabel `panjangSisi`.

2. Menghitung Luas: Program memanggil fungsi `hitungLuasPersegi`, yang menghitung luas dengan rumus `sisi * sisi`. Nilai hasil perhitungan luas dikembalikan dan disimpan dalam variabel `luasPersegi`.

3. Menghitung Keliling: Program kemudian memanggil fungsi `hitungKelilingPersegi`, yang menghitung keliling dengan rumus `4 * sisi`. Hasil perhitungan keliling disimpan dalam variabel `kelilingPersegi`.

4. Menampilkan Hasil: Program mencetak hasil luas dan keliling ke layar menggunakan fungsi `fmt.Printf`, menampilkan nilai yang telah dihitung dengan format yang mudah dibaca.

5. Program Selesai: Setelah menampilkan hasil, program berakhir dan tidak ada langkah lebih lanjut yang dilakukan.
## Unguided
### 1. Program untuk menghitung Faktorial, Permutasi dan Kombinasi
![image](https://github.com/user-attachments/assets/080d9bb9-16b9-439a-8a36-a0180a7e338d)
#### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Fungsi untuk menghitung permutasi P(n, r) = n! / (n-r)!
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r) = n! / (r!(n-r)!)
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	// Input
	fmt.Println("Masukkan empat bilangan asli (a, b, c, d):")
	fmt.Scan(&a, &b, &c, &d)

	// Baris pertama: permutasi dan kombinasi a terhadap c
	perm1 := permutation(a, c)
	comb1 := combination(a, c)

	// Baris kedua: permutasi dan kombinasi b terhadap d
	perm2 := permutation(b, d)
	comb2 := combination(b, d)

	// Output
	fmt.Printf("Permutasi dan Kombinasi untuk %d dan %d: %d %d\n", a, c, perm1, comb1)
	fmt.Printf("Permutasi dan Kombinasi untuk %d dan %d: %d %d\n", b, d, perm2, comb2)
}
```
#### Output
![Cuplikan layar 2024-10-18 203126](https://github.com/user-attachments/assets/d35e57d2-e7a4-4b4a-9845-fc8e7e9f8297)
#### Deskripsi program
Program ini dirancang untuk menghitung permutasi dan kombinasi dari empat bilangan bulat yang dimasukkan oleh pengguna. Pertama, program mendefinisikan fungsi untuk menghitung faktorial, yang digunakan dalam proses perhitungan permutasi dan kombinasi. Setelah itu, program meminta pengguna untuk memasukkan empat bilangan, yang kemudian digunakan untuk menghitung permutasi dan kombinasi untuk dua pasangan angka: (a, c) dan (b, d). Hasil perhitungan permutasi dan kombinasi ini ditampilkan ke layar dalam format yang jelas, memberikan pengguna informasi tentang cara mengatur dan memilih objek dari bilangan yang dimasukkan.
#### Algoritma program
1. Mulai.

2. Definisikan Fungsi:
- Buat fungsi `factorial` untuk menghitung faktorial dari bilangan bulat `n`.
- Buat fungsi `permutation` untuk menghitung permutasi dengan menggunakan fungsi faktorial.
- Buat fungsi `combination` untuk menghitung kombinasi dengan menggunakan fungsi faktorial.

3. Input Pengguna:
- Tampilkan pesan kepada pengguna untuk memasukkan empat bilangan bulat (a, b, c, d).
- Simpan input pengguna ke dalam variabel `a`, `b`, `c`, dan `d`.

4. Hitung Permutasi dan Kombinasi:
- Hitung permutasi dari `a` dan `c` dan simpan hasilnya ke variabel `perm1`.
- Hitung kombinasi dari `a` dan `c` dan simpan hasilnya ke variabel `comb1`.
- Hitung permutasi dari `b` dan `d` dan simpan hasilnya ke variabel `perm2`.
- Hitung kombinasi dari `b` dan `d` dan simpan hasilnya ke variabel `comb2`.

5. Tampilkan Hasil:
- Tampilkan hasil permutasi dan kombinasi untuk pasangan (a, c) ke layar.
- Tampilkan hasil permutasi dan kombinasi untuk pasangan (b, d) ke layar.

6. Selesai.
#### Cara kerja program
1. Pengambilan Input: Program mulai dengan meminta pengguna untuk memasukkan empat bilangan bulat (a, b, c, dan d). Angka-angka ini akan digunakan dalam perhitungan permutasi dan kombinasi.

2. Menghitung Faktorial: Program memiliki fungsi `factorial` yang menghitung faktorial dari bilangan bulat yang diberikan. Faktorial ini merupakan dasar untuk menghitung permutasi dan kombinasi.

3. Menghitung Permutasi: Program menggunakan fungsi `permutation` untuk menghitung permutasi dari bilangan `a` dan `c`, serta `b` dan `d`. Hasil dari perhitungan ini disimpan dalam variabel `perm1` dan `perm2`.

4. Menghitung Kombinasi: Program juga menghitung kombinasi menggunakan fungsi `combination` untuk pasangan yang sama, yaitu `a` dan `c`, serta `b` dan `d`. Hasilnya disimpan dalam variabel `comb1` dan `comb2`.

5. Menampilkan Hasil: Setelah semua perhitungan selesai, program mencetak hasil permutasi dan kombinasi ke layar untuk kedua pasangan angka yang dimasukkan, memberikan informasi tentang cara mengatur dan memilih objek.

6. Program Selesai: Setelah hasil ditampilkan, program berakhir dan tidak ada langkah lebih lanjut yang dilakukan.
### 2. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan
Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.
#### Source Code :
```go
package main

import (
    "fmt"
    "math"
)

// Function to calculate score
func hitungSkor(waktuPengerjaan []int) (int, int) {
    soalSelesai := 0
    totalSkor := 0
    for _, waktu := range waktuPengerjaan {
        if waktu < 301 {
            soalSelesai++
            totalSkor += waktu
        }
    }
    return soalSelesai, totalSkor
}

// Function to determine the winner
func cariPemenang(peserta map[string][]int) (string, int, int) {
    pemenang := ""
    maxSoal := -1
    minSkor := math.MaxInt64
    
    for nama, waktuPengerjaan := range peserta {
        soalSelesai, totalSkor := hitungSkor(waktuPengerjaan)
        if soalSelesai > maxSoal || (soalSelesai == maxSoal && totalSkor < minSkor) {
            pemenang = nama
            maxSoal = soalSelesai
            minSkor = totalSkor
        }
    }
    return pemenang, maxSoal, minSkor
}

func main() {
    var n int
    fmt.Println("Masukkan jumlah peserta:")
    fmt.Scan(&n)
    
    peserta := make(map[string][]int)
    
    for i := 0; i < n; i++ {
        var nama string
        fmt.Printf("Masukkan nama peserta %d:\n", i+1)
        fmt.Scan(&nama)
        
        waktuPengerjaan := make([]int, 8)
        fmt.Printf("Masukkan waktu pengerjaan soal untuk %s:\n", nama)
        for j := 0; j < 8; j++ {
            fmt.Scan(&waktuPengerjaan[j])
        }
        
        peserta[nama] = waktuPengerjaan
    }
    
    // Determine the winner
    pemenang, maxSoal, minSkor := cariPemenang(peserta)
    fmt.Printf("%s %d %d\n", pemenang, maxSoal, minSkor)
}
```
#### Output
![Cuplikan layar 2024-10-18 202848](https://github.com/user-attachments/assets/511b4a97-f20c-43ea-89a6-b20ca5c2b40a)
#### Deskripsi program
Program ini dirancang untuk menghitung dan menentukan pemenang sebuah kompetisi berdasarkan waktu pengerjaan soal peserta. Pertama, program meminta pengguna untuk memasukkan jumlah peserta, diikuti dengan nama dan waktu pengerjaan untuk setiap peserta untuk delapan soal. Dengan menggunakan fungsi `hitungSkor`, program menghitung jumlah soal yang diselesaikan dalam waktu kurang dari 301 detik dan total waktu yang dihabiskan. Kemudian, fungsi `cariPemenang` digunakan untuk menentukan pemenang berdasarkan jumlah soal yang diselesaikan dan skor terendah. Setelah semua data peserta diproses, program mencetak nama pemenang, jumlah soal yang berhasil diselesaikan, dan total skor yang didapat, memberikan informasi yang jelas tentang hasil kompetisi.
#### Algoritma program
1. Mulai.

2. Input Jumlah Peserta: Tanyakan kepada pengguna untuk memasukkan jumlah peserta.

3. Inisialisasi Struktur Data: Buat peta (map) untuk menyimpan nama peserta dan waktu pengerjaan soal mereka.

4. Pengambilan Data Peserta: Untuk setiap peserta dari 1 hingga n, tanyakan nama peserta dan simpan dalam variabel. Inisialisasi array untuk menyimpan waktu pengerjaan delapan soal, kemudian tanyakan waktu pengerjaan untuk setiap soal dan simpan dalam array waktu pengerjaan. Simpan nama peserta dan waktu pengerjaan dalam peta.

5. Menentukan Pemenang: Inisialisasi variabel untuk menyimpan nama pemenang, jumlah soal maksimum yang diselesaikan, dan skor minimum. Untuk setiap peserta dalam peta, panggil fungsi `hitungSkor` untuk mendapatkan jumlah soal yang diselesaikan dan total skor. Jika jumlah soal yang diselesaikan lebih besar dari jumlah maksimum, update nama pemenang, jumlah soal maksimum, dan skor minimum. Jika jumlah soal yang diselesaikan sama dengan jumlah maksimum dan skor lebih kecil dari skor minimum, update nama pemenang dan skor minimum.

6. Tampilkan Hasil: Cetak nama pemenang, jumlah soal yang diselesaikan, dan total skor ke layar.

7. Selesai.
#### Cara kerja program
1. Input Jumlah Peserta: Program dimulai dengan meminta pengguna untuk memasukkan jumlah peserta yang akan berpartisipasi dalam kompetisi.

2. Pengambilan Data Peserta: Program kemudian meminta nama untuk setiap peserta dan waktu pengerjaan mereka untuk delapan soal. Waktu pengerjaan ini disimpan dalam sebuah array yang terkait dengan nama peserta di dalam peta (map).

3. Menghitung Skor: Setelah semua data peserta diinput, program menggunakan fungsi `hitungSkor` untuk menghitung jumlah soal yang diselesaikan dalam waktu kurang dari 301 detik dan total skor yang diperoleh setiap peserta.

4. Menentukan Pemenang: Program kemudian mengevaluasi hasil dari semua peserta. Untuk menentukan pemenang, program membandingkan jumlah soal yang diselesaikan oleh setiap peserta. Jika ada peserta yang menyelesaikan lebih banyak soal, mereka ditetapkan sebagai pemenang. Jika ada beberapa peserta yang menyelesaikan jumlah soal yang sama, pemenang ditentukan berdasarkan skor terendah.

5. Menampilkan Hasil: Setelah proses penentuan pemenang selesai, program mencetak nama pemenang, jumlah soal yang diselesaikan, dan total skor mereka ke layar, memberikan informasi yang jelas tentang hasil kompetisi.

6. Selesai: Program berakhir setelah menampilkan hasil pemenang.
### 3. Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.
#### Source Code :
```go
package main

import "fmt"

// Function to print the sequence
func cetakDeret(n int) {
    for n != 1 {
        fmt.Printf("%d ", n)
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
    }
    fmt.Printf("1\n") // Print the last 1
}

func main() {
    var n int
    fmt.Print("Masukkan nilai awal: ")
    fmt.Scan(&n)

    cetakDeret(n)
}
```
#### Output
![Cuplikan layar 2024-10-18 203051](https://github.com/user-attachments/assets/254b3ccc-c3fb-4188-b039-4402dc7d5283)
#### Deskripsi program
Program ini dirancang untuk mencetak deret berdasarkan aturan yang dikenal sebagai "Collatz conjecture" atau "masalah 3n + 1." Dimulai dengan meminta pengguna untuk memasukkan sebuah nilai awal berupa bilangan bulat, program kemudian menggunakan fungsi `cetakDeret` untuk menghasilkan deret tersebut. Dalam fungsi ini, program melakukan iterasi selama nilai n tidak sama dengan 1. Setiap iterasi mencetak nilai n saat ini, dan jika n adalah bilangan genap, nilai tersebut dibagi 2; jika n adalah bilangan ganjil, n dikalikan dengan 3 dan ditambahkan 1. Proses ini terus berlanjut hingga n mencapai 1, di mana angka 1 juga dicetak sebagai bagian terakhir dari deret. Dengan demikian, program memberikan output deret yang mengikuti aturan Collatz, dimulai dari nilai yang diberikan oleh pengguna hingga berakhir di angka 1.
#### Algoritma program
1. Mulai.

2. Input Nilai Awal: Tanyakan kepada pengguna untuk memasukkan nilai awal (bilangan bulat) dan simpan dalam variabel `n`.

3. Fungsi `cetakDeret`:
- Selama `n` tidak sama dengan 1, lakukan langkah berikut:
- Cetak nilai `n`.
- Jika `n` adalah bilangan genap, ubah `n` menjadi `n / 2`.
- Jika `n` adalah bilangan ganjil, ubah `n` menjadi `3 * n + 1`.

4. Cetak Angka 1: Setelah loop selesai, cetak angka 1 sebagai bagian terakhir dari deret.

5. Selesai.
#### Cara kerja program
1. Input Nilai Awal: Program dimulai dengan meminta pengguna untuk memasukkan sebuah nilai awal berupa bilangan bulat.

2. Fungsi `cetakDeret`:
- Setelah pengguna memasukkan nilai awal, program memanggil fungsi `cetakDeret`, yang menerima nilai tersebut sebagai argumen.
- Di dalam fungsi, program melakukan iterasi selama nilai n tidak sama dengan 1.
- Dalam setiap iterasi, nilai `n` saat ini dicetak ke layar.
- Program memeriksa apakah `n` adalah bilangan genap atau ganjil:
- Jika `n` genap, maka nilai `n` dibagi 2.
- Jika `n` ganjil, maka nilai `n` dikalikan dengan 3 dan ditambahkan 1.
- Proses ini akan terus berlanjut hingga nilai `n` menjadi 1.

3. Menampilkan Hasil: Setelah mencapai angka 1, program mencetak angka 1 sebagai bagian terakhir dari deret.

4. Selesai: Program berakhir setelah mencetak seluruh deret dari nilai awal hingga angka 1.
