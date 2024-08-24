package proceed

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// JoinUsRequest represents the data structure for the form submissions
type JoinUsRequest struct {
	ID          uint   `gorm:"primaryKey"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Description string `json:"description"`
	Email       string `json:"email" binding:"required,email"`
}

// SetupRoutes sets up the routes for handling form submissions and retrieval
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// POST /api/join-us - to submit a form
	router.POST("/api/join-us", func(c *gin.Context) {
		var req JoinUsRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save the form submission to the database
		if err := db.Create(&req).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Form submitted successfully!", "id": req.ID})
	})

	// GET /api/forms - to retrieve all submitted forms
	router.GET("/api/forms", func(c *gin.Context) {
		var forms []JoinUsRequest
		if err := db.Find(&forms).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
			return
		}

		c.JSON(http.StatusOK, forms)
	})

	// GET /api/forms/:id - to retrieve a specific form by ID
	router.GET("/api/forms/:id", func(c *gin.Context) {
		id := c.Param("id")
		var form JoinUsRequest

		if err := db.First(&form, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
			}
			return
		}

		c.JSON(http.StatusOK, form)
	})
}
