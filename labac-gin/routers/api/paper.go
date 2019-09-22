package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"labac/pkg/setting"
	"labac/pkg/util"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	cfg := setting.Cfg
	sec, err := cfg.GetSection("database")
	if err != nil {
		log.Fatal("open section database error %v", err)
	}
	hostname := sec.Key("HOST").MustString("127.0.0.1:6379")
	password := sec.Key("PASSWORD").MustString("")
	db := sec.Key("DB").MustInt(0)

	client = redis.NewClient(&redis.Options{
		Addr:     hostname,
		Password: password,
		DB:       db,
	})
}

func UploadFiles(c *gin.Context) {
	name := c.PostForm("name")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  "文件上传失败",
		})
		return
	}

	fileName := header.Filename
	fileFrag := strings.Split(fileName, ".")
	fileType := fileFrag[len(fileFrag)-1]

	if v, ok := setting.FileTypes[fileType]; ok {
		localFilePath := setting.FilePath + "/" + util.HashName(name)

		err = os.MkdirAll(localFilePath, os.ModePerm)

		localFile, err := os.Create(localFilePath + "/" + fileName)

		if err != nil {
			log.Fatalf("create %s file error %v", localFilePath, err)
			return
		}

		defer localFile.Close()
		io.Copy(localFile, file)

		client.HSet(name, v, localFilePath+"/"+fileName)
		client.HSet(name, v+"Suffix", fileType)

		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "上传成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "文件类型错误",
		})
	}
}

func GetPaperList(c *gin.Context) {
	papers := client.LRange("paperlist", 0, -1)

	type info map[string]string
	infos := make([]info, len(papers.Val()))

	for idx, v := range papers.Val() {
		res := client.HGetAll(v).Val()
		infos[idx] = res
	}
	c.JSON(http.StatusOK, gin.H{
		"data": infos,
	})
}

func DeletePaper(c *gin.Context) {
	name := c.Param("name")
	client.LRem("paperlist", 0, name)
	client.Del(name)

	os.RemoveAll(setting.FilePath + "/" + util.HashName(name))

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "删除成功",
	})
}

func AddPaper(c *gin.Context) {
	paperInfo := make(map[string]interface{})
	err := c.BindJSON(&paperInfo)
	if err != nil {
		log.Fatalf("decode param error :%v", err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "decode error",
		})
	}
	name := paperInfo["name"].(string)
	paperInfo["date"] = paperInfo["date"].(string)[:7]
	paperInfo["create"] = time.Now().Format("2006-01-02 15:04:05")

	papers := client.LRange("paperlist", 0, -1)

	update := util.InList(papers.Val(), name)

	if !update {
		paperInfo["article"] = ""
		paperInfo["ppt"] = ""
		paperInfo["articleSuffix"] = ""
		paperInfo["pptSuffix"] = ""
		client.LPush("paperlist", name)
	}
	client.HMSet(name, paperInfo)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "add paper success",
	})
}

func DownloadFile(c *gin.Context) {
	name := c.Param("name")
	file_type := c.Param("type")[1:]

	info := client.HGetAll(name).Val()
	if filePath, ok := info[file_type]; ok {
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("open file %s error : %v", filePath, err)
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		fileSuffix, _ := info[file_type+"Suffix"]
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name+"."+fileSuffix))
		c.Header("Content-Type", "application/text/plain")
		c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
		c.Writer.Write([]byte(content))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "no such file",
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if username == setting.UserName && password == setting.PassWord {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			log.Fatalf("GenerateToken Error :%v", err)
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "GenerateToken Error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "login success",
			"token": token,
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "username or password wrong",
		})
	}

}
