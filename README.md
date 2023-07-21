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
![ERD](https://firebasestorage.googleapis.com/v0/b/crwn-db-6edbc.appspot.com/o/KP%20(4).png?alt=media&token=33d6199d-34d3-419a-a398-70d723c92ab7)

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
- `email` (varchar(255)): email konsumen
- `password (varchar(255))`: password hash dari konsumen
- `role` (varchar(50)): role dari konsumen ini

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

### Table DetailTransaksi
Tabel "Detail_Transaksi" menyimpan informasi detail untuk setiap transaksi yang dilakukan dalam sistem multifinance.
**Table DetailTransaksi**
- `id_detail` (serial, primary key): ID unik untuk setiap detail transaksi.
- `id_transaksi` (int, foreign key references Transaksi(ID_Transaksi)): ID transaksi yang terkait.
- `id_produk` (int, foreign key references Produk(ID_Produk)): ID produk yang terkait.
- `jumlah_beli` (int): Jumlah barang yang dibeli dalam transaksi.

## API Specification
Berikut contoh API specification dari technical test ini beserta contoh request dan responsenya.

### API Login

Login email dan password.

**Request**

- Method: POST
- Endpoint: /api/login
- Content-Type: application/json

**Contoh Request**
```
POST /api/login
Content-Type: application/json

```

**Contoh Response (Berhasil)**
```json
{
  "status_code": 200,
  "message": "Login Success",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF9rb25zdW1lbiI6MSwiZW1haWwiOiJhbGFuc21pdGhAZXhhbXBsZS5jb20iLCJuYW1lIjoiQWxhbiBTbWl0aCIsInJvbGUiOiJ1c2VyIiwiZ2FqaSI6MTAwMDAwMDB9.9iHI6bUVHbvJzcnThKdNv6Bbhc6R5JWy0MMFRptlpyI"
}
```

**Contoh Response (gagal)**
```json
{
  "status_code": 500,
  "message": "crypto/bcrypt: hashedPassword is not the hash of the given password",
  "data": null
}
```

### API Upload foto

Upload foto KTP dan foto selfie ke Firebase Storage dalam satu request.

**Request**

- Method: POST
- Endpoint: /api/upload/foto
- Content-Type: multipart/form-data

**Contoh Request**
```
POST /api/upload/foto
Content-Type: multipart/form-data

foto_ktp: [file: gambar_foto_ktp.jpg]
foto_selfie: [file: gambar_foto_selfie.jpg]
```


**Contoh Response (Berhasil)**
```json
{
  "status_code": 200,
  "message": "Berhasil upload gambar",
  "data": {
    "url_ktp": "https://firebase-storage-url/foto_ktp.jpg",
    "url_selfie": "https://firebase-storage-url/foto_selfie.jpg"
  }
}
```

**Contoh Response (gagal)**
```json
{
  "status_code": 500,
  "message": "Error"
}
```
### API Create user

Membuat user baru.

**Request**

- Method: POST
- Endpoint: /api/users
- Content-Type: application/json

**Request Body**
```json
{
  "nik": "1234567890",
  "full_name": "John Doe",
  "email": "johndoe@example.com",
  "password": "password123",
  "gaji": 1000,
  "legal_name": "john doe",
  "role": "user",
  "tempat_lahir": "Washington",
  "tanggal_lahir": "2000-24-10",
  "foto_ktp": "https://firebase-storage-url/foto_ktp.jpg",
  "foto_selfie": "https://firebase-storage-url/foto_selfie.jpg"
}
```

**Contoh Response (Berhasil)**
```json
{
  "status_code": 201,
  "message": "Berhasil create user",
  "data": null
}
```

**Contoh Response (gagal)**
```json
{
  "status_code": 500,
  "message": "Error",
  "data": null
}
```
### GetProducts

GetProducts

**Request**

- Method: GET
- Endpoint: /api/products
- Content-Type: application/json



**Contoh Response (Berhasil)**
```json
{
  "status_code": 200,
  "message": "Get Products success",
  "data": [
    {
      "nama_produk": "Laptop"
    },
    {
      "nama_produk": "Smartphone"
    },
    {
      "nama_produk": "TV"
    },
    {
      "nama_produk": "Kulkas"
    },
    {
      "nama_produk": "Motor"
    },
    {
      "nama_produk": "Mobil"
    }
  ]
}
```

**Contoh Response (gagal)**
```json
{
  "status_code": 422,
  "message": "token is malformed: could not JSON decode header: invalid character '\\'' after object key:value pair",
  "data": null
}
```

### GetTransactions

GetTransactions get semua transaksi berdasarkan konsumen id

**Request**

- Method: GET
- Endpoint: /api/products
- Content-Type: application/json


**Contoh Response (Berhasil)**
```json
{
  "status_code": 200,
  "message": "Success Get Transactions",
  "data": [
    {
      "id_transaksi": 1,
      "id_konsumen": 1,
      "nomor_kontrak": "PB14",
      "tanggal_transaksi": "2023-07-21T00:00:00Z",
      "otr": 14000000,
      "admin_fee": 700000,
      "jumlah_cicilan": 4,
      "jumlah_bunga": 840000,
      "nama_asset": "Beli gadget",
      "jenis_transaksi": "Pembelian"
    }
  ]
}
```

**Contoh Response (gagal)**
```json
{
  "status_code": 422,
  "message": "token is malformed: could not JSON decode header: invalid character '\\'' after object key:value pair",
  "data": null
}
```

### Create Transaction

Membuat transaksi baru.

**Request**

- Method: POST
- Endpoint: /api/transactions
- Content-Type: application/json

**Request Body**
```json
{
  "id_konsumen": 2,
  "otr": 19000000,
  "jumlah_cicilan": 4,
  "nama_asset": "Honda Beat",
  "jenis_transaksi": "Pembelian",
  "detail_transaksi": [
    {
      "id_produk": 5,
      "jumlah_beli": 1
    }
  ]
}
```
**Contoh Response (Berhasil)**
```json
{
  "status_code": 201,
  "message": "Success Create Transaction",
  "data": {
    "id_transaksi": 8,
    "id_konsumen": 2,
    "nomor_kontrak": "PB24",
    "tanggal_transaksi": "2023-07-21T10:39:53.874422+07:00",
    "otr": 19000000,
    "admin_fee": 950000,
    "jumlah_cicilan": 4,
    "jumlah_bunga": 1140000,
    "nama_asset": "Honda Scoopy",
    "jenis_transaksi": "Pembelian"
  }
}
```

**Contoh Response (gagal)**
```json
{
  "status_code": 422,
  "message": "token is malformed: could not JSON decode header: invalid character '\\'' after object key:value pair",
  "data": null
}
```

## SetupDB
Berikut query untuk membuat table di postgres. 
```sql
-- Membuat tabel Konsumen
CREATE TABLE public."Konsumen" (
    id_konsumen SERIAL PRIMARY KEY,
    nik VARCHAR(20),
    full_name VARCHAR(255),
    legal_name VARCHAR(255),
    gaji NUMERIC(10, 2),
    tempat_lahir VARCHAR(255),
    tanggal_lahir DATE,
    foto_ktp VARCHAR(255),
    foto_selfie VARCHAR(255),
    role VARCHAR(50),
    email VARCHAR(255),
    password VARCHAR(255)
);

-- Membuat tabel LimitKredit
CREATE TABLE public."LimitKredit" (
    id_limit SERIAL PRIMARY KEY,
    id_konsumen INT REFERENCES "Konsumen"(id_konsumen),
    tenor INT,
    batas_kredit NUMERIC(12, 2)
);

-- Membuat tabel Transaksi
CREATE TABLE public."Transaksi" (
    id_transaksi SERIAL PRIMARY KEY,
    id_konsumen INT REFERENCES "Konsumen"(id_konsumen),
    nomor_kontrak VARCHAR(50),
    tanggal_transaksi DATE,
    otr NUMERIC(12, 2),
    admin_fee NUMERIC(10, 2),
    jumlah_cicilan INT,
    jumlah_bunga NUMERIC(10, 2),
    nama_asset VARCHAR(255),
    jenis_transaksi VARCHAR(50)
);

-- Membuat tabel Produk
CREATE TABLE public."Produk" (
    id_produk SERIAL PRIMARY KEY,
    nama_produk VARCHAR(255)
);

-- Membuat tabel Detail_Transaksi
CREATE TABLE public."DetailTransaksi" (
    id_detail SERIAL PRIMARY KEY,
    id_transaksi INT REFERENCES "Transaksi"(id_transaksi),
    id_produk INT REFERENCES "Produk"(id_produk),
    jumlah_beli INT
);



INSERT INTO "Produk" (nama_produk) VALUES
('Laptop'),
('Smartphone'),
('TV'),
('Kulkas'),
('Motor'),
('Mobil');

```

## Running Project
Untuk running project saya menggunakan `fresh`. Buka terminal pada project kredit-plus kemudian ketik `fresh`

## .env & kredit-plus.json
Contoh .env yang dipakai disini:
```
DB_USER=
DB_PASSWORD=
DB_ADDRESS=
DB_PORT=
DB_NAME=
SECRET_KEY=
FIREBASE_BUCKET=
FIREBASE_SERVICE_JSON=
```
value dari `FIREBASE_BUCKET`, `FIREBASE_SERVICE_JSON` & field yang lainnya silahkan disesuaikan.

Contoh kredit-plus.json file tersebut dari setting firebase:
```
{
  "type": "service_account",
  "project_id": "",
  "private_key_id": "",
  "private_key": "-----BEGIN PRIVATE KEY-----\contoh private key disiniI\n-----END PRIVATE KEY-----\n",
  "client_email": "",
  "client_id": "",
  "auth_uri": "",
  "token_uri": "",
  "auth_provider_x509_cert_url": "",
  "client_x509_cert_url": "",
  "universe_domain": "googleapis.com"
}
```
untuk semua field diatas harus di generate dari firebasenya.
