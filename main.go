package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "BDDMidAzir184"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "CATech"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

func createUser() {

}

func getUser() {

}

func putUser() {

}

func main() {

	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}

	defer db.Close()

	if db.HasTable(&User{}) {
		db.DropTable(&User{})
		db.CreateTable(&User{})

	} else {
		db.CreateTable(&User{})
	}
	db.LogMode(true)

	r := gin.Default()

	r.POST("user/create", func(c *gin.Context) {
		user := User{}
		err := c.BindJSON(&user)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		user.Token = "string"
		db.NewRecord(user)
		db.Create(&user)
		c.JSON(http.StatusOK, gin.H{
			"token": user.Token,
		})

	})

	r.GET("user/get", func(c *gin.Context) {
		token := c.GetHeader("x-token")
		user := User{}
		//db.Where("token LIKE ?", token).Find(&user)
		db.Where("token = ?", token).First(&user)
		c.JSON(http.StatusOK, gin.H{
			"name": user.Name,
		})
	})

	r.PUT("user/update", func(c *gin.Context) {
		token := c.GetHeader("x-token")
		user := User{}
		err := c.BindJSON(&user)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		name := user.Name
		db.Where("token = ?", token).First(&user)
		user.Name = name
		db.Save(&user)
	})

	r.POST("/gacha/draw", func(c *gin.Context) {

	})

	r.Run(":8080")

}
