package controllers

import (
	"apotek-management/config"
	"apotek-management/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllStok(c *gin.Context) {
	var stok []models.Stok
	config.DB.Find(&stok)
	c.JSON(http.StatusOK, stok)
}

func CreateStok(c *gin.Context) {
    var input models.Stok
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "details": err.Error()})
        return
    }

    if err := input.Validate(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, input)
}
func UpdateStok(c *gin.Context) {
	var stok models.Stok
	id := c.Param("id")

	if err := config.DB.First(&stok, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	var input models.Stok
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&stok).Updates(input)
	c.JSON(http.StatusOK, stok)
}

func DeleteStok(c *gin.Context) {
	var stok models.Stok
	id := c.Param("id")

	if err := config.DB.First(&stok, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	config.DB.Delete(&stok)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
