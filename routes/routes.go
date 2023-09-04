package routes

import (
    "user-api/controllers"
    "github.com/gin-gonic/gin"
)

// SetupRoutes menginisialisasi rute API
func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.GET("/users", controllers.GetAllUsers)
        api.GET("/users/:id", controllers.GetUserByID)
    }
}
