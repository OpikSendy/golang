# 📄 01 - Product Requirement Document (PRD)

## 📌 Nama Produk
**OrderPulse POS & Live Payment Simulator (Next.js + TanStack Query)**

## 🎯 Ringkasan Produk
**OrderPulse** adalah aplikasi Point-of-Sale (POS) kasir modern dan dashboard monitoring pesanan real-time yang terhubung langsung secara *full-stack* dengan backend Golang Clean Architecture yang sudah ter-deploy live di Railway (`https://golangpr.up.railway.app`).

Aplikasi ini mendemonstrasikan alur transaksi end-to-end: dari pembuatan pesanan kasir (*Checkout*), simulasi pembayaran QRIS/Virtual Account melalui *Payment Webhook*, hingga pemantauan status pesanan yang otomatis ter-update secara instan (*real-time*).

---

## 👥 Pengguna Sasaran (Target Personas)

1. **Kasir / Outlet Staff**:
   * Memilih menu makanan & minuman dari katalog.
   * Menginput nama pelanggan dan memproses checkout order baru.
   * Menampilkan QRIS / instruksi pembayaran kepada pelanggan.
2. **Pelanggan (Customer)**:
   * Melakukan simulasi pembayaran tagihan pesanan.
3. **Manager / Kitchen Admin**:
   * Memantau antrean pesanan yang masuk secara real-time.
   * Memantau status pembayaran (*Pending* vs *Paid*).
   * Melihat total omset/revenue harian.

---

## 🚀 Fitur Utama & Prioritas (Prioritized Features)

### 🔴 P0 (Must Have - Core Features)
* **Katalog Menu & Keranjang (Cart)**:
  * Grid menu interaktif dengan filter kategori (Kopi, Makanan, Non-Kopi).
  * Drawer keranjang belanja dengan kalkulasi otomatis subtotal dan item list.
  * Form input nama customer dan nominal amount.
  * Tombol **Checkout** yang memanggil API `POST https://golangpr.up.railway.app/api/v1/orders`.
* **Interactive Payment Modal & Webhook Simulator**:
  * Modal popup pembayaran menampilkan nominal, ID pesanan, dan QR Code QRIS dinamis.
  * Tombol **"⚡ Simulasikan Pembayaran Berhasil"** yang menembak API `POST https://golangpr.up.railway.app/api/v1/webhooks/payment` di background.
  * Efek visual transisi status pesanan dari *Pending* $\rightarrow$ *Paid* dengan checklist animasi.
* **Real-time Order Tracker Dashboard**:
  * Tabel & kartu antrean pesanan menggunakan **TanStack React Query** dengan auto-polling (interval 3-5 detik) dan manual refetch.
  * Badge indikator status (*Pending: Kuning*, *Paid: Hijau*).

### 🟡 P1 (Should Have)
* **Metrik Statistik Cepat**:
  * Total Omset Sukses (IDR), Total Transaksi, Jumlah Pending vs Paid.
* **Filter & Pencarian**:
  * Filter order berdasarkan status (*All*, *Pending*, *Paid*) dan pencarian nama customer.
* **Toast & Notification Feedback**:
  * Toast notifikasi elegan saat order baru dibuat atau pembayaran sukses.

### 🟢 P2 (Nice to Have)
* **Print Struk Pesanan**:
  * Fitur simulasi cetak struk kasir (*Receipt Print Preview*).
* **Dark / Light Glassmorphism UI**:
  * Tema visual premium dengan efek glassmorphism dan micro-animations.

---

## 📏 Kriteria Penerimaan (Acceptance Criteria)

1. **Checkout Flow**:
   * Ketika kasir mengklik checkout dengan nama customer terisi, order berhasil dibuat di database PostgreSQL live dan mengembalikan respons HTTP 201 dengan `id` order baru.
2. **Webhook Payment Flow**:
   * Ketika tombol simulasi bayar ditekan, request webhook dikirim ke backend Railway dan status order di database PostgreSQL berubah menjadi `paid`.
   * UI modal menampilkan feedback sukses tanpa perlu reload browser.
3. **Dashboard Real-time**:
   * Daftar order di dashboard langsung merefleksikan perubahan status dalam hitungan detik berkat cache invalidation TanStack Query.
