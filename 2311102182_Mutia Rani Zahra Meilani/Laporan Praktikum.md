
<p align="center">
  <strong>LAPORAN PRAKTIKUM</strong>
  <br>
  <strong>ALGORITMA DAN PEMROGRAMAN 2</strong>
  <br>
</p>

<br>

<p align="center">
  <strong>MODUL IV</strong>
  <br>
  <strong>PROSEDUR</strong>
  <br>
</p>

<br>

<p align="center">
  <img src="https://github.com/user-attachments/assets/296eb3c2-bd6b-401f-8256-3661150ec39e" alt="Logo" width="200"/>
</p>

<br>

<p align="center">
  <strong>Disusun Oleh :</strong>
  <br>
  Mutia Rani Zahra Meilani
  <br>
  2311102182
  <br>
  S1IF1105
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu :</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong>
  <br>
  <strong>FAKULTAS INFORMATIKA</strong>
  <br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong>
  <br>
  <strong>2024</strong>
</p>

## <strong> DASAR TEORI </strong>

<span style="font-size:16px"><strong> ── PENGERTIAN PROSEDUR</strong></span>
<br>
Dalam konteks pemrograman, prosedur (atau fungsi tanpa nilai kembali) adalah sekumpulan blok kode yang dikelompokkan dan diberi nama. Ketika kita memanggil nama prosedur itu, instruksi di dalamnya akan dieksekusi. Di Golang, prosedur yang tidak mengembalikan nilai disebut sebagai fungsi tanpa nilai kembali. Prosedur kerap digunakan ketika ingin mengeksekusi blok kode yang sama berulang kali tanpa harus menulis blok kode tersebut. Prosedur kerap digunakan dikarenakan :

1. Modularitas: Membagi kode menjadi prosedur yang lebih kecil membuat kode lebih mudah dibaca, dipahami, dan dikelola.
2. Reusabilitas: Prosedur dapat dipanggil berulang kali dari berbagai bagian program, menghindari duplikasi kode.
3. Abstraksi: Prosedur memungkinkan kita untuk menyembunyikan detail implementasi dan hanya fokus pada apa yang prosedur itu lakukan.

<span style="font-size:16px"><strong> ── SINTAKS DASAR PROSEDUR</strong></span>

Prosedur dalam Go dideklarasikan menggunakan kata kunci func, diikuti dengan nama fungsi, dan parameter (jika ada). Sintaks dasar deklarasi prosedur adalah sebagai berikut:
```go
func namaProsedur(parameter1 tipeData1, parameter2 tipeData2, ...) {
    // Blok kode yang akan dieksekusi
}
```

- `func`: Kata kunci untuk mendeklarasikan sebuah fungsi.

- `namaProsedur`: Nama yang diberikan untuk fungsi tersebut.

- `parameter`: Variabel yang menerima nilai saat fungsi dipanggil.

- `tipeData`: Tipe data dari parameter.

- `blok kode`: Kumpulan instruksi yang akan dieksekusi ketika fungsi dipanggil.

Contoh Penggunaan :
```go
func sapa(nama string) {
    fmt.Println("Halo,", nama, "!")
}
```

Contoh Penggunaan :
```go
sapa("Budi") // Output: Halo, Budi!
```

**Prosedur Dengan Return Value**

Walaupun prosedur tidak mengembalikan nilai, kita bisa mengubah nilai variabel global dari dalam prosedur. Namun, praktik ini umumnya tidak disarankan karena dapat membuat kode menjadi sulit dipahami dan diuji.

Contoh Penggunaan :
```go
package main

import "fmt"

var jumlah int

func tambah(a, b int) {
    jumlah = a + b
}

func main() {
    tambah(5, 3)
    fmt.Println("Hasil penjumlahan:", jumlah)
}
```
Program tersebut akan tetap dapat dijalankan dengan hasil penjumlahan sesuai dengan baris kode pada prosedur tambah

**Prosedur Merubah Variabel**

