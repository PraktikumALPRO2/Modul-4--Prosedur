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

```

### Output:


### Full code Screenshot:


### Deskripsi Program : 


### Algoritma Program


### Cara Kerja Program:



### 2. Buatlah program beserta prosedur untuk menghitung luas persegi dan keliling persegi

### Source Code :
```go

```

### Output:


### Full code Screenshot:

### Deskripsi Program : 


### Algoritma Program

### Cara Kerja Program:



## Unguided 

### 1. Program untuk menghitung Faktorial, Permutasi dan Kombinasi

### Source Code :
```go


```

### Output:


### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program


### Cara Kerja Program:


### 2. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan
**Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.**

### Source Code :
```go


```

### Output:

### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program:

### 3. Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.

### Source Code :
```go

```

### Output:


### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program:

