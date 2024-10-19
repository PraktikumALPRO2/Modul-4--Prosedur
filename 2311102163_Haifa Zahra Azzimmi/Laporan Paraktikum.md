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
  Haifa Zahra Azzimmi / 2311102163<br>
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
**1. Definisi Prosedur**

Prosedur merupakan sejenis fungsi yang tidak menghasilkan nilai balik. Tidak seperti fungsi umum yang mengembalikan hasil, prosedur menjalankan tindakan tertentu, seperti menghasilkan output atau memodifikasi data.

**2. Deklarasi Prosedur**

Deklarasi prosedur adalah proses di mana kita menentukan nama prosedur, parameter yang mungkin diterima, dan blok kode yang akan dijalankan ketika prosedur dipanggil. Dalam bahasa Go, prosedur dapat dideklarasikan seperti fungsi, tetapi tanpa nilai kembalian.

Contoh deklarasi prosedur dalam Go:
```go

// Deklarasi prosedur
func tampilkanPesan(pesan string) {
    fmt.Println(pesan)
}

```

**3. Cara Pemanggilan Prosedur**

Memanggil prosedur dalam pemrograman berarti menjalankan atau mengeksekusi blok kode yang telah dideklarasikan dalam prosedur tersebut. 

```go
func main() {
    // Memanggil prosedur tampilkanPesan
    tampilkanPesan("Hello, World!")
}

```

**4. Parameter**

a. Parameter formal 

Parameter formal adalah variabel yang dideklarasikan dalam definisi fungsi atau prosedur untuk menerima nilai saat fungsi atau prosedur tersebut dipanggil. Mereka menentukan jenis dan jumlah argumen yang fungsi atau prosedur harapkan untuk diterima.

Contoh dalam bahasa Go:

```go
// Deklarasi prosedur dengan parameter formal
func cetakDetail(nama string, umur int) {
    fmt.Printf("Nama: %s, Umur: %d\n", nama, umur)
}

```
b. Parameter Aktual

Parameter Aktual adalah nilai atau argumen nyata yang diberikan ke prosedur atau fungsi saat dipanggil. Parameter ini menggantikan parameter formal yang dideklarasikan dalam definisi prosedur atau fungsi.

Contohnya dalam bahasa Go:
package main

import "fmt"
```go
// Deklarasi prosedur dengan parameter formal
func cetakDetail(nama string, umur int) {
    fmt.Printf("Nama: %s, Umur: %d\n", nama, umur)
}

func main() {
    // Parameter aktual
    cetakDetail("Budi", 25)
    cetakDetail("Siti", 30)
}

```
c. Pass by Value

Pass by value adalah konsep pemrograman di mana nilai argumen dikirim ke fungsi, dan perubahan dalam fungsi tidak mempengaruhi nilai asli di luar fungsi. Jadi, salinan nilai dikirimkan.

Contoh dalan Go:
```go
func ubahNilai(x int) {
    x = x * 2
}

func main() {
    a := 10
    ubahNilai(a)
    fmt.Println(a) // Output: 10
}

```
d. Pass by Reference

Pass by reference adalah konsep pemrograman di mana alamat memori argumen dikirim ke fungsi, sehingga perubahan dalam fungsi akan mempengaruhi data asli. Jadi, fungsi bekerja langsung dengan data asli, bukan salinannya.

Contoh dalan Go:

 ```go
func ubahNilai(x *int) {
    *x = *x * 2
}

func main() {
    a := 10
    ubahNilai(&a)
    fmt.Println(a) // Output: 20
}
```




## Guided 

### 1. Buatlah program beserta prosedur yang digunakan untuk menghitung nilai faktorial dan permutasi

### Source Code :
```go
package main

import "fmt"

// Prosedur utama
func main() {
	var a, b int

	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	// Memilih mana yang lebih besar antara a dan b untuk menghitung permutasi
	if a >= b {
		fmt.Printf("Permutasi dari %dP%d adalah: %d\n", a, b, permutasi(a, b))
	} else {
		fmt.Printf("Permutasi dari %dP%d adalah: %d\n", b, a, permutasi(b, a))
	}
}

// Prosedur untuk menghitung faktorial dari suatu bilangan
func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Prosedur untuk menghitung permutasi nPr
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

```

