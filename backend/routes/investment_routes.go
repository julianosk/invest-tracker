package routes

import (
    "github.com/gin-gonic/gin"
    "investment-management-app/controllers"
)

func SetupRoutes(router *gin.Engine, investments *controllers.Investments) {
    investmentGroup := router.Group("/api/investments")
    {
        investmentGroup.POST("/", investments.CreateInvestment)
        investmentGroup.GET("/", investments.GetInvestments)
        investmentGroup.PUT("/:id", investments.UpdateInvestment)
    }
}