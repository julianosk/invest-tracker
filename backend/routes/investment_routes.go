package routes

import (
    "github.com/gin-gonic/gin"
    "investment-management-app/controllers"
)

func SetupRoutes(router *gin.Engine) {
    investmentGroup := router.Group("/api/investments")
    {
        investmentGroup.POST("/", controllers.CreateInvestment)
        investmentGroup.GET("/", controllers.GetInvestments)
        investmentGroup.PUT("/:id", controllers.UpdateInvestment)
    }
}