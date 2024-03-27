package lnurls

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lnurl-demo/api"
	"lnurl-demo/boltDB"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func RouterRunOnPhone() {
	router := setupRouterOnPhone()
	err := router.Run(":9090")
	if err != nil {
		return
	}
}

func setupRouterOnPhone() *gin.Engine {
	router := gin.Default()

	router.POST("/addInvoice", func(c *gin.Context) {

		id := uuid.New().String()
		amountStr := c.PostForm("amount")
		amountInt, _ := strconv.Atoi(amountStr)
		result := true
		if amountInt <= 0 {
			result = false
		}
		var invoiceStr string
		if result {
			invoiceStr = api.AddInvoice(int64(amountInt), "")
		}
		if invoiceStr == "" {
			result = false
		}

		err := boltDB.InitPhoneDB()
		if err != nil {
			fmt.Printf("%s InitPhoneDB err :%v\n", api.GetTimeNow(), err)
		}

		db, err := bolt.Open("phone.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			fmt.Printf("%s bolt.Open :%v\n", api.GetTimeNow(), err)
		}
		defer func(db *bolt.DB) {
			err := db.Close()
			if err != nil {
				fmt.Printf("%s db.Close :%v\n", api.GetTimeNow(), err)
			}
		}(db)
		s := &boltDB.PhoneStore{DB: db}

		if result {
			invoiceStr = strings.ToUpper(invoiceStr)
			err = s.CreateOrUpdateInvoice("invoices", &boltDB.Invoice{
				ID: id,
				//PubKey:     "4",
				Amount:     amountInt,
				InvoiceStr: invoiceStr,
			})
			if err != nil {
				fmt.Printf("%s CreateOrUpdateInvoice err :%v\n", api.GetTimeNow(), err)
				result = false
			}
		} else {
			id = ""
		}

		c.JSON(http.StatusOK, gin.H{
			"time":    api.GetTimeNow(),
			"id":      id,
			"amount":  amountInt,
			"invoice": invoiceStr,
			"result":  result,
		})
	})

	return router
}
