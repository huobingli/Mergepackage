package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"Mergepackage/server/service/config"

	"github.com/gin-gonic/gin"
)

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

//启动
func main() {
	router := gin.Default()
	cur_path, _ := os.Getwd()
	fmt.Println(cur_path)
	cfi, err := config.Load_config_new("./conf.toml")
	if err != nil {
		fmt.Println("load conf failed !!!")
		return
	} else {
		fmt.Println("load conf succeed")
	}

	// 加载CORS中间件
	router.Use(Cors())

	// 加载打包页面
	router.LoadHTMLFiles("./html/index.html")
	//加载静态资源，例如网页的css、js
	router.Static("/html", "./html")
	//加载单个静态文件
	router.StaticFile("./favicon.ico", "./html/favicon.ico")
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// 增加接口
	router.GET("/CallTestPackageMaker", CallTestPackageMaker)
	router.GET("/CallPageParse", CallPageParse)
	router.GET("/CallZipPackage", CallZipPackage)
	router.GET("/GetDirFiles", GetDirFiles)

	router.POST("/CallMergePackage2", CallMergePackage2)
	// 静态文件服务
	fmt.Println(cfi.Zip_dir)
	fmt.Println(cfi.Xd_dir)
	fmt.Println(cfi.Jrzd_dir)
	router.StaticFS("/zipDir", http.Dir(cfi.Zip_dir))
	router.StaticFS("/xdDir", http.Dir(cfi.Xd_dir))
	router.StaticFS("/jrzdDir", http.Dir(cfi.Jrzd_dir))
	router.Run(":7001")
}

func CallTestPackageMaker(c *gin.Context) {
	jrzd := c.Query("jrzd")
	xiadan := c.Query("xiadan")
	fmt.Println("CallTestPackageMaker:", xiadan, jrzd)

	var out bytes.Buffer
	bat_cmd := "testPackageMaker.bat " + jrzd + " " + xiadan
	fmt.Println(bat_cmd)
	cmd := exec.Command("cmd.exe", "/c", bat_cmd)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("%s", out.String())
}

func CallPageParse(c *gin.Context) {
	// cmd := exec.Command("python.exe", "PackPageParser.py")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Printf("combined out:\n%s\n", string(out))
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// fmt.Printf("combined out:\n%s\n", string(out))
	// c.JSON(http.StatusOK, gin.H{"data": string(out)})

	// bat_cmd := "test.bat " + "1111"
	// fmt.Println(bat_cmd)
	// cmd := exec.Command("cmd.exe", "/c", bat_cmd)
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(out)
	// c.JSON(http.StatusOK, gin.H{"data": string(out)})

	var out bytes.Buffer
	bat_cmd := "test.bat " + "1111 222"
	fmt.Println(bat_cmd)
	cmd := exec.Command("cmd.exe", "/c", bat_cmd)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("%s", out.String())
	c.JSON(http.StatusOK, gin.H{"data": out.String()})
}

func CallZipPackage(c *gin.Context) {
	jrzd := c.Query("jrzd")
	fmt.Println("CallZipPackage:", jrzd)

	var out bytes.Buffer
	bat_cmd := "zip_package.bat " + jrzd
	fmt.Println(bat_cmd)
	cmd := exec.Command("cmd.exe", "/c", bat_cmd)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("%s", out.String())
}

func GetDirFiles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "文件读取失败"})
}

func CallMergePackage2(c *gin.Context) {
	hq := c.PostForm("hq")
	fmt.Println("hq:", hq)

	qs := c.PostForm("qs")
	fmt.Println("qs:", qs)

	broker := c.PostForm("broker")
	fmt.Println("broker", broker)

	//var out bytes.Buffer
	bat_cmd := "test.bat " + hq + " " + qs + " " + broker
	fmt.Println(bat_cmd)
	// cmd := exec.Command("cmd.exe", "/c", bat_cmd)
	// cmd.Stdout = &out
	// err := cmd.Run()
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println("%s", out.String())

}
