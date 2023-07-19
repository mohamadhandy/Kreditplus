# Technical Test Kredit Plus Golang
![general_architecture](https://firebasestorage.googleapis.com/v0/b/crwn-db-6edbc.appspot.com/o/Kredit%20plus.png?alt=media&token=edd2db01-8276-43a5-8846-580026883d5d)
Berikut merupakan gambar dari general architecture untuk technical test Kredit Plus menggunakan golang nantinya. Saya akan coba jelaskan salah dua dari API ini yaitu `API upload` dan `API create user` untuk memberikan gambaran yang lebih jelas dari architecture ini. 
### API Upload
- User mengirimkan request dengan `formdata` yang berisi gambar.
- Handler mengakses Firebase untuk mengirim gambar ke Firebase.
   - Jika proses upload gambar gagal, Handler memberikan respons berupa error (HTTP 500).
   - Jika proses upload gambar berhasil, Handler memberikan respons API dengan `link` dan status codenya 200.
- Dari sini, user memanggil API `create user` setelah mendapatkan `link` gambar.
### API Create User
- Setelah menerima request dari user, handler meneruskannya ke use case.
- Pada Usecase akan menangani logika bisnis jika diperlukan.
- Use case kemudian memanggil kode repositori.
- Di repositori, proses insert akan dilakukan di PostgreSQL.
   - Apakah berhasil atau tidak, repositori memberikan response.
- Repositori mengirimkan respon kembali ke use case.
- Use case mengembalikannya ke handler.
- Handler akan memberikan response kepada user.

Dengan memisahkan `API upload` dan `API create user` tujuannya adalah `separation of concern` sehingga memudahkan developer mengembangkan masing-masing API secara scalable.
## Gambaran Umum ERD
![ERD](https://firebasestorage.googleapis.com/v0/b/crwn-db-6edbc.appspot.com/o/KP%20(1).png?alt=media&token=001ae373-0e13-47eb-93ed-39389e5cf5a9)

### Table Konsumen
Tabel "Konsumen" menyimpan informasi tentang konsumen yang terdaftar dalam sistem multifinance.

**Table Konsumen**
- `ID_Konsumen` (serial, primary key): ID unik untuk setiap konsumen.
- `NIK` (varchar(20)): Nomor Induk Kependudukan (NIK) konsumen.
- `Full_Name` (varchar(255)): Nama lengkap konsumen.
- `Legal_Name` (varchar(255)): Nama lengkap sesuai dokumen resmi konsumen.
- `Gaji` (numeric(10, 2)): Jumlah gaji konsumen.
- `Tempat_Lahir` (varchar(255)): Tempat lahir konsumen.
- `Tanggal_Lahir` (date): Tanggal lahir konsumen.
- `Foto_KTP` (varchar(255)): URL atau path untuk foto KTP konsumen.
- `Foto_Selfie` (varchar(255)): URL atau path untuk foto selfie konsumen.

### Table Limit
Tabel "Limit" menyimpan informasi tentang batas kredit dan tenor yang dimiliki oleh konsumen.
**Table Limit**
- `ID_Limit` (serial, primary key): ID unik untuk setiap entri limit.
- `ID_Konsumen` (int, foreign key references Konsumen(ID_Konsumen)): ID konsumen yang terkait.
- `Tenor` (int): Tenor kredit dalam bulan.
- `Batas_Kredit` (numeric(12, 2)): Batas kredit yang diberikan kepada konsumen.

### Table Produk
Tabel "Produk" menyimpan informasi tentang produk yang tersedia dalam sistem multifinance.
**Table Produk**
- `ID_Produk` (serial, primary key): ID unik untuk setiap produk.
- `Nama_Produk` (varchar(255)): Nama produk yang tersedia dalam sistem.

### Table Transaksi
Tabel "Transaksi" menyimpan informasi tentang transaksi yang dilakukan oleh konsumen dalam sistem multifinance.
**Table Transaksi**
- `ID_Transaksi` (serial, primary key): ID unik untuk setiap transaksi.
- `ID_Konsumen` (int, foreign key references Konsumen(ID_Konsumen)): ID konsumen yang terkait.
- `ID_Produk` (int, foreign key references Produk(ID_Produk)): ID produk yang terkait.
- `Nomor_Kontrak` (varchar(50)): Nomor kontrak untuk transaksi.
- `Tanggal_Transaksi` (date): Tanggal transaksi.
- `OTR` (numeric(12, 2)): Harga kendaraan atau aset lainnya sebelum diskon.
- `Admin_Fee` (numeric(10, 2)): Biaya administrasi transaksi.
- `Jumlah_Cicilan` (int): Jumlah cicilan yang akan dilakukan.
- `Jumlah_Bunga` (numeric(10, 2)): Jumlah bunga yang diterapkan pada transaksi.
- `Nama_Asset` (varchar(255)): Nama aset atau kendaraan yang dibeli.
- `Jenis_Transaksi` (varchar(50)): Jenis transaksi yang dilakukan, misalnya "Pembelian" atau "Pembiayaan".

### Table Detail_Transaksi
Tabel "Detail_Transaksi" menyimpan informasi detail untuk setiap transaksi yang dilakukan dalam sistem multifinance.
**Table Detail_Transaksi**
- `ID_Detail` (serial, primary key): ID unik untuk setiap detail transaksi.
- `ID_Transaksi` (int, foreign key references Transaksi(ID_Transaksi)): ID transaksi yang terkait.
- `Jumlah_Beli` (int): Jumlah barang yang dibeli dalam transaksi.
