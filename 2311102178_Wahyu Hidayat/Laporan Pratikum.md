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
Fungsi adalah blok kode yang dapat dijalankan secara independen dalam sebuah program. Fungsi digunakan untuk memecah program menjadi bagian-bagian yang lebih kecil dan lebih terorganisir. Setiap fungsi dirancang untuk melakukan tugas tertentu dan dapat dipanggil dari bagian lain dari program untuk menghindari penulisan ulang kode yang sama. Dengan kata lain, fungsi memberikan cara untuk membuat program lebih modular dan efisien dalam hal pemeliharaan serta penggunaan ulang kode [1]. 

Fungsi juga meningkatkan keterbacaan program karena pengembang dapat memisahkan logika yang kompleks ke dalam beberapa fungsi yang lebih kecil dan lebih mudah dikelola [2].

### Deklarasi Fungsi
Dalam bahasa Go, deklarasi fungsi dimulai dengan kata kunci func diikuti dengan nama fungsi, parameter yang diperlukan (opsional), dan tipe pengembalian (opsional). Fungsi dideklarasikan dengan sintaks sebagai berikut:
```go
func namaFungsi(parameter1 tipe1, parameter2 tipe2) tipePengembalian {
    // blok kode yang akan dieksekusi
    return nilaiPengembalian
}

```
#### Penjelasan:
namaFungsi: Nama yang digunakan untuk memanggil fungsi.

parameter: Nilai masukan yang diterima oleh fungsi, dapat terdiri dari nol atau lebih parameter.

tipePengembalian: Tipe data yang dikembalikan oleh fungsi, bisa berupa tipe data tunggal atau lebih dari satu (multi-return values).

blok kode: Kumpulan perintah yang akan dijalankan ketika fungsi dipanggil.

return: Mengembalikan hasil dari fungsi ke pemanggil (opsional, tergantung tipe pengembalian).

## Contoh Deklarasi Fungsi dalam Golang
Contoh sederhana deklarasi fungsi dalam Golang:

```go
func sapa(nama string) string {
    return "Halo, " + nama
}

```
Pada contoh di atas, fungsi sapa menerima parameter nama bertipe string dan mengembalikan nilai bertipe string yang merupakan pesan sapaan [3].

### Pemanggilan Fungsi
Pemanggilan fungsi adalah cara untuk mengeksekusi kode yang berada di dalam sebuah fungsi. Fungsi yang telah dideklarasikan bisa dipanggil kapan saja selama berada dalam lingkup yang benar. Untuk memanggil fungsi, kita cukup menuliskan nama fungsi diikuti oleh kurung buka dan tutup beserta argumen yang diperlukan (jika ada) [1].

Contoh pemanggilan fungsi dari contoh sebelumnya:

```go
func main() {
    hasil := sapa("Andi") // Memanggil fungsi 'sapa' dengan argumen "Andi"
    fmt.Println(hasil)    // Output: Halo, Andi
}
```
Dalam contoh di atas, fungsi sapa dipanggil dalam fungsi main dengan memberikan nilai "Andi" sebagai argumen. Hasil dari pemanggilan fungsi tersebut disimpan dalam variabel hasil yang kemudian dicetak ke layar [3].

#### Fungsi dengan Nilai Kembali Lebih dari Satu
Bahasa Go mendukung fungsi dengan pengembalian lebih dari satu nilai (multi-return). Ini berguna ketika kita perlu mengembalikan beberapa hasil dari satu fungsi. Contoh fungsi dengan beberapa nilai kembali:

```go
func hitung(a int, b int) (int, int) {
    penjumlahan := a + b
    pengurangan := a - b
    return penjumlahan, pengurangan
}

```
#### Pemanggilan fungsi:
```go
func hitung(a int, b int) (int, int) {
    penjumlahan := a + b
    pengurangan := a - b
    return penjumlahan, pengurangan
}

```
Pada contoh di atas, fungsi hitung mengembalikan dua nilai, yaitu hasil penjumlahan dan pengurangan dari dua bilangan. Saat memanggil fungsi tersebut, kita bisa menangkap kedua nilai pengembaliannya sekaligus [3].






## II. GUIDED
## 1. Program Perhitungan Permutasi Dua Bilangan Bulat

