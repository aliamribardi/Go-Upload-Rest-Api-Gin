package Controllers

import (
	"Go-Upload-Rest-Api-Gin/Database"
	"net/http"
	"Go-Upload-Rest-Api-Gin/Models"
	"github.com/gin-gonic/gin"
)

func Upload (c *gin.Context) {
	// Get the File
	name_file, err := c.FormFile("image")
	file_name := name_file.Filename

	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Failed to Upload",
		})
		return
	}

	upload := Models.Upload{NameFile: file_name}

	Database.DB.Create(&upload)

	err = c.SaveUploadedFile(name_file, "Assets/Upload/Image/" + name_file.Filename)

	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Failed to Upload",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"image": "Assets/Upload/Image/" + name_file.Filename,
		"LocalName": name_file.Filename,
		"message": "success",
	})

}

func UploadApi (c *gin.Context) {
	// var upload Models.Upload

	// Get the File
	name_file, err := c.FormFile("image")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Failed to Upload Image"})
		return
	}

	c.SaveUploadedFile(name_file, "Assets/Upload/Image/" + name_file.Filename)

}