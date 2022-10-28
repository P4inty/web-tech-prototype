package file

import (
	"log"
	"net/http"
	"time"
	"webtech/prototype/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
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
	file.CreatedAt = time.Now()
	file.UpdatedAt = time.Now()

	_, _, err := db.Db.Collection("files").Add(c, file)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uri": file.Uri})
}

func Upload(c *gin.Context) {
	/*
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
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"message": "Unable to save the file",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"message": "Your file has been successfully uploaded.",
				})
	*/
}

func All(c *gin.Context) {

	var files []File

	iter := db.Db.Collection("files").Documents(c)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong on our side",
			})
			break
		}
		var f File
		if err := doc.DataTo(&f); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong on our side",
			})
			break
		}
		files = append(files, f)
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})

}

func Download(c *gin.Context) {
	/*
		var file File
		uri := c.Param("uri")

		if err := db.DB.First(&file, "uri = ?", uri).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Cant find a file with the specified uri",
			})
			return
		}

		path, err := os.Getwd()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong on our side",
			})
			return
		}

		matches, err := filepath.Glob(path + "/public/upload/" + uri + "*")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong on our side",
			})
		}

		if len(matches) == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Could not find a file with the given uri",
			})
		}
		extension := filepath.Ext(matches[0])
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+file.Name+extension)
		c.Header("Content-Type", "application/octet-stream")
		c.File(matches[0])
	*/
}
