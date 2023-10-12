package sqlite

import (
	"database/sql"

	"github.com/cheatsnake/botyard/internal/entities/bot"
	"github.com/cheatsnake/botyard/pkg/exterr"
)

func (s *Storage) AddBot(newBot *bot.Bot) error {
	q1 := `SELECT COUNT(*) FROM bots WHERE name = ?`

	var total int
	err := s.conn.QueryRow(q1, newBot.Name).Scan(&total)
	if err != nil {
		return exterr.ErrorInternalServer("Can't count bots from database.")
	}

	if total != 0 {
		return exterr.ErrorBadRequest("Bot with this name already exists.")
	}

	q2 := `INSERT INTO bots (id, name, description, avatar) VALUES (?, ?, ?, ?)`
	_, err = s.conn.Exec(q2, newBot.Id, newBot.Name, newBot.Description, newBot.Avatar)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't save bot to database.")
	}

	return nil
}

func (s *Storage) GetBot(id string) (*bot.Bot, error) {
	q := `SELECT id, name, description, avatar FROM bots WHERE id = ?`

	var b bot.Bot

	err := s.conn.QueryRow(q, id).Scan(&b.Id, &b.Name, &b.Description, &b.Avatar)
	if err == sql.ErrNoRows {
		return nil, exterr.ErrorNotFound("Bot not found.")
	}
	if err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't get bot from database.")
	}

	return &b, nil
}

func (s *Storage) GetAllBots() ([]*bot.Bot, error) {
	q := `SELECT id, name, description, avatar FROM bots`

	rows, err := s.conn.Query(q)
	if err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't get bots from database.")
	}
	defer rows.Close()

	bots := []*bot.Bot{}

	for rows.Next() {
		var b bot.Bot
		if err := rows.Scan(&b.Id, &b.Name, &b.Description, &b.Avatar); err != nil {
			// TODO logger
			return nil, exterr.ErrorInternalServer("Can't retrive bot from database.")
		}
		bots = append(bots, &b)
	}

	if err := rows.Err(); err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't retrive bots from database.")
	}

	return bots, nil
}

func (s *Storage) EditBot(editedBot *bot.Bot) error {
	q := `UPDATE bots SET name = ?, description = ?, avatar = ? WHERE id = ?`

	_, err := s.conn.Exec(q, editedBot.Name, editedBot.Description, editedBot.Avatar, editedBot.Id)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't save bot to database.")
	}

	return nil
}

func (s *Storage) DeleteBot(id string) error {
	q := `DELETE FROM bots WHERE id = ?`

	_, err := s.conn.Exec(q, id)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't delete bot from database.")
	}

	return nil
}

func (s *Storage) SaveCommand(cmd *bot.Command) error {
	q := `INSERT OR REPLACE INTO bot_commands (id, bot_id, alias, description) VALUES (?, ?, ?, ?)`

	_, err := s.conn.Exec(q, cmd.Id, cmd.BotId, cmd.Alias, cmd.Description)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't save command to database.")
	}

	return nil
}

func (s *Storage) GetCommands(botId string) ([]*bot.Command, error) {
	q := `SELECT id, alias, description FROM bot_commands WHERE bot_id = ?`

	rows, err := s.conn.Query(q, botId)
	if err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't get commands from database.")
	}
	defer rows.Close()

	cmds := []*bot.Command{}

	for rows.Next() {
		cmd := bot.Command{
			BotId: botId,
		}

		if err := rows.Scan(&cmd.Id, &cmd.Alias, &cmd.Description); err != nil {
			// TODO logger
			return nil, exterr.ErrorInternalServer("Can't retrive command from database.")
		}

		cmds = append(cmds, &cmd)
	}

	if err := rows.Err(); err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't retrive commands from database.")
	}

	return cmds, nil
}

func (s *Storage) GetCommand(id string) (*bot.Command, error) {
	q := `SELECT bot_id, alias, description FROM bot_commands WHERE id = ?`
	cmd := &bot.Command{
		Id: id,
	}

	err := s.conn.QueryRow(q, id).Scan(&cmd.BotId, &cmd.Alias, &cmd.Description)
	if err == sql.ErrNoRows {
		return nil, exterr.ErrorNotFound("Command not found.")
	}
	if err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't get command from database.")
	}

	return cmd, nil
}

func (s *Storage) DeleteCommand(id string) error {
	q := `DELETE FROM bot_commands WHERE id = ?`
	_, err := s.conn.Exec(q, id)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't delete bot command from database.")
	}

	return nil
}

func (s *Storage) DeleteCommandsByBot(botId string) error {
	q := `DELETE FROM bot_commands WHERE bot_id = ?`

	_, err := s.conn.Exec(q, botId)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't delete bot commands from database.")
	}

	return nil
}

func (s *Storage) GetKey(botId string) (*bot.Key, error) {
	q := `SELECT token FROM bot_keys WHERE bot_id = ?`
	bk := &bot.Key{
		BotId: botId,
	}

	err := s.conn.QueryRow(q, botId).Scan(&bk.Token)
	if err == sql.ErrNoRows {
		return nil, exterr.ErrorNotFound("Key not found.")
	}
	if err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't get key from database.")
	}

	return bk, nil
}

func (s *Storage) SaveKey(botKey *bot.Key) error {
	q := `INSERT OR REPLACE INTO bot_keys (bot_id, token) VALUES (?, ?)`

	_, err := s.conn.Exec(q, botKey.BotId, botKey.Token)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't save key to database.")
	}

	return nil
}

func (s *Storage) DeleteKey(botId string) error {
	q := `DELETE FROM bot_keys WHERE bot_id = ?`

	_, err := s.conn.Exec(q, botId)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't delete bot key from database.")
	}

	return nil
}

func (s *Storage) GetWebhook(botId string) (*bot.Webhook, error) {
	q := `SELECT url, secret FROM bot_webhooks WHERE bot_id = ?`
	bw := &bot.Webhook{
		BotId: botId,
	}

	err := s.conn.QueryRow(q, botId).Scan(&bw.Url, &bw.Secret)
	if err == sql.ErrNoRows {
		return nil, exterr.ErrorNotFound("Webhook not found.")
	}
	if err != nil {
		// TODO logger
		return nil, exterr.ErrorInternalServer("Can't get webhook from database.")
	}

	return bw, nil
}

func (s *Storage) SaveWebhook(webhook *bot.Webhook) error {
	q := `INSERT OR REPLACE INTO bot_webhooks (bot_id, url, secret) VALUES (?, ?, ?)`

	_, err := s.conn.Exec(q, webhook.BotId, webhook.Url, webhook.Secret)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't save webhook to database.")
	}

	return nil
}

func (s *Storage) DeleteWebhook(botId string) error {
	q := `DELETE FROM bot_webhooks WHERE bot_id = ?`

	_, err := s.conn.Exec(q, botId)
	if err != nil {
		// TODO logger
		return exterr.ErrorInternalServer("Can't delete bot webhook from database.")
	}

	return nil
}
