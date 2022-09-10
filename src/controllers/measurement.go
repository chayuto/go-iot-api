package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type CreateMeasurementInput struct {
// 	Title  string `json:"title" binding:"required"`
// 	Author string `json:"author" binding:"required"`
// }

// type UpdateMeasurementInput struct {
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// }

// GET /books
// Find all books
func FindMeasurements(c *gin.Context) {
	// var books []models.Measurement
	// models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GET /books/:id
// Find a book
func FindMeasurement(c *gin.Context) {
	// Get model if exist
	// var book models.Measurement
	// if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// POST /books
// Create new book
func CreateMeasurement(c *gin.Context) {
	// // Validate input
	// var input CreateMeasurementInput
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Create book
	// book := models.Measurement{Title: input.Title, Author: input.Author}
	// models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
