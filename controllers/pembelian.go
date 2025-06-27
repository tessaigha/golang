package controllers

import (
	"go-kasir-api/database"
	"go-kasir-api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPembelian(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, kode_barang, nama_barang, harga_beli, jumlah, tanggal_beli,supplier FROM pembelian")
	if err != nil {
		log.Println("Gagal ambil data pembelian:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
		return
	}
	defer rows.Close()

	var daftarPembelian []models.Pembelian
	for rows.Next() {
		var p models.Pembelian
		if err := rows.Scan(&p.ID, &p.Kode_Barang, &p.Nama_Barang, &p.Harga_Beli, &p.Jumlah, &p.Tanggal_Beli, &p.Supplier); err != nil {
			log.Println("Gagal scan:", err)
			continue
		}
		daftarPembelian = append(daftarPembelian, p)
	}

	c.JSON(http.StatusOK, daftarPembelian)
}

func AddPembelian(c *gin.Context) {
	var newPembelian models.Pembelian

	if err := c.BindJSON(&newPembelian); err != nil {
		log.Println("Gagal bind json:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal bind json"})
		return
	}

	result, err := database.DB.Exec("INSERT INTO pembelian (kode_barang,nama_barang,harga_beli,jumlah,tanggal_beli,supplier) VALUES (?,?,?,?,?,?)", newPembelian.Kode_Barang, newPembelian.Nama_Barang, newPembelian.Harga_Beli, newPembelian.Jumlah, newPembelian.Tanggal_Beli, newPembelian.Supplier)

	if err != nil {
		log.Println("Gagal insert data pembelian:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal insert data"})
		return
	}

	insertedID, err := result.LastInsertId()
	if err == nil {
		newPembelian.ID = int(insertedID)
	}
}
