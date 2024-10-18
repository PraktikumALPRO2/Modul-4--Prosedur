<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL III</strong></h2>
<h2 align="center"><strong> FUNGSI </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Wahyu Hidayat / 2311102178<br>
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
Prosedur adalah blok kode yang tidak mengembalikan nilai tetapi dirancang untuk menjalankan tugas tertentu. Berbeda dengan fungsi yang mengembalikan hasil, prosedur hanya menjalankan sekumpulan instruksi tanpa memberikan nilai kembali ke pemanggil. Prosedur berguna untuk memecah program menjadi bagian yang lebih sederhana dan mudah dipahami, terutama ketika tugas tidak memerlukan pengembalian nilai [1].

Prosedur membantu meningkatkan keterbacaan dan organisasi program karena memungkinkan pembagian logika program menjadi beberapa bagian yang terpisah namun tetap koheren dalam menjalankan tugas tertentu [2].

### Deklarasi Prosedur
Dalam bahasa pemrograman, deklarasi prosedur biasanya sangat mirip dengan fungsi, namun prosedur tidak memiliki tipe pengembalian. Pada beberapa bahasa seperti Pascal atau Go (dimana prosedur disebut sebagai "fungsi tanpa pengembalian"), prosedur dideklarasikan menggunakan kata kunci yang sama dengan fungsi namun tanpa tipe pengembalian.

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
package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    if a >= b {
        permutasi(a, b)
    } else {
        permutasi(b, a)
    }
}

// Prosedur faktorial untuk menghitung dan menampilkan faktorial
func faktorial(n int) {
    var hasil int = 1
    var i int
    for i = 1; i <= n; i++ {
        hasil = hasil * i
    }
    fmt.Println("Faktorial dari", n, "adalah:", hasil)
}

// Prosedur permutasi untuk menghitung dan mencetak hasil permutasi
func permutasi(n, r int) {
    faktorialN := 1
    faktorialNR := 1
    // Hitung faktorial n
    for i := 1; i <= n; i++ {
        faktorialN *= i
    }
    // Hitung faktorial (n-r)
    for i := 1; i <= (n - r); i++ {
        faktorialNR *= i
    }
    // Cetak hasil permutasi
    fmt.Println("Permutasi dari", n, "dan", r, "adalah:", faktorialN/faktorialNR)
}
```
#### Screenshoot Source Code

![Screenshot 2024-10-18 160722](https://github.com/user-attachments/assets/1c196384-aed0-45f5-ab10-29f209cb065e)



#### Screenshoot Output
![Screenshot 2024-10-18 160728](https://github.com/user-attachments/assets/85c7508d-0715-4cbe-a245-d88e7ef9a819)



#### Deskripsi Program
Program ini menghitung permutasi dari dua bilangan bulat positif, yaitu jumlah cara untuk menyusun sejumlah objek dari total objek yang tersedia. Program meminta dua input bilangan dari pengguna, dan jika bilangan pertama lebih kecil dari bilangan kedua, keduanya akan dibalik agar perhitungan permutasi tetap valid. Setelah menerima input, program akan menggunakan dua prosedur: satu untuk menghitung faktorial dan satu lagi untuk menghitung permutasi berdasarkan rumus matematika. Hasil permutasi kemudian dicetak di layar tanpa mengembalikan nilai, sehingga pengguna dapat langsung melihat hasil perhitungan.

#### Algoritma Program
1. Input: Minta pengguna untuk memasukkan dua bilangan bulat positif (a, b).
2. Periksa: Jika a >= b, panggil permutasi(a, b), jika tidak, panggil permutasi(b, a).
3. Prosedur Faktorial:
- Terima parameter n.
- Hitung faktorial dari n.
- Cetak hasil faktorial.
4. Prosedur Permutasi:
- Terima parameter n dan r.
- Hitung faktorial dari n dan (n - r).
- Cetak hasil permutasi (faktorial(n) / faktorial(n - r)).

#### Cara Kerja
1. Input:
- Program meminta pengguna untuk memasukkan dua bilangan bulat positif. Nilai yang dimasukkan disimpan dalam variabel a dan b.
2. Perbandingan:
- Program memeriksa apakah nilai a lebih besar atau sama dengan nilai b.
- Jika ya, program akan melanjutkan untuk menghitung permutasi dengan menggunakan a dan b.
- Jika tidak, program akan membalik nilai a dan b dan kemudian melanjutkan untuk menghitung permutasi.
3. Hitung Permutasi:
- Program memanggil prosedur yang bernama permutasi, menggunakan nilai a dan b (atau b dan a jika dibalik) sebagai argumen.
4. Prosedur Faktorial:
- Dalam prosedur faktorial, program menghitung dan mencetak nilai faktorial dari bilangan yang diberikan dengan mengalikan semua angka dari satu hingga bilangan tersebut.
5. Prosedur Permutasi:
- Dalam prosedur permutasi, program menghitung faktorial dari bilangan total dan faktorial dari selisih antara total dan jumlah objek yang disusun.
Hasil permutasi kemudian dicetak berdasarkan pembagian antara faktorial total dan faktorial selisih.

## 2. Program Penghitungan Luas dan Keliling Persegi

#### Source Code
```go
package main

