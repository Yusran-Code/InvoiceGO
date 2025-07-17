# 🧾 InvoiceGO - PDF Invoice Generator in Go

**InvoiceGO** A supporting system for PT PERTAMINA’s main application, providing an efficient solution for generating invoices. Built with Go (Golang) that generates and downloads PDF invoices and LO Laporan operational(Operations Report). This project is perfect for showcasing backend skills, following clean architecture and a modular structure.

---

## 🚀 Features

- ✅ Generate invoice PDF from form input
- ✅ Preview invoice in browser
- ✅ Download invoice as PDF
- ✅ Input validation: numbers only, required fields
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


## 🧑‍💻 Usage

1. Visit `http://localhost:8080/setup`
2. Fill in the invoice form
3. Click "Preview" to see the invoice
4. Click "Download" to get PDF





