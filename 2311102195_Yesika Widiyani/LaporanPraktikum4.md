[Dondar Poster.pdf](https://github.com/user-attachments/files/17437057/Dondar.Poster.pdf)[23 OKTOBER 2024.pdf](https://github.com/user-attachments/files/17437051/23.OKTOBER.2024.pdf)<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>
<h2 align="center"><strong>MODUL IV</strong></h2>
<h2 align="center"><strong> PROSEDUR </strong></h2>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Yesika Widiyani<br>
  2311102195<br>
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

# DASAR TEORI
## Prosedur
Dalam bahasa pemrograman Go, prosedur diimplementasikan melalui fungsi (functions). Prosedur ini merupakan sekumpulan instruksi yang dikemas dalam sebuah fungsi untuk menyelesaikan suatu tugas tertentu. Dengan menggunakan prosedur, kode program bisa lebih terstruktur, modular, dan mudah dipelihara.

### **1. Deklarasi Fungsi sebagai Prosedur**
Fungsi di Go mendefinisikan prosedur dengan sintaks yang sederhana. Deklarasi fungsi dimulai dengan kata kunci `func`, diikuti oleh nama fungsi, parameter (jika ada), dan tipe nilai yang dikembalikan (opsional).

#### **Sintaks Dasar:**
```go
func namaFungsi(parameter1 tipe, parameter2 tipe) tipeKembalian {
    // blok kode yang berisi instruksi prosedur
    return nilaiKembalian
}
```

- **`func`**: Kata kunci untuk mendefinisikan sebuah fungsi.
- **`namaFungsi`**: Nama dari prosedur yang digunakan untuk memanggilnya.
- **`parameter1`, `parameter2`**: Daftar parameter yang diterima oleh prosedur, masing-masing memiliki tipe data yang ditentukan.
- **`tipeKembalian`**: Tipe data dari nilai yang dikembalikan oleh prosedur (opsional).
- **`return`**: Perintah untuk mengembalikan nilai dari prosedur (jika ada).

### **2. Prosedur Tanpa Pengembalian Nilai**
Prosedur dalam Go dapat berupa fungsi yang tidak mengembalikan nilai (void). Prosedur ini hanya mengeksekusi sejumlah instruksi tanpa memberikan hasil apapun.

#### **Contoh:**
```go
func cetakPesan(pesan string) {
    fmt.Println(pesan)
}
```
Pada contoh di atas, prosedur `cetakPesan` menerima satu parameter `pesan` bertipe `string` dan mencetaknya tanpa mengembalikan nilai.

### **3. Prosedur dengan Pengembalian Nilai**
Prosedur dapat mengembalikan satu atau lebih nilai. Ini berguna untuk melakukan operasi dan mengembalikan hasilnya.

#### **Contoh Prosedur Mengembalikan Satu Nilai:**
```go
func Tambah(a int, b int) int {
    return a + b
}
```
Fungsi `Tambah` merupakan prosedur yang menerima dua parameter bertipe `int` dan mengembalikan hasil penjumlahan keduanya.

#### **Contoh Prosedur Mengembalikan Banyak Nilai:**
```go
func hitung(a int, b int) (int, int) {
    penjumlahan := a + b
    perkalian := a * b
    return penjumlahan, perkalian
}
```
Pada contoh ini, prosedur `hitung` mengembalikan dua nilai: hasil penjumlahan dan perkalian dari dua angka yang diterima sebagai parameter.

### **4. Prosedur Anonim**
Go mendukung pembuatan prosedur tanpa nama, yang disebut **fungsi anonim**. Fungsi anonim dapat digunakan sebagai parameter untuk fungsi lain atau dijalankan secara langsung.

#### **Contoh Fungsi Anonim:**
```go
func() {
    fmt.Println("Ini adalah prosedur anonim")
}()
```
Fungsi anonim ini langsung dijalankan setelah didefinisikan karena diakhiri dengan `()`.

### **5. Prosedur dan Goroutine**
Salah satu fitur utama dalam Go adalah kemampuannya untuk menjalankan prosedur secara bersamaan (concurrency) menggunakan **goroutines**. Dengan menggunakan goroutines, prosedur dapat dijalankan di thread yang terpisah, yang membuat aplikasi lebih efisien dalam penggunaan sumber daya.

#### **Contoh Penggunaan Goroutine dengan Prosedur:**
```go
func cetakPesan(pesan string) {
    fmt.Println(pesan)
}

func main() {
    go cetakPesan("Pesan 1")
    go cetakPesan("Pesan 2")
    time.Sleep(time.Second)
}
```
Dua goroutine dipanggil untuk menjalankan prosedur `cetakPesan` secara paralel.

### **6. Prosedur dengan Error Handling**
Go memiliki sistem penanganan error yang eksplisit, di mana sebuah prosedur dapat mengembalikan nilai error bersama dengan nilai lain untuk memberi tahu jika terjadi kesalahan dalam eksekusi.

#### **Contoh Prosedur dengan Pengembalian Error:**
```go
func bagi(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("pembagian dengan nol")
    }
    return a / b, nil
}
```
Pada contoh di atas, prosedur `bagi` mengembalikan dua nilai: hasil pembagian dan error jika terjadi kesalahan (seperti pembagian dengan nol).

### **Kesimpulan**
Prosedur dalam bahasa Go adalah fungsi yang digunakan untuk mengeksekusi sekelompok instruksi dengan tujuan tertentu. Go menyediakan banyak fleksibilitas dalam hal penanganan pengembalian nilai, error, serta mendukung concurrency dengan goroutines. Prosedur dapat dibuat tanpa mengembalikan nilai, mengembalikan beberapa nilai, dan bisa dijalankan secara paralel untuk efisiensi.

Prosedur dalam Go sangat penting untuk membuat kode yang modular dan terstruktur, serta sangat bermanfaat dalam membangun aplikasi skala besar yang cepat dan dapat diandalkan.

------
# GUIDED
## ## Guided - 1 - 3.1
### Study Case
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap
𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑.

Catatan: permutasi (P) dan kombinasi (C) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan menggunakan persamaan berikut!
<br>
![soalgu1](https://github.com/user-attachments/assets/2e18508b-42ec-48f9-888e-0f4d4cc674c4)

### Source Code
![Guided1](https://github.com/user-attachments/assets/03a6d7d4-8ad7-4eef-a1b9-9c4aa9833166)

### Screenshot Output
![ssoug1](https://github.com/user-attachments/assets/7b21d230-809f-4bf9-a53c-fac54088d392)

### Deskripsi Program
Program di atas dibuat untuk menghitung **permutasi** dari dua bilangan `a` dan `b`. Permutasi adalah operasi yang menghitung banyaknya cara menyusun `r` objek dari `n` objek yang tersedia, di mana urutan sangat diperhatikan. Program ini meminta input dua bilangan (`a` dan `b`) dari pengguna dan kemudian menghitung permutasi dengan formula:

P(n,r) = n!/(n-r)!

di mana `n!` adalah faktorial dari `n`, yaitu hasil perkalian dari semua bilangan bulat positif dari 1 hingga `n`.

#### Algoritma Program

1. **Input dari Pengguna:**
   - Program pertama kali meminta dua bilangan `a` dan `b` dari pengguna.
   - Jika `a >= b`, program akan menghitung permutasi dengan `a` sebagai `n` dan `b` sebagai `r`. Jika tidak, posisi `a` dan `b` akan ditukar sehingga `b` dianggap sebagai `n` dan `a` sebagai `r`.

2. **Perhitungan Faktorial:**
   - Fungsi `faktorial(n int) int` menghitung faktorial dari `n`. 
   - Fungsi ini menggunakan sebuah loop dari 1 hingga `n`, di mana setiap iterasi mengalikan nilai sebelumnya dengan `i` untuk menghitung faktorial.

3. **Perhitungan Permutasi:**
   - Fungsi `permutasi(n, r int) int` menghitung permutasi menggunakan rumus : Permutasi = faktorial(n)/faktorial(n-r).
   - Fungsi ini menggunakan nilai faktorial dari `n` dan `n - r`, lalu membaginya untuk mendapatkan hasil permutasi.

4. **Output:**
   - Hasil dari perhitungan permutasi ditampilkan ke layar menggunakan `fmt.Println`.

#### Cara Kerja Program

1. Program dimulai dengan mendeklarasikan dua variabel `a` dan `b` sebagai bilangan bulat (`int`).
2. Program meminta input dua bilangan dari pengguna dengan menggunakan `fmt.Scan(&a, &b)`.
3. Program mengecek apakah nilai `a` lebih besar atau sama dengan `b`. Jika ya, permutasi dihitung dengan `a` sebagai `n` dan `b` sebagai `r`. Jika tidak, posisi `a` dan `b` ditukar.
4. Fungsi `permutasi` dipanggil dengan parameter `n` dan `r`, dan hasilnya dihitung dengan membagi faktorial `n` dengan faktorial `(n - r)`.
5. Setelah hasil permutasi diperoleh, hasil tersebut dicetak ke layar.

## Guided - 2
### Study Case
Membuat program dengan bahasa go untuk mencari sebuah Luas dan Keliling Persegi.

### Source Code
![guided2](https://github.com/user-attachments/assets/4797d404-e886-47a3-b7fe-b54314249c4b)

### Screenshot Output
![Uploading Dondar Poster.png…]()

### Deskripsi Program

------

# UNGUIDED
## Unguided - 1
### Study Case
### Source Code
### Deskripsi Program

## Unguided - 2
### Study Case
### Source Code
### Deskripsi Program

# Unguided - 3
### Study Case
### Source Code
### Deskripsi Program
