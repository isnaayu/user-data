package controllers

import (
    "net/http"
    "user-api/models"
    "user-api/utils"

    "github.com/gin-gonic/gin"
)

// GetAllUsers mengambil semua pengguna
func GetAllUsers(c *gin.Context) {
    users := models.GetAllUsers()
    utils.RespondJSON(c, http.StatusOK, users)
}

// GetUserByID mengambil pengguna berdasarkan ID
func GetUserByID(c *gin.Context) {
    userID := c.Param("id")
    user, err := models.GetUserByID(userID)
    if err != nil {
        utils.RespondError(c, http.StatusNotFound, "Pengguna tidak ditemukan")
        return
    }
    utils.RespondJSON(c, http.StatusOK, user)
}
