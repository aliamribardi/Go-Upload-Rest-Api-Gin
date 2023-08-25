package Controllers

import (
	"Go-Upload-Rest-Api-Gin/Database"
	"net/http"
	"Go-Upload-Rest-Api-Gin/Models"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	var upload []Models.Upload

	Database.DB.Find(&upload)

	if upload == nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{"message": "Data not Found"})
		return
	}
	// data := string
	// for i := 0 ; i < len(upload) ; i++ {
	// 	data[i]
	// }

	c.HTML(http.StatusOK, "index.html", gin.H{"image": upload})

}

func Upload (c *gin.Context) {
	// Get the File
	name_file, err := c.FormFile("image")
	file_name := "Assets/Upload/Image/" + name_file.Filename

	if err != nil {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"message": "Failed to Upload",
		})
		return
	}

	upload := Models.Upload{NameFile: file_name}

	Database.DB.Create(&upload)

	err = c.SaveUploadedFile(name_file, "Assets/Upload/Image/" + name_file.Filename)

	if err != nil {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"message": "Failed to Upload",
		})
		return
	}

	c.HTML(http.StatusOK, "upload.html", gin.H{
		"image": "Assets/Upload/Image/" + name_file.Filename,
		"LocalName": name_file.Filename,
		"message": "success",
	})

}

func UploadApi (c *gin.Context) {
	// Get the File
	name_file, err := c.FormFile("image")
	file_name := "Assets/Upload/Image/" + name_file.Filename

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Upload",
		})
		return
	}

	upload := Models.Upload{NameFile: file_name}

	Database.DB.Create(&upload)

	err = c.SaveUploadedFile(name_file, "Assets/Upload/Image/" + name_file.Filename)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Upload",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"image": "Assets/Upload/Image/" + name_file.Filename,
	})

}