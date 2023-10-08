package filehandlers

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"

	"github.com/cheatsnake/botyard/internal/config"
	"github.com/cheatsnake/botyard/internal/entities/file"
	"github.com/cheatsnake/botyard/internal/services/fileservice"
	"github.com/cheatsnake/botyard/internal/tools/ulid"
	"github.com/cheatsnake/botyard/pkg/exterr"
	"github.com/cheatsnake/botyard/pkg/extlib"

	"github.com/gofiber/fiber/v2"
)

const fileTag = "file"

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

type Handlers struct {
	service *fileservice.Service
}

func New(s *fileservice.Service) *Handlers {
	return &Handlers{
		service: s,
	}
}

func (h *Handlers) UploadFiles(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return exterr.ErrorBadRequest(err.Error())
	}

	files := form.File[fileTag]

	if len(files) == 0 {
		return exterr.ErrorBadRequest("No files found to upload.")
	}

	if len(files) > config.Global.Limits.Message.MaxAttachedFiles {
		return exterr.ErrorBadRequest(
			fmt.Sprintf(
				"Too many files, max allowed amount is %d.",
				config.Global.Limits.Message.MaxAttachedFiles,
			),
		)
	}

	result := make([]*file.File, 0, len(files))

	// TODO parallel upload
	for _, f := range files {
		filePath, fileName, contentType, fileSize, err := fileSaver(c, f)
		if err != nil {
			return err
		}

		newFile, err := h.service.AddFile(filePath, fileName, contentType, fileSize)
		if err != nil {
			return err
		}

		result = append(result, newFile)
	}

	return c.JSON(result)
}

func fileSaver(c *fiber.Ctx, file *multipart.FileHeader) (string, string, string, int, error) {
	contentType := file.Header["Content-Type"][0]
	ext, err := extlib.ExtensionFromContentType(contentType)
	if err != nil {
		return "", "", "", 0, err
	}

	fileType, ok := knownContentTypes[contentType]
	if !ok {
		fileType = "files"
	}

	maxFileSize := defineMaxFileSize(fileType)

	if file.Size > maxFileSize {
		return "", "", "", 0, exterr.ErrorBadRequest(
			fmt.Sprintf("Max allowed size for %s are %d bytes, but got %d.",
				fileType, maxFileSize, file.Size,
			))
	}

	filePath := path.Join(".", os.Getenv("FILES_FOLDER"), fileType, ulid.New()+ext)
	if err := c.SaveFile(file, filePath); err != nil {
		return filePath, file.Filename, contentType, int(file.Size), err
	}

	return filePath, file.Filename, contentType, int(file.Size), nil
}

func defineMaxFileSize(fileType string) int64 {
	if fileType == "images" {
		return int64(config.Global.Limits.File.MaxImageSize)
	}

	if fileType == "audios" {
		return int64(config.Global.Limits.File.MaxAudioSize)
	}

	if fileType == "videos" {
		return int64(config.Global.Limits.File.MaxVideoSize)
	}

	return int64(config.Global.Limits.File.MaxFileSize)
}
