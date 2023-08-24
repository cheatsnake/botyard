package message

import (
	"botyard/internal/tools/ulid"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	testChatId := ulid.New()
	testSenderId := ulid.New()
	testBody := "test"
	testFileIds := []string{ulid.New(), ulid.New()}

	t.Run("check id", func(t *testing.T) {
		testMsg := New(testChatId, testSenderId, testBody, testFileIds)
		got := ulid.Verify(testMsg.Id)
		if got != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, nil)
		}
	})

	t.Run("check chat id", func(t *testing.T) {
		testMsg := New(testChatId, testSenderId, testBody, testFileIds)
		expect := testChatId
		got := testMsg.ChatId
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
		}

		got2 := ulid.Verify(testMsg.ChatId)
		if got2 != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got2, nil)
		}
	})

	t.Run("check sender id", func(t *testing.T) {
		testMsg := New(testChatId, testSenderId, testBody, testFileIds)
		expect := testSenderId
		got := testMsg.SenderId
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
		}

		got2 := ulid.Verify(testMsg.SenderId)
		if got2 != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got2, nil)
		}
	})

	t.Run("check body", func(t *testing.T) {
		testMsg := New(testChatId, testSenderId, testBody, testFileIds)
		expect := testBody
		got := testMsg.Body
		if got != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
		}
	})

	t.Run("check file ids", func(t *testing.T) {
		testMsg := New(testChatId, testSenderId, testBody, testFileIds)

		for i, fileId := range testFileIds {
			expect := fileId
			got := testMsg.FileIds[i]
			if got != expect {
				t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, expect)
			}
		}

		for i := range testMsg.FileIds {
			got := ulid.Verify(testMsg.FileIds[i])
			if got != nil {
				t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, nil)
			}
		}
	})

	t.Run("check timestamp", func(t *testing.T) {
		testTimestamp := time.Now()
		testMsg := New(testChatId, testSenderId, testBody, testFileIds)
		got := testTimestamp.Before(testMsg.Timestamp)

		if !got {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testMsg, got, true)
		}
	})
}