Prosedur juga dapat secara sepihak mengubah nilai dari sebuah variabel menggunakan referensi. namun hal ini juga umumnya tidak disarankan karena dapat membuat kode menjadi sulit dipahami dan diuji.

Contoh Penggunaan :
```go
package main

import "fmt"

func tambah(a, b int, jumlah *int) {
	*jumlah = a + b
}

func main() {
	var jumlah int
    tambah(5, 3, &jumlah)
    fmt.Println("Hasil penjumlahan:", jumlah)
}
```


## <strong> GUIDED </strong>

### ── Guided 1

#### Source Code

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(permutasi(a, b))
	} else {
		fmt.Println(permutasi(b, a))
	}
}
func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
```

#### Output

![image](https://github.com/user-attachments/assets/8fb14f6e-33ed-4562-84cf-812bd803747b)

Deskripsi Program :
Program ini digunakan untuk menghitung permutasi dari dua bilangan bulat yang diberikan. Program ini akan meminta pengguna untuk memasukkan dua bilangan bulat, kemudian akan menghitung permutasi dari bilangan yang lebih besar terhadap bilangan yang lebih kecil. Program ini menggunakan fungsi permutasi yang akan menghitung faktorial, sehingga hasil akhirnya berupa jumlah permutasi dari kedua bilangan yang sudah diinputkan oleh pengguna.

### ── Guided 2

#### Source Code

```go
package main

import "fmt"

func luas(a int)int{
	return a*a
}

func keliling(a int)int{
	return 4*a
}

func main(){
	var input int
	fmt.Scan(&input)
	fmt.Println("Luas Persegi : ", luas(input))
	fmt.Println("Keliling Persegi : ", keliling(input))
}
```

#### Output

![image](https://github.com/user-attachments/assets/fea20b11-1f7a-4ecd-a8e3-cf63cb78b07a)

Deskripsi Program :
Program ini digunakan untuk menghitung luas dan keliling dari sebuah persegi. Program ini akan meminta pengguna untuk memasukan panjang sisi dari sebuah persegi. Program ini memiliki 2 fungsi yang berguna untuk menghitung luas dan keliling persegi dari panjang sisi yang dimasukan oleh pengguna.

## <strong>  UNGUIDED </strong>

### ── Unguided 1

#### Study Case

Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)

**Masukan** terdiri dari empat buah bilangan asli `a`, `b`, `c`, dan `d` yang dipisahkan oleh spasi, dengan syarat `a ≥ c` dan `b ≥ d`.

**Keluaran** terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi `a` terhadap `c`, sedangkan baris kedua adalah hasil permutasi dan kombinasi `b` terhadap `d`.

**Catatan**: permutasi (P) dan kombinasi (C) dari `n` terhadap `r` (`n ≥ r`) dapat dihitung dengan menggunakan persamaan berikut!

$P(n, r) = \frac{n!}{(n-r)!}$ ,  Sedangkan  $C(n, r) = \frac{n!}{r!(n-r)!}$

#### Source Code

```go
package main

import "fmt"

func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(x, y int) {
	fmt.Print(faktorial(x) / faktorial(x-y))
}

func kombinasi(x, y int)  {
	fmt.Print(faktorial(x) / (faktorial(y) * faktorial(x-y)))
}

func main() {
	var a, b, c, d int
	fmt.Print("Input Nilai : ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Print("Permutasi A dan C : ")
		permutasi(a, c)
		fmt.Println(" ")
		fmt.Print("Kombinasi A dan C : ")
		kombinasi(a, c)
		fmt.Println(" ")
		fmt.Print("Permutasi B dan D : ")
		permutasi(b, d)
		fmt.Println(" ")
		fmt.Print("Kombinasi B dan D : ")
		kombinasi(b, d)
	} else {
		fmt.Println("Angka yang diberikan tidak valid.")
	}
}

