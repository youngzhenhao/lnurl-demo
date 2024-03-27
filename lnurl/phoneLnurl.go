package lnurl

import (
	"github.com/gin-gonic/gin"
	"lnurl-demo/api"
	"net/http"
	"strconv"
)

func RouterRunOnPhone() {
	router := setupRouterOnPhone()
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func setupRouterOnPhone() *gin.Engine {
	router := gin.Default()
	router.POST("/addInvoice", func(c *gin.Context) {
		amountStr := c.PostForm("amount")
		amountInt, _ := strconv.Atoi(amountStr)
		invoiceStr := api.AddInvoice(int64(amountInt), "")
		c.JSON(http.StatusOK, gin.H{
			"invoice": invoiceStr,
		})
	})
	return router
}
