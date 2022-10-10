package file

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"webtech/prototype/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadMeta(c *gin.Context) {
	var meta MetaData
	if err := c.ShouldBindJSON(&meta); err != nil {
		log.Default().Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var file = Meta2File(&meta)
	var uri = uuid.New().String()
	file.Uri = uri

	if err := db.DB.Create(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uri": file.Uri})
}

func Upload(c *gin.Context) {
	upload, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	uri := c.Param("uri")

	var file File

	if err := db.DB.First(&file, "uri = ?", uri).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Cant find a file with the specified uri",
		})
		return
	}

	extension := filepath.Ext(upload.Filename)

	path, err := os.Getwd()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong on our side",
		})
		return
	}

	if err := c.SaveUploadedFile(upload, path+"/public/upload/"+uri+extension); err != nil {
		log.Default().Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}