```

#### Output

![image](https://github.com/user-attachments/assets/aa8a03c9-36cd-434d-b107-898a0ff7a8e0)

#### Deskripsi Program :

Program ini menghitung permutasi dan kombinasi berdasarkan input pengguna. Pertama, program meminta empat angka, yaitu `a`, `b`, `c`, dan `d`. Jika nilai `a` lebih besar atau sama dengan `c` serta `b` lebih besar atau sama dengan `d`, maka program akan menghitung dan mencetak hasil permutasi dan kombinasi untuk dua set bilangan: set pertama adalah `a` dan `c`, sedangkan set kedua adalah `b` dan `d`. Penghitungan permutasi dilakukan dengan rumus faktorial `x! / (x - y)!`, sedangkan kombinasi dihitung dengan rumus `x! / (y! * (x - y)!)`. Faktorial sendiri dihitung menggunakan fungsi `faktorial`, yang mengalikan bilangan secara berurutan dari 1 hingga nilai input yang diberikan. Jika kondisi nilai input tidak terpenuhi (misalnya, `a` kurang dari `c` atau `b` kurang dari `d`), maka program akan menampilkan pesan bahwa "Angka yang diberikan tidak sesuai."

### ── Unguided 2

#### Study Case

Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya.

Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.

```bash
prosedure hitungSkor(in/out soal, skor : integer)
```
Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit).

Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.

#### Source Code

```go
package main

import "fmt"

func kalkulasiSkor(waktu []int) (int, int) {
	jumlahSelesai := 0
	totalWaktu := 0
	for _, t := range waktu {
		if t < 301 {
			jumlahSelesai++
			totalWaktu += t
		}
	}
	return jumlahSelesai, totalWaktu
}

func hitungSkor(dataPeserta map[string][]int) {
	namaPemenang := ""
	soalMaksimum := -1
	skorMinimum := 0

	for nama, waktu := range dataPeserta {
		soalSelesai, totalWaktu := kalkulasiSkor(waktu)
		if soalSelesai > soalMaksimum || (soalSelesai == soalMaksimum && totalWaktu < skorMinimum) {
			namaPemenang = nama
			soalMaksimum = soalSelesai
			skorMinimum = totalWaktu
		}
	}
	fmt.Println("====================")
	fmt.Printf("Winner : %s \n", namaPemenang)
	fmt.Printf("Soal Diselesaikan : %d \n", soalMaksimum)
	fmt.Printf("Skor : %d \n", skorMinimum)
}

func main() {
	selesai := false
	dataPeserta := make(map[string][]int)

	if !selesai {
		for {
			waktu := [8]int{}
			var namaPeserta string
			fmt.Print("Masukkan Nama Peserta : ")
			fmt.Scan(&namaPeserta)

			if namaPeserta == "selesai" {
				selesai = true
				break
			}
			fmt.Printf("Waktu Penyelesaian Soal : ")
			for i := 0; i < 8; i++ {
				fmt.Scan(&waktu[i])
			}
			dataPeserta[namaPeserta] = waktu[:]
		}
	}

	hitungSkor(dataPeserta)
}


