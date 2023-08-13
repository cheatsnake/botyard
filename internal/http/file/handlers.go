package file

import (
	"botyard/internal/entities/file"
	"botyard/internal/storage"
	"botyard/internal/tools/ulid"
	"botyard/pkg/extlib"
	"errors"
	"fmt"
	"mime/multipart"
	"path"

	"github.com/gofiber/fiber/v2"
)

type handlers struct {
	service *service
}

func Handlers(s storage.Storage) *handlers {
	return &handlers{
		service: &service{
			store: s,
		},
	}
}

func (h *handlers) LoadFiles(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	files := form.File[fileFieldName]

	if len(files) > maxFilesAmount {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("too many files, max allowed amount is %d", maxFilesAmount),
		)
	}

	result := make([]*file.File, 0, len(files))

	// TODO parallel upload
	for _, f := range files {
		filePath, contentType, err := fileSaver(c, f)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		newFile := file.New(filePath, contentType)
		err = h.service.store.AddFile(newFile)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
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
