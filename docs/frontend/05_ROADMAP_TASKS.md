# 🗺️ 05 - Development Roadmap & Task Milestones

Rencana pengerjaan frontend **OrderPulse POS (Next.js 15 + SSR + TanStack Query)** yang dipecah menjadi 5 modul terstruktur:

---

## 📦 Modul FE 1: Inisialisasi & Setup Arsitektur Dasar
* [ ] Inisialisasi proyek Next.js 15 (App Router) + TypeScript menggunakan runtime **Bun**.
* [ ] Instalasi dependensi inti: `@tanstack/react-query`, `@tanstack/react-query-devtools`, `lucide-react`, `clsx`, `tailwind-merge`.
* [ ] Setup `QueryClientProvider` di `src/components/providers.tsx` dan integrasi ke `layout.tsx`.
* [ ] Setup tema styling dasar (Design System Tokens, Glassmorphism, Dark Slate Color Palette).
* [ ] Setup Navbar dengan indikator status live backend (`GET /health`).

---

## 📦 Modul FE 2: API Client Layer & Custom TanStack Query Hooks
* [ ] Pembuatan TypeScript Interfaces (`src/types/order.ts`) untuk Entity `Order`, DTO `CreateOrderInput`, dan `PaymentWebhookPayload`.
* [ ] Pembuatan API fetcher functions terpusat di `src/lib/api.ts`.
* [ ] Pembuatan Custom TanStack Hooks:
  * `useOrders()`: Fetch list orders dengan konfigurasi stale time & refetch interval.
  * `useCreateOrder()`: Mutation hook untuk checkout order baru + auto invalidate cache.
  * `usePaymentWebhook()`: Mutation hook untuk trigger simulasi pembayaran webhook.

---

## 📦 Modul FE 3: Halaman Cashier POS (Menu Catalog & Cart Drawer)
* [ ] Pembuatan katalog menu interaktif dengan data mock representatif (Kopi, Makanan, Minuman) beserta harga dan gambar/icon.
* [ ] Filter kategori menu (All, Coffee, Food, Beverage).
* [ ] Drawer / Panel Keranjang Belanja (*Cart*):
  * Tambah/kurang kuantitas item.
  * Hitung subtotal dan kalkulasi otomatis.
  * Input Form nama pelanggan.
* [ ] Integrasi tombol **Checkout** dengan mutasi `useCreateOrder`.

---

## 📦 Modul FE 4: Interactive Payment Modal & Webhook Simulator
* [ ] Pembuatan komponen Dialog / Modal Pembayaran dinamis:
  * Menampilkan Detail Order (ID, Nama Customer, Total Tagihan).
  * Menampilkan Visual QRIS Code / Virtual Account perbankan.
* [ ] Tombol Aksi **"⚡ Simulasikan Pembayaran Berhasil"**:
  * Mengirim `POST /api/v1/webhooks/payment` via `usePaymentWebhook`.
  * Menampilkan animasi checklist hijau sukses pembayaran (*Idempotent feedback*).
* [ ] Opsi Cetak Struk Ringkas (*Receipt Preview*).

---

## 📦 Modul FE 5: Real-time Order Tracker Dashboard & SSR Integration
* [ ] Pembuatan halaman SSR `/orders` dengan TanStack Hydration Boundary.
* [ ] Widget Ringkasan Statistik:
  * Total Omset Sukses (IDR).
  * Total Transaksi, Total Pesanan Pending vs Paid.
* [ ] Tabel & Kartu Daftar Antrean Pesanan dengan badge status real-time.
* [ ] Filter status pesanan (*All*, *Pending*, *Paid*) dan fitur Search.
* [ ] Atomic Commits dan persiapan deployment Vercel/Railway.
