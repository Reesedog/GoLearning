package controllers

import (
	"memo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MemoController struct {
	DB *gorm.DB
}

func (ctrl *MemoController) CreateMemo(c *gin.Context) {
	var memo models.Memo
	if err := c.BindJSON(&memo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctrl.DB.Create(&memo)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ctrl *MemoController) GetMemos(c *gin.Context) {
	var memos []models.Memo
	ctrl.DB.Find(&memos)
	c.JSON(http.StatusOK, memos)
}

func (ctrl *MemoController) DeleteMemo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	println(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid memo ID"})
		return
	}

	if err := ctrl.DB.Delete(&models.Memo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Memo deleted"})
}
