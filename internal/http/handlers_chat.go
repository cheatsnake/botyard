package http

import (
	"botyard/internal/chat"
	"botyard/pkg/extlib"
	"botyard/pkg/ulid"
	"errors"
	"fmt"
	"mime/multipart"
	"path"

	"github.com/gofiber/fiber/v2"
)

const filesFolder = "stock"
const maxFilesAmount = 10
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

type chatBody struct {
	chat.Chat
	Id struct{} `json:"-"`
}

type messageBody struct {
	chat.Message
	Id        struct{} `json:"-"`
	Timestamp struct{} `json:"-"`
}

func (s *Server) createChat(c *fiber.Ctx) error {
	b := new(chatBody)

	if err := c.BodyParser(b); err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	chat := chat.New(b.MemberIds, s.Storage)
	err := s.Storage.AddChat(chat)
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusCreated).JSON(chat)
}

func (s *Server) sendMessage(c *fiber.Ctx) error {
	b := new(messageBody)

	if err := c.BodyParser(b); err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	chat, err := s.Storage.GetChat(b.ChatId)
	if err != nil {
		return newErr(err, fiber.StatusNotFound)
	}

	err = chat.SendMessage(b.SenderId, b.Body, b.FileIds)
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	return c.JSON(response{Message: "message sended"})
}

func (s *Server) getMessages(c *fiber.Ctx) error {
	chatId := c.Params("id", "")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)

	chat, err := s.Storage.GetChat(chatId)
	if err != nil {
		return newErr(err, fiber.StatusNotFound)
	}

	result, err := chat.GetMessages(page, limit)
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	return c.JSON(result)
}

func (s *Server) loadFiles(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	files := form.File[fileFieldName]

	if len(files) > maxFilesAmount {
		return newErr(
			fmt.Errorf("too many files, max allowed amount is %d", maxFilesAmount),
			fiber.StatusBadRequest,
		)
	}

	result := make([]*chat.File, 0, len(files))

	// TODO parallel upload
	for _, file := range files {
		filePath, contentType, err := fileSaver(c, file)
		if err != nil {
			return newErr(err, fiber.StatusBadRequest)
		}

		f := chat.NewFile(filePath, contentType)
		err = s.Storage.AddFile(f)
		if err != nil {
			return newErr(err, fiber.StatusBadRequest)
		}

		result = append(result, f)
	}

	return c.JSON(result)
}

func (s *Server) clearChat(c *fiber.Ctx) error {
	chatId := c.Params("id", "")

	chat, err := s.Storage.GetChat(chatId)
	if err != nil {
		return newErr(err, fiber.StatusNotFound)
	}

	err = chat.Clear()
	if err != nil {
		return newErr(err, fiber.StatusBadRequest)
	}

	return c.JSON(response{Message: "chat cleared"})
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
