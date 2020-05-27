package main

import (
	"net/http"
	"reverseiplookup/resolver"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
)

func main() {
	r := gin.Default()

	storage := SetupDB()
	resolv := resolver.NewResolver(storage)

	go func() {
		s := gocron.NewScheduler()
		s.Every(1).Hours().Do(resolv.UpdateValid, 100)
		<-s.Start()
	}()

	r.GET("/ip/:address", func(c *gin.Context) {
		addr := c.Param("address")
		list, err := resolv.IPLookup(addr)
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
		list, err := resolv.HostLookup(addr)
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
