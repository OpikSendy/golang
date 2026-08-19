# 🗄️ 03 - Database Entity Relationship & State Diagram

## 📊 Database Schema (PostgreSQL via Backend GORM)

Aplikasi Frontend berinteraksi dengan tabel `orders` yang di-manage oleh backend Golang:

```mermaid
erDiagram
    ORDERS {
        bigint id PK "Primary Key (Auto Increment)"
        varchar customer_name "Nama Pelanggan (Max 100)"
        varchar item_name "Nama Menu / Pesanan (Max 150)"
        numeric amount "Total Harga Pesanan (12, 2)"
        varchar status "Status Pesanan (pending, paid, cancelled)"
        timestamptz created_at "Waktu Pesanan Dibuat"
        timestamptz updated_at "Waktu Pesanan Diperbarui"
    }
```

---

## 🔄 Lifecycle & State Transition Pesanan

```mermaid
stateDiagram-v2
    [*] --> Pending: POST /api/v1/orders (Kasir Checkout)
    
    Pending --> Paid: POST /api/v1/webhooks/payment (Simulasi Bayar Berhasil)
    Pending --> Cancelled: POST /api/v1/webhooks/payment (Simulasi Bayar Batal)
    
    Paid --> [*]: Pesanan Selesai / Siap Disajikan
    Cancelled --> [*]: Pesanan Dibatalkan
    
    note right of Paid
        Idempotency Check:
        Jika webhook dengan status 'paid'
        dikirim ulang, status tidak berubah
        dan tidak terjadi duplikasi.
    end note
```
