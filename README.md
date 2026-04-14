# Himaose Project - Backend API

Backend API untuk platform manajemen kegiatan Himpunan Mahasiswa Oseanografi. Dibangun menggunakan **Go (Golang)** dengan framework **Fiber** dan **GORM** sebagai ORM-nya.

## Tech Stack

* **Language:** Go (Golang)
* **Framework:** [Gofiber/Fiber v2](https://gofiber.io/)
* **ORM:** [GORM](https://gorm.io/)
* **Database:** MySQL
* **Authentication:** JWT (JSON Web Token)

## Folder Structure

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

## Database Migration
### Create Migration    
<pre>
migrate create -ext sql -dir database/migrations -seq nama_file_migrasi
</pre>

### Run Migration
<pre>
migrate -path database/migrations/ -database "mysql://root:@tcp(localhost:3306)/himaose" up
</pre>
