package utils

import (
    "github.com/gin-gonic/gin"
)

// RespondJSON mengirimkan respons JSON
func RespondJSON(c *gin.Context, status int, payload interface{}) {
    c.JSON(status, payload)
}

// RespondError mengirimkan respons error
func RespondError(c *gin.Context, status int, message string) {
    c.JSON(status, gin.H{"error": message})
}
