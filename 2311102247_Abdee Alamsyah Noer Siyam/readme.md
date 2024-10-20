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
  <strong>PROSEDUR
	</strong>
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
  Abdee Alamsyah Noer Siyam
  <br>
  2311102247
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


## <strong> Dasar Teori </strong>

<strong><h2>PENGERTIAN PROSEDUR</h2></strong>

Prosedur dalam bahasa Go (atau Golang) adalah fungsi yang tidak mengembalikan nilai, sering digunakan ketika kita hanya ingin mengeksekusi sejumlah perintah tanpa memerlukan hasil yang dikembalikan. Konsep prosedur ini sering ditemukan dalam banyak bahasa pemrograman, dan di Go, bentuk ini diimplementasikan melalui fungsi yang tidak menggunakan pernyataan return untuk mengembalikan nilai apa pun. Dengan kata lain, fungsi ini menjalankan aksi atau proses tertentu, tetapi tidak memberikan keluaran balik kepada pemanggil.

Dalam konteks pemrograman, prosedur sangat berguna untuk menyederhanakan kode yang perlu melakukan tugas berulang atau kompleks tanpa memerlukan nilai hasil, seperti mencetak pesan, mengubah status, atau memodifikasi data eksternal (seperti menulis ke file atau database).

**Elemen-elemen dari sebuah Prosedur**

Berikut adalah elemen-elemen dari sebuah prosedur dalam Go:

1. Deklarasi Fungsi: Sama seperti fungsi pada umumnya, prosedur dideklarasikan menggunakan kata kunci func, diikuti dengan nama prosedur, parameter yang diperlukan, dan badan fungsi. Namun, dalam prosedur, tidak ada tipe data hasil pengembalian setelah tanda kurung parameter.

2. Tidak Ada Nilai Return: Karena prosedur tidak mengembalikan nilai, tidak ada tipe data pengembalian yang dideklarasikan, dan tidak ada pernyataan return di dalamnya (meskipun return tanpa nilai bisa digunakan untuk menghentikan eksekusi fungsi).

3. Tujuan Utama: Prosedur umumnya digunakan untuk melakukan tindakan seperti logging, memodifikasi data secara global, atau melakukan operasi input/output yang tidak memerlukan hasil langsung.

<strong>Sintaks Prosedur
</strong>

Dalam Go, prosedur dideklarasikan menggunakan kata kunci func, diikuti oleh nama prosedur, daftar parameter (jika ada). Berikut adalah struktur dasar dari sebuah prosedur:

```go
package main

import "fmt"

// Prosedur ini tidak mengembalikan nilai apapun
func cetakPesan() {
    fmt.Println("Halo, ini adalah prosedur tanpa return!")
}

func main() {
    // Memanggil prosedur
    cetakPesan()
}

```

Penjelasan : 
- cetakPesan: Ini adalah prosedur sederhana yang mencetak pesan "Halo, ini adalah prosedur tanpa return!" ke layar. Tidak ada parameter yang diterima, dan tidak ada nilai yang dikembalikan.

- fmt.Println: Digunakan untuk menampilkan teks ke layar.

- Ketika prosedur dipanggil di dalam main, ia hanya menjalankan instruksi yang ada di dalamnya tanpa memberikan output atau nilai kembali ke main.

Contoh :

```go
package main

import "fmt"

// Prosedur yang menerima parameter batas tetapi tidak mengembalikan nilai
func cetakAngkaSampai(batas int) {
    for i := 1; i <= batas; i++ {
        fmt.Println(i)
    }
}

func main() {
    // Memanggil prosedur dengan batas 5
    cetakAngkaSampai(5)
}
```
Penjelasan :

- cetakAngkaSampai: Prosedur ini menerima satu parameter batas dengan tipe data int dan mencetak bilangan dari 1 hingga angka tersebut. Meskipun prosedur ini melakukan iterasi dan melakukan pencetakan, tidak ada nilai yang dikembalikan ke pemanggil.

- fmt.Println: Dipanggil di dalam loop untuk mencetak setiap angka.

Ketika prosedur ini dipanggil dengan parameter 5 dalam fungsi main, ia mencetak angka dari 1 sampai 5 tanpa memberikan hasil atau output balik.

<strong><h2>Kapan Menggunakan Prosedur</h2></strong>

Prosedur ini cocok untuk digunakan ketika :

- Mencetak output: Fungsi seperti fmt.Println digunakan untuk mencetak pesan ke layar.

- Membuat efek samping: Fungsi dapat mengubah nilai variabel global atau melakukan operasi input/output.