### Output :
![image](https://github.com/user-attachments/assets/c365c126-0c75-4d81-b7da-0566db4d4b39)

### Full code Screenshot :
![image](https://github.com/user-attachments/assets/483dc846-be7c-4769-9a40-f9a4dc490d26)


### Deskripsi Program : 
Program ini meminta pengguna untuk memasukkan dua nilai, a dan b. Kemudian, program memilih nilai yang lebih besar antara a dan b untuk menghitung permutasi nPr (dimana n adalah nilai yang lebih besar dan r adalah nilai yang lebih kecil). Program ini menggunakan dua prosedur: satu untuk menghitung faktorial (faktorial) dan satu untuk menghitung permutasi (permutasi). Hasil permutasi kemudian dicetak ke layar.
- Jika a lebih besar atau sama dengan b, program menghitung permutasi aPb.
- Jika b lebih besar, program menghitung permutasi bPa.

### Algoritma Program :
1. Mulai
2. Minta pengguna memasukkan dua nilai a dan b.
3. Baca input pengguna.
4. Jika a ≥ b, hitung permutasi aPb.
5. Jika b > a, hitung permutasi bPa.
6. Cetak hasil permutasi yang dihitung.
7. Selesai

### Cara Kerja Program :
Program dimulai dengan meminta pengguna memasukkan dua nilai, a dan b. Program kemudian membandingkan kedua nilai tersebut untuk menentukan mana yang lebih besar. Selanjutnya, program menggunakan nilai yang lebih besar sebagai n dan nilai yang lebih kecil sebagai r dalam perhitungan permutasi nPr. Fungsi faktorial digunakan untuk menghitung faktorial dari sebuah bilangan, dan hasil ini digunakan dalam fungsi permutasi untuk menghitung permutasi nPr. Hasil perhitungan permutasi kemudian dicetak ke layar.

### 2. Buatlah program beserta prosedur untuk menghitung luas persegi dan keliling persegi

### Source Code :
```go
package main

import (
	"fmt"
)

// Prosedur utama
func main() {
	var panjangSisi float64

	// Meminta input panjang sisi persegi
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&panjangSisi)

	// Menghitung luas dan keliling persegi menggunakan prosedur terpisah
	luas := hitungLuas(panjangSisi)
	keliling := hitungKeliling(panjangSisi)

	// Menampilkan hasil
	fmt.Printf("Luas persegi: %g\n", luas)
	fmt.Printf("Keliling persegi: %g\n", keliling)
}

// Prosedur untuk menghitung luas persegi
func hitungLuas(sisi float64) float64 {
	return sisi * sisi
}

// Prosedur untuk menghitung keliling persegi
func hitungKeliling(sisi float64) float64 {
	return 4 * sisi
}

```

### Output :
![image](https://github.com/user-attachments/assets/62de6da9-67c7-43db-934b-7029ac640ff4)

### Full code Screenshot :
![image](https://github.com/user-attachments/assets/69f4a885-4d5f-4d6c-9314-264f0928e95e)

### Deskripsi Program : 
rogram ini meminta pengguna untuk memasukkan panjang sisi dari sebuah persegi. Kemudian, program menghitung luas dan keliling persegi menggunakan prosedur terpisah, dan menampilkan hasilnya.

### Algoritma Program :
1. Mulai
2. Minta pengguna memasukkan panjang sisi persegi.
3. Baca input pengguna.
4. Hitung luas persegi menggunakan prosedur hitungLuas.
5. Hitung keliling persegi menggunakan prosedur hitungKeliling.
6. Tampilkan hasil luas dan keliling persegi.
7. Selesai

### Cara Kerja Program :
1. Program dimulai dan meminta pengguna untuk memasukkan panjang sisi dari sebuah persegi.
2. Input panjang sisi dibaca dan disimpan dalam variabel panjangSisi.
3. Program memanggil prosedur hitungLuas dengan panjangSisi sebagai argumen untuk menghitung luas persegi (sisi * sisi).
4. Program memanggil prosedur hitungKeliling dengan panjangSisi sebagai argumen untuk menghitung keliling persegi (4 * sisi).
5. Hasil perhitungan luas dan keliling kemudian ditampilkan ke layar.

## Unguided 

### 1. Program untuk menghitung Faktorial, Permutasi dan Kombinasi
![image](https://github.com/user-attachments/assets/9ba253ad-cf25-47b1-9e25-594bba4eeea3)


### Source Code :
```go
//Haifa Zahra Azzimmi
//2311102163

package main

import (
    "fmt"
)

// Fungsi untuk menghitung faktorial dari sebuah angka
func hitungFaktorial(x int) int {
    if x == 0 || x == 1 {
        return 1
    }
    hasil := 1
    for i := 2; i <= x; i++ {
        hasil *= i
    }
    return hasil
}

// Fungsi untuk menghitung permutasi dari n dan r
func hitungPermutasi(n, r int) int {
    return hitungFaktorial(n) / hitungFaktorial(n-r)
}

// Fungsi untuk menghitung kombinasi dari n dan r
func hitungKombinasi(n, r int) int {
    return hitungFaktorial(n) / (hitungFaktorial(r) * hitungFaktorial(n-r))
}

func main() {
    var angka1, angka2, angka3, angka4 int

    // Meminta pengguna memasukkan 4 angka
    fmt.Println("Masukkan empat angka dengan syarat angka1 >= angka3 dan angka2 >= angka4:")
    fmt.Scan(&angka1, &angka2, &angka3, &angka4)

    // Mengecek apakah syarat angka1 >= angka3 dan angka2 >= angka4 terpenuhi
    if angka1 < angka3 || angka2 < angka4 {
        fmt.Println("Syarat angka1 >= angka3 dan angka2 >= angka4 tidak terpenuhi!")
        return
    }

    // Menghitung permutasi dan kombinasi untuk dua pasangan angka
    permutasiA := hitungPermutasi(angka1, angka3)
    kombinasiA := hitungKombinasi(angka1, angka3)
    permutasiB := hitungPermutasi(angka2, angka4)
    kombinasiB := hitungKombinasi(angka2, angka4)

    // Menampilkan hasil perhitungan
    fmt.Printf("Hasil permutasi dari %d terhadap %d: %d, kombinasi: %d\n", angka1, angka3, permutasiA, kombinasiA)
    fmt.Printf("Hasil permutasi dari %d terhadap %d: %d, kombinasi: %d\n", angka2, angka4, permutasiB, kombinasiB)
}

```

### Output :
![image](https://github.com/user-attachments/assets/182071fe-0f5a-45c1-86a4-fb59c1903275)


### Full code Screenshot :
![image](https://github.com/user-attachments/assets/72871c8a-bace-4f3f-863a-e090a2c9c12b)

### Deskripsi Program : 
Program ini meminta pengguna untuk memasukkan empat angka dengan syarat angka1 ≥ angka3 dan angka2 ≥ angka4. Kemudian, program menghitung dan menampilkan hasil permutasi dan kombinasi untuk dua pasangan angka yang dimasukkan. Menggunakan fungsi untuk menghitung faktorial, permutasi, dan kombinasi.

### Algoritma Program :
1. Mulai
2. Minta pengguna memasukkan empat angka dengan syarat angka1 ≥ angka3 dan angka2 ≥ angka4.
3. Baca input pengguna.
4. Periksa apakah syarat angka1 ≥ angka3 dan angka2 ≥ angka4 terpenuhi.
   - Jika syarat tidak terpenuhi, tampilkan pesan kesalahan dan hentikan program.
5. Jika syarat terpenuhi, hitung permutasi dan kombinasi untuk pasangan angka (angka1, angka3) dan (angka2, angka4).
6. Tampilkan hasil perhitungan permutasi dan kombinasi untuk kedua pasangan angka tersebut.
7. Selesai

### Cara Kerja Program :
Program ini meminta pengguna memasukkan empat angka dengan syarat angka1 ≥ angka3 dan angka2 ≥ angka4. Jika syarat terpenuhi, program menghitung dan menampilkan hasil permutasi dan kombinasi dari dua pasangan angka tersebut menggunakan fungsi yang sudah didefinisikan untuk menghitung faktorial, permutasi, dan kombinasi. Jika syarat tidak terpenuhi, program akan menampilkan pesan kesalahan dan berhenti. Jadi, program ini melakukan operasi matematika pada input angka sesuai dengan aturan yang diberikan

### 2. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan
**Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.**

### Source Code :
```go
package main

import (
    "fmt"
    "math"
)

// Fungsi untuk menghitung skor
func hitungSkor(waktuPengerjaan []int) (int, int) {
    jumlahSoalSelesai := 0
    totalWaktu := 0
    for _, waktu := range waktuPengerjaan {
        if waktu < 301 {
            jumlahSoalSelesai++
            totalWaktu += waktu
        }
    }
    return jumlahSoalSelesai, totalWaktu
}

// Fungsi untuk mencari pemenang
func cariPemenang(peserta map[string][]int) (string, int, int) {
    namaPemenang := ""
    jumlahSoalMax := -1
    skorMin := math.MaxInt64

    for nama, waktuPengerjaan := range peserta {
        jumlahSoalSelesai, totalWaktu := hitungSkor(waktuPengerjaan)
        if jumlahSoalSelesai > jumlahSoalMax || (jumlahSoalSelesai == jumlahSoalMax && totalWaktu < skorMin) {
            namaPemenang = nama
            jumlahSoalMax = jumlahSoalSelesai
            skorMin = totalWaktu
        }
    }
    return namaPemenang, jumlahSoalMax, skorMin
}

func main() {
    var jumlahPeserta int
    fmt.Println("Silakan masukkan jumlah peserta:")
    fmt.Scan(&jumlahPeserta)
    
    peserta := make(map[string][]int)
    
    for i := 0; i < jumlahPeserta; i++ {
        var nama string
        fmt.Printf("Masukkan nama untuk peserta ke-%d:\n", i+1)
        fmt.Scan(&nama)
        
        waktuPengerjaan := make([]int, 8)
        fmt.Printf("Masukkan waktu pengerjaan untuk setiap soal oleh %s:\n", nama)
        for j := 0; j < 8; j++ {
            fmt.Scan(&waktuPengerjaan[j])
        }
        
        peserta[nama] = waktuPengerjaan
    }
    
    // Mencari pemenang
    pemenang, maxSoal, minSkor := cariPemenang(peserta)
    fmt.Printf("Pemenangnya adalah %s dengan jumlah soal %d dan total waktu %d\n", pemenang, maxSoal, minSkor)
}

```

### Output :
![image](https://github.com/user-attachments/assets/fc642d7f-d0dd-4d97-8e99-169107af97af)

### Full code Screenshot:
![image](https://github.com/user-attachments/assets/5876b3ba-8d61-4d34-8158-e77fa5f8d2a1)

### Deskripsi Program : 
Program ini meminta pengguna untuk memasukkan jumlah peserta dan waktu pengerjaan setiap soal oleh masing-masing peserta. Program ini kemudian menghitung skor (jumlah soal yang selesai dalam waktu kurang dari 301 detik dan total waktu pengerjaan) untuk setiap peserta dan menentukan pemenang berdasarkan jumlah soal yang selesai dan total waktu pengerjaan terendah.
### Algoritma Program :
1. Mulai
2. Minta pengguna memasukkan jumlah peserta.
3. Baca input pengguna.
   - Untuk setiap peserta:
   - Minta nama peserta.
   - Baca nama peserta.
   - Minta waktu pengerjaan untuk setiap soal (8 soal).
   - Baca waktu pengerjaan.
   - Simpan data waktu pengerjaan dalam peta (map).
4. Hitung skor untuk setiap peserta:
   - Hitung jumlah soal yang selesai dalam waktu kurang dari 301 detik.
   - Hitung total waktu pengerjaan untuk soal yang selesai.
5. Tentukan pemenang berdasarkan:
   - Jumlah soal yang selesai terbanyak.
   - Jika jumlah soal sama, pilih peserta dengan total waktu pengerjaan terendah.
7. Tampilkan pemenang dan hasil perhitungan.
8. Selesai
   
### Cara Kerja Program :
1. Program dimulai dengan meminta pengguna memasukkan jumlah peserta.
2. Untuk setiap peserta, pengguna diminta memasukkan nama dan waktu pengerjaan untuk setiap soal.
3. Data waktu pengerjaan disimpan dalam peta (map) dengan nama peserta sebagai kunci.
4. Program menghitung skor untuk setiap peserta dengan menghitung jumlah soal yang selesai dalam waktu kurang dari 301 detik dan total waktu pengerjaan.
5. Program mencari pemenang berdasarkan jumlah soal yang selesai terbanyak. Jika terdapat beberapa peserta dengan jumlah soal yang sama, peserta dengan total waktu pengerjaan terendah akan dipilih sebagai pemenang.
6. Nama pemenang, jumlah soal yang selesai, dan total waktu pengerjaan ditampilkan ke layar.
   
### 3. Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.
![image](https://github.com/user-attachments/assets/40d94ec4-0290-4ac1-8260-f7442596d4e8)

### Source Code :
```go
//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
)

func cetakDeret(n int) {
	// Selama n belum bernilai 1, lanjutkan mencetak deret
	for n != 1 {
		// Cetak angka saat ini diikuti oleh spasi
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			// Jika angka genap, bagi dengan 2
			n /= 2
		} else {
			// Jika angka ganjil, kalikan dengan 3 dan tambahkan 1
			n = 3*n + 1
		}
	}
	// Terakhir, cetak angka 1
	fmt.Println(n)
}

func main() {
	// Deklarasi variabel untuk masukan
	var nilaiAwal int

	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai suku awal: ")
	fmt.Scan(&nilaiAwal)

	// Panggil fungsi cetakDeret dengan nilai yang dimasukkan pengguna
	cetakDeret(nilaiAwal)
}

```

### Output:
![image](https://github.com/user-attachments/assets/6e2c4107-d2bb-4453-b933-db0c2291b7f0)

### Full code Screenshot:
![image](https://github.com/user-attachments/assets/2f20aa95-f8f6-4122-9cf6-e8f331647f32)

### Deskripsi Program : 
Program ini meminta pengguna untuk memasukkan nilai awal dari sebuah deret. Selanjutnya, program mencetak deret angka berdasarkan aturan berikut: jika angka genap, bagi dengan 2; jika angka ganjil, kalikan dengan 3 dan tambahkan 1. Program terus mencetak angka hingga mencapai 1

### Algoritma Program :
1. Minta pengguna memasukkan nilai awal dari sebuah deret.
2. Baca input pengguna.
3.Selama nilai belum bernilai 1:
  - Cetak nilai saat ini diikuti dengan spasi.
  - Jika nilai genap, bagi dengan 2.
  - Jika nilai ganjil, kalikan dengan 3 dan tambahkan 1.
4. Cetak angka 1.
5. Selesai

### Cara Kerja Program :
1. Program dimulai dan meminta pengguna untuk memasukkan nilai awal dari sebuah deret.
2. Input pengguna dibaca dan disimpan dalam variabel nilaiAwal.
3. Program memanggil fungsi cetakDeret dengan nilaiAwal sebagai argumen.
4. Dalam fungsi cetakDeret, selama nilai belum bernilai 1, program akan:
  - Mencetak nilai saat ini diikuti dengan spasi.
  - Jika nilai genap, program akan membagi nilai tersebut dengan 2.
  - Jika nilai ganjil, program akan mengalikan nilai tersebut dengan 3 dan menambahkannya dengan 1.
5. Program mencetak angka 1 sebagai akhir dari deret.
6. Program selesai.

