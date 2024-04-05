package memeselector

import (
	"io"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestGetRandomMeme(t *testing.T) {
	expectedContent := []byte("expected")
	mapFS := fstest.MapFS{
		"images/titi": &fstest.MapFile{Data: expectedContent},
		//"images/tata": &fstest.MapFile{},
		//"images/tutu": &fstest.MapFile{},
	}

	actual, err := GetRandomMeme(mapFS)
	if err != nil {
		t.Errorf("test failed, failed to call GetRandomMeme: %v", err)
		return
	}
	actualContent, err := io.ReadAll(actual)

	if err != nil {
		t.Errorf("test failed, failed to read meme: %v", err)
		return
	}
	if string(expectedContent) != string(actualContent) {
		t.Errorf("test failed, unexpected content")
		return
	}
}

func TestSelectMemeWhenImageDirIsEmpty(t *testing.T) {
	mapFS := fstest.MapFS{
		"images/":          &fstest.MapFile{Mode: fs.ModeDir},
	}

	_, err := selectMeme(mapFS)
	if err == nil {
		t.Errorf("test failed, expected error when directory is empty")
		return
	}
}

func TestSelectMemeWhenImagesDirContainsOnlyDirs(t *testing.T) {
	mapFS := fstest.MapFS{
		"images/":          &fstest.MapFile{Mode: fs.ModeDir},
		"images/subfolder": &fstest.MapFile{Mode: fs.ModeDir},
	}

	_, err := selectMeme(mapFS)
	if err == nil {
		t.Errorf("test failed, expected error when directory is empty")
		return
	}
}

func TestSelectMemeWithOneFile(t *testing.T) {
	mapFS := fstest.MapFS{
		"images/":          &fstest.MapFile{Mode: fs.ModeDir},
		"images/subfolder": &fstest.MapFile{Mode: fs.ModeDir},
		"images/titi": &fstest.MapFile{Data: []byte{}},
	}

	_, err := selectMeme(mapFS)
	if err != nil {
		t.Errorf("test failed, expected error when directory is empty")
		return
	}
}