- Memisahkan logika: Membagi kode menjadi fungsi-fungsi kecil membuat kode lebih terstruktur dan mudah dipahami.

Perbedaan dengan Fungsi yang Mengembalikan Nilai:

| Fitur           | Fungsi Tanpa Nilai Kembali                  | Fungsi dengan Nilai Kembali       |
|-----------------|---------------------------------------------|-----------------------------------|
| Tujuan Utama    | Melakukan aksi atau mengubah state          | Menghasilkan output               |
| Nilai Kembali   | Tidak ada                                   | Ada                               |
| Penggunaan      | Mencetak, mengubah variabel global, dll.    | Perhitungan, pengambilan data, dll. |

<strong><h2>Mengapa Menggunakan Prosedur</h2></strong>

Beberapa alasan utama mengapa programmer lebih suka menggunakan prosedur adalah :

- Modularitas: Memisahkan logika program menjadi fungsi-fungsi kecil meningkatkan keterbacaan dan pemeliharaan kode.

- Reusabilitas: Fungsi dapat dipanggil berulang kali dari berbagai bagian program.

- Abstraksi: Menyembunyikan detail implementasi dan hanya mengekspos fungsionalitas yang diperlukan.

```go
func luasPersegi(panjang int, lebar int) int {
    return panjang * lebar
}
```

## <strong> Guided </strong>

### <h2>GUIDED 1</h2>

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

![guided1](https://github.com/user-attachments/assets/5f441d5f-7d9b-4637-9b90-a3943cfc056c)

#### Deskripsi Program : 

Program ini berfungsi untuk menghitung nilai permutasi dari dua bilangan bulat yang diinput oleh pengguna. Permutasi adalah konsep matematika yang menghitung jumlah cara berbeda untuk mengatur sejumlah elemen dengan memperhatikan urutan. Program ini dimulai dengan meminta pengguna untuk memasukkan dua bilangan bulat `a` dan `b`. Setelah menerima input, program memeriksa apakah nilai `a` lebih besar atau sama dengan `b`. Jika kondisi ini terpenuhi, program akan menghitung permutasi `P(a, b)`, yang berarti menghitung permutasi dari `a` elemen diambil sebanyak `b` elemen. Jika `b` lebih besar dari `a`, maka program menghitung permutasi `P(b, a)` untuk memastikan bahwa nilai yang lebih besar selalu digunakan sebagai total elemen, sesuai dengan aturan permutasi di mana `n` (jumlah total elemen) harus lebih besar atau sama dengan `r` (jumlah elemen yang diambil). Fungsi utama program ini adalah `permutasi`, yang menggunakan rumus permutasi:  $P(n, r) = \frac{n!}{(n - r)!}$. Untuk menghitung nilai faktorial `n!`, program memanfaatkan fungsi `faktorial`, yang menghitung hasil perkalian dari semua bilangan bulat positif hingga `n`. Fungsi `faktorial` bekerja dengan menggunakan perulangan `for`, di mana nilai `hasil` terus dikalikan dengan bilangan yang berjalan dari 1 hingga `n`. Kemudian, fungsi `permutasi` memanggil fungsi `faktorial` dua kali: sekali untuk menghitung faktorial dari `n`, dan sekali lagi untuk menghitung faktorial dari `n - r`. Setelah itu, hasil pembagian antara kedua faktorial ini adalah nilai permutasi yang akan dikembalikan oleh fungsi tersebut. Pada akhirnya, program mencetak hasil permutasi ke layar menggunakan fungsi `fmt.Println`. Program ini memastikan bahwa nilai yang lebih besar digunakan sebagai total elemen dalam permutasi, dan pengguna mendapatkan hasil yang valid dari perhitungan permutasi tergantung pada input yang diberikan.

### <h2>GUIDED 2</h2>

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

![guided2](https://github.com/user-attachments/assets/ea3329ef-1c51-48d1-8511-5333c53aea60)

#### Deskripsi Program : 

Program ini digunakan untuk menghitung luas dan keliling sebuah persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna. Program dimulai dengan meminta pengguna untuk memasukkan panjang sisi persegi melalui input. Setelah input diberikan, dua fungsi utama akan dijalankan: fungsi `luas` dan fungsi `keliling`. Fungsi `luas` menerima satu parameter `a` (panjang sisi persegi) dan mengembalikan hasil dari `a * a`, yang merupakan rumus untuk menghitung luas persegi. Fungsi `keliling` juga menerima satu parameter `a` dan menghitung keliling persegi dengan rumus `4 * a`, karena keliling persegi adalah hasil dari empat kali panjang sisinya. Setelah kedua perhitungan selesai, program akan mencetak hasil luas dan keliling persegi ke layar menggunakan `fmt.Println`.

## <strong> Unguided </strong>

### <h2>UNGUIDED 1</h2>

#### Question

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

	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		permutasi(a, c)
		fmt.Print(" ")
		kombinasi(a, c)
		fmt.Println(" ")
		permutasi(b, d)
		fmt.Print(" ")
		kombinasi(b, d)
	} else {
		fmt.Println("Angka yang diberikan tidak sesuai.")
	}
}
```

#### Output

![unguided2](https://github.com/user-attachments/assets/d15493f0-bab1-4646-8e11-7099aa7804fd)


#### Deskripsi Program : 

Program ini dirancang untuk menghitung permutasi dan kombinasi dari dua pasang bilangan yang diinput oleh pengguna. Program dimulai dengan menerima empat bilangan bulat: `a`, `b`, `c`, dan `d`. Setelah input diterima, program melakukan pengecekan apakah nilai `a` lebih besar atau sama dengan `c`, dan `b` lebih besar atau sama dengan `d`. Jika kondisi ini terpenuhi, program menghitung permutasi dan kombinasi untuk pasangan bilangan `(a, c)` dan `(b, d)`. Perhitungan permutasi dilakukan dengan rumus $P(x, y) = \frac{x!}{(x - y)!}$, sedangkan kombinasi dihitung dengan rumus $C(x, y) = \frac{x!}{y!(x - y)!}$. Kedua perhitungan ini menggunakan fungsi `faktorial`, yang mengembalikan hasil perkalian berturut-turut dari angka 1 hingga bilangan input. Hasil dari perhitungan permutasi dan kombinasi untuk setiap pasangan bilangan ditampilkan secara bergantian. Jika kondisi awal tidak terpenuhi, program akan mencetak pesan kesalahan "Angka yang diberikan tidak sesuai" sebagai penanda bahwa input tidak valid untuk perhitungan.

### <h2>UNGUIDED 2</h2>

#### Question

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
	fmt.Println(namaPemenang, soalMaksimum, skorMinimum)
}	

func main() {
	selesai := false
	dataPeserta := make(map[string][]int)

	if !selesai {
		for {
			waktu := [8]int{}
			var namaPeserta string

			fmt.Scan(&namaPeserta)

			if namaPeserta == "selesai" {
        selesai = true
        break
      }
			for i:= 0; i < 8; i++ {
				fmt.Scan(&waktu[i])
			}
			dataPeserta[namaPeserta] = waktu[:]
		}
	}

	hitungSkor(dataPeserta)
}
```

