# 🧾 InvoiceGO - PDF Invoice Generator in Go

**InvoiceGO** is an internal supporting tool for PT PERTAMINA’s main application, designed to efficiently generate PDF invoices. Built with Go (Golang), it supports both invoice generation and download. The system also extracts raw Excel data from the main web application to generate Laporan Operasional (Operational Reports), enhanced with additional fields such as vendor name headers and total price/tax summaries.

---

## 🚀 Features

- ✅ Google Outh Login
- ✅ Generate invoice PDF from form input
- ✅ Preview invoice in browser
- ✅ Download invoice as PDF
- ✅ Input validation: numbers only, required fields
- ✅ Extract raw excel data LO
- ✅ Generate LO with Add vendor name heading & total cell
- ✅ JSON error response from API
- ✅ Clean & scalable folder structure
- ✅ Ready to extend with database integration

---


---

## ⚙️ Installation

```bash
git clone https://github.com/yourusername/InvoiceGO.git
cd InvoiceGO
go mod tidy
go run cmd/main.go

## 🧑‍💻 Usage

1. Visit `http://localhost:8080/setup`
2. Fill in the invoice form
3. Click "Preview" to see the invoice
4. Click "Download" to get PDF

## 📡 API Endpoints

| Method | Endpoint         | Description                            | Auth Required |
|--------|------------------|----------------------------------------|----------------|
| GET    | `/`              | Show login or landing page             | ❌             |
| GET    | `/index`         | Dashboard / Home (after login)         | ✅             |
| GET    | `/setup`         | Setup invoice form                     | ✅             |
| POST   | `/generate`      | Generate and preview invoice           | ✅             |
| POST   | `/generate-pdf`  | Generate and download invoice as PDF   | ✅             |
| GET    | `/lo`            | Show LO (Letter of Offer) form         | ✅             |
| POST   | `/previewLo`     | Preview LO PDF                         | ✅             |
| POST   | `/downloadLo`    | Download LO PDF                        | ✅             |
| GET    | `/static/...`    | Serve static files (CSS, JS, etc.)     | ❌             |








