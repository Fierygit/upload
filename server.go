package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	uploadDir := "upload_files"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		panic(err)
	}

	r := gin.Default()

	// 设置更大的文件上传限制（针对multipart表单）
	r.MaxMultipartMemory = 8 << 30 // 8GB

	r.PUT("/upload/*filename", func(c *gin.Context) {
		// 记录请求开始时间
		startTime := time.Now()
		log.Printf("[%s] 开始处理上传请求", startTime.Format("15:04:05.000"))

		fileName := strings.TrimPrefix(c.Param("filename"), "/")
		fileName = strings.ReplaceAll(fileName, "/", "_")
		if fileName == "" {
			fileName = "file"
		}

		timestamp := time.Now().Format("20060102_150405")
		saveName := fmt.Sprintf("%s_%s", timestamp, fileName)
		savePath := filepath.Join(uploadDir, saveName)

		// 创建文件
		log.Printf("创建文件: %s", savePath)
		out, err := os.Create(savePath)
		if err != nil {
			log.Printf("文件创建失败: %v", err)
			c.String(500, "无法保存文件: %v", err)
			return
		}
		defer out.Close()

		// 流式写入
		log.Printf("开始写入数据")
		written, err := out.ReadFrom(c.Request.Body)
		if err != nil {
			log.Printf("写入失败: 已写入 %d 字节，错误: %v", written, err)
			c.String(500, "保存失败: %v", err)
			return
		}
		log.Printf("写入完成: 总计 %d 字节", written)

		// 生成响应
		downloadURL := fmt.Sprintf("http://%s/download/%s", c.Request.Host, saveName)
		log.Printf("准备返回响应: %s", downloadURL)
		c.String(200, "%s\n", downloadURL)
		log.Printf("[%s] 响应已发送", time.Now().Format("15:04:05.000"))
	})

	r.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filePath := filepath.Join(uploadDir, filename)
		c.FileAttachment(filePath, filename)
	})

	// 使用自定义HTTP服务器配置
	srv := &http.Server{
		Addr:         ":7070",
		Handler:      r,
		ReadTimeout:  60 * time.Minute,
		WriteTimeout: 60 * time.Minute,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
