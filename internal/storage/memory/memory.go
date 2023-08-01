// Package memory represents a build-in in-memory database for development
// and testing purposes. Don't use it in production!
package memory

import (
	"botyard/internal/bot"
	"botyard/internal/chat"
	"botyard/internal/user"
	"botyard/pkg/extlib"
	"errors"
	"sync"

	"golang.org/x/exp/slices"
)

type Storage struct {
	bots     []*bot.Bot
	users    []*user.User
	chats    []*chat.Chat
	messages []*chat.Message
	files    []*chat.File
	mu       sync.Mutex
}

func New() *Storage {
	return &Storage{
		bots:     make([]*bot.Bot, 0, 10),
		users:    make([]*user.User, 0, 10),
		chats:    make([]*chat.Chat, 0, 10),
		messages: make([]*chat.Message, 0, 100),
		files:    make([]*chat.File, 0, 10),
	}
}

func (s *Storage) AddBot(bot *bot.Bot) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	candidate, _ := s.GetBot(bot.Name)
	if candidate != nil {
		return errors.New("bot with this name already exists")
	}

	s.bots = append(s.bots, bot)
	return nil
}

func (s *Storage) GetBot(id string) (*bot.Bot, error) {
	for _, bot := range s.bots {
		if bot.Id == id {
			return bot, nil
		}
	}

	return nil, errors.New("bot not found")
}

func (s *Storage) AddUser(user *user.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.users = append(s.users, user)
	return nil
}

func (s *Storage) GetUser(id string) (*user.User, error) {
	for _, user := range s.users {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, errors.New("bot not found")
}

func (s *Storage) AddChat(chat *chat.Chat) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.chats = append(s.chats, chat)
	return nil
}

func (s *Storage) GetChat(id string) (*chat.Chat, error) {
	for _, chat := range s.chats {
		if chat.Id == id {
			return chat, nil
		}
	}

	return nil, errors.New("chat not found")
}

func (s *Storage) FindChat(userId, botId string) (*chat.Chat, error) {
	for _, chat := range s.chats {
		if slices.Contains(chat.MemberIds, userId) && slices.Contains(chat.MemberIds, botId) {
			return chat, nil
		}
	}

	return nil, errors.New("chat not found")
}

func (s *Storage) AddMessage(msg *chat.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.messages = append(s.messages, msg)

	return nil
}

func (s *Storage) GetMessagesByChat(chatId string, page, limit int) (int, []*chat.Message, error) {
	if page == 0 || limit == 0 {
		return 0, nil, nil
	}

	chatMsgs := extlib.SliceFilter(extlib.SliceReverse(s.messages), 10, func(m *chat.Message) bool {
		return m.ChatId == chatId
	})

	if page == 1 && limit >= len(chatMsgs) {
		msgs := make([]chat.Message, len(chatMsgs))
		for i, msg := range chatMsgs {
			msgs[i] = *msg
		}

		return len(chatMsgs), chatMsgs, nil
	}

	return len(chatMsgs), extlib.SlicePaginate(chatMsgs, page, limit), nil
}

func (s *Storage) DeleteMessagesByChat(chatId string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	filtered := extlib.SliceFilter(s.messages, len(s.messages), func(m *chat.Message) bool {
		return m.ChatId != chatId
	})

	s.messages = filtered
	return nil
}

func (s *Storage) AddFile(file *chat.File) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.files = append(s.files, file)
	return nil
}

func (s *Storage) GetFile(id string) (*chat.File, error) {
	for _, file := range s.files {
		if file.Id == id {
			return file, nil
		}
	}

	return nil, errors.New("file not found")
}

func (s *Storage) GetFiles(ids []string) ([]*chat.File, error) {
	files := make([]*chat.File, 0, len(ids))

	for _, file := range s.files {
		if slices.Contains(ids, file.Id) {
			files = append(files, file)
		}
	}

	return files, nil
}

func (s *Storage) DeleteFile(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	idx := slices.IndexFunc(s.files, func(f *chat.File) bool {
		return f.Id == id
	})

	if idx == -1 {
		return errors.New("file not found")
	}

	s.messages = extlib.SliceRemoveElement(s.messages, idx)
	return nil
}
