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
├── config/         
├── controllers/   
├── database/    
├── middleware/     
├── models/         
├── repository/
├── routes/     
├── services/       
├── utils/          
├── main.go         
└── .env            
</pre>

<pre>
git clone https://github.com/farhandwiyan/project-himaose.git
</pre>

<pre>
# Membuat file migrasi baru
migrate create -ext sql -dir database/migrations -seq nama_file_migrasi

# Menjalankan migrasi
migrate -path database/migrations/ -database "mysql://user:password@tcp(localhost:3306)/nama_db" up
</pre>
