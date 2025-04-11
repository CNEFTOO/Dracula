package app

import "github.com/gin-gonic/gin"

func InitRouter(g *gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return nil
}
