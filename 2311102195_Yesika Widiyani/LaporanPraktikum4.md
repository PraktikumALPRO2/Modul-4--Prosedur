<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
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
## Guided - 1
### Study Case
Buatlah sebuah program beserta prosedur yang digunakan untuk menghitung nilai faktorial dan permutasi.**

*Masukan : Terdiri dari dua buah bilangan positif a dan b*

*Keluaran : Berupa sebuah bilangan bulat yang menyatakan nilai a permutasi b apabila >= b atau b permutasi a untuk kemungkinan yang lain*

### Source Code
![g1](https://github.com/user-attachments/assets/a2ad0f31-1a99-43ce-a305-72f3a7605ea3)

### Screenshot Code
![image](https://github.com/user-attachments/assets/706c544f-3caa-431e-a0fb-30116dcbae5e)

### Deskripsi Program
Program ini digunakan untuk menghitung permutasi dari dua bilangan bulat yang dimasukkan oleh pengguna. Permutasi adalah cara untuk menghitung berapa banyak cara yang mungkin untuk menyusun sejumlah elemen yang dipilih dari himpunan yang lebih besar.

**Algoritma Program**
Inisialisasi variabel: Program mendeklarasikan dua variabel a dan b untuk menyimpan bilangan bulat yang dimasukkan oleh pengguna.
Input: Program meminta pengguna memasukkan dua angka dengan memanggil prosedur bacaInput.
Perbandingan: Program membandingkan nilai a dan b. Jika a >= b, maka program akan menghitung permutasi dari a sebagai n dan b sebagai r. Sebaliknya, jika b > a, program menghitung permutasi dari b sebagai n dan a sebagai r.
Perhitungan Faktorial: Program menghitung faktorial dari n dan n - r menggunakan prosedur faktorial.
Perhitungan Permutasi: Program menghitung permutasi dengan memanggil prosedur hitungPermutasi dan menampilkan hasilnya.
Output: Program mencetak hasil perhitungan permutasi ke layar.

**Cara Kerja Program:**
Program dimulai dengan memanggil prosedur bacaInput untuk membaca input pengguna (dua bilangan a dan b).
Program kemudian menentukan apakah a >= b atau sebaliknya melalui prosedur permutasi.
Setelah itu, program memanggil hitungPermutasi untuk menghitung permutasi dan mencetak hasilnya ke layar.

## Guided - 2
### Study Case
Membuat program dengan bahasa go untuk mencari sebuah Luas dan Keliling Persegi.

### Source Code
![guided2](https://github.com/user-attachments/assets/4797d404-e886-47a3-b7fe-b54314249c4b)

### Screenshot Output
![ssoug2](https://github.com/user-attachments/assets/24f8024d-1559-4c5b-8073-e127a5f39fe5)

### Deskripsi Program
Program ini digunakan untuk menghitung luas dan keliling dari sebuah persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna. Program ini telah diubah ke dalam bentuk prosedur, sehingga tiap bagian seperti input, perhitungan luas, dan perhitungan keliling dikelola oleh prosedur yang terpisah. Ini membantu membuat kode lebih modular dan mudah dikelola.

**Algoritma Program:**
1. Program pertama-tama meminta input dari pengguna berupa panjang sisi persegi. Input ini disimpan dalam variabel `sisi` yang bertipe `float64` untuk mendukung angka desimal.
2. Setelah input diterima, program menghitung **luas** persegi dengan rumus:
   Luas = sisi x sisi
3. Kemudian, program menghitung **keliling** persegi dengan rumus:
  Keliling = 4 x sisi
4. Setelah perhitungan selesai, hasil **luas** dan **keliling** ditampilkan ke layar dengan format dua angka desimal menggunakan fungsi `fmt.Printf`.

**Cara Kerja Program:**
1. Program dimulai dengan mendeklarasikan variabel `sisi` yang akan menyimpan input dari pengguna. 
2. Program menampilkan pesan "Masukkan panjang sisi persegi:" untuk meminta pengguna memasukkan nilai panjang sisi persegi. Fungsi `fmt.Scan(&sisi)` digunakan untuk membaca input dari pengguna.
3. Setelah nilai sisi dimasukkan, program menghitung luas persegi dengan mengalikan nilai `sisi` dengan dirinya sendiri (sisi * sisi). Hasil perhitungan ini disimpan dalam variabel `luas`.
4. Program juga menghitung keliling persegi dengan mengalikan nilai `sisi` dengan 4 (4 * sisi). Hasilnya disimpan dalam variabel `keliling`.
5. Setelah perhitungan selesai, program menampilkan hasil perhitungan luas dan keliling menggunakan fungsi `fmt.Printf` untuk memformat hasil ke dua angka di belakang koma.
6. Output menampilkan nilai luas dan keliling yang sudah dihitung berdasarkan input panjang sisi yang diberikan oleh pengguna.

------

# UNGUIDED
## Unguided - 1
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

## Unguided - 2
### Study Case
Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.
![image](https://github.com/user-attachments/assets/c92445af-a30b-46d8-b6f8-4003870afefc)

Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit).
Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.

Keterangan:

Astuti menyelesaikan 6 soal dalam waktu 287 menit, sedangkan Bertha 7 soal dalam waktu 294 menit. Karena Bertha menyelesaikan lebih banyak, maka Bertha menang. Jika keduanya menyelesaikan sama banyak, maka pemenang adalah yang menyelesaikan dengan total waktu paling kecil.

### Source Code
![un2](https://github.com/user-attachments/assets/9b29dc67-8041-4e7e-b29c-06f4f3cc77af)

### Screenshot Output
![image](https://github.com/user-attachments/assets/02726aab-5961-4778-9a93-9163259dcdfc)

### Deskripsi Program
Program ini dirancang untuk menentukan pemenang kompetisi pemrograman tingkat nasional. Setiap peserta diberikan 8 soal yang harus diselesaikan dalam waktu maksimum 5 jam (300 menit). Peserta yang berhasil menyelesaikan jumlah soal terbanyak dalam waktu terpendek adalah pemenangnya. Program ini menggunakan prosedur `hitungSkor` untuk menghitung jumlah soal yang diselesaikan dan total waktu yang dibutuhkan setiap peserta. Input nama peserta dan waktu pengerjaan dilakukan dalam fungsi utama, sementara hasil perhitungan ditampilkan setelah semua peserta diinput.

### **Algoritma Program**

1. **Inisialisasi Variabel**:
   - Deklarasikan variabel untuk menyimpan nama pemenang, jumlah soal yang diselesaikan oleh pemenang, dan total waktu pemenang.
   
2. **Menerima Input**:
   - Dalam loop, terus terima input nama peserta dan waktu pengerjaan hingga pengguna memasukkan kata "Selesai".
   - Jika input nama peserta adalah "Selesai", keluar dari loop.

3. **Membaca Waktu Pengerjaan**:
   - Untuk setiap peserta, baca 8 angka yang merepresentasikan waktu pengerjaan masing-masing soal.
   
4. **Menghitung Skor**:
   - Panggil prosedur `hitungSkor`, yang menerima waktu pengerjaan dan mengembalikan jumlah soal yang diselesaikan dan total waktu yang dibutuhkan.
   
5. **Menentukan Pemenang**:
   - Bandingkan jumlah soal yang diselesaikan dengan pemenang sebelumnya. Jika peserta saat ini menyelesaikan lebih banyak soal, atau jika sama banyak tetapi dengan waktu lebih singkat, maka peserta ini menjadi pemenang sementara.
   
6. **Output**:
   - Setelah semua peserta dimasukkan, tampilkan nama pemenang, jumlah soal yang diselesaikan, dan total waktu yang digunakan.

### **Cara Kerja Program**

1. **Input Peserta**:
   - Program meminta pengguna untuk memasukkan nama peserta diikuti oleh waktu pengerjaan 8 soal. Peserta dapat menginput waktu dalam menit. Jika seorang peserta tidak menyelesaikan suatu soal, waktu yang digunakan dianggap 301 menit.

2. **Memanggil Prosedur `hitungSkor`**:
   - Prosedur ini menganalisis waktu yang diinput. Setiap waktu yang lebih dari 300 menit diabaikan (dianggap tidak menyelesaikan soal), dan hanya soal yang diselesaikan dalam 300 menit atau kurang yang dihitung. Total waktu dari soal-soal yang diselesaikan dijumlahkan.

3. **Menentukan Pemenang**:
   - Program memeriksa apakah jumlah soal yang diselesaikan oleh peserta saat ini lebih banyak dari pemenang sebelumnya. Jika sama, program membandingkan total waktu untuk menentukan pemenang. Ini dilakukan hingga semua peserta diinput.

4. **Menampilkan Hasil**:
   - Setelah pengguna memasukkan "Selesai", program mencetak nama pemenang, jumlah soal yang diselesaikan, dan total waktu yang digunakan.

# Unguided - 3
### Study Case
3)	Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah ½n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah:
4)	
22 11  34  17  52  26  13  40  20  10  5  16  8  4  2  1
 
Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1.
Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.

![image](https://github.com/user-attachments/assets/3ec0d166-c7a7-4d35-8da1-35927869e010)

Masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000.

Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dan dipisahkan oleh sebuah spasi.

### Source Code
![un3](https://github.com/user-attachments/assets/08047a25-08a6-4bae-8938-aaec7565b0b9)

### Screenshot Output
![image](https://github.com/user-attachments/assets/5dbe1a6f-97c5-4724-b507-77a0af763ecf)

### Deskripsi Program
Program ini dirancang untuk menghasilkan dan mencetak deret bilangan berdasarkan aturan yang didefinisikan dalam masalah. Deret dimulai dengan bilangan bulat positif \( n \). Jika \( n \) genap, suku berikutnya adalah \( \frac{n}{2} \); jika ganjil, suku berikutnya adalah \( 3n + 1 \). Proses ini terus berlangsung hingga suku terakhir mencapai nilai 1. Program ini menggunakan prosedur `cetakDeret` untuk menghitung dan mencetak deret berdasarkan nilai awal yang diberikan oleh pengguna.

#### **Algoritma Program**
1. **Inisialisasi**:
   - Mulai dengan meminta input dari pengguna untuk bilangan bulat positif \( n \) (kurang dari 1.000.000).

2. **Validasi Input**:
   - Periksa apakah \( n \) valid (harus lebih besar dari 0 dan kurang dari 1.000.000). Jika tidak valid, tampilkan pesan kesalahan dan akhiri program.

3. **Prosedur `cetakDeret`**:
   - Menerima parameter \( n \).
   - Dalam loop, lakukan hal berikut:
     - Cetak nilai \( n \).
     - Jika \( n \) genap, ubah \( n \) menjadi \( \frac{n}{2} \).
     - Jika \( n \) ganjil, ubah \( n \) menjadi \( 3n + 1 \).
   - Akhiri loop ketika \( n \) menjadi 1, dan cetak 1 sebagai suku terakhir.

4. **Output**:
   - Cetak deret bilangan yang dihasilkan di baris yang sama, dipisahkan oleh spasi.

#### **Cara Kerja Program**
1. **Input**:
   - Program meminta pengguna untuk memasukkan bilangan bulat positif.
   - Pengguna memasukkan nilai, yang disimpan dalam variabel \( n \).

2. **Validasi**:
   - Program memeriksa nilai \( n \):
     - Jika \( n \) lebih kecil atau sama dengan 0 atau lebih besar atau sama dengan 1.000.000, program akan menampilkan pesan kesalahan dan menghentikan eksekusi.

3. **Menghasilkan Deret**:
   - Program memanggil prosedur `cetakDeret` dengan nilai \( n \):
     - Di dalam prosedur, selama nilai \( n \) tidak sama dengan 1, program mencetak nilai \( n \) dan memperbarui nilainya sesuai dengan aturan (genap atau ganjil).
     - Setelah mencapai 1, program juga mencetak angka 1 sebagai suku terakhir.

4. **Output**:
   - Hasil deret bilangan ditampilkan di layar sebagai output akhir program, di mana semua suku dipisahkan oleh spasi.
