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

### Definisi Prosedur<br/>
  Prosedur adalah serangkaian instruksi yang dipecah dari program yang lebih besar untuk mengurangi kompleksitas kode. Ketika prosedur dipanggil dalam program
  utama, ia akan memberikan efek langsung pada jalannya program. Suatu subprogram disebut sebagai prosedur jika:<br/>
  1. Tidak memiliki deklarasi tipe nilai yang dikembalikan, dan<br/>
  2. Tidak menggunakan kata kunci return dalam isinya.<br/>
  Prosedur mirip dengan instruksi dasar seperti penugasan (assignment) atau instruksi dari paket seperti fmt.Scan atau fmt.Print. Oleh karena itu, nama prosedur
  sebaiknya menggunakan kata kerja atau sesuatu yang menggambarkan suatu tindakan, misalnya cetak, hitungRerata, cariNilai, belok, mulai, dll.<br/>
### Deklarasi Prosedur<br/>
  ![image](https://github.com/user-attachments/assets/6d1c5b52-393a-4886-842e-2dd1b8eb2f2d)>br/>
  Penulisan deklarasi ini berada di luar blok yang dari program utama atau fungsi main() pada suatu program C, dan bisa ditulis sebelum atau setelah dari blok
  program utama tersebut.<br/>
  Contoh deklarasi prosedur mencetak n nilai pertama dari deret Fibonacci<br/>
  ![image](https://github.com/user-attachments/assets/af00ed5f-67b8-46aa-a0bf-1b51097fa66f)<br/>
### Cara Pemanggilan Prosedur
  Seperti yang telah dijelaskan, prosedur hanya akan dijalankan jika dipanggil, baik secara langsung maupun tidak langsung oleh program utama. Pemanggilan tidak
  langsung berarti prosedur tersebut dipanggil melalui subprogram lain sebagai perantara.<br/>
  Cara memanggil prosedur cukup sederhana, yaitu dengan menuliskan nama prosedur dan menyertakan parameter atau argumen yang diperlukan. Misalnya, prosedur
  cetakNFibo dipanggil dengan menuliskan namanya dan memberikan sebuah variabel atau nilai integer tertentu sebagai argumen untuk parameter n.<br/>
  Contohnya:<br/>
  ![image](https://github.com/user-attachments/assets/bd95646f-17c7-4e16-9b2e-b3dfc1c8e892)<br/>
  ![image](https://github.com/user-attachments/assets/74830e5c-c03f-46a1-b2de-182201893a22)<br/>
### Contoh Program dengan Prosedur<br/>
   Berikut ini adalah contoh penulisan prosedur pada suatu program lengkap. Buatlah sebuah program beserta prosedur yang digunakan untuk menampilkan suatu pesan
   error, warning atau informasi berdasarkan masukan dari user.<br/>
   Masukan: terdiri dari sebuah bilangan bulat flag (0 s.d. 2) dan sebuah string M.<br/>
   Keluaran: berupa string pesan M beserta jenis pesannya, yaitu error, warning atau informasi berdasarkan nilai flag 0, 1 dan 2 secara berturut-turut.<br/>
   ![image](https://github.com/user-attachments/assets/57ce1970-8232-4a64-8ab5-979580a393c4)<br/>
   Penulisan argumen pada parameter cetakPesan(pesan, bilangan) harus sesuai tipe data pada func cetakPesan (M string, flag int), yaitu string kemudian integer.<br/>
### Parameter
   Subprogram yang dipanggil dapat berinteraksi dengan pemanggilnya melalui argumen yang diberikan ke parameter yang dideklarasikan dalam subprogram. Terdapat dua
   jenis parameter berdasarkan posisinya dalam program, yaitu parameter formal dan parameter aktual.<br/>
   ![image](https://github.com/user-attachments/assets/801d26f8-a078-4d3b-8c86-52cc28f19c77)<br/>
   - Parameter Fromal<br/>
     Parameter formal adalah parameter yang ditulis pada saat deklarasi suatu subprogram, parameter ini berfungsi sebagai petunjuk bahwa argumen apa saja yang
     diperlukan pada saat pemanggilan subprogram. Sebagai contoh parameter jari_jari, tinggi pada deklarasi fungsi volumeTabung adalah parameter formal (teks
     berwarna merah). Artinya ketika memanggil volumeTabung maka kita harus mempersiapkan dua integer (berapapun nilainya) sebagai jari_jari dan tinggi.<br/>
   - Parameter Aktual<br/>
     Sedangkan parameter aktual adalah argumen yang digunakan pada bagian parameter saat pemanggilan suatu subprogram. Banyaknya argumen dan tipe data yang
     terdapat pada parameter aktual harus mengikuti parameter formal. Sebagai contoh argumen r, t, 15, 14 dan 100 pada contoh kode di atas (teks berwarna biru)
     adalah parameter aktual, yang menyatakan nilai yang kita berikan sebagai jari-jari dan tinggi.<br/>
Selain itu, parameter juga diklasifikasikan berdasarkan cara alokasi memorinya, yaitu pass by value dan pass by reference.<br/>
1. Pass by Value<br/>
Pada pass by value, nilai dari parameter aktual disalin ke variabel lokal (parameter formal) dalam subprogram. Ini berarti parameter aktual dan formal memiliki alamat memori yang berbeda. Subprogram dapat menggunakan nilai pada parameter formal untuk berbagai proses, namun tidak dapat mengembalikan informasi tersebut ke pemanggil melalui parameter aktual, karena pemanggil tidak memiliki akses ke memori subprogram. Pass by value dapat diterapkan baik pada fungsi maupun prosedur.
Dalam notasi pseudocode, semua parameter formal pada fungsi secara default menggunakan pass by value, sedangkan pada prosedur menggunakan kata kunci "in" saat mendeklarasikan parameter formal. Di bahasa pemrograman Go, sama seperti pada pseudocode fungsi, tidak ada kata kunci khusus untuk parameter formal pada fungsi dan prosedur.<br/>
2. Pass by Reference (Pointer)<br/>
Pada pass by reference, parameter formal bertindak sebagai pointer yang menyimpan alamat memori dari parameter aktual. Hal ini berarti perubahan yang terjadi pada parameter formal akan berdampak langsung pada parameter aktual. Setelah subprogram selesai dijalankan, nilai akhir dari parameter tersebut dapat diakses oleh pemanggil. Pass by reference lebih cocok digunakan pada prosedur. Dalam pseudocode, parameter pass by reference pada prosedur ditulis dengan kata kunci "in/out", sedangkan di bahasa Go menggunakan simbol asterik (*) sebelum tipe data pada parameter formal.<br/>
Catatan: Sebaiknya parameter pada fungsi menggunakan pass by value, karena fungsi dapat mengembalikan nilai ke pemanggil tanpa memberikan dampak langsung pada program. Meskipun demikian, pass by reference juga dapat digunakan jika diperlukan. Sementara itu, pass by reference lebih ideal digunakan pada prosedur, karena prosedur tidak dapat mengembalikan nilai secara langsung. Dengan pass by reference, prosedur dapat seolah-olah mengirimkan nilai ke pemanggil.<br/>

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

Kode di atas menghitung dan mencetak luas serta keliling sebuah persegi. Program ini terdiri dari dua prosedur utama, yaitu `hitungLuas` dan `hitungKeliling`. Prosedur `hitungLuas` menerima panjang sisi persegi sebagai parameter dan menghitung luasnya menggunakan rumus 
sisi × sisi, kemudian mencetak hasilnya dengan format desimal hingga dua angka di belakang koma. Prosedur `hitungKeliling` juga menerima panjang sisi persegi sebagai parameter, lalu menghitung keliling menggunakan rumus 4 × sisi, dan mencetak hasil keliling dengan format serupa. Pada bagian utama (`main`), program meminta pengguna untuk memasukkan panjang sisi persegi. Setelah itu, prosedur `hitungLuas` dan `hitungKeliling` dipanggil untuk menghitung dan menampilkan hasil perhitungan luas dan keliling persegi secara berurutan.

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

Kode di atas untuk menghitung permutasi dan kombinasi dari dua pasang bilangan integer. Program ini menggunakan tiga prosedur utama: `factorial`, `permutation`, dan `combination`. Prosedur `factorial` menghitung faktorial dari bilangan `n` dan menyimpan hasilnya dalam variabel yang diterima sebagai pointer. Prosedur `permutation` menghitung permutasi 𝑛𝑃𝑟 dengan memanfaatkan hasil faktorial dari `n` dan `n - r`. Hasil perhitungan permutasi disimpan dalam variabel hasil yang juga dikirim sebagai pointer. Prosedur `combination` menghitung kombinasi 𝑛𝐶𝑟 dengan memanfaatkan faktorial dari `n`, `r`, dan `n - r`, kemudian hasilnya juga disimpan dalam variabel hasil menggunakan pointer.<br/>
Pada bagian utama (`main`), program meminta input dari pengguna berupa empat bilangan integer: `a`, `b`, `c`, dan `d`. Program kemudian menghitung permutasi dan kombinasi dari `a` terhadap `c`, serta dari `b` terhadap `d`. Hasil perhitungan untuk kedua pasangan bilangan tersebut dicetak secara terpisah. Output pertama menampilkan hasil permutasi dan kombinasi untuk bilangan `a` dan `c`, sedangkan output kedua menampilkan hasil perhitungan untuk `b` dan `d`. Program ini memanfaatkan pointer untuk mengelola hasil perhitungan agar dapat dimodifikasi langsung di dalam prosedur.<br/>

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

Kode di atas digunakan untuk menentukan pemenang dari sebuah kompetisi pemrograman di mana peserta menyelesaikan maksimal 8 soal dalam waktu 5 jam (301 menit per soal). Program ini menggunakan prosedur `hitungSkor` untuk menghitung jumlah soal yang berhasil diselesaikan dalam waktu kurang dari 301 menit dan total waktu yang diperlukan oleh peserta.<br/>
Pada bagian utama (`main`), program menerima input berupa nama peserta dan waktu pengerjaan untuk 8 soal. Jika nama yang diinput adalah "selesai", program akan berhenti menerima data. Untuk setiap peserta, waktu pengerjaan setiap soal dimasukkan, kemudian dihitung jumlah soal yang diselesaikan dan total waktu pengerjaannya dengan memanggil prosedur `hitungSkor`. Program akan terus memperbarui informasi pemenang sementara berdasarkan siapa yang menyelesaikan lebih banyak soal. Jika dua peserta menyelesaikan jumlah soal yang sama, peserta dengan total waktu pengerjaan lebih sedikit akan menjadi pemenang.<br/>
Setelah semua peserta selesai diinput, program akan mencetak nama pemenang, jumlah soal yang diselesaikan, dan total waktu yang dibutuhkan untuk menyelesaikan soal-soal tersebut.

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

Kode di atas adalah program yang mencetak deret bilangan berdasarkan aturan dari "Collatz Conjecture" atau yang sering dikenal sebagai masalah 3n + 1. Prosedur `cetakDeret` menerima satu bilangan bulat `n` sebagai input dan mencetak deretnya. Jika `n` adalah bilangan genap, maka `n` dibagi 2. Jika `n` adalah bilangan ganjil, maka dihitung dengan rumus 3𝑛+1. Proses ini berulang hingga nilai `n` menjadi 1, yang juga akan dicetak sebagai elemen terakhir dari deret.<br/>
Pada bagian utama (`main`), program meminta pengguna memasukkan sebuah bilangan awal `n`. Setelah input diterima, program memanggil prosedur `cetakDeret` untuk mencetak deret bilangan yang dihasilkan hingga mencapai angka 1. Setiap angka dalam deret dicetak dengan spasi sebagai pemisah.

### KESIMPULAN
### REFERENSI
