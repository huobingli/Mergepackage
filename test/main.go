package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

type conf struct {
	File_path  string
	Ret_string string
}

var cf conf

// cors中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func Load_config(path string) error {
	if _, err := toml.DecodeFile(path, &cf); err != nil {
		return err
	}
	return nil
}

//启动
func main() {
	router := gin.Default()
	cur_path, _ := os.Getwd()
	fmt.Println(cur_path)
	err := Load_config("./config.toml")
	if err != nil {
		fmt.Println("load conf failed !!!")
		return
	} else {
		fmt.Println("load conf succeed")
	}

	// 加载CORS中间件
	router.Use(Cors())

	// 加载打包页面
	// router.LoadHTMLFiles("./html/index.html")
	// //加载静态资源，例如网页的css、js
	// router.Static("/html", "./html")
	// //加载单个静态文件
	// router.StaticFile("./favicon.ico", "./html/favicon.ico")
	// router.GET("/", func(context *gin.Context) {
	// 	context.HTML(http.StatusOK, "index.html", nil)
	// })

	// 增加接口
	router.GET("/api/checkversion2", GetUrl)
	router.GET("/GetDirFiles", GetDirFiles)

	router.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	// 静态文件服务
	fmt.Println(cf.File_path)
	fmt.Println(cf.Ret_string)
	router.StaticFS("/Path", http.Dir(cf.File_path))
	router.Run(":7001")
}

func GetDirFiles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "文件读取失败"})
}

type HevoUpdateInfo struct {
	UpdateUrl         string
	UpdateDescription string
	IsUpdate          string
	IsSilence         string
	ForceUpdate       string
	UpdateSize        string
	MD5Value          string
	TargetVersion     string
	ComponentItems    string
}

type Hevo struct {
	HevoUpdate      HevoUpdateInfo
	ProtocolVersion string
}

func GetUrl(c *gin.Context) {
	c.XML(http.StatusOK, Hevo{
		HevoUpdate{
			UpdateUrl:         "http://www.baidu.com",
			UpdateDescription: "更新描述",
			IsUpdate:          "true",
			IsSilence:         "true",
			ForceUpdate:       "true",
			UpdateSize:        "100",
			MD5Value:          "123123321",
			TargetVersion:     "1.0.0",
			ComponentItems:    "1.0.0.1",
		},
		ProtocolVersion: "3.3",
	})
}
