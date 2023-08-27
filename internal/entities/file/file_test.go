package file

import (
	"testing"
)

func TestNew(t *testing.T) {
	testPath := "/test/file"
	testMimeType := "text/plain"

	t.Run("check path", func(t *testing.T) {
		testFile, err := New(testPath, testMimeType)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testFile, err.Error(), nil)
		}

		expect := testPath
		got := testFile.Path
		if got != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testFile, got, expect)
		}
	})

	t.Run("check empty path", func(t *testing.T) {
		expect := errPathIsEmpty
		testFile, err := New("", testMimeType)
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFile, nil, expect)
		}

		got := err.Error()
		if got != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testFile, got, expect)
		}
	})

	t.Run("check mime type", func(t *testing.T) {
		testFile, err := New(testPath, testMimeType)
		if err != nil {
			t.Errorf("%#v\ngot: %s,\nexpect: %v\n", testFile, err.Error(), nil)
		}

		expect := testMimeType
		got := testFile.MimeType
		if got != expect {
			t.Errorf("%#v\ngot: %s,\nexpect: %s\n", testFile, got, expect)
		}
	})

	t.Run("check invalid mime type", func(t *testing.T) {
		expect := errInvalidMimeType
		testFile, err := New(testPath, "test")
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFile, nil, expect)
		}

		if err.Error() != expect {
			t.Errorf("%#v\ngot: %v,\nexpect: %v\n", testFile, err.Error(), expect)
		}
	})

}
