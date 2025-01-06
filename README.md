1. [Go](https://go.dev/dl/) versi 1.21 atau lebih baru.
2. [Docker](https://www.docker.com/products/docker-desktop) untuk menjalankan container MySQL.
3. [Docker Compose](https://docs.docker.com/compose/install/) untuk mengelola container Docker.

Jalankan MySQL menggunakan Docker Compose:
docker-compose up -d

Menjalankan Migrasi Database:
go run ./cmd/main.go

run 
go run ./cmd/main.go


-nb:
untuk pakek method
POST /stok
ex:
{
  "nama_obat": "Paracetamol",
  "jumlah": 100,
  "harga": 5000.00
}
