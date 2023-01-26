package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.POST("/streamupload", streamuploadHandler)
	if err := engine.Run("0.0.0.0:19090"); err != nil {
		panic(err)
	}
}

// streamuploadHandler is a handler for uploading a file.
func streamuploadHandler(c *gin.Context) {
	uploadedFileName := filepath.Join(getWorkingDir(), "file")

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
	c.JSON(200, map[string]string{"message": "ok stream"})
}

// getWorkingDir returns the current working directory.
func getWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
