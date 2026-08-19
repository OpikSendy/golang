# 📡 04 - API Contract & Specifications

## 🌐 Live Backend Base URL
```text
https://golangpr.up.railway.app
```

---

## 📋 Endpoint Specifications

### 1. Server Health Check
* **Endpoint**: `GET /health`
* **Response 200 OK**:
```json
{
  "status": "ok",
  "message": "Mini Order Management API is running healthy! 🚀"
}
```

---

### 2. Create Order
* **Endpoint**: `POST /api/v1/orders`
* **Headers**: `Content-Type: application/json`
* **Request Body**:
```json
{
  "customer_name": "string (min 3, max 100)",
  "item_name": "string (min 2, max 150)",
  "amount": "number (gt 0)"
}
```
* **Response 201 Created**:
```json
{
  "success": true,
  "message": "Pesanan berhasil dibuat",
  "data": {
    "id": 4,
    "customer_name": "Dewi Lestari",
    "item_name": "Croissant Cokelat",
    "amount": 25000,
    "status": "pending",
    "created_at": "2026-08-19T13:30:00Z",
    "updated_at": "2026-08-19T13:30:00Z"
  }
}
```

---

### 3. Get All Orders
* **Endpoint**: `GET /api/v1/orders`
* **Response 200 OK**:
```json
{
  "success": true,
  "message": "Daftar pesanan berhasil diambil",
  "total": 3,
  "data": [
    {
      "id": 4,
      "customer_name": "Dewi Lestari",
      "item_name": "Croissant Cokelat",
      "amount": 25000,
      "status": "pending",
      "created_at": "2026-08-19T13:30:00Z",
      "updated_at": "2026-08-19T13:30:00Z"
    }
  ]
}
```

---

### 4. Get Order By ID
* **Endpoint**: `GET /api/v1/orders/:id`
* **Response 200 OK**:
```json
{
  "success": true,
  "message": "Detail pesanan berhasil ditemukan",
  "data": {
    "id": 4,
    "customer_name": "Dewi Lestari",
    "item_name": "Croissant Cokelat",
    "amount": 25000,
    "status": "pending",
    "created_at": "2026-08-19T13:30:00Z",
    "updated_at": "2026-08-19T13:30:00Z"
  }
}
```

---

### 5. Payment Webhook Simulator
* **Endpoint**: `POST /api/v1/webhooks/payment`
* **Headers**: `Content-Type: application/json`
* **Request Body**:
```json
{
  "order_id": 4,
  "payment_status": "paid",
  "transaction_id": "TRX-WEB-20260819"
}
```
* **Response 200 OK**:
```json
{
  "success": true,
  "message": "Status pesanan berhasil diperbarui melalui Webhook",
  "transaction_id": "TRX-WEB-20260819",
  "data": {
    "id": 4,
    "customer_name": "Dewi Lestari",
    "item_name": "Croissant Cokelat",
    "amount": 25000,
    "status": "paid",
    "created_at": "2026-08-19T13:30:00Z",
    "updated_at": "2026-08-19T13:30:30Z"
  }
}
```
