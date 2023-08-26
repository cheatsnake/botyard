package file

import (
	"botyard/internal/tools/ulid"
	"testing"
)

func TestNew(t *testing.T) {
	testPath := "/test/file"
	testMimeType := "application/test"

	t.Run("check id", func(t *testing.T) {
		testFile := New(testPath, testMimeType)
		got := ulid.Verify(testFile.Id)
		if got != nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFile, got, nil)
		}
	})

	t.Run("check path", func(t *testing.T) {
		testFile := New(testPath, testMimeType)
		expect := testPath
		got := testFile.Path
		if got != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testFile, got, expect)
		}
	})

	t.Run("check mime type", func(t *testing.T) {
		testFile := New(testPath, testMimeType)
		expect := testMimeType
		got := testFile.MimeType
		if got != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testFile, got, expect)
		}
	})

}
