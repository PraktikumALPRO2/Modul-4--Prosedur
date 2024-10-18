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

#### Source Code
```go
package main

import (
	"fmt"
)

// Prosedur untuk menghitung dan menampilkan faktorial
func factorial(n int) {
	if n == 0 {
		fmt.Printf("Faktorial dari %d adalah: 1\n", n)
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Printf("Faktorial dari %d adalah: %d\n", n, result)
}

// Prosedur untuk menghitung dan menampilkan permutasi P(n, r)
func permutation(n, r int) {
	factN := 1
	for i := 1; i <= n; i++ {
		factN *= i
	}
	factNR := 1
	for i := 1; i <= (n - r); i++ {
		factNR *= i
	}
	p := factN / factNR
	fmt.Printf("Permutasi P(%d, %d) adalah: %d\n", n, r, p)
}

// Prosedur untuk menghitung dan menampilkan kombinasi C(n, r)
func combination(n, r int) {
	factN := 1
	for i := 1; i <= n; i++ {
		factN *= i
	}
	factR := 1
	for i := 1; i <= r; i++ {
		factR *= i
	}
	factNR := 1
	for i := 1; i <= (n - r); i++ {
		factNR *= i
	}
	c := factN / (factR * factNR)
	fmt.Printf("Kombinasi C(%d, %d) adalah: %d\n", n, r, c)
}

func main() {
	// Input
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan a, b, c, d: ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// Hasil permutasi dan kombinasi untuk a dan c
	permutation(a, c)
	combination(a, c)

	// Hasil permutasi dan kombinasi untuk b dan d
	permutation(b, d)
	combination(b, d)
}

```
#### Screenshoot Source Code
![Screenshot 2024-10-18 162708](https://github.com/user-attachments/assets/da362146-d0da-4304-8a43-c62990bdb543)

#### Screenshoot Output
![Screenshot 2024-10-18 162714](https://github.com/user-attachments/assets/8746109d-b1bc-4fa0-86b9-d3668edae206)

#### Deskripsi Program
Program ini menghitung dan menampilkan faktorial, permutasi, dan kombinasi berdasarkan input empat bilangan bulat yang dimasukkan oleh pengguna. Setelah meminta pengguna untuk memasukkan bilangan a, b, c, dan d, program memanggil prosedur untuk menghitung permutasi P(a, c) dan C(a, c), serta P(b, d) dan C(b, d). Hasil dari setiap perhitungan ditampilkan langsung ke layar. Program ini dirancang untuk memberikan informasi yang jelas dan cepat tentang faktorial, permutasi, dan kombinasi dari bilangan yang diberikan.

#### Algoritma Program
1. Input:
- Minta pengguna untuk memasukkan empat bilangan bulat: a, b, c, dan d.
2. Hitung dan Tampilkan Permutasi dan Kombinasi:
- Panggil prosedur permutation dengan argumen a dan c untuk menghitung permutasi P(a, c) dan tampilkan hasilnya.
- Panggil prosedur combination dengan argumen a dan c untuk menghitung kombinasi C(a, c) dan tampilkan hasilnya.
- Panggil prosedur permutation dengan argumen b dan d untuk menghitung permutasi P(b, d) dan tampilkan hasilnya.
- Panggil prosedur combination dengan argumen b dan d untuk menghitung kombinasi C(b, d) dan tampilkan hasilnya.
3. Selesai:
- Program selesai setelah semua hasil perhitungan ditampilkan.

#### Cara Kerja
1. Input: Pengguna diminta memasukkan empat bilangan bulat: a, b, c, dan d.
2. Hitung: Program menghitung dan menampilkan:
- Permutasi P(a, c)
- Kombinasi C(a, c)
- Permutasi P(b, d)
- Kombinasi C(b, d)
3. Output: Semua hasil perhitungan ditampilkan di layar.

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
#### Screenshoot Source Code
![Screenshot 2024-10-18 171334](https://github.com/user-attachments/assets/593b23ce-b6d5-42ed-86ca-d007a15747e4)

#### Screenshoot Output
![Screenshot 2024-10-18 171340](https://github.com/user-attachments/assets/a39e3dba-3718-460e-a1f0-bc473974674a)

#### Deskripsi Program
Program ini menentukan pemenang kompetisi pemrograman berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan. Setiap peserta diberi 8 soal, dan jika waktu pengerjaan soal lebih dari 301 menit, soal tersebut dianggap tidak terselesaikan. Program membaca input nama peserta beserta waktu pengerjaan soal, lalu menghitung jumlah soal yang diselesaikan dan total waktu yang diperlukan. Pemenang ditentukan berdasarkan siapa yang menyelesaikan soal terbanyak, dan jika sama, yang menyelesaikan dengan waktu paling sedikit menang.

#### Algoritma Program Penentuan Pemenang Kompetisi Pemrograman:
1. Mulai Program
2. Inisialisasi variabel:
   - `pemenangNama` untuk menyimpan nama pemenang.
   - `maxSoalDiselesaikan` untuk menyimpan jumlah soal maksimal yang berhasil diselesaikan oleh peserta.
   - `minWaktu` untuk menyimpan waktu total minimal peserta dalam menyelesaikan soal.
3. Ulangi (loop) proses input:
   - Minta input dari pengguna berupa nama peserta diikuti oleh 8 waktu pengerjaan soal (dalam menit).
   - Jika input adalah "Selesai", keluar dari loop.
   - Pisahkan input menjadi `nama peserta` dan `waktu pengerjaan soal` dalam array.
4. Proses tiap peserta:
   - Inisialisasi variabel `totalSoal` dan `totalWaktu` untuk menyimpan jumlah soal yang diselesaikan dan waktu total pengerjaan peserta.
   - Loop 8 soal:
     - Jika waktu pengerjaan soal kurang dari 301 menit, tambahkan soal ke `totalSoal` dan tambahkan waktu ke `totalWaktu`.
5. Bandingkan hasil peserta dengan pemenang sementara:
   - Jika jumlah soal yang diselesaikan peserta lebih banyak daripada `maxSoalDiselesaikan`, peserta ini menjadi pemenang sementara.
   - Jika jumlah soal sama, bandingkan waktu total. Peserta dengan waktu lebih sedikit menjadi pemenang.
6. Ulangi langkah 3-5 untuk setiap peserta hingga semua peserta telah dimasukkan.
7. Tampilkan pemenang:
   - Cetak nama pemenang, jumlah soal yang diselesaikan, dan waktu total.
8. Selesai Program.


#### Cara Kerja
1. Program dimulai dengan menginisialisasi variabel untuk menyimpan nama pemenang, jumlah soal maksimal yang diselesaikan, dan waktu total minimal.
2. Program kemudian meminta input pengguna dalam bentuk satu baris yang berisi nama peserta diikuti dengan waktu pengerjaan 8 soal. Jika input adalah "Selesai", program berhenti meminta input.
3. Untuk setiap peserta, program membaca nama dan waktu pengerjaan soal. Waktu pengerjaan setiap soal yang lebih dari atau sama dengan 301 menit dianggap tidak terselesaikan. Jika waktu pengerjaan kurang dari 301 menit, soal tersebut dihitung sebagai soal yang diselesaikan, dan waktu pengerjaannya ditambahkan ke total waktu.
4. Setelah menghitung jumlah soal yang diselesaikan dan total waktu untuk setiap peserta, program membandingkan hasil peserta tersebut dengan hasil pemenang sementara:
- Jika peserta baru menyelesaikan lebih banyak soal, maka peserta baru ini menjadi pemenang sementara.
- Jika jumlah soal sama dengan pemenang sementara, peserta dengan waktu total lebih sedikit menjadi pemenang sementara.
5. Proses ini berulang sampai semua peserta dimasukkan.
6. Setelah semua peserta diinput, program menampilkan nama pemenang, jumlah soal yang diselesaikan, dan waktu total yang dihabiskan.
7. Program selesai setelah menampilkan hasil pemenang.

## 3. Program untuk Mencetak Deret Bilangan Berdasarkan Algoritma 3n+1

#### Source Code
```go
package main

import (
	"fmt"
)

func cetakDeret(n int) {
	// Selama n bukan 1, teruskan cetak deret
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2 // Jika n genap
		} else {
			n = 3*n + 1 // Jika n ganjil
		}
	}
	// Cetak 1 sebagai elemen terakhir
	fmt.Print(1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai awal: ")
	fmt.Scan(&n)

	// Panggil fungsi untuk mencetak deret
	cetakDeret(n)
}


```
#### Screenshoot Source Code
![Screenshot 2024-10-18 172458](https://github.com/user-attachments/assets/f282917e-7d64-4018-943e-9a4602510e13)

#### Screenshoot Output
![Screenshot 2024-10-18 172502](https://github.com/user-attachments/assets/5b631b76-0a99-4608-898f-6cae76bec825)

#### Deskripsi Program
Program ini dirancang untuk mencetak deret bilangan berdasarkan aturan algoritma 3n+1, yang juga dikenal sebagai Collatz Conjecture. Program dimulai dengan sebuah bilangan bulat positif sebagai input. Jika bilangan tersebut genap, maka akan dibagi dua; namun jika ganjil, akan dihitung menggunakan rumus tiga kali bilangan ditambah satu. Proses ini terus berulang hingga bilangan mencapai satu, yang menjadi bilangan terakhir dalam deret. Program ini menggunakan sebuah prosedur untuk melakukan perhitungan dan mencetak hasilnya dalam satu baris dengan setiap elemen dipisahkan oleh spasi.

#### Algoritma Program
1. Mulai
2. Minta pengguna untuk memasukkan bilangan bulat positif n.
3. Cetak nilai n.
4. Lakukan perulangan selama nilai n belum sama dengan satu:
- Jika n adalah bilangan genap, bagi nilai n dengan dua.
- Jika n adalah bilangan ganjil, ubah nilai n menjadi tiga kali n ditambah satu.
- Cetak nilai n setelah dilakukan perhitungan.
5. Akhiri perulangan ketika nilai n menjadi satu.
6. Cetak nilai satu sebagai elemen terakhir deret.
7. Selesai

#### Cara Kerja
1. Input Pengguna:
- Program dimulai dengan meminta pengguna memasukkan sebuah bilangan bulat positif. Nilai ini disimpan dalam variabel n, yang akan menjadi titik awal deret.
2. Fungsi Perhitungan Deret:
- Program menggunakan fungsi cetakDeret(n) untuk melakukan perhitungan deret. Setiap kali fungsi ini dipanggil, ia akan mencetak bilangan awal, kemudian menghitung elemen-elemen berikutnya dalam deret sesuai aturan berikut:
  - Jika bilangan genap, maka bilangan tersebut dibagi dua.
  - Jika bilangan ganjil, maka dihitung dengan rumus tiga kali bilangan ditambah satu.
- Hasil perhitungan dicetak pada layar, dengan elemen-elemen deret dipisahkan oleh spasi.
3. Pengulangan Sampai Nilai Satu:
- Perhitungan berulang terus menerus hingga nilai n mencapai satu. Ini sesuai dengan aturan Collatz Conjecture, di mana setiap bilangan pada akhirnya akan turun menjadi satu.
- Setiap nilai yang dihasilkan dari perhitungan (baik hasil pembagian atau penambahan) dicetak dalam satu baris hingga deret selesai.
4. Akhir Program:
- Ketika nilai n sudah menjadi satu, program mencetak "1" sebagai elemen terakhir dan berhenti.

### Kesimpulan
Penggunaan prosedur dalam pemrograman adalah alat yang sangat berguna untuk menyusun kode secara terstruktur dan efisien. Dengan memahami cara mendeklarasikan dan menggunakan prosedur, programmer dapat meningkatkan keterbacaan dan keorganisasian kode mereka, serta memudahkan proses debugging dan pengembangan di masa mendatang. Prosedur memberikan cara untuk membagi program menjadi bagian yang lebih kecil dan lebih mudah dikelola, menjadikannya komponen penting dalam pemrograman modern.

## Referensi 
[1] Go.dev. (n.d.). Introduction to Go functions. Go Documentation.

[2] Zakhour, M., et al. (2010). The Java programming language. Addison-Wesley.