import (
	"fmt"
)

// Prosedur untuk menghitung dan menampilkan luas persegi
func hitungLuas(sisi float64) {
	luas := sisi * sisi
	fmt.Printf("Luas persegi: %.2f\n", luas)
}

// Prosedur untuk menghitung dan menampilkan keliling persegi
func hitungKeliling(sisi float64) {
	keliling := 4 * sisi
	fmt.Printf("Keliling persegi: %.2f\n", keliling)
}

func main() {
	var sisi float64

	fmt.Print("Sisi Persegi: ")
	fmt.Scan(&sisi)

	// Panggil prosedur untuk menghitung luas dan keliling
	hitungLuas(sisi)
	hitungKeliling(sisi)
}

```
#### Screenshoot Source Code
![Screenshot 2024-10-18 162007](https://github.com/user-attachments/assets/205a1f11-5b54-4806-836c-aa0b3e26aba9)



#### Screenshoot Output
![Screenshot 2024-10-18 162017](https://github.com/user-attachments/assets/7819802a-6c4f-41ac-b1ff-2db968f152d2)




#### Deskripsi Program
Program ini menghitung luas dan keliling persegi berdasarkan sisi yang dimasukkan oleh pengguna. Setelah meminta input berupa panjang sisi persegi, program memanggil dua prosedur: hitungLuas untuk menghitung dan menampilkan luas persegi, serta hitungKeliling untuk menghitung dan menampilkan keliling persegi. Kedua prosedur ini tidak mengembalikan nilai, tetapi langsung mencetak hasil perhitungan ke layar, sehingga pengguna dapat dengan mudah melihat hasilnya. Program ini dirancang untuk memberikan informasi yang cepat dan jelas tentang luas dan keliling persegi yang diberikan.

#### Algoritma Program
1. Input:
- Minta pengguna untuk memasukkan panjang sisi persegi.
2. Panggil Prosedur Hitung Luas:
- Hitung luas persegi dengan rumus sisi dikali sisi.
- Tampilkan hasil luas ke layar.
3. Panggil Prosedur Hitung Keliling:
- Hitung keliling persegi dengan rumus empat kali sisi.
- Tampilkan hasil keliling ke layar.
4. Selesai:
- Program selesai setelah menampilkan hasil luas dan keliling.

#### Cara Kerja
1. Input: Program meminta pengguna untuk memasukkan panjang sisi persegi dan menyimpannya dalam variabel sisi.
2. Hitung Luas: Program memanggil prosedur hitungLuas, yang menghitung luas persegi dan mencetak hasilnya.
3. Hitung Keliling: Program memanggil prosedur hitungKeliling, yang menghitung keliling persegi dan mencetak hasilnya.
4. Output: Program menampilkan hasil luas dan keliling persegi kepada pengguna.

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

## 2. Program Komposisi Tiga Fungsi Matematika

#### Source Code
```go
package main

