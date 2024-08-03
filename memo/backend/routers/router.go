package routers

import (
    "time"
    "memo/controllers"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    memoController := &controllers.MemoController{DB: db}

    router.POST("/memo", memoController.CreateMemo)
    router.GET("/memos", memoController.GetMemos)

    return router
}
