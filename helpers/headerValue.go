package helpers

import "github.com/gin-gonic/gin"

func GetContextType(c *gin.Context) string {
	return c.Request.Header.Get("Content-type")
}
