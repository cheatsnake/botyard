package handlers

import (
	"botyard/internal/entities/file"
	"botyard/internal/services"
	"botyard/internal/tools/ulid"
	"botyard/pkg/extlib"
	"errors"
	"fmt"
	"mime/multipart"
	"path"

	"github.com/gofiber/fiber/v2"
)

const filesFolder = "stock"
const maxFiles = 10
const fileFieldName = "file"

var maxFileSizes = map[string]int64{
	"images": 2 * 1024 * 1024,  // 2 MB
	"audios": 5 * 1024 * 1024,  // 5 MB
	"videos": 25 * 1024 * 1024, // 25 MB
	"files":  10 * 1024 * 1024, // 10 MB
}

var knownContentTypes = map[string]string{
	"image/gif":     "images",
	"image/jpeg":    "images",
	"image/png":     "images",
	"image/svg+xml": "images",
	"image/webp":    "images",

	"video/mp4":       "videos",
	"video/webm":      "videos",
	"video/ogg":       "videos",
	"video/quicktime": "videos",
	"video/x-flv":     "videos",

	"audio/mpeg": "audios",
	"audio/ogg":  "audios",
	"audio/wav":  "audios",
	"audio/aac":  "audios",
}

type FileHandlers struct {
	service *services.FileService
}

func NewFileHandlers(s *services.FileService) *FileHandlers {
	return &FileHandlers{
		service: s,
	}
}

func (h *FileHandlers) LoadMany(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return extlib.ErrorBadRequest(err.Error())
	}

	files := form.File[fileFieldName]

	if len(files) > maxFiles {
		return extlib.ErrorBadRequest(
			fmt.Sprintf("too many files, max allowed amount is %d", maxFiles),
		)
	}

	result := make([]*file.File, 0, len(files))

	// TODO parallel upload
	for _, f := range files {
		filePath, contentType, err := fileSaver(c, f)
		if err != nil {
			return extlib.ErrorBadRequest(err.Error())
		}

		newFile, err := file.New(filePath, contentType)
		if err != nil {
			return extlib.ErrorBadRequest(err.Error())
		}

		err = h.service.AddFile(newFile)
		if err != nil {
			return extlib.ErrorBadRequest(err.Error())
		}

		result = append(result, newFile)
	}

	return c.JSON(result)
}

func fileSaver(c *fiber.Ctx, file *multipart.FileHeader) (string, string, error) {
	contentType := file.Header["Content-Type"][0]
	ext, err := extlib.ExtensionFromContentType(contentType)
	if err != nil {
		return "", "", err
	}

	fileType, ok := knownContentTypes[contentType]
	if !ok {
		fileType = "files"
	}

	maxFileSize, ok := maxFileSizes[fileType]
	if !ok {
		return "", "", errors.New("failed to recognize the file type")
	}

	if file.Size > maxFileSize {
		return "", "", fmt.Errorf(
			"max allowed size for %s are %d bytes, but got %d",
			fileType, maxFileSize, file.Size,
		)
	}

	filePath := path.Join(".", filesFolder, fileType, ulid.New()+ext)
	if err := c.SaveFile(file, filePath); err != nil {
		return filePath, contentType, err
	}

	return filePath, contentType, nil
}
