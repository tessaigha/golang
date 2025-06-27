package controllers

import (
    "go-kasir-api/database"
    "go-kasir-api/models"
    "net/http"
    "log"
    "strconv"

    "github.com/gin-gonic/gin"
)


func GetBarang(c *gin.Context) {
    rows, err := database.DB.Query("SELECT id, nama, harga, stok FROM barang")
    if err != nil {
        log.Println("Gagal ambil data barang:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
        return
    }
    defer rows.Close()

    var daftarBarang []models.Barang
    for rows.Next() {
        var b models.Barang
        if err := rows.Scan(&b.ID, &b.Nama, &b.Harga, &b.Stok); err != nil {
            log.Println("Gagal scan:", err)
            continue
        }
        daftarBarang = append(daftarBarang, b)
    }

    c.JSON(http.StatusOK, daftarBarang)
}

func AddBarang(c *gin.Context) {
    var newBarang models.Barang

    // Ambil data JSON dari body request
    if err := c.ShouldBindJSON(&newBarang); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
        return
    }

    // Simpan data ke database MySQL
    result, err := database.DB.Exec("INSERT INTO barang (nama, harga, stok) VALUES (?, ?, ?)",
        newBarang.Nama, newBarang.Harga, newBarang.Stok)

    if err != nil {
        log.Println("Gagal insert ke database:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
        return
    }

    // Ambil ID yang baru dibuat
    insertedID, err := result.LastInsertId()
    if err == nil {
        newBarang.ID = int(insertedID)
    }

    // Kirim response sukses
    c.JSON(http.StatusCreated, newBarang)
}
func DeleteBarang(c *gin.Context) {
    id := c.Param("id")

    result, err := database.DB.Exec("DELETE FROM barang WHERE id = ?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus barang"})
        return
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Barang tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Barang berhasil dihapus"})
}
func UpdateBarang(c *gin.Context) {
    id := c.Param("id")
    var barang models.Barang

    if err := c.ShouldBindJSON(&barang); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Format tidak valid"})
        return
    }

    result, err := database.DB.Exec("UPDATE barang SET nama = ?, harga = ?, stok = ? WHERE id = ?",
        barang.Nama, barang.Harga, barang.Stok, id)

    // Tambahkan log
    log.Println("ID:", id, "Data:", barang)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update"})
        return
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "ID tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, barang)
}

func atoi(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}
func GetBarangByID(c *gin.Context) {
    id := c.Param("id")
    var b models.Barang
    err := database.DB.QueryRow("SELECT id, nama, harga, stok FROM barang WHERE id = ?", id).
        Scan(&b.ID, &b.Nama, &b.Harga, &b.Stok)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Barang tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, b)
}


