package controllers

import (
	"invest-tracker/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Investments struct {
	db *gorm.DB
}

func NewInvestmentsController(db *gorm.DB) *Investments {
	return &Investments{db: db}
}

func (i *Investments) CreateInvestment(c *gin.Context) {
	var investment models.Investment
	if err := c.BindJSON(&investment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	investment.CreatedAt = time.Now()
	if err := i.db.Create(&investment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, investment)
}

func (i *Investments) GetInvestments(c *gin.Context) {
	var investments []models.Investment
	if err := i.db.Find(&investments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, investments)
}

func (i *Investments) UpdateInvestment(c *gin.Context) {
	id := c.Param("id")
	var investment models.Investment
	if err := c.BindJSON(&investment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := i.db.Model(&models.Investment{}).Where("id = ?", id).Updates(investment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, investment)
}
