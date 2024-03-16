package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type FileUpload struct {
	Files []*multipart.FileHeader `form:"files[]" binding:"required"`
}

func HandleFilesUpload(c *gin.Context) {
	var form FileUpload

	// Parse the form data including the files
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Iterate through each file
	for _, fileHeader := range form.Files {
		// Access the original file name
		originalFileName := fileHeader.Filename

		// Save the file
		err := c.SaveUploadedFile(fileHeader, "uploads/"+originalFileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// Process the uploaded file as needed
		fmt.Printf("Original File Name: %s, File saved: %s\n", originalFileName, originalFileName)
	}

	// Send a response
	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})
}
