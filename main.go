package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/web/", "./web")
	router.POST("/streamupload", streamUploadHandler)
	router.POST("/formupload", formUploadHandler)
	if err := router.Run("0.0.0.0:19090"); err != nil {
		panic(err)
	}
}

// streamUploadHandler is a handler for uploading a file.
func streamUploadHandler(c *gin.Context) {
	fileName := c.Query("filename")
	if fileName == "" {
		fileName = "file"
	}
	uploadedFileName := filepath.Join(getWorkingDir(), "uploads", fileName)

	if c.GetHeader("Content-Type") != "application/octet-stream" {
		err := fmt.Errorf("required octet-stream")
		c.AbortWithStatusJSON(400, map[string]string{"message": err.Error()})
		return
	}
	file, err := os.Create(uploadedFileName)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Create: " + err.Error()})
		return
	}
	defer file.Close()
	_, err = io.Copy(file, c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Copy: " + err.Error()})
		return
	}
	c.JSON(200, map[string]string{"message": "ok", "uploadedFileName": uploadedFileName})

}

// formUploadHandler is a handler for uploading a file.
func formUploadHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileName := filepath.Base(file.Filename)
	uploadedFileName := filepath.Join(getWorkingDir(), "uploads", fileName)
	c.SaveUploadedFile(file, uploadedFileName)
	c.JSON(200, map[string]string{"message": "ok", "uploadedFileName": uploadedFileName})
}

// getWorkingDir returns the current working directory.
func getWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
