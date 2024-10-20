<h2 align="center">LAPORAN PAKTIKUM ALGORITMA DAN PEMROGRAMAN 2</h2>
<h2 align="center">MODUL 3</h2>
<h2 align="center">FUNGSI</h2>

<p align="center"><img src=https://github.com/user-attachments/assets/37e9c953-078b-4ef4-97e7-a5d25344e50b alt="Logo" width="300"/><p>

<p align="center">Disusun Oleh : </p>
<p align="center">Irshad Benaya Fardeca / 2311102199</p>
<p align="center">IF-11-05</p>
<br></br>
<p align="center">Dosen Pengampu : </p>
<p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>
<br></br>
<h3 align="center">PROGRAM STUDI S1 TEKNIK INFORMATIKA</h3>
<h3 align="center">FAKULTAS INFORMATIKA</h3>
<h3 align="center">TELKOM UNIVERSITY PURWOKERTO</h3>
<h3 align="center">2024</p>

<br></br>

#### I. DASAR TEORI
#### Procedure
Prosedure dalam Golang adalah blok kode yang dapat dipanggil berulang kali untuk melakukan tugas tertentu. 
Ciri-ciri procedure:
-Membentuk suatu instruksi baru
-Nama fungsi membentuk kata kerja
-Tidak memiliki tipe data Tidak terdapat "return"
-Tidak dapat menggunakan parameter secara langsung

Perbedaan parameter aktual dengan parameter formal
1. Parameter aktual Merupakan variabel atau eskpresi dalam parameter yang digunakan untuk memanggil sebuah sub program
Contoh:
```go
func volumeTabung(jari_jari, tinggi int){
...
}
```
jari_jari pada deklarasi func volume tabung merupakan parameter formal

2. Parameter formal Merupakan variabel yang ada didalam parameter list pada sebuah program
Contoh:
```go
func main(){
...
volumeTabung(14, 100)
}
```
14 dan 100 adalah parameter aktual

<br></br>