#### Source Code
```go
package main 

import "fmt"
 
func main(){
	var a,b int 
	fmt.Scan(&a, &b) 
	if a >= b {
		fmt.Println(permutasi(a,b))
	}else{
		fmt.Println(permutasi(b,a))
	}
}
func faktorial(n int) int{ 
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ { 
		hasil = hasil * i
	}

	return hasil
}
func permutasi(n,r int) int {
	return faktorial(n) / faktorial(n-r)
}
```
#### Screenshoot Source Code
![Screenshot 2024-10-13 165455](https://github.com/user-attachments/assets/876a0dc8-2ef3-44db-a9ae-11e12592d7e3)



#### Screenshoot Output
![Screenshot 2024-10-13 165444](https://github.com/user-attachments/assets/757e7796-17de-4103-ad21-3c2386e5038c)



#### Deskripsi Program
Program ini menghitung permutasi dua bilangan bulat positif, yaitu jumlah cara untuk menyusun sejumlah objek dari total objek yang tersedia. Program meminta dua input bilangan dan menghitung permutasi berdasarkan rumus matematika permutasi. Jika nilai bilangan pertama lebih kecil dari yang kedua, maka keduanya dibalik agar perhitungan tetap valid.

#### Algoritma Program
1. Input dua bilangan bulat.
2. Periksa apakah bilangan pertama lebih besar atau sama dengan bilangan kedua.
3. Jika benar, hitung permutasi menggunakan rumus permutasi.
4. Jika salah, tukar urutan bilangan dan hitung permutasi.
5. Hitung faktorial dari bilangan pertama dan faktorial dari selisih bilangan pertama dan kedua.
6. Bagikan hasil faktorial untuk mendapatkan nilai permutasi.
7. Cetak hasil perhitungan.

#### Cara Kerja
- Fungsi faktorial(n int): Fungsi ini menghitung faktorial dengan mengalikan bilangan dari satu hingga bilangan tersebut.
- Fungsi permutasi(n, r int): Fungsi ini menghitung permutasi dengan membagi faktorial bilangan pertama dengan faktorial selisih antara bilangan pertama dan kedua.
- Proses Input: Program menerima dua bilangan bulat dari pengguna, kemudian memeriksa apakah bilangan pertama lebih besar dari atau sama dengan bilangan kedua.
- Proses Output: Hasil perhitungan permutasi dicetak ke layar berdasarkan input pengguna.

## 2. Program Penghitungan Luas dan Keliling Persegi

#### Source Code
```go
package main

import (
	"fmt"
)

func hitungLuas(sisi float64) float64 {
	return sisi * sisi
}

func hitungKeliling(sisi float64) float64 {
	return 4 * sisi
}

func main() {
	var sisi float64

	fmt.Print("Sisi Persegi: ")
	fmt.Scan(&sisi)

	luas := hitungLuas(sisi)
	keliling := hitungKeliling(sisi)

	fmt.Printf("Luas persegi: %.2f\n", luas)
	fmt.Printf("Keliling persegi: %.2f\n", keliling)
}
```
#### Screenshoot Source Code
![Screenshot 2024-10-13 171043](https://github.com/user-attachments/assets/500889b0-56d7-4943-80fe-dba39e2e3e8f)



#### Screenshoot Output
![Screenshot 2024-10-13 171037](https://github.com/user-attachments/assets/80d7b38d-9263-4c39-aa41-c496c38b1c7e)




#### Deskripsi Program
Program ini digunakan untuk menghitung luas dan keliling dari sebuah persegi. Pengguna memasukkan panjang sisi persegi, lalu program akan menghitung dan menampilkan hasil luas dan kelilingnya. Perhitungan dilakukan dengan rumus sederhana untuk luas dan keliling persegi.

#### Algoritma Program
1. Input panjang sisi persegi dari pengguna.
2. Hitung luas persegi menggunakan rumus: panjang sisi × panjang sisi.
3. Hitung keliling persegi menggunakan rumus: 4 × panjang sisi.
4. Tampilkan hasil perhitungan luas dan keliling dengan format dua angka desimal.

#### Cara Kerja
- Fungsi hitungLuas(sisi float64): Fungsi ini menghitung luas persegi dengan mengalikan panjang sisi dengan dirinya sendiri.
- Fungsi hitungKeliling(sisi float64): Fungsi ini menghitung keliling persegi dengan mengalikan panjang sisi dengan 4.
- Proses Input: Pengguna diminta untuk memasukkan nilai panjang sisi persegi.
- Proses Output: Program menampilkan hasil perhitungan luas dan keliling persegi dengan dua angka desimal untuk lebih presisi.

## III. UNGUIDED
## 1. Program Penghitung Permutasi dan Kombinasi

#### Source Code
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

// Fungsi untuk menghitung permutasi P(n, r)
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r)
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	// Input
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan a, b, c, d: ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// Hasil permutasi dan kombinasi untuk a dan c
	p_ac := permutation(a, c)
	c_ac := combination(a, c)

	// Hasil permutasi dan kombinasi untuk b dan d
	p_bd := permutation(b, d)
	c_bd := combination(b, d)

	// Output
	fmt.Printf("%d %d\n", p_ac, c_ac)
	fmt.Printf("%d %d\n", p_bd, c_bd)
}

```
#### Screenshoot Source Code
![Screenshot 2024-10-13 171753](https://github.com/user-attachments/assets/e441a9a2-e5b2-4917-ad79-d6393e5d4edb)




#### Screenshoot Output
![Screenshot 2024-10-13 171736](https://github.com/user-attachments/assets/7fe92782-099e-4ef4-9175-93d7548731c7)





#### Deskripsi Program
Program ini dibuat untuk menghitung nilai permutasi dan kombinasi dari empat bilangan bulat yang diberikan sebagai input. Pengguna akan memasukkan empat bilangan bulat, yaitu a, b, c, dan d. Program ini kemudian menghitung hasil permutasi dan kombinasi dari a terhadap c, serta b terhadap d, dan menampilkan hasilnya dalam dua baris. Baris pertama menunjukkan hasil perhitungan permutasi dan kombinasi dari a terhadap c, dan baris kedua menunjukkan hasil perhitungan permutasi dan kombinasi dari b terhadap d.

#### Algoritma Program
1. Input:
- Program menerima empat bilangan bulat a, b, c, dan d.
2. Proses:
- Untuk menghitung permutasi dan kombinasi, terlebih dahulu dilakukan perhitungan faktorial dari bilangan-bilangan yang dibutuhkan.
- Rumus yang digunakan:
  - Permutasi dihitung dengan membagi faktorial n dengan faktorial dari selisih n dan r.
  - Kombinasi dihitung dengan membagi faktorial n dengan hasil perkalian faktorial r dan faktorial dari selisih n dan r.
- Permutasi dan kombinasi dihitung untuk pasangan a terhadap c, serta b terhadap d.
3. Output:
- Hasil permutasi dan kombinasi dari a terhadap c ditampilkan di baris pertama.
- Hasil permutasi dan kombinasi dari b terhadap d ditampilkan di baris kedua.

#### Cara Kerja
1. Input Data:
- Program meminta pengguna memasukkan empat bilangan bulat. Sebagai contoh, jika pengguna memasukkan "5 10 3 10", maka a adalah 5, b adalah 10, c adalah 3, dan d adalah 10.
2. Menghitung Faktorial:
- Faktorial adalah hasil perkalian berurutan dari suatu bilangan dengan bilangan yang lebih kecil hingga satu. Misalnya, faktorial dari 5 adalah 5 x 4 x 3 x 2 x 1 = 120.
- Faktorial dihitung untuk bilangan-bilangan yang relevan dalam menghitung permutasi dan kombinasi.
3. Menghitung Permutasi dan Kombinasi:
- Permutasi dihitung dengan membagi faktorial dari a dengan faktorial dari selisih a dan c. Contoh: untuk a = 5 dan c = 3, permutasi dihitung sebagai 5 faktorial dibagi dengan 2 faktorial, yang hasilnya adalah 60.
- Kombinasi dihitung dengan membagi faktorial dari a dengan hasil kali faktorial dari c dan faktorial dari selisih a dan c.
4. Menampilkan Hasil:
- Hasil perhitungan permutasi dan kombinasi dari a terhadap c ditampilkan terlebih dahulu.
- Kemudian, hasil perhitungan permutasi dan kombinasi dari b terhadap d ditampilkan di baris berikutnya.





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
[1] A. A. Demailly, "Efficient Function Calls and Modular Programming in Go," Journal of Software Engineering, vol. 34, no. 2, pp. 145-153, 2022.

[2] R. C. Martin, Clean Code: A Handbook of Agile Software Craftsmanship, Pearson Education, 2009.

[3] A. Donovan and B. W. Kernighan, The Go Programming Language, Addison-Wesley, 2015.
