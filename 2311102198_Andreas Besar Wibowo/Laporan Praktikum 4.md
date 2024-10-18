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
  Andreas Besar Wibowo / 2311102198<br>
  S1 IF11-05
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
## I. Dasar Teori
### Definisi Procedure
Prosedur merupakan sub program yang melakukan satu atau lebih proses tanpa mengembalikan nilai dari hasil dari proses yang dijalankan. Namun, pada prosedur diperbolehkan mencetak hasil luaran yang dihasilkan dari proses yang dijalankan. Contohnya, mencetak hasil dari proses aritmatika[2]. Prosedur dapat dianggap sebagai pembentukan suatu instruksi baruyang dibuat untuk mempermudah pemahaman algoritma program yang lebih besar. Kedudukannya sama seperti instruksi dasar yang sudah ada sebelumnya (assignment) dan/atau instruksi yang berasal dari paket (fmt), seperti fmt.Scandan fmt.Print. Karena itu selalu pilih nama prosedur yang berbentuk kata kerja atau sesuatu yang merepresentasikan proses sebagai nama dari prosedur[1].

### Parameter
Antar pemanggil subprogram dan subprogramnya dapat saling berkomunikasi melalui argumen yang diberikan melalui parameter yang dideklarasikan pada subprogram. Ada dua jenis pengiriman informasi: 
- Pengiriman nilai ke subprogram. Nilai akan disalin dari pemanggil ke suatu variabel lokal di lokasi subprogram. Subprogram dapat menggunakan nilai tersebut untuk apapun, tetapi tidak dapat mengembalikan informasinya ke pemanggil karena pemanggil tidak dapat mengakses memori yang digunakan subprogram.
- Pengiriman alamat dari suatu variabel yang dimiliki pemanggil ke subprogram. Subprogram mengambil data langsung dari lokasi tersebut. Karena lokasinya diarea pemanggil, setiap perubahan data dilokasi tersebut yang dilakukan oleh subprogram, nilai terakhirnya akan dapat diketahui oleh si pemanggil setelah subprogram tersebut selesai eksekusi[1].

Ketika sebuah prosedur dipanggil, maka pada dasarnya dapat melakukan proses pertukaran data antara program utama dan sub rogram. Proses pertukaran ini dapat dilakukan dengan penggunaan parameter. Pada prosedur terdapat parameter yaitu sebagai berikut:
1. Parameter Aktual, Parameter aktual merupakan parameter yang disertakan pada saat prosedur dipanggil untuk dijalankan pada program utama, atau dapat sering jenis ini sebagai argumen.

2. Parameter Formal, Parameter formal merupakan paramater yang dituliskan pada saat melakukan pendefinisian sebuah prosedur. Parameter ini ditulis pada bagian kode sub program untuk menerima nilai dari parameter aktual pada program utama. Terdapat tiga jenis parameter formal yang dapat digunakan pada prosedur, diantaranya:
- Parameter masukkan (input), Parameter masukan merupakan parameter yang menerima nilai dari parameter aktual untuk dilakukan pemrosesan pada prosedur.
- Parameter keluaran (output), Parameter keluaran merupakan parameter yang digunakan untuk menampung hasil dari pemrosesan dari sebuah prosedur, selanjutnya hasilnya dikirimkan ke program utama program.
- Parameter masukan dan keluaran (input/ output), Parameter yang menerima nilai dari parameter aktual untuk diproses dalam prosedur kemudian nilai yang dihasilkan dari pemrosesan tersebut ditampung ke dalam parameter output. parameter output yang dihasilkan oleh prosedur dapat dimanfaatkan pada instruksi kode berikutnya[2].

## II. GUIDED
**1. Buatlah sebuah program beserta prosedur yang digunakan untuk menghitung nilai faktorial dan permutasi.**

*Masukan : Terdiri dari dua buah bilangan positif a dan b*

*Keluaran : Berupa sebuah bilangan bulat yang menyatakan nilai a permutasi b apabila >= b atau b permutasi a untuk kemungkinan yang lain*

