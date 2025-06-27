package database

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kasirdb")
    if err != nil {
        log.Fatal("Gagal koneksi ke MySQL:", err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatal("Ping ke database gagal:", err)
    }

    log.Println("Koneksi ke MySQL berhasil!")
}
