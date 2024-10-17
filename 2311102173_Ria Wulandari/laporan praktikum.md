
<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL III</strong></h2>
<h2 align="center"><strong> PROSEDUR </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Ria Wulandari / 2311102173<br>
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
#### 1. Definisi Prosedur
Dalam konteks pemrograman, prosedur atau fungsi adalah blok kode yang dirancang untuk melakukan tugas tertentu. Prosedur dapat menerima input (parameter), melakukan operasi, dan dapat mengembalikan nilai. Fungsi membantu dalam organisasi kode, membuatnya lebih terstruktur, dan memudahkan pemeliharaan serta penggunaan kembali.
#### 2. Deklarasi Prosedur
Deklarasi fungsi di Go dilakukan menggunakan kata kunci `func`, diikuti oleh nama fungsi, parameter (jika ada), tipe pengembalian (jika ada), dan blok kode yang berisi logika fungsi tersebut. Berikut adalah sintaks umum untuk mendeklarasikan fungsi:
```go
func namaFungsi(parameter1 tipe1, parameter2 tipe2, ...) returnType {
    // blok kode
    return nilai
}
```
##### Contoh deklarasi fungsi
Berikut adalah contoh deklarasi fungsi sederhana yang menjumlahkan dua bilangan:
```go
func tambah(a int, b int) int {
    return a + b
}
```
#### 3. Cara Pemanggilan Prosedur
Untuk memanggil fungsi yang telah dideklarasikan, Anda cukup menyebutkan nama fungsi dan memberikan argumen yang sesuai dengan parameter yang telah ditentukan. Berikut adalah contoh cara memanggil fungsi:
```go
func main() {
    hasil := tambah(3, 5) // Memanggil fungsi tambah dengan argumen 3 dan 5
    fmt.Println("Hasil penjumlahan:", hasil)
}
```
#### 4. Parameter Fungsi
Parameter dalam fungsi Go dapat dibagi menjadi dua kategori: parameter formal dan parameter aktual.
##### a.Parameter Formal
Parameter formal adalah parameter yang didefinisikan dalam deklarasi fungsi. Parameter ini bertindak sebagai variabel yang akan menerima nilai saat fungsi dipanggil. Dalam contoh fungsi `tambah`, `a` dan `b` adalah parameter formal.
```go
func tambah(a int, b int) int {
    return a + b
}
```
##### b.  Parameter Aktual
Parameter aktual adalah nilai yang sebenarnya diberikan kepada parameter formal saat fungsi dipanggil. Dalam pemanggilan fungsi `tambah(3, 5)`, nilai `3` dan `5` adalah parameter aktual yang diberikan kepada fungsi.
```go
func main() {
    hasil := tambah(3, 5) // 3 dan 5 adalah parameter aktual
}
```

## Guided
1. 
