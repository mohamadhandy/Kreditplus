package handlers

import (
	"context"
	"fmt"
	"io"
	"kredit-plus/helper"
	"kredit-plus/models"
	"kredit-plus/usecases"
	"net/http"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type KonsumenHandlerInterface interface {
	CreateUser(c *gin.Context)
	UploadImageKonsumen(c *gin.Context)
}

type konsumenHandler struct {
	KonsumenUseCase usecases.KonsumenUsecaseInterface
}

func InitKonsumenHandler(u usecases.KonsumenUsecaseInterface) KonsumenHandlerInterface {
	return &konsumenHandler{
		KonsumenUseCase: u,
	}
}

func (h *konsumenHandler) CreateUser(c *gin.Context) {
	konsumen := models.KonsumenRequest{}
	c.BindJSON(&konsumen)
	konsumenResult := h.KonsumenUseCase.CreateUser(konsumen)
	c.JSON(konsumenResult.StatusCode, konsumenResult)
}

func (h *konsumenHandler) UploadImageKonsumen(c *gin.Context) {
	// Parse the request to get the MultipartForm
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error parsing the request: %s", err.Error()))
		return
	}

	// Get the slice of files sent by the client with the key "image"
	files := c.Request.MultipartForm.File["image"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, "No images uploaded.")
		return
	}

	uploadImages := []string{}

	// Loop through the uploaded files
	for _, file := range files {
		// Generate a unique filename for the uploaded image
		filename := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))

		// Create a new Google Cloud Storage client
		ctx := context.Background()
		storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICE_JSON")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to create Google Cloud Storage client: %s", err.Error()))
			return
		}

		// Open the file
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to open file: %s", err.Error()))
			return
		}
		defer src.Close()

		// Create a new bucket handle
		bucketName := os.Getenv("FIREBASE_BUCKET")
		bucket := storageClient.Bucket(bucketName)

		// Create an object writer using the bucket and filename
		obj := bucket.Object(filename)
		w := obj.NewWriter(ctx)

		// Copy the file data to the object writer
		if _, err := io.Copy(w, src); err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to upload file to Firebase Storage: %s", err.Error()))
			return
		}

		// Close the object writer to ensure the data is flushed to Firebase Storage
		if err := w.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to close object writer: %s", err.Error()))
			return
		}

		// Generate the download URL for the uploaded image
		downloadURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s", bucketName, filename)
		result, err := helper.FetchFirebaseImage(downloadURL)
		fmt.Println("res", result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		uploadImages = append(uploadImages, result)
	}

	newImageURL := helper.ImageURL{
		FotoKTP:    uploadImages[0],
		FotoSelfie: uploadImages[1],
	}
	response := helper.Response{
		Data:       newImageURL,
		StatusCode: http.StatusOK,
		Message:    "Upload image success",
	}
	c.JSON(http.StatusOK, response)
}
