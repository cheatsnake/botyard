package message

import (
	"botyard/internal/tools/ulid"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	testChatId := ulid.New()
	testSenderId := ulid.New()
	testBody := "test"
	testFileIds := []string{ulid.New(), ulid.New()}

	t.Run("check id", func(t *testing.T) {
		testMsg, err := New(testChatId, testSenderId, testBody, testFileIds)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, err.Error(), nil)
		}
	})

	t.Run("check chat id", func(t *testing.T) {
		testMsg, err := New(testChatId, testSenderId, testBody, testFileIds)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, err.Error(), nil)
		}

		expect := testChatId
		got := testMsg.ChatId
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
		}
	})

	t.Run("check too long body", func(t *testing.T) {
		bodies := []string{strings.Repeat("a", maxBodyLen+1), strings.Repeat("B", maxBodyLen*2)}
		for _, b := range bodies {
			testMsg, err := New(testChatId, testSenderId, b, testFileIds)
			expect := errBodyTooLong
			got := err.Error()
			if got != expect {
				t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
			}
		}
	})

	t.Run("check too many files", func(t *testing.T) {
		fileIds := [][]string{
			strings.Split(strings.Repeat("a", maxFiles+1), ""),
			strings.Split(strings.Repeat("a", maxFiles*2), ""),
		}
		for _, fi := range fileIds {
			testMsg, err := New(testChatId, testSenderId, testBody, fi)
			expect := errTooManyFiles
			got := err.Error()
			if got != expect {
				t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
			}
		}
	})

	t.Run("check sender id", func(t *testing.T) {
		testMsg, err := New(testChatId, testSenderId, testBody, testFileIds)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, err.Error(), nil)
		}

		expect := testSenderId
		got := testMsg.SenderId
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
		}
	})

	t.Run("check body", func(t *testing.T) {
		testMsg, err := New(testChatId, testSenderId, testBody, testFileIds)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, err.Error(), nil)
		}

		expect := testBody
		got := testMsg.Body
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
		}
	})

	t.Run("check file ids", func(t *testing.T) {
		testMsg, err := New(testChatId, testSenderId, testBody, testFileIds)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, err.Error(), nil)
		}

		for i, fileId := range testFileIds {
			expect := fileId
			got := testMsg.FileIds[i]
			if got != expect {
				t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
			}
		}
	})

	t.Run("check timestamp", func(t *testing.T) {
		testTimestamp := time.Now()
		testMsg, err := New(testChatId, testSenderId, testBody, testFileIds)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, err.Error(), nil)
		}

		expect := true
		got := testTimestamp.Before(testMsg.Timestamp)

		if !got {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
		}
	})
}
