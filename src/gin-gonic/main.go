package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Birthday string

func (b *Birthday) UnmarshalParam(param string) error {
	*b = Birthday(strings.Replace(param, "-", "/", -1))
	return nil
}

func main() {
	router := gin.Default()
	var request struct {
		Birthday Birthday `form:"birthday"`
	}

	router.GET("/test/:name", func(c *gin.Context) {
		_ = c.BindQuery(&request)

		name := c.Param("name")
		par := c.Query("lastname")
		c.JSON(http.StatusOK, gin.H{"message": "Hello " + name + " " + par + " " + string(request.Birthday)})
	})

	router.Run(":3000")

}
