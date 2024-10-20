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
  Aditya Sulistiawan / 2311102193<br>
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
Prosedur adalah suatu program yang terpisah dalam blok sendiri yang berfungsi sebagai seubprogram (program bagian)[81]–[90]. Prosedur diawali dengan kata cadangan PROCEDURE  di dalam bagian deklarasi prosedur[91]–[100]. Prosedur dipanggil dan digunakan di dalam blok program lainnya dengan menyebutkan judul prosedurnya

### Deklarasi Prosedur
Deklarasi prosedur adalah sebuah deklarasi yang dibuat dalam program agar bisa digunakan lagi dalam program, jadi deklarasi prosedur merupakan sebuah SUB PROGRAM yang bisa dipanggil sewaktu-waktu dalam program bila dibutuhkan. Dalam sebuah program yang terstruktur, Deklarasi Prosedur ini sangat dibutuhkan agar program menjadi lebih simpel.

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
/// Aditya Sulistiawan
/// 2311102193
/// 11-05

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
#### Screenshoot Source Code
![SC](https://github.com/user-attachments/assets/15142e99-b517-49d6-b026-301a94284227)





#### Screenshoot Output
![gui1](https://github.com/user-attachments/assets/47055452-c971-47ed-96dc-051f7ca2ab65)




#### Deskripsi Program
Program ini ditulis dalam bahasa Go dan bertujuan untuk menghitung permutasi dari dua bilangan bulat yang dimasukkan oleh pengguna. Permutasi adalah cara untuk menghitung jumlah cara untuk mengatur sejumlah objek yang berbeda. Dalam konteks ini, program menghitung permutasi ( nPr ), di mana ( n ) adalah jumlah total objek dan ( r ) adalah jumlah objek yang akan diambi

#### Algoritma Program
1. Mulai
2. Deklarasi Variabel: Deklarasikan dua variabel integer ( a ) dan ( b ).
3. Input: Minta pengguna untuk memasukkan dua bilangan bulat ( a ) dan ( b ).
4. Perbandingan:
	Jika ( a \geq b ):
	 Panggil fungsi permutasi(a, b).
	Jika ( b > a ):
	 Panggil fungsi permutasi(b, a).
5. Fungsi faktorial(n):
	Inisialisasi variabel hasil dengan 1.
	Untuk setiap ( i ) dari 1 hingga ( n ):
	Kalikan hasil dengan ( i ).
	Kembalikan hasil.
6. Fungsi permutasi(n, r):
	Hitung hasil permutasi dengan rumus ( \frac{n!}{(n-r)!} ) menggunakan fungsi faktorial.
	Cetak hasil permutasi.
7. Selesai
#### Cara Kerja
1. Input Pengguna: Program meminta pengguna untuk memasukkan dua bilangan bulat, ( a ) dan ( b ).

2. Perbandingan Nilai: Program membandingkan nilai ( a ) dan ( b ):
	Jika ( a ) lebih besar atau sama dengan ( b ), program akan menghitung permutasi ( a ) dan ( b ) (dalam hal ini, ( n = a ) dan ( r = b )).
	Jika ( b ) lebih besar dari ( a ), program akan menghitung permutasi ( b ) dan ( a ) (dalam hal ini, ( n = b ) dan ( r = a )).
3. Menghitung Faktorial: Program menggunakan fungsi faktorial untuk menghitung faktorial dari bilangan yang diberikan.
4. Menghitung Permutasi: Fungsi permutasi menghitung nilai permutasi dengan rumus: [ nPr = \frac{n!}{(n-r)!} ] Hasilnya kemudian dicetak ke layar.

## 2. Program Penghitungan Luas dan Keliling Persegi
#### Source Code
```go
package main

import "fmt"


func hitungLuas(sisi float64) {
	luas := sisi * sisi
	fmt.Printf("Luas persegi adalah: %.2f\n", luas)
}


func hitungKeliling(sisi float64) {
	keliling := 4 * sisi
	fmt.Printf("Keliling persegi adalah: %.2f\n", keliling)
}

func main() {
	var sisi float64

	fmt.Print("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi)

	hitungLuas(sisi)      
	hitungKeliling(sisi)
}
```
#### Screenshoot Source Code
![SC](https://github.com/user-attachments/assets/2bc740a7-ebfa-433c-a4c6-2d476c52b509)


#### Screenshoot Output
![Gui2](https://github.com/user-attachments/assets/2b9b8ad4-2551-41e1-91bc-5ad23af9d2d3)


#### Deskripsi Program
Program ini ditulis dalam bahasa Go dan bertujuan untuk menghitung luas dan keliling dari sebuah persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna. Program ini menggunakan prosedur untuk melakukan perhitungan dan mencetak hasilnya secara langsung.

#### Algoritma Program
1. Mulai
2. Deklarasi Variabel: Deklarasikan variabel sisi bertipe float64.
3. Input: Minta pengguna untuk memasukkan panjang sisi persegi.
4. Panggil Prosedur hitungLuas(sisi):
	-Hitung luas persegi dan cetak hasilnya.
5. Panggil Prosedur hitungKeliling(sisi):
	-Hitung keliling persegi dan cetak hasilnya.
6. Selesai

#### Cara Kerja
1. Input Pengguna: Program meminta pengguna untuk memasukkan panjang sisi persegi.
2. Menghitung Luas: Program menggunakan prosedur hitungLuas untuk menghitung luas persegi dengan rumus: [ \text{Luas} = \text{sisi} \times \text{sisi} ] Hasilnya dicetak ke layar.
3. Menghitung Keliling: Program menggunakan prosedur hitungKeliling untuk menghitung keliling persegi dengan rumus: [ \text{Keliling} = 4 \times \text{sisi} ] Hasilnya juga dicetak ke layar.

## III. UNGUIDED
## 1. Program Penghitung Permutasi dan Kombinasi
1. Minggu Ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kullah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, Iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakan kaltan membantu Jonas? (tidak tentunya ya :p)

Masukan : Terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a >= c dan b >= d

Keluaran : Terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
#### Source Code
```go
// Aditya Sulistiawan
// 2311102193
// S1IF11-05

package main

import "fmt"

func main() {
	var a, b, c, d int

	// fungsi untuk menginput bilangan a, b, c, dan d
	fmt.Print("Bilangan a: ")
	fmt.Scan(&a)

	fmt.Print("Bilangan b: ")
	fmt.Scan(&b)

	fmt.Print("Bilangan c: ")
	fmt.Scan(&c)

	fmt.Print("Bilangan d: ")
	fmt.Scan(&d)

	// Untuk menjelaskan syarat bahwa a>= c dan b>=d
	if a >= c && b >= d {
		var permutasiAC, kombinasiAC, permutasiBD, kombinasiBD int

		// Untuk menampilkan hasil permutasi dan kombinasi dari bilangan a dan c
		permutasi(a, c, &permutasiAC)
		kombinasi(a, c, &kombinasiAC)
		fmt.Printf("\nPermutasi (a, c): %d\n", permutasiAC)
		fmt.Printf("Kombinasi (a, c): %d\n", kombinasiAC)

		// Untuk menampilkan hasil permutasi dan kombinasi dari bilangan b dan d
		permutasi(b, d, &permutasiBD)
		kombinasi(b, d, &kombinasiBD)
		fmt.Printf("\nPermutasi (b, d): %d\n", permutasiBD)
		fmt.Printf("Kombinasi (b, d): %d\n", kombinasiBD)

	} else {
		// Untuk yang tidak memenuhi syarat
		fmt.Println("Input tidak sesuai dengan syarat yang ada")
	}
}

// Prosedur untuk menghitung faktorial
func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur untuk menghitung permutasi
func permutasi(n, r int, hasil *int) {
	var faktorialN, faktorialNR int
	faktorial(n, &faktorialN)
	faktorial(n-r, &faktorialNR)
	*hasil = faktorialN / faktorialNR
}

// Prosedur untuk menghitung kombinasi
func kombinasi(n, r int, hasil *int) {
	var faktorialN, faktR, faktorialNR int
	faktorial(n, &faktorialN)
	faktorial(r, &faktR)
	faktorial(n-r, &faktorialNR)
	*hasil = faktorialN / (faktR * faktorialNR)
}
```
#### Screenshoot Source Code
![SC](https://github.com/user-attachments/assets/1cbfe556-e0ac-4c60-989b-f8103159a2d5)


#### Screenshoot Output
![Ungui1](https://github.com/user-attachments/assets/d6c9a071-d91f-42c4-89b6-182a7bfd1cee)


#### Deskripsi Program
Program ini ditulis dalam bahasa Go dan bertujuan untuk menghitung permutasi dan kombinasi dari empat bilangan bulat yang dimasukkan oleh pengguna: ( a ), ( b ), ( c ), dan ( d ). Program ini juga memeriksa syarat bahwa ( a ) harus lebih besar atau sama dengan ( c ) dan ( b ) harus lebih besar atau sama dengan ( d ) sebelum melakukan perhitungan. Jika syarat tidak terpenuhi, program akan memberikan pesan kesalahan.

#### Algoritma Program
1. Mulai
2. Deklarasi Variabel: Deklarasikan variabel integer ( a, b, c, d ).
3. Input: Minta pengguna untuk memasukkan bilangan ( a, b, c, d ).
4. Pemeriksaan Syarat:
	-Jika ( a \geq c ) dan ( b \geq d ):
	-Deklarasikan variabel untuk menyimpan hasil permutasi dan kombinasi.
	-Panggil fungsi permutasi(a, c, &permutasiAC) dan kombinasi(a, c, &kombinasiAC).
	-Cetak hasil permutasi dan kombinasi untuk ( (a, c) ).
	-Panggil fungsi permutasi(b, d, &permutasiBD) dan kombinasi(b, d, &kombinasiBD).
	-Cetak hasil permutasi dan kombinasi untuk ( (b, d) ).
	Jika tidak:
	-Cetak "Input tidak sesuai dengan syarat yang ada".
5. Fungsi faktorial(n, hasil):
	Inisialisasi hasil dengan 1.
	Untuk setiap ( i ) dari 1 hingga ( n ):
	Kalikan hasil dengan ( i ).
6. Fungsi permutasi(n, r, hasil):
	Hitung faktorial ( n ) dan ( n-r ).
	Hitung hasil permutasi dengan rumus ( nPr ) dan simpan di hasil.
7. Fungsi kombinasi(n, r, hasil):
	Hitung faktorial ( n ), ( r ), dan ( n-r ).
	Hitung hasil kombinasi dengan rumus ( nCr ) dan simpan di hasil.
8. Selesai

#### Cara Kerja
1. Input Pengguna: Program meminta pengguna untuk memasukkan empat bilangan bulat: ( a ), ( b ), ( c ), dan ( d ).
2. Pemeriksaan Syarat: Program memeriksa apakah ( a \geq c ) dan ( b \geq d ):
	-Jika syarat terpenuhi, program akan menghitung permutasi dan kombinasi untuk pasangan ( (a, c) ) dan ( (b, d) ).
	-Jika syarat tidak terpenuhi, program akan mencetak pesan kesalahan.
3. Menghitung Faktorial: Fungsi faktorial digunakan untuk menghitung faktorial dari bilangan yang diberikan.
4. Menghitung Permutasi: Fungsi permutasi menghitung permutasi menggunakan rumus: [ nPr = \frac{n!}{(n-r)!} ]
5. Menghitung Kombinasi: Fungsi kombinasi menghitung kombinasi menggunakan rumus: [ nCr = \frac{n!}{r!(n-r)!} ]
6. Menampilkan Hasil: Hasil dari permutasi dan kombinasi dicetak ke layar.

## 2. Program Penentuan Pemenang Kompetisi Pemrograman Berdasarkan Skor dan Waktu Penyelesaian
Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. 

Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.

Keterangan:

Astuti menyelesaikan 6 soal dalam waktu 287 menit, sedangkan Bertha 7 soal dalam waktu 294 menit. Karena Bertha menyelesaikan lebih banyak, maka Bertha menang. Jika keduanya menyelesaikan sama banyak, maka pemenang adalah yang menyelesaikan dengan total waktu paling kecil.
#### Source Code
```go
// Aditya Sulistiawan
// 2311102193

package main

import (
    "fmt"
)

// Struct untuk menyimpan data peserta
type Peserta struct {
    nama       string
    waktu      [8]int
    soalBenar  int
    totalWaktu int
}

// Prosedur untuk input jumlah peserta
func inputJumlahPeserta(n *int) {
    fmt.Print("Jumlah Peserta: ")
    fmt.Scan(n)
}

// Prosedur untuk input data peserta
func inputDataPeserta(nomorPeserta int, peserta *Peserta) {
    // Input nama peserta
    fmt.Printf("\nNama peserta %d: ", nomorPeserta+1)
    fmt.Scan(&peserta.nama)

    // Input waktu pengerjaan
    fmt.Print("Waktu Pengerjaan Soal (8 soal): ")
    for j := 0; j < 8; j++ {
        fmt.Scan(&peserta.waktu[j])
    }
}

// Prosedur untuk menghitung total soal yang dikerjakan dan total waktu
func hitungSkor(waktu [8]int, soal *int, totalWaktu *int) {
    *soal = 0
    *totalWaktu = 0
    for i := 0; i < 8; i++ {
        if waktu[i] <= 300 { // jika waktu pengerjaan kurang dari 300 menit, soal selesai
            *soal++
            *totalWaktu += waktu[i] // hanya tambahkan waktu soal yang selesai
        }
    }
}

// Prosedur untuk menentukan pemenang
func tentukanPemenang(peserta Peserta, pemenangSekarang *Peserta) {
    if peserta.soalBenar > pemenangSekarang.soalBenar || 
       (peserta.soalBenar == pemenangSekarang.soalBenar && peserta.totalWaktu < pemenangSekarang.totalWaktu) {
        *pemenangSekarang = peserta
    }
}

// Prosedur untuk menampilkan hasil akhir
func tampilkanHasil(pemenang Peserta) {
    fmt.Printf("\nNama pemenang: %s\n", pemenang.nama)
    fmt.Printf("Jumlah soal yang selesai: %d\n", pemenang.soalBenar)
    fmt.Printf("Total waktu yang dihabiskan: %d menit\n", pemenang.totalWaktu)
}

func main() {
    // Input jumlah peserta
    var n int
    inputJumlahPeserta(&n)

    // Inisialisasi data pemenang sementara
    pemenang := Peserta{
        soalBenar: -1,
        totalWaktu: 1000,
    }

    // Proses setiap peserta
    for i := 0; i < n; i++ {
        // Input data peserta
        var peserta Peserta
        inputDataPeserta(i, &peserta)

        // Hitung skor peserta
        hitungSkor(peserta.waktu, &peserta.soalBenar, &peserta.totalWaktu)

        // Tentukan pemenang sementara
        tentukanPemenang(peserta, &pemenang)
    }

    // Tampilkan hasil akhir
    tampilkanHasil(pemenang)
}
```
#### Screenshoot Source Code
![SC](https://github.com/user-attachments/assets/20ef0b84-bdcd-4012-b87f-b9ec74cb20e1)


#### Screenshoot Output
![Ungi2](https://github.com/user-attachments/assets/23afb799-6495-42f4-968a-750ea9cfe386)


#### Deskripsi Program
Program ini adalah aplikasi untuk menentukan pemenang dalam sebuah kompetisi pemrograman. Kompetisi ini melibatkan beberapa peserta yang masing-masing mengerjakan 8 soal. Pemenang ditentukan berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan.

#### Algoritma Program Penentuan Pemenang Kompetisi Pemrograman:
1. Mulai
2. Deklarasi variabel n (jumlah peserta), pemenang (struct Peserta)
3. Panggil prosedur inputJumlahPeserta(&n)
4. Inisialisasi pemenang dengan soalBenar = -1 dan totalWaktu = 1000
5. Untuk i = 0 sampai n-1, lakukan:
   a. Deklarasi variabel peserta (struct Peserta)
   b. Panggil prosedur inputDataPeserta(i, &peserta)
   c. Panggil prosedur hitungSkor(peserta.waktu, &peserta.soalBenar, &peserta.totalWaktu)
   d. Panggil prosedur tentukanPemenang(peserta, &pemenang)
6. Panggil prosedur tampilkanHasil(pemenang)
7. Selesai

#### Cara Kerja
1. Program meminta input jumlah peserta.
2. Untuk setiap peserta:
   - Meminta input nama peserta.
   - Meminta input waktu pengerjaan untuk 8 soal.
   - Menghitung jumlah soal yang berhasil diselesaikan (waktu <= 300 menit) dan total waktu pengerjaan.
   - Membandingkan dengan pemenang sementara dan memperbarui jika lebih baik.
3. Setelah semua peserta diproses, program menampilkan nama pemenang, jumlah soal yang diselesaikan, dan total waktu pengerjaan.


## 3. Program untuk Mencetak Deret Bilangan Berdasarkan Algoritma 3n+1

Sklena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah Van, tetapi jika ganjil maka suku berikutnya bernilal 3n+-1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilal 1. Sebagai contoh Jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah

22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1 

Untuk suku awal sampal dengan 1000000, diketahui deret selalu mencapal suku dengan nilal 1 Buat program sklena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat da dalam prosedur cetak Deret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.

prosedure cetakDeret (in n : integer )
Masukan : berupa satu bilangan integer positif yang lebih kecil dari 1000000

Keluaran : terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dan dipisahkan oleh sebuah spasi


#### Source Code
```go
//Aditya Ssulistiawannn

package main

import (
	"fmt"
)

func cetakDeret(n int) {
	// Mencetak nilai awal
	fmt.Print(n)

	// Menghitung dan mencetak deret
	for n != 1 {
		if n%2 == 0 {
			n = n / 2 // Jika genap, bagi dua
		} else {
			n = 3*n + 1 // Jika ganjil, 3n + 1
		}
		fmt.Print(" ", n) // Mencetak suku berikutnya dengan spasi
	}
}

func main() {
	var nilaiAwal int
	fmt.Print("Masukkan bilangan bulat positif (kurang dari 1000000): ")
	fmt.Scan(&nilaiAwal)

	if nilaiAwal > 0 && nilaiAwal < 1000000 {
		cetakDeret(nilaiAwal)
		fmt.Println() // Mencetak newline setelah deret
	} else {
		fmt.Println("Masukkan bilangan bulat positif yang valid.")
	}
}

```
#### Screenshoot Source Code
![SC](https://github.com/user-attachments/assets/6049ced0-586f-4b98-b0f0-17ca95154383)


#### Screenshoot Output
![Ungui3](https://github.com/user-attachments/assets/bd2149ca-939b-4c87-bfa2-8ad5f81c28d6)


#### Deskripsi Program
Program ini mengimplementasikan algoritma Collatz Conjecture (juga dikenal sebagai problema 3n + 1). Program menerima input bilangan bulat positif dari pengguna dan menghasilkan deret bilangan berdasarkan aturan Collatz sampai mencapai nilai 1.

#### Algoritma Program
Fungsi main:
1. Mulai
2. Deklarasi variabel nilaiAwal sebagai integer
3. Tampilkan "Masukkan bilangan bulat positif (kurang dari 1000000): "
4. Baca input ke nilaiAwal
5. Jika nilaiAwal > 0 DAN nilaiAwal < 1000000, maka:
   a. Panggil fungsi cetakDeret(nilaiAwal)
   b. Cetak baris baru
6. Jika tidak, tampilkan "Masukkan bilangan bulat positif yang valid."
7. Selesai

Fungsi cetakDeret(n int):
1. Cetak n (nilai awal)
2. Selama n ≠ 1, lakukan:
   a. Jika n mod 2 = 0, maka:
      n = n / 2
   b. Jika tidak, maka:
      n = 3 * n + 1
   c. Cetak " " (spasi) diikuti nilai n
3. Kembali ke fungsi pemanggil
#### Cara Kerja
1. Program meminta pengguna memasukkan bilangan bulat positif kurang dari 1.000.000.
2. Program memeriksa apakah input valid (positif dan kurang dari 1.000.000).
3. Jika input valid, program memanggil fungsi cetakDeret untuk menghasilkan dan mencetak deret Collatz.
4. Fungsi cetakDeret mencetak nilai awal, kemudian menerapkan aturan Collatz:
   - Jika bilangan genap, bagi dengan 2.
   - Jika bilangan ganjil, kalikan dengan 3 dan tambah 1.
5. Proses berlanjut sampai nilai mencapai 1.
6. Setiap nilai dalam deret dicetak dengan spasi sebagai pemisah.

### Kesimpulan
Dalam pemrograman, prosedur adalah alat yang sangat bermanfaat untuk menyusun kode secara terstruktur dan efisien. Dengan memahami cara mendeklarasikan dan menggunakan prosedur, programmer dapat meningkatkan keterbacaan dan keorganisasian kode mereka, serta memudahkan proses debugging dan pengembangan di masa mendatang. Ini menjadikannya komponen penting dalam pemrograman.

## Referensi 
[1] Muhaz,  DEKLARASI PROSEDUR (PROCEDURE DECLARATION)

[2] olang Official Documentation. (n.d.). A Tour of Go.