```

#### Output

![image](https://github.com/user-attachments/assets/c694c0aa-ca98-4994-a448-95ced3a62cbb)

#### Deskripsi Program :

Program ini bertujuan untuk menentukan pemenang dalam sebuah kompetisi berdasarkan waktu yang dihabiskan peserta untuk menyelesaikan soal-soal pemrograman. Program menerima input berupa nama peserta dan waktu yang diperlukan untuk menyelesaikan masing-masing soal. Data ini disimpan dalam sebuah map dengan nama peserta sebagai kunci dan array waktu sebagai nilai. Fungsi `kalkulasiSkor` digunakan untuk menghitung jumlah soal yang berhasil diselesaikan peserta (waktu di bawah 301 menit) serta total waktu yang dibutuhkan untuk soal-soal yang berhasil dikerjakan. Fungsi `hitungSkor` kemudian menentukan pemenang berdasarkan dua kriteria: peserta yang menyelesaikan soal terbanyak, dan jika ada peserta dengan jumlah soal yang sama, yang memiliki total waktu terkecil. Input dari pengguna diterima hingga mereka memasukkan kata "selesai", yang menandakan akhir dari input peserta. Setelah input selesai, program akan menampilkan nama pemenang, jumlah soal yang diselesaikan, dan total waktu yang dihabiskan untuk soal yang berhasil dikerjakan. Dari sisi kode, program terdiri dari tiga bagian utama. Pertama, fungsi `kalkulasiSkor` yang melakukan penghitungan soal yang diselesaikan dan total waktu yang dihabiskan. Kedua, fungsi `hitungSkor` yang mencari pemenang dengan mengevaluasi setiap peserta berdasarkan hasil dari `kalkulasiSkor`. Ketiga, fungsi `main` yang mengelola input peserta dan waktu pengerjaan soal, lalu memanggil fungsi `hitungSkor` untuk menampilkan pemenangnya.

### ── Unguided 3

#### Study Case

Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah ½n, tetapi jika ganjil maka suku berikutnya bernilai 3n + 1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah:

**22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1**

Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1.

Buat program Skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.

```bash
prosedure cetakDeret(in n : integer)
```
**Masukan** berupa satu bilangan integer positif yang lebih kecil dari 1000000

**Keluaran** terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dipisahkan oleh sebuah spasi

| Masukan | Keluaran |
|---------|----------|
| 22      | 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1 |


#### Source Code

```go
package main

import "fmt"

func cetakDeret(n int) {

	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2 
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(n)
}

func main() {
	var n int
	fmt.Print("Input Nilai : ")
	fmt.Scan(&n)

	if(n < 1000000){
		fmt.Print("Suku Deret : ")
		cetakDeret(n)
	}else {
		fmt.Println("Inputan harus antara 1 sampai 1,000,000!")
	}
}
```

#### Output

![image](https://github.com/user-attachments/assets/fa5ed712-3efc-4975-9905-0e2bcd72568d)

#### Deskripsi Program :

Program ini bertujuan untuk mencetak deret bilangan berdasarkan aturan dari Skiena dan Revilla dalam Programming Challenges. Deret dimulai dengan sebuah bilangan bulat n, dan untuk setiap bilangan dalam deret, jika bilangan tersebut genap, maka bilangan berikutnya adalah setengah dari bilangan itu. Jika bilangan tersebut ganjil, bilangan berikutnya dihitung dengan rumus 3n + 1. Deret berakhir ketika nilai bilangan mencapai 1. Program ini terdiri dari dua fungsi utama: cetakDeret dan main. Fungsi cetakDeret menerima input berupa integer n dan mencetak deret bilangan sesuai aturan di atas. Jika n adalah bilangan genap, nilai n dibagi dua, sedangkan jika ganjil, nilai n dihitung dengan rumus 3n + 1. Fungsi ini terus berjalan hingga nilai n menjadi 1, dan deret akan ditampilkan di terminal. Pada fungsi main, program pertama kali membaca input dari pengguna berupa bilangan integer. Jika nilai yang diinput kurang dari 1.000.000, maka fungsi cetakDeret akan dipanggil untuk mencetak deret bilangan. Namun, jika inputan lebih besar dari 1.000.000, program akan menampilkan pesan error yang menyatakan bahwa inputan harus berada dalam rentang 1 hingga 1.000.000.

## <strong> REFERENSI </strong>

#### [1] Santanapurba, Harja. “DASAR-DASAR BAHASA PEMROGRAMAN GOLANG.” Pilkomedia, 2022.https://pta.pilkommedia.org/progress/upload/AhmadFaisal_A1C615001_Dasar-DasarBahasaPemrogramanGolang.pdf

#### [2] Udacity Team. “Functions in Go (Golang).” Udacity Blog, 2023. https://www.udacity.com/blog/2023/05/functions-in-go-golang.html.

#### [3] MySkill.ID. “Function on Golang.” Medium, 2023. https://medium.com/@myskill.id/function-on-golang-27b6577e9cbe.
