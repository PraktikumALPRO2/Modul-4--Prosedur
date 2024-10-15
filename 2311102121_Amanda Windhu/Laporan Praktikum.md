# <h2 align="center">LAPORAN PRAKTIKUM</h2>
# <h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>
# <h2 align="center">MODUL 4</h2>
# <h2 align="center">PROSEDUR</h2>
<p align="center">
    <img src="https://github.com/user-attachments/assets/3ccfed0b-72d1-4349-ac08-c4dce379c827" alt="Gambar">
</p>
 <h3  align="center" >Disusun Oleh : </h3>
<p align="center">Amanda Windhu Gustyas - 2311102121</p>
<p align="center">IF-11-05</p>
 <h3 <p align="center" >Dosen Pengampu : </h3> </p>
 <p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>

# <h3 align="center"> PROGRAM STUDI S1 TEKNIK INFORMATIKA </h3>
# <h3 align="center"> FAKULTAS INFORMATIKA </h3>
# <h3 align="center"> TELKOM UNIVERSITY PURWOKERTO </h3>
# <h3 align="center"> 2024 </h3>

## I. DASAR TEORI



## II. GUIDED

### 1. Contoh Program dengan Function

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b) // Membaca input dari pengguna untuk dua nilai integer a dan b
	if a >= b { 
		// Jika a lebih besar atau sama dengan b, panggil prosedur permutasi dengan parameter (a, b)
		permutasi(a, b)
	} else { 
		// Jika b lebih besar dari a, panggil prosedur permutasi dengan parameter (b, a)
		permutasi(b, a)
	}
}

func faktorial(n int) int {
	var hasil int = 1
	// Loop untuk menghitung faktorial dari n
	for i := 1; i <= n; i++ {
		hasil *= i // Mengalikan hasil dengan nilai i pada setiap iterasi
	}
	return hasil // Mengembalikan hasil faktorial
}

func permutasi(n, r int) {
	// Menghitung permutasi nPr dan langsung mencetak hasilnya
	hasil := faktorial(n) / faktorial(n-r)
	fmt.Println(hasil) // Mencetak hasil permutasi
}
```
## Output: ![image](https://github.com/user-attachments/assets/6ed035eb-545b-4835-a4f0-ec0d64af7862)

Kode di atas

### 2. Menghitung Luas dan Keliling Persegi

```go
package main

import "fmt"

// Prosedur untuk menghitung dan mencetak luas persegi
func hitungLuas(sisi float64) {
    luas := sisi * sisi
    fmt.Printf("Luas persegi: %.2f\n", luas) // Mencetak hasil luas
}

// Prosedur untuk menghitung dan mencetak keliling persegi
func hitungKeliling(sisi float64) {
    keliling := 4 * sisi
    fmt.Printf("Keliling persegi: %.2f\n", keliling) // Mencetak hasil keliling
}

func main() {
    var sisi float64

    fmt.Print("Masukkan panjang sisi persegi: ")
    fmt.Scan(&sisi)

    hitungLuas(sisi)      // Memanggil prosedur hitungLuas
    hitungKeliling(sisi)  // Memanggil prosedur hitungKeliling
}
```
## Output: ![image](https://github.com/user-attachments/assets/d33e1c01-7b2e-48ab-926c-58807500733d)

Kode di atas

## III. UNGUIDED

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)<br/> Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. <br/> Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. <br/> Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!<br/> ![image](https://github.com/user-attachments/assets/5b90c7e3-9f76-45eb-bf14-8f1bca637918) <br/> Selesaikan program tersebut dengan memanfaatkan procedure yang diberikan berikut ini!<br/>![image](https://github.com/user-attachments/assets/d7a28bc4-25bd-4c1d-9091-e058e26a1407)


### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. <br/> Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.<br/> ![image](https://github.com/user-attachments/assets/b8f89dce-7575-4a33-ad7a-61ece8188938)<br/> Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban, maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). <br/> Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.<br/>![image](https://github.com/user-attachments/assets/3117bbb6-fc8f-45e5-9027-890dea0cb6dc)<br/>Keterangan:<br/> Astuti menyelesaikan 6 soal dalam waktu 287 menit, sedangkan Bertha 7 soal dalam waktu 294 menit. Karena Bertha menyelesaikan lebih banyak, maka Bertha menang. Jika keduanya menyelesaikan sama banyak, maka pemenang adalah yang menyelesaikan dengan total waktu paling kecil.<br/>











### KESIMPULAN
### REFERENSI
