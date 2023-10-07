package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	// conn is database connection reference
	conn *sql.DB
}

func New(path string) (*Storage, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return &Storage{conn: conn}, nil
}

func (s *Storage) InitTables() error {
	botTable := `CREATE TABLE IF NOT EXISTS bots (
		id TEXT NOT NULL CHECK(LENGTH(id) <= 64),
		name TEXT NOT NULL CHECK(LENGTH(name) <= 32),
		description TEXT CHECK(LENGTH(description) <= 512),
		avatar TEXT CHECK(LENGTH(avatar) <= 256),
		PRIMARY KEY (id)
	);`

	commandTable := `CREATE TABLE IF NOT EXISTS bot_commands (
		bot_id TEXT NOT NULL CHECK(LENGTH(bot_id) <= 64),
		alias TEXT NOT NULL CHECK(LENGTH(alias) <= 32),
		description TEXT CHECK(LENGTH(description) <= 256),
		PRIMARY KEY (bot_id, alias)
	);`

	keyTable := `CREATE TABLE IF NOT EXISTS bot_keys (
		bot_id TEXT NOT NULL CHECK(LENGTH(bot_id) <= 64),
		token TEXT NOT NULL CHECK(LENGTH(token) <= 32),
		PRIMARY KEY (bot_id)
	);`

	webhookTable := `CREATE TABLE IF NOT EXISTS bot_webhooks (
		bot_id TEXT NOT NULL CHECK(LENGTH(bot_id) <= 64),
		url TEXT NOT NULL CHECK(LENGTH(url) <= 128),
		secret TEXT CHECK(LENGTH(secret) <= 64),
		PRIMARY KEY (bot_id)
	);`

	userTable := `CREATE TABLE IF NOT EXISTS users (
		id TEXT NOT NULL CHECK(LENGTH(id) <= 64),
		nickname TEXT NOT NULL,
		PRIMARY KEY (id)
	);`

	chatTable := `CREATE TABLE IF NOT EXISTS chats (
		id TEXT NOT NULL CHECK(LENGTH(id) <= 64),
		user_id TEXT NOT NULL CHECK(LENGTH(user_id) <= 64),
		bot_id TEXT NOT NULL CHECK(LENGTH(bot_id) <= 64),
		PRIMARY KEY (id),
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (bot_id) REFERENCES bots(id)
	);`

	messageTable := `CREATE TABLE IF NOT EXISTS messages (
		id TEXT NOT NULL CHECK(LENGTH(id) <= 64),
		chat_id TEXT NOT NULL CHECK(LENGTH(chat_id) <= 64),
		sender_id TEXT NOT NULL CHECK(LENGTH(sender_id) <= 64),
		body TEXT,
		attachment_ids TEXT,
		timestamp INT64 NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (chat_id) REFERENCES chats(id)
	);`

	fileTable := `CREATE TABLE IF NOT EXISTS files (
		id TEXT NOT NULL CHECK(LENGTH(id) <= 64),
		path TEXT NOT NULL CHECK(LENGTH(path) <= 256),
		name TEXT NOT NULL CHECK(LENGTH(name) <= 128),
		mime_type TEXT NOT NULL CHECK(LENGTH(mime_type) <= 64),
		size INT64,
		PRIMARY KEY (id)
	);`

	tables := []string{botTable, commandTable, keyTable, webhookTable,
		userTable, chatTable, messageTable, fileTable}

	for _, t := range tables {
		_, err := s.conn.Exec(t)
		if err != nil {
			return fmt.Errorf("can't create table: %w", err)
		}
	}

	return nil
}
