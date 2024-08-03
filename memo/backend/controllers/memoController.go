package controllers

import (
    "memo/models"
    "net/http"

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
