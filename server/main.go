package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

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
	if err := load_config("./conf.toml"); err != nil {
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
	router.POST("/upload_file", HandleUploadFile)
	router.POST("/upload_muti_file", HandleUploadMutiFile)
	router.GET("/download", HandleDownloadFile)
	router.GET("/getUrl", Get)
	router.GET("/file", HandleShowFile)
	router.GET("/CallMergePackage", CallMergePackage)
	router.GET("/CallZipPackage", CallZipPackage)

	// 文件服务器
	// path := "D:\\gitProject\\Mergepackage\\server\\package"
	fmt.Println(cf.Zip_dir)
	fmt.Println(cf.Xd_dir)
	fmt.Println(cf.Jrzd_dir)
	router.StaticFS("/zipDir", http.Dir(cf.Zip_dir))
	router.StaticFS("/xdDir", http.Dir(cf.Xd_dir))
	router.StaticFS("/jrzdDir", http.Dir(cf.Jrzd_dir))
	router.Run(":7001")
}

func CallMergePackage(c *gin.Context) {
	jrzd := c.Query("jrzd")
	xiadan := c.Query("xiadan")
	fmt.Println("CallMergePackage:", xiadan, jrzd)

	var out bytes.Buffer
	bat_cmd := "merge_package.bat " + jrzd + " " + xiadan
	fmt.Println(bat_cmd)
	cmd := exec.Command("cmd.exe", "/c", bat_cmd)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("%s", out.String())
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

//错误
func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(c *gin.Context) {
	url := c.Param("url")
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	c.JSON(200, gin.H{"status": 1, "result": result.String(), "message": "Success"})
}

// HandleUploadFile 上传单个文件
func HandleUploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
		return
	}
	fmt.Println(header.Filename)
	fmt.Println(string(content))
	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}

// HandleUploadMutiFile 上传多个文件
func HandleUploadMutiFile(c *gin.Context) {
	// 限制上传文件大小
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 4<<20)
	// 限制放入内存的文件大小
	err := c.Request.ParseMultipartForm(4 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
		return
	}
	formdata := c.Request.MultipartForm
	files := formdata.File["file"]
	for _, v := range files {
		file, err := v.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}
		fmt.Println(v.Filename)
		fmt.Println(string(content))
	}
	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}

// HandleDownloadFile 下载文件
func HandleDownloadFile(c *gin.Context) {
	//content := c.Query("content")
	//content = "hello world, 我是一个文件，" + content

	filename := c.Query("content")
	path := "D:\\ci\\autoBuild\\cd-tool\\output\\build\\"
	filepath := path + filename
	c.Writer.WriteHeader(http.StatusOK)
	// 返回文件名
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/text/plain")
	c.File(filepath)
}

func HandleShowFile(c *gin.Context) {
	path := "D:\\ci\\autoBuild\\cd-tool\\output\\build\\"
	fileName := path + c.Query("name")
	fmt.Println(fileName)
	c.File(fileName)

	//http.FileServer(http.Dir(BaseUploadPath))
	//http.FileServer()

}
