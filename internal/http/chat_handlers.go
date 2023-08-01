package http

import (
	"botyard/internal/chat"
	"botyard/pkg/extlib"
	"botyard/pkg/ulid"
	"fmt"
	"mime/multipart"
	"path"

	"github.com/gofiber/fiber/v2"
)

const (
	maxImageSize   = 2 * 1024 * 1024  // 2 MB
	maxAudioSize   = 5 * 1024 * 1024  // 5 MB
	maxVideoSize   = 25 * 1024 * 1024 // 25 MB
	maxFileSize    = 10 * 1024 * 1024 // 10 MB
	maxFilesAmount = 10
)

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

	// TODO add checks for file extensions
	images := form.File["image"]
	videos := form.File["video"]
	audios := form.File["audio"]
	files := form.File["file"]

	totalFiles := len(images) + len(videos) + len(audios) + len(files)
	if totalFiles > maxFilesAmount {
		return newErr(
			fmt.Errorf("too many files, max allowed amount is %d", maxFilesAmount),
			fiber.StatusRequestEntityTooLarge,
		)
	}

	result := make([]*chat.File, 0, totalFiles)

	// TODO goroutines for parallel upload
	for _, file := range images {
		filePath, contentType, err := fileSaver(c, file, "images", maxImageSize)
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

	for _, file := range videos {
		filePath, contentType, err := fileSaver(c, file, "videos", maxVideoSize)
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

	for _, file := range audios {
		filePath, contentType, err := fileSaver(c, file, "audios", maxAudioSize)
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

	for _, file := range files {
		filePath, contentType, err := fileSaver(c, file, "files", maxFileSize)
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

func fileSaver(c *fiber.Ctx, file *multipart.FileHeader, fileType string, maxSize int) (string, string, error) {
	var filePath, contentType string

	if file.Size > maxImageSize {
		return filePath, contentType, fmt.Errorf(
			"max allowed size for %s is %d bytes, but got %d",
			fileType, maxSize, file.Size,
		)
	}

	contentType = file.Header["Content-Type"][0]
	ext, err := extlib.ExtensionFromContentType(contentType)
	if err != nil {
		return filePath, contentType, err
	}

	filePath = path.Join(".", "store", fileType, ulid.New()+ext)
	if err := c.SaveFile(file, filePath); err != nil {
		return filePath, contentType, err
	}

	return filePath, contentType, nil
}
