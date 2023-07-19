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
![ERD](https://firebasestorage.googleapis.com/v0/b/crwn-db-6edbc.appspot.com/o/KP%20(2).png?alt=media&token=8fabdf22-0d51-4cb0-a087-71e46aa45224)

### Table Konsumen
Tabel "Konsumen" menyimpan informasi tentang konsumen yang terdaftar dalam sistem multifinance.

**Table Konsumen**
- `id_konsumen` (serial, primary key): ID unik untuk setiap konsumen.
- `nik` (varchar(20)): Nomor Induk Kependudukan (NIK) konsumen.
- `full_name` (varchar(255)): Nama lengkap konsumen.
- `legal_name` (varchar(255)): Nama lengkap sesuai dokumen resmi konsumen.
- `gaji` (numeric(10, 2)): Jumlah gaji konsumen.
- `tempat_lahir` (varchar(255)): Tempat lahir konsumen.
- `tanggal_lahir` (date): Tanggal lahir konsumen.
- `foto_ktp` (varchar(255)): URL atau path untuk foto KTP konsumen.
- `foto_selfie` (varchar(255)): URL atau path untuk foto selfie konsumen.

### Table LimitKredit
Tabel "LimitKredit" menyimpan informasi tentang batas kredit dan tenor yang dimiliki oleh konsumen.
**Table LimitKredit**
- `id_limit` (serial, primary key): ID unik untuk setiap entri limit.
- `id_konsumen` (int, foreign key references Konsumen(ID_Konsumen)): ID konsumen yang terkait.
- `tenor` (int): Tenor kredit dalam bulan.
- `batas_kredit` (numeric(12, 2)): Batas kredit yang diberikan kepada konsumen.

### Table Produk
Tabel "Produk" menyimpan informasi tentang produk yang tersedia dalam sistem multifinance.
**Table Produk**
- `id_produk` (serial, primary key): ID unik untuk setiap produk.
- `nama_produk` (varchar(255)): Nama produk yang tersedia dalam sistem.

### Table Transaksi
Tabel "Transaksi" menyimpan informasi tentang transaksi yang dilakukan oleh konsumen dalam sistem multifinance.
**Table Transaksi**
- `id_transaksi` (serial, primary key): ID unik untuk setiap transaksi.
- `id_konsumen` (int, foreign key references Konsumen(ID_Konsumen)): ID konsumen yang terkait.
- `nomor_kontrak` (varchar(50)): Nomor kontrak untuk transaksi.
- `tanggal_transaksi` (date): Tanggal transaksi.
- `otr` (numeric(12, 2)): Harga kendaraan atau aset lainnya sebelum diskon.
- `admin_fee` (numeric(10, 2)): Biaya administrasi transaksi.
- `jumlah_cicilan` (int): Jumlah cicilan yang akan dilakukan.
- `jumlah_bunga` (numeric(10, 2)): Jumlah bunga yang diterapkan pada transaksi.
- `nama_asset` (varchar(255)): Nama aset atau kendaraan yang dibeli.
- `jenis_transaksi` (varchar(50)): Jenis transaksi yang dilakukan, misalnya "Pembelian" atau "Pembiayaan".

### Table Detail_Transaksi
Tabel "Detail_Transaksi" menyimpan informasi detail untuk setiap transaksi yang dilakukan dalam sistem multifinance.
**Table Detail_Transaksi**
- `id_detail` (serial, primary key): ID unik untuk setiap detail transaksi.
- `id_transaksi` (int, foreign key references Transaksi(ID_Transaksi)): ID transaksi yang terkait.
- `id_produk` (int, foreign key references Produk(ID_Produk)): ID produk yang terkait.
- `jumlah_beli` (int): Jumlah barang yang dibeli dalam transaksi.
