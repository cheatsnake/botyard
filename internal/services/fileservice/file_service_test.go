package fileservice

import (
	"testing"

	"github.com/cheatsnake/botyard/internal/entities/file"
	mock "github.com/cheatsnake/botyard/internal/storage/_mock"
	"github.com/cheatsnake/botyard/internal/tools/ulid"
)

func TestAddFile(t *testing.T) {
	fs := New(mock.FileStore())

	t.Run("add a new file", func(t *testing.T) {
		testFile, err := fs.AddFile("/test", "test.txt", "text/plain", 13)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFile, err, nil)
		}
	})

	t.Run("add a new file with invalid data", func(t *testing.T) {
		testFile, err := fs.AddFile("/test", "test.txt", "invalid type", 13)
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFile, err, "error")
		}
	})
}

func TestGetFiles(t *testing.T) {
	fs := New(mock.FileStore())

	t.Run("get files", func(t *testing.T) {
		testAttachmentIds := []string{ulid.New(), ulid.New()}
		testFiles, err := fs.GetFiles(testAttachmentIds)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFiles, err, nil)
		}

		if testFiles == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFiles, testFiles, []*file.File{})
		}
	})
}

func TestDeleteFile(t *testing.T) {
	fs := New(mock.FileStore())

	t.Run("delete files", func(t *testing.T) {
		testFileId := ulid.New()
		err := fs.DeleteFiles([]string{testFileId})
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}
