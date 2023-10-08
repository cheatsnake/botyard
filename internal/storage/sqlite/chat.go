package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/cheatsnake/botyard/internal/entities/chat"
	"github.com/cheatsnake/botyard/pkg/exterr"
)

func (s *Storage) AddChat(chat *chat.Chat) error {
	q := `INSERT INTO chats (id, bot_id, user_id) VALUES (?, ?, ?)`

	_, err := s.conn.Exec(q, chat.Id, chat.BotId, chat.UserId)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't save chat to database.")
	}

	return nil
}

func (s *Storage) GetChat(id string) (*chat.Chat, error) {
	q := `SELECT bot_id, user_id FROM chats WHERE id = ?`
	c := chat.Chat{
		Id: id,
	}

	err := s.conn.QueryRow(q, id).Scan(&c.BotId, &c.UserId)
	if err == sql.ErrNoRows {
		return nil, exterr.ErrorNotFound("Chat not found.")
	}
	if err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't get chat from database.")
	}

	return &c, nil
}

func (s *Storage) GetChats(userId, botId string) ([]*chat.Chat, error) {
	q := `SELECT id FROM chats WHERE user_id = ? AND bot_id = ?`

	rows, err := s.conn.Query(q, userId, botId)
	if err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't get chats from database.")
	}
	defer rows.Close()

	chats := []*chat.Chat{}

	for rows.Next() {
		c := chat.Chat{
			UserId: userId,
			BotId:  botId,
		}
		if err := rows.Scan(&c.Id); err != nil {
			// TODO logger
			return nil, exterr.ErrorInternalServer("Can't retrive chat from database.")
		}
		chats = append(chats, &c)
	}

	if err := rows.Err(); err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't retrive chats from database.")
	}

	return chats, nil
}

func (s *Storage) DeleteChat(id string) error {
	q := `DELETE FROM chats WHERE id = ?`

	_, err := s.conn.Exec(q, id)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't delete chat from database.")
	}

	return nil
}

func (s *Storage) AddMessage(msg *chat.Message) error {
	q := `INSERT INTO messages (id, chat_id, sender_id, body, attachment_ids, timestamp) VALUES (?, ?, ?, ?, ?, ?)`
	attachmentIds := strings.Join(msg.AttachmentIds, ",")

	_, err := s.conn.Exec(q, msg.Id, msg.ChatId, msg.SenderId, msg.Body, attachmentIds, msg.Timestamp)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't save message to database.")
	}

	return nil
}

func (s *Storage) GetMessage(id string) (*chat.Message, error) {
	q := `SELECT chat_id, sender_id, body, attachment_ids, timestamp FROM messages WHERE id = ?`
	attachmentIds := ""
	msg := chat.Message{
		Id: id,
	}

	err := s.conn.QueryRow(q, id).Scan(&msg.ChatId, &msg.SenderId, &msg.Body, &attachmentIds, &msg.Timestamp)
	if err == sql.ErrNoRows {
		return nil, exterr.ErrorNotFound("Message not found.")
	}
	if err != nil {
		// TODO logger
		fmt.Println(err)
		return nil, exterr.ErrorInternalServer("Can't get message from database.")
	}

	if len(attachmentIds) > 0 {
		msg.AttachmentIds = strings.Split(attachmentIds, ",")
	}

	return &msg, nil
}

func (s *Storage) GetMessagesByChat(chatId, senderId string, page, limit int, since int64) (int, []*chat.Message, error) {
	q1 := `SELECT COUNT(*) FROM messages WHERE chat_id = ?`
	q2 := `SELECT id, sender_id, body, attachment_ids, timestamp FROM messages WHERE chat_id = ?`
	queryArgs := []interface{}{chatId}

	if senderId != "" {
		cond := " AND sender_id = ?"
		q1 += cond
		q2 += cond
		queryArgs = append(queryArgs, senderId)
	}

	if since != 0 {
		cond := " AND timestamp > ?"
		q1 += cond
		q2 += cond
		queryArgs = append(queryArgs, since)
	}

	var total int
	err := s.conn.QueryRow(q1, queryArgs...).Scan(&total)
	if err != nil {
		return total, nil, exterr.ErrorInternalServer("Can't count messages from database.")
	}

	q2 += " ORDER BY timestamp DESC LIMIT ? OFFSET ?"
	q2 = fmt.Sprintf("SELECT * FROM (%s) ORDER BY timestamp", q2) // reorder

	queryArgs = append(queryArgs, limit)
	queryArgs = append(queryArgs, (page-1)*limit)

	rows, err := s.conn.Query(q2, queryArgs...)
	if err != nil {
		// TODO logger
		return total, nil, exterr.ErrorInternalServer("Can't get messages from database.")
	}
	defer rows.Close()

	msgs := []*chat.Message{}

	for rows.Next() {
		attachmentIds := ""
		msg := chat.Message{
			ChatId: chatId,
		}

		if err := rows.Scan(&msg.Id, &msg.SenderId, &msg.Body, &attachmentIds, &msg.Timestamp); err != nil {
			// TODO logger
			return total, nil, exterr.ErrorInternalServer("Can't retrive message from database.")
		}

		if len(attachmentIds) > 0 {
			msg.AttachmentIds = strings.Split(attachmentIds, ",")
		}

		msgs = append(msgs, &msg)
	}

	if err := rows.Err(); err != nil {
		// TODO logger
		return total, nil, exterr.ErrorInternalServer("Can't retrive messages from database.")
	}

	return total, msgs, nil
}

func (s *Storage) DeleteMessagesByChat(chatId string) error {
	q := `DELETE FROM messages WHERE chat_id = ?`

	_, err := s.conn.Exec(q, chatId)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't delete messages from database.")
	}

	return nil
}