#### Source Code
```go
package main

import "fmt"

func main() {
    var a, b int // mendeklarasikan 2 variabel 
    fmt.Scan(&a, &b)  // untuk input dari pengguna
    if a >= b { // jika a >= b, akan ada permutasi
        permutasi(a, b) // memanggil prosedur permutasi
    } else { // jika b >= a, akan ada permutasi
        permutasi(b, a) // memanggil prosedur permutasi
    }
}

// prosedur untuk menghitung faktorial 
func faktorial(n int, hasil *int) {
    *hasil = 1
    for i := 1; i <= n; i++ {
        *hasil *= i
    }
}

// prosedur untuk menghitung dan mencetak permutasi
func permutasi(n, r int) {
    var faktorialN, faktorialNR int

    // menghitung faktorial n
    faktorial(n, &faktorialN)

    // menghitung faktorial (n-r)
    faktorial(n-r, &faktorialNR)

    // menghitung permutasi 
    hasil := faktorialN / faktorialNR
    fmt.Println(hasil)
}
```
#### Screenshoot Source Code
![Guided1 carbon](https://github.com/user-attachments/assets/405ba559-f38c-48aa-a0a6-f2630a7030e9)

#### Screenshoot Output
![Guided1 go](https://github.com/user-attachments/assets/bfb33dd4-877e-4934-921f-c8383684d8be)

#### Deskripsi Program
Program diatas merupakan program untuk menghitung dan mencetak hasil dari permutasi dari dua bilangan yang diinputkan oleh user. Program meminta user untuk mengisi 2 bilangan dan menentukan bilangan mana yang lebih besar lalu menghitung permutasi dengan rumus yang sudah tersedia.

Algortima dari program tersebut sebagai berikut :
1. Deklarasikan dua variable yaitu 'a' dan 'b'.
2. Ambil input dari user untuk mengisi variable.
3. Jika "a lebih besar dari b", maka permutasi (a,b).
4. Jika "b lebih besar dari a", maka permutasi (b,a).
5. Terdapat prosedur dalam 'permutasi' : hitung faktorial n, faktorial nr, dan menghitung permutasi.
6. Mencetak hasil permutasi.

Cara kerja dari program tersebut yaitu program mendeklarasikan 2 variabel untuk menyimpan inputan dari user. Setelah itu, program memeriksa dari 2 bilangan itu mana yang lebih besar. Dengan memanggil prosedur 'permutasi', programakan menghitung faktorial dari nilai yang lebih besar dan selisih dari 2 bilangan. Hasilnya akan muncul di output. Prosedur 'faktorial' akan digunakan untuk menghitung faktorial yang digunakan juga dalam menghitung permutasi.

**2. Buatlah sebuah program beserta prosedur yang digunakan untuk menghitung luas dan keliling persegi.**

*Masukan : Terdiri dari bilangan yang menjelaskan sisi persegi*

*Keluaran : Berupa sebuah bilangan bulat yang menyatakan luas dan keliling persegi sesuai sisi yang diinputkan user*

#### Source Code
```go
package main

import (
	"fmt"
)

// prosedur untuk menghitung luas persegi
func hitungLuas(sisi float64, luas *float64) {
	*luas = sisi * sisi
}

// prosedur untuk menghitung keliling persegi
func hitungKeliling(sisi float64, keliling *float64) {
	*keliling = 4 * sisi
}

func main() {
	var sisi, luas, keliling float64

	// untuk inputan dari user
	fmt.Print("Sisi Persegi: ")
	fmt.Scan(&sisi)

	// menghitung luas persegi menggunakan prosedur hitungLuas
	hitungLuas(sisi, &luas)
	// menghitung keliling persegi menggunakan prosedur hitungKeliling
	hitungKeliling(sisi, &keliling)

	// menampilkan hasil dari perhitungan luas dan keliling persegi
	fmt.Printf("Luas persegi: %.2f\n", luas)
	fmt.Printf("Keliling persegi: %.2f\n", keliling)
}
```
#### Screenshoot Source Code
![Guided2 carbon](https://github.com/user-attachments/assets/21aff99e-48a5-4581-9fa6-8f986fbd8008)

#### Screenshoot Output
![Guided2 go](https://github.com/user-attachments/assets/b6e06763-03a2-46c7-afc8-1d15f4017f00)

#### Deskripsi Program
Program diatas merupakan program untuk menghitug luas dan keliling persegi dari sisi yang diinputkan user. Program meminta user untuk memasukkan sisi dari persegi, lalu program akan menghitungnya dengan menggunakan prosedur terpisah.

Algortima dari program tersebut sebagai berikut :
1. Deklarasikan variabel 'sisi', 'luas', dan 'keliling' dengan tipe float64.
2. Minta input 'sisi' dari user.
3. Panggil prosedur 'hitungLuas' untuk menghitung luas persegi.
4. Panggil prosedur 'hitungKeliling' untuk menghitung keliling persegi.
5. Tampilkan hasil perhitungan luas dan keliling persegi.

Cara kerja dari program tersebut yaitu program mendeklarasikan variabel untuk menyimpan sisi, luas, dan keliling. Setelah user memasukkan sisinya, program akan memanggil prosedur 'hitungLuas' dan 'hitungKeliling' untuk menghitung luas dan keliling. Prosedur tersebut menggunakan pointer. Hasil dari perhitungan akan ditampilkan dalam output.

## III. UNGUIDED
**1. Minggu Ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kullah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, Iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakan kaltan membantu Jonas? (tidak tentunya ya :p)**

*Masukan : Terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a >= c dan b >= d*

*Keluaran : Terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF11-05

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
![Unguided1 carbon](https://github.com/user-attachments/assets/005060e3-1fae-474e-99bf-8f8ce268133e)

#### Screenshoot Output
![Unguided1 go](https://github.com/user-attachments/assets/4571feb2-df8e-4243-9507-34b7010c4641)

#### Deskripsi Program
Program diatas merupakan program untuk menghitung permutasi dan kombinasi yang nilainya diambil dari inputan user. Program meminta user untuk mengisi variabel 'a', 'b', 'c', dan 'd' dengan syarat yang sudah disediakan. Jika syarat terpenuhi program akan berjalan dengan pasangan '(a,c)' dan '(b,d)'. Jika syarat tidak terpenuhi, program akan berhenti.

Algortima dari program tersebut sebagai berikut :
1. Deklarasikan 4 variabel 'a', 'b', 'c', dan 'd'.
2. Minta user untuk menginputkan 4 variabel tersebut.
3. Program harus memeriksa syarat.
4. Panggil prosedur `permutasi(a, c, &permutasiAC)` untuk menghitung permutasi `(a, c)`.
5. Panggil prosedur `kombinasi(a, c, &kombinasiAC)` untuk menghitung kombinasi `(a, c)`.
6. Tampilkan hasil permutasi dan kombinasi AC.
7. Panggil prosedur `permutasi(b, d, &permutasiBD)` untuk menghitung permutasi `(b, d)`.
8. Panggil prosedur `kombinasi(b, d, &kombinasiBD)` untuk menghitung kombinasi `(b, d)`.
9. Tampilkan hasil permutasi dan kombinasi BD.

Cara kerja dari program tersebut yaitu program mendeklarasikan variable untuk menyimpan input dari pengguna. Setelah itu, program akan memeriksa bilangan dari syarat yang sudah tersedia. Jika syarat terpenuhi, program akan memanggil prosedur 'permutasi' untuk menghitung permutasi dan prosedur 'kombinasi' untuk menghitung kombinasi. Hasil dari perhitungannya akan di tampilkan dalam output.

**2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 Jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya.**

**Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.**

```go
prosedure hitungSkor (in/out soal : integer)
```
**Setiap barts masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 Integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesalkan soal. Jika tidak berhasil atau tidak mengirimkan Jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit).**

**Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesalkan, dan nilal yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.**

*Keterangan : Astuti menyelesaikan 6 soal dalam waktu 287 menit, sedangkan Bertha 7 soal dalam waktu 294 menit. Karena Bertha menyelesaikan lebih banyak, maka Bertha menang. jike keduanya menyelesaikan sama banyak, maka pemenang adalah yang menyelesaikan dengan total waktu paling kecil.*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF11-05

package main

import (
	"fmt"
)

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

func main() {
	var n int
	fmt.Print("Jumlah Peserta: ")
	fmt.Scan(&n)

	var pemenang string
	var maximalSoal int
	var minimalWaktu int = 1000 // nilai awal 
	var namaPeserta string
	var waktu [8]int

	for i := 0; i < n; i++ {
		fmt.Printf("\nNama peserta %d: ", i+1)
		fmt.Scan(&namaPeserta)

		// Input waktu pengerjaan soal
		fmt.Print("Waktu Pengerjaan Soal (8 soal): ")
		for j := 0; j < 8; j++ {
			fmt.Scan(&waktu[j]) // input waktu pengerjaan untuk setiap soal
		}

		// Hitung soal yang selesai dan total waktu yang digunakan
		var soal int
		var totalWaktu int
		hitungSkor(waktu, &soal, &totalWaktu)

		// pemenang berdasarkan jumlah soal yang selesai dan waktu tersingkat
		if soal > maximalSoal || (soal == maximalSoal && totalWaktu < minimalWaktu) {
			pemenang = namaPeserta
			maximalSoal = soal
			minimalWaktu = totalWaktu
		}
	}

	// Output hasil akhir
	fmt.Printf("\nNama pemenang: %s\n", pemenang)
	fmt.Printf("Jumlah soal yang selesai: %d\n", maximalSoal)
	fmt.Printf("Total waktu yang dihabiskan: %d menit\n", minimalWaktu)
}
```
#### Screenshoot Source Code
![Unguided2 carbon](https://github.com/user-attachments/assets/9a67d591-0f88-4325-98c0-079da04d1573)

#### Screenshoot Output
![Unguided2 go](https://github.com/user-attachments/assets/62d55bab-1ca9-4aee-9fbb-cca2c125cd5d)

#### Deskripsi Program
Program diatas merupakan program untuk menghitung jumlah soal yang berhasil di selesaikan dalam waktu tertentu. Program meminta user untuk memasukkan jumlah peserta, lalu mencatat waktu yang di habiskan untuk mengerjakan 8 soal. Setelah semuanya diinput, pada akhir akan mencetak nama pemenang dengan jumlah soal.

Algortima dari program tersebut sebagai berikut :
1. Deklarasikan variabel 'n' untuk jumlah peserta.
2. Minta input dari pengguna.
3. Deklarasikan variabel untuk menyimpan nama pemenang, maksimal soal, dan minimal waktu.
4. Setiap peserta diinputkan nama peserta, waktu pengerjaan, dan jumlah soal yang diselesaikan dan total waktu dengan prosedur 'hitungSkor'.
5. Tentukan nama pemenang berdasarkan jumlah soal dan waktu tersingkat.
6. Tampilkan nama pemenang, jumlah soal, dan total waktu yang dihabiskan.

Cara kerja dari program tersebut yaitu mendeklarasikan variabel untuk menyimpan data peserta. Setelah menginputkan jumlah peserta, program meminta nama dan waktu dengan prosedur 'hitungSkor' digunakan untuk menghitung jumlah soal dan total waktu. Program membandingkan hasil untuk menentukan pemenang berdasarkan persyaratan. Hasil akhir ditampilkan di layar.

**3. Sklena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah Van, tetapi jika ganjil maka suku berikutnya bernilal 3n+-1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilal 1. Sebagai contoh Jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah**

```go
22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1 
```
**Untuk suku awal sampal dengan 1000000, diketahui deret selalu mencapal suku dengan nilal 1
Buat program sklena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat da dalam prosedur cetak Deret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.**

```go
prosedure cetakDeret (in n : integer )
```

*Masukan : berupa satu bilangan integer positif yang lebih kecil dari 1000000*

*Keluaran : terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dan dipisahkan oleh sebuah spasi*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF11-05

package main

import (
	"fmt"
)

// Prosedur untuk deret bilangan
func cetakDeret(n int) {
	fmt.Print("Suku dan deret: ")
	for n != 1 { // Deret akan berhenti jika nilai n = 1
		fmt.Printf("%d ", n) // cetak deret diatur oleh spasi 
		if n%2 == 0 {        // Jika n = genap
			n = n / 2
		} else {             // Jika n = ganjil
			n = 3*n + 1
		}
	}
	fmt.Println(1) // untuk suku terakhir
}

func main() {
	var n int
	fmt.Print("Bilangan: ")
	fmt.Scan(&n)

	// Validasi input dari user
	if n <= 0 || n >= 1000000 {
		fmt.Println("Nilai harus positif dan kurang dari 1000000.")
	} else {

		// untuk cetak deret
		cetakDeret(n)
	}
}
```
#### Screenshoot Source Code
![Unguided3 carbon](https://github.com/user-attachments/assets/7d981896-110c-4d64-bce1-f626e33377f8)

#### Screenshoot Output
![Unguided3 go](https://github.com/user-attachments/assets/23edaded-6d17-4714-905a-e7fd6f76b25a)

#### Deskripsi Program
Program diatas merupakan program untuk mencetak deret dengan aturan jika bilangan genap, maka akan dibagi 2, dan jika ganjil, maka akan dikalikan 3 dan ditambahkan 1. Program ini akan terus berproses hingga mencapai angka 1. Program juga perlu mevalidasi input dan memastikan bilangan yang dimasukkan sesuai dengan ketentuan.

Algortima dari program tersebut sebagai berikut :
1. Deklarasikan variabel 'n' untuk menyimpan bilangan input.
2. Minta inputan user untuk bilangan 'n'.
3. Validasi input bahwa 'n' adalah bilangan positif dan kurang dari 1.000.000.
4. Di dalam prosedurnya menggunakan loop untuk mencetak nilai 'n' sampai 1.

Cara kerja dari program ini yaitu mendeklarasikan variabel untuk menyimpan input dari user. Setelah itu, program mevalidasi untuk memastikan bilangan tersebut sesuai dengan syarat yang ada. Jika "iya" program akan memanggil prosedur 'cetakDeret' yang mencetak nilai 'n' hingga bernilai 1. Hasilnya akan dicetak dalam satu baris dan dipisahkan oleh spasi.

## Referensi
[1] Course Hero. Notasi algoritma: Penulisan dalam bahasa Go procedure subparams kamus deklarasi. Diakses dari https://www.coursehero.com/file/p5u9347q/Notasi-algoritma-Penulisan-dalam-bahasa-Go-procedure-subparams-kamus-deklarasi/

[2] UHW Perbanas. Modul Digital Mata Kuliah: Algoritma dan Pemrograman. Diakses dari www.perbanas.ac.id

