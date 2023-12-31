# Test Stamps Indonesia

Berikut merupakan hasil test initial screening untuk proses pendafataran di Stamps Indonesia.

Repository milik Muhammad Wyndham Haryata Permana.

## Isi Program
Program di repository ini bertujuan untuk menjawab 2 pertanyaan yang di lampirkan di email. Jawaban dari kedua pertanyaan tersebut saya rangkum menjadi 1 buat program yang dapat dipilih ketika menjalankan programnya.

## Screenshot Program
Screenshot hasil pengerjaan program FooBar

<img src="image-4.png" width="550">

Screenshot hasil pengerjaan program cuaca Jakarta 5 hari kedepan

<img src="image-5.png" width="550">

## Limitasi Program
### Program Foo Bar
Limitasi:
- Hanya melihat array dari 1 s/d 100
- Tidak melihat bilangan imajiner

### Program Open Weather Map
Limitasi:
- Waktu cuaca yagn diambil adalah cuaca pukul 13.00 (jam 1 siang).
- Limitasi API yang hanya bisa melihat forecast menyebabkan jam eksekusi program mempengaruhi hasil.
- Apabila program dieksekusi sebelum jam 13.00 maka cuaca hari ini akan tercantum.
- Apabila program dieksekusi setelah jam 13.00 maka cuaca akan dimulai dari hari besoknya.

## Sebelum Menjalankan Program
Sebelum menjalankan program, diperlukan 1 variable berupa `API Key` untuk https://openweathermap.org. Demi alasan keamanan, saya tidak akan mencancumkan `API Key` saya dalam repository ini. Untuk mengatur `API Key`. Untuk menambahkan `API Key`, dapat dilakukan langkah-langkah berikut:

1. Membuat file baru bernama `.env`
2. Menuliskan baris berikut dalam file `.env`:
```
OPEN_WEATHER_API_KEY=<<Taruh API Key disini>
```

## Menjalankan Program
Untuk menjalankan program, saya telah menyiapkan beberapa executable yang dapat dijalankan dari beberapa sistem. Pertama-tama anda perlu membuka `Terminal`/`Powershell` di direktori proyek ini. Kemudian anda dapat mencoba menjalankan perintah berikut (pilih salah satu sesuai Sistem Operasi anda):
```bash
make run-windows # Untuk sistem operasi Windows

make run-linux # Untuk sistem operasi Linux

make run-mac # Untuk sistem operasi Mac OS
```

Apabila program gagal dijalankan (kemungkinan besar dikarenakan `cmake` tidak terinstall di sistem), anda dapat menjalankan program secara manual dengan melakukan perintah berikut:
### Windows
Anda dapat langsung menjalankan file `stamps-id.exe` yang berada di folder `/output/windows` atau menjalankan perintah berikut:
```powershell
.\output\windows\stamps-id.exe
```
### Mac OS
Anda dapat menjalankan perintah berikut:
```zsh
./output/mac/stamps-id
```
### Linux
Anda dapat menjalankan perintah berikut:
```bash
./output/linux/stamps-id
```

## Menjalankan Program Dari Kode

### Persiapan
Beberapa program yang perlu diinstall agar dapat menjalankan program langsung dari kode:
- Go Programming Langugage (https://go.dev/doc/install)
- VS Code (https://code.visualstudio.com/)

Beberapa program yang perlu diinstall khusus di Windows:
- Chocolatey (https://chocolatey.org/install)
- GNU Make (https://community.chocolatey.org/packages/make)

### Menjalankan Program
Untuk menjalankan program, ikuti langkah berikut:
- Clone Repository ke local
- Buka VSCode
- Buka Folder Repository
- Buka Sidebar `Run And Debug` (`Cmd + Shift + D` di Mac atau `Ctrl + Shift + D` di Windows / Linux)
- Tekan `Launch Package`
