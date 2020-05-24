package main

import (
	"net/http"
	"reverseiplookup/resolver"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ip/:address", func(c *gin.Context) {
		addr := c.Param("address")
		list, err := resolver.IPLookup(addr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.JSON(200, gin.H{
			"message": list,
		})
	})
	r.GET("/domain/:address", func(c *gin.Context) {
		addr := c.Param("address")
		list, err := resolver.HostLookup(addr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.JSON(200, gin.H{
			"message": list,
		})
	})
	r.Run()
}