#### II. GUIDED
##### 1. Buatlah sebuah program beserta fungsi yang digunakan untuk menghitung nilai faktorial dan permutasi. Masukan terdiri dari dua buah bilangan positif a dan b. Keluaran berupa sebuah bilangan bulat yang menyatakan nilai a permutasi b apabila a≥ b atau b pemutasi a untuk kemungkinan yang lain.
##### Source code
```go
package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var fak1, fak2 int
	faktorial(n, &fak1)
	faktorial(n-r, &fak2)
	*hasil = fak1 / fak2
}

func main() {
	var a, b int
	var simpan int
	fmt.Scan(&a, &b)
	if a >= b {
		permutasi(a, b, &simpan)
		fmt.Print(simpan)
	} else {
		permutasi(b, a, &simpan)
		fmt.Print(simpan)
	}
}
```
##### Screenshoot Output
![Screenshot 2024-10-20 213504](https://github.com/user-attachments/assets/eb0a3265-2ae4-40ef-99dc-579b1ef494f1)

##### Deskripsi Program
Program ini digunakan untuk menghitung permutasi dari dua bilangan bulat yang diberikan oleh user. Setelah menerima input, program akan secara otomatis menentukan urutan perhitungan permutasi yang paling efisien. Proses perhitungan dilakukan melalui fungsi faktorial yang menghitung faktorial dari suatu bilangan dan fungsi permutasi yang menerapkan rumus permutasi dengan memanfaatkan hasil faktorial. Hasil akhir berupa jumlah permutasi yang mungkin dari kedua bilangan tersebut.

##### 2. Program menghitung luas dan keliling persegi
##### Source code
```go
package main

import "fmt"

func Luas(s int) {
	var luas int
	luas = s * s
	fmt.Print(luas)
}
func Keliling(s int) {
	var keliling int
	keliling = s * 4
	fmt.Print(keliling)
}
func main() {
	var s int
	fmt.Scan(&s)
	fmt.Print("Luas Persegi : ")
	Luas(s)
	fmt.Print("\nKeliling Persegi : ")
	Keliling(s)
}
```
##### Screenshoot Output
![Screenshot 2024-10-20 204218](https://github.com/user-attachments/assets/a8169435-c932-4a29-9265-193a0e745c44)

##### Deskripsi Program
Program ini digunakan untuk menghitung luas dan keliling sebuah persegi. User akan diminta untuk memasukkan panjang sisi persegi, kemudian program akan menjalankan fungsi Luas dan Keliling untuk menghitung nilai masing-masing. Fungsi Luas akan mengalikan sisi dengan dirinya sendiri untuk mendapatkan luas, sedangkan fungsi Keliling akan mengalikan sisi dengan 4 untuk mendapatkan keliling.

<br></br>


#### III. UNGUIDED

##### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat ac dan b >= d Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n >= r) dapat dihitung dengan menggunakan persamaan berikut!
![Screenshot 2024-10-20 205855](https://github.com/user-attachments/assets/bebedffd-292f-445e-8750-d721eb3f4ca2)

```go
package main

import (
	"fmt"
)

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var fak1, fak2 int
	faktorial(n, &fak1)
	faktorial(n-r, &fak2)
	*hasil = fak1 / fak2
}

func kombinasi(n, r int, hasil *int) {
	var fak1, fak3, fak2 int
	faktorial(n, &fak1)
	faktorial(r, &fak3)
	faktorial(n-r, &fak2)
	*hasil = fak1 / (fak3 * fak2)
}

func main() {
	var a, b, c, d, simpan int
	fmt.Scanln(&a, &b, &c, &d)
	if a >= c && b >= d {
		permutasi(a, c, &simpan)
		fmt.Println(simpan)
		kombinasi(a, c, &simpan)
		fmt.Println(simpan)
		permutasi(b, d, &simpan)
		fmt.Println(simpan)
		kombinasi(b, d, &simpan)
		fmt.Println(simpan)
	} else {
		fmt.Println("Tidak memenuhi kondisi")
	}
}
```
##### Screenshoot Output
![Screenshot 2024-10-20 213754](https://github.com/user-attachments/assets/086c4373-43eb-4c86-a180-9e5b9782952d)

##### Deskripsi Program
Program ini digunakan untuk menghitung permutasi dan kombinasi dari dua pasang bilangan bulat yang diberikan oleh user. Setelah menerima input empat bilangan, program akan memeriksa apakah pasangan bilangan pertama lebih besar sama dengan pasangan bilangan kedua. Jika kondisi ini terpenuhi, maka program akan menghitung permutasi dan kombinasi dari kedua pasangan tersebut secara bergantian. Perhitungan dilakukan dengan menggunakan fungsi-fungsi untuk menghitung faktorial, permutasi, dan kombinasi. Hasil akhir berupa jumlah permutasi dan kombinasi dari setiap pasangan bilangan kemudian ditampilkan secara berurutan.

##### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.
prosedure hitungSkor (in/out soal, skor : integer)
Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilal adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.
Keterangan:
Astuti menyelesaikan 6 soal dalam waktu 287 menit, sedangkan Bertha 7 soal dalam waktu 294 menit. Karena Bertha menyelesaikan lebih banyak, maka Bertha menang. Jika keduanya menyelesaikan sama banyak, maka pemenang adalah yang menyelesaikan dengan total waktu paling kecil.
```go
package main

import (
	"fmt"
)

func hitungSkor(waktu []int, soal *int, skor *int) {
	*soal = 0
	*skor = 0

	for _, t := range waktu {
		if t <= 300 {
			*soal++
			*skor += t
		}
	}
}

func main() {

	var namaPemenang string
	maxSoal := -1
	minWaktu := 301 * 8

	for i := 0; i < 100; i++ {
		var nama string
		var waktu [8]int
		var soal, skor int

		fmt.Print("Masukkan nama peserta: ")
		fmt.Scan(&nama)

		if nama == "selesai" {
			break
		}

		fmt.Print("Masukkan waktu pengerjaan 8 soal: ")
		for j := 0; j < 8; j++ {
			fmt.Scan(&waktu[j])
		}

		hitungSkor(waktu[:], &soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			maxSoal = soal
			minWaktu = skor
			namaPemenang = nama
		}
	}

	fmt.Printf("Pemenang: %s, Soal Diselesaikan: %d, Total Waktu: %d menit\n",
		namaPemenang, maxSoal, minWaktu)
}
```
##### Screenshoot Output
![Screenshot 2024-10-20 222647](https://github.com/user-attachments/assets/778918c0-a9f8-412a-9bca-b38cf1b00dd9)

##### Deskripsi Program
Program ini digunakan untuk menentukan pemenang dalam sebuah kompetisi yang berbasis waktu. Program akan terus meminta pengguna memasukkan nama peserta dan waktu pengerjaan untuk setiap soal hingga pengguna memasukkan kata "selesai". Setiap peserta akan dinilai berdasarkan jumlah soal yang dikerjakan dalam waktu kurang dari 5 menit dan total waktu pengerjaan. Peserta dengan jumlah soal terbanyak dan waktu tercepat akan dinyatakan sebagai pemenang.

##### 3. Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah 1/2n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah: 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1
Untuk suku awal sampal dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1. Buat program sklena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetak Deret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal. 
prosedure cetak Deret(in n : integer) 
Masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000. Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dan dipisahkan oleh sebuah spasi.
```go
package main

import (
	"fmt"
)

func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat : ")
	fmt.Scan(&n)

	if n <= 0 || n >= 1000000 {
		fmt.Println("Bilangan harus kurang dari 1.000.000")
		return
	}

	cetakDeret(n)
}
```
##### Screenshoot Output
![Screenshot 2024-10-20 205334](https://github.com/user-attachments/assets/95ef956d-31d2-47ce-a457-d7c895452057)

##### Deskripsi Program
Program ini digunakan untuk menghasilkan deret bilangan yang mengikuti aturan khusus, pertama user akan diminta untuk menginputkan bilangan bulat. Jika bilangan tersebut genap, maka bilangan tersebut dibagi dua. Sebaliknya, jika bilangan tersebut ganjil, maka bilangan tersebut dikalikan tiga dan ditambah satu. Proses ini berulang terus hingga nilai bilangan mencapai 1. Setiap bilangan yang dihasilkan dalam proses ini akan ditampilkan. Program ini juga memiliki batasan input: bilangan yang dimasukkan harus kurang dari 1.000.000. Batasan ini berfungsi untuk menghindari hasil yang tidak terduga atau waktu eksekusi yang terlalu lama.


### Referensi
[1] MODUL 4 PROSEDUR
[2] Perbedaan function dengan procedure pada golang. Diakses pada 20 Oktober 2024, dari https://www.studocu.com/id/document/universitas-telkom/algorithm-and-programming/perbedaan-function-dengan-procedure-pada-golang/31410467
[3] Doxsey, C. (2016). Introducing Go: Build Reliable, Scalable Programs. United States: O'Reilly Media.