import (
	"fmt"
)

// Function f(x) = x^2
func f(x int) int {
	return x * x
}

// Function g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Function h(x) = x + 1
func h(x int) int {
	return x + 1
}

// Main function to read input and calculate the compositions
func main() {
	var a, b, c int
	fmt.Println("Masukkan nilai a, b, c dipisahkan oleh spasi:")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	// Calculate f(g(h(a)))
	result1 := f(g(h(a)))
	// Calculate g(h(f(b)))
	result2 := g(h(f(b)))
	// Calculate h(f(g(c)))
	result3 := h(f(g(c)))

	// Output the results
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}


```
#### Screenshoot Source Code
![Screenshot 2024-10-13 173415](https://github.com/user-attachments/assets/793a8c82-e2eb-4fd9-82c7-71dc324bad83)





#### Screenshoot Output
![Screenshot 2024-10-13 173421](https://github.com/user-attachments/assets/4a00b51d-322e-4730-b265-fa9f703621b1)




#### Deskripsi Program
Program ini mengimplementasikan tiga fungsi matematika sederhana: fungsi kuadrat, fungsi pengurangan, dan fungsi penambahan. Program kemudian menerima tiga input dari pengguna, menghitung komposisi dari ketiga fungsi tersebut, dan menampilkan hasilnya. Komposisi yang dihitung adalah:
1. Fungsi f(g(h(a)))
2. Fungsi g(h(f(b)))
3. Fungsi h(f(g(c)))

#### Algoritma Program
1. Definisikan tiga fungsi:
- Fungsi f(x) yang mengembalikan hasil kuadrat dari input.
- Fungsi g(x) yang mengurangi 2 dari input.
- Fungsi h(x) yang menambah 1 pada input.
2. Baca tiga nilai input dari pengguna (a, b, c).
3. Hitung hasil komposisi fungsi untuk setiap input:
- Komposisi f(g(h(a))):
  - Pertama, hitung h(a) (nilai a ditambah 1).
  - Kedua, hitung g(h(a)) (hasil dari h(a) dikurangi 2).
  - Ketiga, hitung f(g(h(a))) (kuadratkan hasil dari g(h(a))).
- Komposisi g(h(f(b))):
  - Pertama, hitung f(b) (kuadratkan nilai b).
  - Kedua, hitung h(f(b)) (hasil dari f(b) ditambah 1).
  - Ketiga, hitung g(h(f(b))) (kurangi hasil dari h(f(b)) dengan 2).
- Komposisi h(f(g(c))):
  - Pertama, hitung g(c) (nilai c dikurangi 2).
  - Kedua, hitung f(g(c)) (kuadratkan hasil dari g(c)).
  - Ketiga, hitung h(f(g(c))) (tambahkan 1 pada hasil dari f(g(c))).
4. Tampilkan hasil dari setiap komposisi di layar.


#### Cara Kerja
1. Pengguna memasukkan tiga angka a, b, dan c yang dipisahkan oleh spasi.
2. Program membaca angka tersebut dan menghitung tiga hasil komposisi fungsi berdasarkan urutan berikut:
- Untuk angka pertama a, program akan menghitung fungsi f(g(h(a))).
- Untuk angka kedua b, program menghitung g(h(f(b))).
- Untuk angka ketiga c, program menghitung h(f(g(c))).
3. Setelah ketiga hasil diperoleh, program akan menampilkan setiap hasil di layar, masing-masing pada baris yang berbeda.






## 3. Menentukan Posisi Titik terhadap Dua Lingkaran

#### Source Code
```go
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

