package services

import (
	"botyard/internal/entities/file"
	"botyard/internal/tools/ulid"
	"testing"
)

func TestFileService(t *testing.T) {
	mockStore := &mockFileStore{}
	fileService := NewFileService(mockStore)

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

type mockFileStore struct{}

func (mfs *mockFileStore) AddFile(file *file.File) error {
	return nil
}

func (mfs *mockFileStore) GetFile(id string) (*file.File, error) {
	return &file.File{Id: id}, nil
}

func (mfs *mockFileStore) GetFiles(ids []string) ([]*file.File, error) {
	return []*file.File{}, nil
}

func (mfs *mockFileStore) DeleteFile(id string) error {
	return nil
}
