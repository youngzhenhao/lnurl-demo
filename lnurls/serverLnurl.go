package lnurls

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lnurl-demo/api"
	"lnurl-demo/boltDB"
	"net/http"
	"time"
)

func RouterRunOnServer() {
	router := setupRouterOnServer()
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func setupRouterOnServer() *gin.Engine {
	router := gin.Default()

	router.POST("/upload/user", func(c *gin.Context) {
		//id := c.Param("id")
		id := uuid.New().String()
		name := c.PostForm("name")
		ip := c.PostForm("ip")
		result := true
		if id == "" || name == "" {
			result = false
		}
		user := &boltDB.User{
			ID:   id,
			Name: name,
			IP:   ip,
		}
		err := boltDB.InitServerDB()
		if err != nil {
			fmt.Printf("%s InitServerDB err :%v\n", api.GetTimeNow(), err)
		}
		db, err := bolt.Open("server.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			fmt.Printf("%s bolt.Open :%v\n", api.GetTimeNow(), err)
		}
		defer func(db *bolt.DB) {
			err := db.Close()
			if err != nil {
				fmt.Printf("%s db.Close :%v\n", api.GetTimeNow(), err)
			}
		}(db)
		s := &boltDB.ServerStore{DB: db}
		if result {
			err = s.CreateOrUpdateUser("users", user)
			if err != nil {
				fmt.Printf("%s CreateOrUpdateUser err :%v\n", api.GetTimeNow(), err)
				result = false
			}
		}
		var lnurlStr string
		serverDomainOrIp := api.GetEnv("SERVER_DOMAIN_OR_IP")
		if result {
			// TODO: /pay
			lnurlStr = Encode("http://" + serverDomainOrIp + "/pay?id=" + id)
		} else {
			id = ""
		}
		c.JSON(http.StatusOK, gin.H{
			"time":   api.GetTimeNow(),
			"id":     id,
			"name":   name,
			"ip":     ip,
			"result": result,
			"lnurl":  lnurlStr,
			//"url": Decode(lnurlStr),
		})
	})

	router.POST("/pay", func(c *gin.Context) {
		id := c.Query("id")
		amount := c.PostForm("amount")
		result := true

		err := boltDB.InitServerDB()
		if err != nil {
			fmt.Printf("%s InitServerDB err :%v\n", api.GetTimeNow(), err)
		}
		db, err := bolt.Open("server.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			fmt.Printf("%s bolt.Open :%v\n", api.GetTimeNow(), err)
		}
		defer func(db *bolt.DB) {
			err := db.Close()
			if err != nil {
				fmt.Printf("%s db.Close :%v\n", api.GetTimeNow(), err)
			}
		}(db)
		s := &boltDB.ServerStore{DB: db}
		user, err := s.ReadUser("users", id)
		if err != nil {
			fmt.Printf("%s ReadUser err :%v\n", api.GetTimeNow(), err)
		}
		// TODO: send POST request to call phone's /addInvoice
		PostPhoneToAddInvoice(user.IP, amount)

		c.JSON(http.StatusOK, gin.H{
			"time":   api.GetTimeNow(),
			"id":     id,
			"result": result,
		})
	})

	router.POST("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{})
	})

	return router
}
