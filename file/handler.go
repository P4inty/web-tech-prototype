package file

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"
	"webtech/prototype/db"

	"cloud.google.com/go/storage"
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

	upload, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	uri := c.Param("uri")

	file := findFile(c, uri)

	if file == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong on our side",
		})
		return
	}

	extension := filepath.Ext(upload.Filename)

	object := db.Bu.Object(uri + extension)
	writer := object.NewWriter(c)
	defer writer.Close()

	src, err := upload.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong on our side",
		})
		return
	}
	defer src.Close()

	if _, err := io.Copy(writer, src); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})

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
			log.Panicln(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong on our side",
			})
			break
		}
		var f File
		if err := doc.DataTo(&f); err != nil {
			log.Panicln(err.Error())
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

	uri := c.Param("uri")
	file := findFile(c, uri)

	if file == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Cant find a file with the specified uri",
		})
		return
	}

	iter := db.Bu.Objects(c, &storage.Query{
		Prefix: uri,
	})

	var oAttrs *storage.ObjectAttrs
	for {
		attrs, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong on our side",
			})
			break
		}

		oAttrs = attrs
	}

	o, err := db.Bu.Object(oAttrs.Name).NewReader(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong on our side",
		})
	}
	defer o.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(o)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+file.Name+filepath.Ext(oAttrs.Name))
	c.Header("Content-Type", "application/octet-stream")
	c.Writer.Write(buf.Bytes())

}

func findFile(c *gin.Context, uri string) *File {
	var file File
	iter := db.Db.Collection("files").Where("Uri", "==", uri).Documents(c)
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

		if err := doc.DataTo(&file); err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong on our side",
			})
			break
		}
	}
	return &file
}
