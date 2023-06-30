# Test Stamps Indonesia

Berikut merupakan hasil test initial screening untuk proses pendafataran di Stamps Indonesia.

Repository milik Muhammad Wyndham Haryata Permana.

## Isi Program
Program di repository ini bertujuan untuk menjawab 2 pertanyaan yang di lampirkan di email. Jawaban dari kedua pertanyaan tersebut saya rangkum menjadi 1 buat program yang dapat dipilih ketika menjalankan programnya.

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
```zsh
./output/mac/stamps-id
```
### Linux
```bash
./output/linux/stamps-id
```
