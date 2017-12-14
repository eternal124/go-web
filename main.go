//package main
//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func mainPage(c *gin.Context){
//	c.HTML(http.StatusOK, "index.html", nil)
//}
//
//func main(){
//	router := gin.Default()
//	router.LoadHTMLGlob("E:/MyGo/src/app/cli/html/*")
//
//	router.GET("/main", mainPage)
//
//	router.Run(":9090")
//}


package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"app/server/apis"
	"log"
	"strconv"
	"app/server/database"
)

func main() {
	router := gin.Default()    //获得路由实例

	//添加中间件
	router.Use(Middleware)

	//加载模板路径
	router.LoadHTMLGlob("E:/MyGo/src/app/templates/*")

	//注册接口
	router.GET("/get", GetHandler)
	router.GET("/post", PostHandler)
	router.POST("/post", PostHandler)
	router.PUT("/put", PutHandler)
	router.DELETE("/delete", DeleteHandler)
	//监听端口

	defer database.SqlDB.Close()

	router.Run(":8000")
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}

func GetHandler(c *gin.Context) {
	users, err := apis.Retrieve()
	if err != nil {
		log.Fatal("get error")
		c.String(http.StatusNotFound, "get error")
	}
	fmt.Println(users)
	c.HTML(http.StatusOK, "post.html", gin.H{
		"users": users,
	})
}
func PostHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "get.html", nil)
	} else {
		name := c.PostForm("name")
		password := c.PostForm("password")
		icon := c.PostForm("icon")
		info := c.PostForm("info")
		academy := c.PostForm("academy")
		level, err := strconv.ParseInt(c.PostForm("level"), 10, 64)
		if err != nil {
			fmt.Println("数字转换失败", err)
		}
		user := apis.User{
			Id: 3,
			Name: name,
			Password: password,
			Icon: icon,
			Info: info,
			Academy: academy,
			Level: level,
		}
		var users []apis.User
		users = append(users, user)
		count, err := apis.Create(users)
		fmt.Printf("create user %d", count)

		c.String(http.StatusOK, "It works")
	}
}

func PutHandler(c *gin.Context) {
	c.String(http.StatusOK, "put")
}

func DeleteHandler(c *gin.Context) {
	c.String(http.StatusOK, "delete")
}