#### Output

![unguided2](https://github.com/user-attachments/assets/874e7cce-aaf4-4f5a-8b6d-e8ce6934871d)

Deskripsi Program :

Program ini bertujuan untuk menghitung dan menentukan pemenang dari suatu kompetisi berdasarkan waktu penyelesaian soal oleh peserta. Setiap peserta mengerjakan delapan soal, dan waktu yang dihabiskan untuk setiap soal dicatat. Program ini akan mencetak nama pemenang, jumlah soal yang berhasil diselesaikan, dan total waktu yang dihabiskan untuk menyelesaikan soal-soal tersebut. Program dimulai dengan mendeklarasikan fungsi `kalkulasiSkor`, yang menerima slice dari integer `waktu` dan mengembalikan dua nilai: jumlah soal yang berhasil diselesaikan dalam waktu kurang dari 301 detik dan total waktu yang dihabiskan untuk soal-soal tersebut. Dalam fungsi ini, setiap waktu yang kurang dari 301 akan dihitung, dan jika memenuhi syarat, `jumlahSelesai` dan `totalWaktu` akan diperbarui. Fungsi kedua, `hitungSkor`, menerima sebuah peta `dataPeserta`, yang berisi nama peserta dan waktu yang mereka habiskan untuk menyelesaikan soal. Fungsi ini akan menentukan pemenang berdasarkan jumlah soal yang diselesaikan dan total waktu yang dihabiskan. Pemenang akan ditentukan dengan membandingkan jumlah soal yang berhasil diselesaikan; jika ada peserta dengan jumlah soal yang sama, pemenang ditentukan berdasarkan waktu penyelesaian yang lebih sedikit. Dalam fungsi `main`, program menginisialisasi map `dataPeserta` untuk menyimpan waktu penyelesaian setiap peserta. Program akan terus meminta input nama peserta dan waktu untuk setiap soal hingga nama peserta yang dimasukkan adalah "selesai". Setelah menerima semua input, program akan memanggil fungsi `hitungSkor` untuk menentukan dan mencetak pemenang bersama dengan jumlah soal yang diselesaikan dan total waktu yang dihabiskan. 

### <h2>UNGUIDED 3</h2>

#### Question

Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah ½n, tetapi jika ganjil maka suku berikutnya bernilai 3n + 1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah:

**22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1**

Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1.

Buat program Skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.

```bash
prosedure cetakDeret(in n : integer)
```
**Masukan** berupa satu bilangan integer positif yang lebih kecil dari 1000000

**Keluaran** terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dipisahkan oleh sebuah spasi

| Masukan | Keluaran 																	 |
|---------|--------------------------------------------|
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
	fmt.Scan(&n)

	if(n < 1000000){
		cetakDeret(n)
	}else {
		fmt.Println("Inputan harus antara 1 sampai 1,000,000!")
	}
}
```

#### Output

![unguided3](https://github.com/user-attachments/assets/8225442b-7b8c-4e5e-8d97-1960910f3dc1)

Deskripsi Program :

Program ini dirancang untuk menghitung dan mencetak deret berdasarkan algoritma Collatz, yang juga dikenal sebagai "masalah 3n + 1." Deret dimulai dari bilangan bulat positif `n` yang dimasukkan oleh pengguna. Algoritma ini mengikuti aturan di mana, jika `n` adalah genap, maka nilainya dibagi dua; jika `n` adalah ganjil, maka dihitung sebagai `3n + 1`. Proses ini berlanjut hingga nilai `n` menjadi 1, dengan setiap nilai yang dihasilkan dicetak ke layar. Fungsi utama program, `cetakDeret`, menerima parameter `n` dan mencetak nilai `n` dalam loop selama `n` tidak sama dengan 1. Di dalam fungsi ini, terdapat pengecekan untuk menentukan apakah `n` genap atau ganjil dan melakukan perhitungan sesuai dengan aturan tersebut. Fungsi `main` meminta pengguna untuk memasukkan bilangan bulat `n` dan memeriksa apakah nilai tersebut kurang dari 1.000.000. Jika ya, maka fungsi `cetakDeret` dipanggil untuk mencetak deret; jika tidak, program akan mencetak pesan kesalahan yang menyatakan bahwa input harus berada dalam rentang 1 sampai 1.000.000. Dengan demikian, program ini tidak hanya menampilkan deret Collatz, tetapi juga memastikan bahwa input yang diberikan sesuai dengan batasan yang telah ditetapkan.

## <strong> Kesimpulan </strong>

Kesimpulan mengenai prosedur dalam bahasa pemrograman Go (Golang) adalah bahwa prosedur, yang dikenal sebagai fungsi tanpa nilai kembalian, merupakan alat penting untuk menyusun dan mengorganisir kode dalam aplikasi Go. Prosedur digunakan untuk menjalankan serangkaian perintah atau logika tertentu tanpa mengembalikan nilai, sehingga memudahkan pengelolaan alur program dan meningkatkan keterbacaan kode. Prosedur dapat menerima argumen, yang memungkinkan penggunaannya untuk melakukan operasi berdasarkan data yang diberikan. Dengan mendefinisikan prosedur, pengembang dapat mengurangi pengulangan kode, meningkatkan modularitas, dan memfasilitasi pemeliharaan kode. Penggunaan prosedur juga mendukung prinsip pemrograman terstruktur, di mana program dibangun dari blok-blok yang lebih kecil dan dapat digunakan kembali. Secara keseluruhan, prosedur dalam Golang memberikan fleksibilitas dan efisiensi dalam pengembangan aplikasi, serta membantu dalam mengorganisir logika program dengan cara yang lebih teratur dan terstruktur.

## <strong> Referensi </strong>

#### [1] Sanjaya, Rizqi Hafian. [2024]. Cara Menggunakan Method dan Function | Tutorial Golang Bahasa Indonesia. Diakses melalui website https://buildwithangga.com/tips/cara-menggunakan-method-dan-function-tutorial-golang-bahasa-indonesia 

#### [2] Revou. Apa itu Function dalam Pemrograman?. Diakses melalui website https://revou.co/kosakata/function

#### [3] Yahya, Adin. [2023]. AFunction dan Method Bahasa Pemrograman Golang. Diakses melalui website https://adinyahya.com/function-dan-method-bahasa-pemrograman-golang/
