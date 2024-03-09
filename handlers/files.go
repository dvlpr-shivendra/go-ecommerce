package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFilesUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
		files := form.File["files[]"]

		for _, file := range files {
			fmt.Println(file.Filename)

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, "storage/" + file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}