// Fungsi untuk memeriksa apakah suatu titik berada di dalam lingkaran
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	// Input koordinat untuk dua lingkaran dan titik sembarang
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	// Input untuk Lingkaran 1
	fmt.Print("Masukkan pusat lingkaran 1 (cx1 cy1) dan radius r1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input untuk Lingkaran 2
	fmt.Print("Masukkan pusat lingkaran 2 (cx2 cy2) dan radius r2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input untuk Titik Sembarang
	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scan(&x, &y)

	// Cek posisi titik terhadap lingkaran
	dalamLingkaran1 := didalam(cx1, cy1, r1, x, y)
	dalamLingkaran2 := didalam(cx2, cy2, r2, x, y)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
#### Screenshoot Source Code
![Screenshot 2024-10-13 174426](https://github.com/user-attachments/assets/2f14133f-26bb-41c3-9d7f-5c8626486c07)




#### Screenshoot Output
![Screenshot 2024-10-13 174439](https://github.com/user-attachments/assets/4fd769c7-ba9b-47b5-aec1-fb918fbbc0b1)






#### Deskripsi Program
Program ini berfungsi untuk menentukan apakah sebuah titik sembarang berada di dalam lingkaran pertama, di dalam lingkaran kedua, di dalam kedua lingkaran, atau di luar keduanya. Input yang diperlukan adalah koordinat pusat dan jari-jari dari dua lingkaran serta koordinat titik sembarang. Program kemudian menghitung jarak titik tersebut dari pusat setiap lingkaran dan memeriksa apakah jaraknya lebih kecil atau sama dengan jari-jari lingkaran.

#### Algoritma Program
1. Menerima input dari pengguna:
- Koordinat pusat dan jari-jari lingkaran pertama.
- Koordinat pusat dan jari-jari lingkaran kedua.
- Koordinat titik sembarang yang ingin diperiksa.
2. Menghitung jarak antara titik sembarang dan pusat lingkaran pertama menggunakan rumus jarak Euclidean.
3. Menghitung jarak antara titik sembarang dan pusat lingkaran kedua.
4. Membandingkan hasil jarak dengan jari-jari lingkaran:
- Jika jarak titik ke pusat lingkaran pertama kurang dari atau sama dengan jari-jari lingkaran pertama, maka titik berada di dalam lingkaran pertama.
- Jika jarak titik ke pusat lingkaran kedua kurang dari atau sama dengan jari-jari lingkaran kedua, maka titik berada di dalam lingkaran kedua.
- Jika jarak titik ke kedua lingkaran kurang dari atau sama dengan masing-masing jari-jari, maka titik berada di dalam kedua lingkaran.
- Jika jarak titik ke kedua lingkaran lebih besar dari jari-jari, maka titik berada di luar kedua lingkaran.
5. Menampilkan hasilnya, apakah titik berada di dalam lingkaran pertama, kedua, keduanya, atau di luar.

#### Cara Kerja
1. Input: Pengguna memasukkan tiga set data:
- Koordinat dan jari-jari lingkaran pertama.
- Koordinat dan jari-jari lingkaran kedua.
- Koordinat titik sembarang.
2. Penghitungan Jarak: Program menghitung jarak dari titik sembarang ke pusat lingkaran menggunakan rumus jarak Euclidean (akar kuadrat dari jumlah kuadrat selisih koordinat).
3. Pengecekan Posisi:
- Jika jarak titik ke pusat lingkaran lebih kecil atau sama dengan jari-jari lingkaran, titik tersebut dianggap berada di dalam lingkaran.
- Program melakukan hal ini untuk kedua lingkaran.
4. Hasil: Program menentukan apakah titik tersebut berada di dalam lingkaran pertama, di dalam lingkaran kedua, di dalam keduanya, atau di luar keduanya, lalu mencetak hasilnya.

### Kesimpulan
Fungsi adalah alat penting dalam pemrograman yang memungkinkan kita untuk membuat kode lebih modular, terorganisir, dan dapat digunakan kembali. Dengan membagi program ke dalam beberapa fungsi, kita dapat menyederhanakan pengembangan dan pemeliharaan kode. Di Go, deklarasi fungsi mencakup nama fungsi, parameter, dan tipe pengembalian, serta pemanggilan fungsi dilakukan dengan menyebutkan nama fungsi beserta argumennya [2][3].

## Referensi 
[1] Go.dev. (n.d.). Introduction to Go functions. Go Documentation.

[2] Zakhour, M., et al. (2010). The Java programming language. Addison-Wesley.
