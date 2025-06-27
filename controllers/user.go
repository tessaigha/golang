package controllers

import (
	"go-kasir-api/database"
	"go-kasir-api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, nama, username, email FROM user")
	if err != nil {
		log.Println("Gagal ambil data user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
		return
	}
	defer rows.Close()

	var daftarUser []models.User
	for rows.Next() {
		var b models.User
		if err := rows.Scan(&b.ID, &b.Nama, &b.Username, &b.Email); err != nil {
			log.Println("Gagal scan:", err)
			continue
		}
		daftarUser = append(daftarUser, b)
	}

	c.JSON(http.StatusOK, daftarUser)
}

func AddUser(c *gin.Context) {
	var newUser models.User

	// Ambil data JSON dari body request
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	// Simpan data ke database MySQL
	result, err := database.DB.Exec("INSERT INTO user (nama, username, email) VALUES (?, ?, ?)",
		newUser.Nama, newUser.Username, newUser.Email)

	if err != nil {
		log.Println("Gagal insert ke database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	insertedID, err := result.LastInsertId()
	if err == nil {
		newUser.ID = int(insertedID)
	}

	// Kirim response sukses
	c.JSON(http.StatusCreated, newUser)
}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	result, err := database.DB.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus user"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data User berhasil dihapus"})
}
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format tidak valid"})
		return
	}

	result, err := database.DB.Exec("UPDATE user SET nama = ?, username = ?, email = ? WHERE id = ?",
		user.Nama, user.Username, user.Email, id)

	// Tambahkan log
	log.Println("ID:", id, "Data:", user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update"})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var b models.User
	err := database.DB.QueryRow("SELECT id, nama, username, email FROM user WHERE id = ?", id).
		Scan(&b.ID, &b.Nama, &b.Username, &b.Email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data User tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, b)
}
