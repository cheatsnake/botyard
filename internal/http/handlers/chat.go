package handlers

import (
	"botyard/internal/chat"
	"botyard/internal/storage"
	"botyard/internal/tools/ulid"
	"botyard/pkg/extlib"
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

type Chat struct {
	store storage.Storage
}

func NewChat(store storage.Storage) *Chat {
	return &Chat{
		store: store,
	}
}

func (ch *Chat) CreateChat(c *fiber.Ctx) error {
	userId := fmt.Sprintf("%v", c.Locals("userId"))
	b := new(struct {
		BotId string `json:"botId"`
	})

	if err := c.BodyParser(b); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err := ch.store.GetBot(b.BotId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	chat := chat.New(userId, b.BotId, ch.store)
	err = ch.store.AddChat(chat)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(chat)
}

func (ch *Chat) SendMessage(c *fiber.Ctx) error {
	b := new(struct {
		chat.Message
		Id        struct{} `json:"-"`
		Timestamp struct{} `json:"-"`
	})

	if err := c.BodyParser(b); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	chat, err := ch.store.GetChat(b.ChatId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	err = chat.SendMessage(b.SenderId, b.Body, b.FileIds)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(response{Message: "message sended"})
}

func (ch *Chat) GetMessages(c *fiber.Ctx) error {
	chatId := c.Params("id", "")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)

	chat, err := ch.store.GetChat(chatId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	result, err := chat.GetMessages(page, limit)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(result)
}

func (ch *Chat) LoadFiles(c *fiber.Ctx) error {

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

	result := make([]*chat.File, 0, len(files))

	// TODO parallel upload
	for _, file := range files {
		filePath, contentType, err := fileSaver(c, file)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		f := chat.NewFile(filePath, contentType)
		err = ch.store.AddFile(f)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result = append(result, f)
	}

	return c.JSON(result)
}

func (ch *Chat) ClearChat(c *fiber.Ctx) error {
	chatId := c.Params("id", "")

	chat, err := ch.store.GetChat(chatId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	err = chat.Clear()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
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
