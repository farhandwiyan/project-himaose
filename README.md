# 🌊 Himaose Project - Backend API

Backend API untuk platform manajemen kegiatan Himpunan Mahasiswa Oseanografi. Dibangun menggunakan **Go (Golang)** dengan framework **Fiber** dan **GORM** sebagai ORM-nya.

## 🚀 Tech Stack

* **Language:** Go (Golang)
* **Framework:** [Gofiber/Fiber v2](https://gofiber.io/)
* **ORM:** [GORM](https://gorm.io/)
* **Database:** MySQL / PostgreSQL
* **Authentication:** JWT (JSON Web Token)
* **Validation:** Go-Playground Validator

## 📁 Folder Structure

<pre>
.
├── config/         # Konfigurasi database & environment
├── controllers/    # Handler untuk HTTP request
├── middleware/     # JWT & Logging middleware
├── models/         # Struct database & JSON schema
├── repository/     # Query langsung ke database
├── services/       # Logic bisnis & validasi
├── utils/          # Helper (response formatter, hash, dll)
├── main.go         # Entry point aplikasi
└── .env            # Konfigurasi variabel lingkungan
</pre>

<pre>
git clone https://github.com/farhandwiyan/project-himaose.git
</pre>
