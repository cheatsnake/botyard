package services

import (
	"botyard/internal/entities/file"
	mock "botyard/internal/storage/_mock"
	"botyard/internal/tools/ulid"
	"testing"
)

func TestFileService(t *testing.T) {
	fileService := NewFileService(mock.FileStore())

	t.Run("add a new file", func(t *testing.T) {
		testFile, err := file.New("/test", "text/plain")
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFile, err, nil)
		}

		err = fileService.AddFile(testFile)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFile, err, nil)
		}
	})

	t.Run("get files", func(t *testing.T) {
		testFileIds := []string{ulid.New(), ulid.New()}
		testFiles, err := fileService.GetFiles(testFileIds)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFiles, err, nil)
		}

		if testFiles == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFiles, testFiles, []*file.File{})
		}
	})

	t.Run("delete a file", func(t *testing.T) {
		testFileId := ulid.New()
		err := fileService.DeleteFile(testFileId)
		if err != nil {
			t.Errorf("got: %v,\nexpect: %v\n", err, nil)
		}
	})
}
