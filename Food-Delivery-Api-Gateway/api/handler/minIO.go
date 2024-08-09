package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

// @Summary Upload a file to MinIO
// @Description Upload a file to MinIO
// @Tags MinIO
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "File to upload"
// @Param filename formData string true "Filename"
// @Success 200 {object} string "Success"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /minio/upload [post]
func (h *Handler) UploadFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	bucketName := "testbucket"
	location := "us-east-1" 

	exists, err := h.miniIO.BucketExists(c, bucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if bucket exists"})
		return
	}

	// Create the bucket if it doesn't exist
	if !exists {
		err = h.miniIO.MakeBucket(c, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bucket"})
			return
		}
	}

	objectName := "images/" + time.Now().Format("20060102150405") + "-" + c.Request.FormValue("filename")

	_, err = h.miniIO.PutObject(c, bucketName, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	presignedURL, err := h.miniIO.PresignedGetObject(c, bucketName, objectName, 24*time.Hour, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": presignedURL.String()})
}
