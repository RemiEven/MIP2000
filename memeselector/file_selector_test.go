package memeselector

import (
	"testing"
	"testing/fstest"
)

func TestSelectMeme(t *testing.T) {
	mapFS := fstest.MapFS{
		"images/titi": &fstest.MapFile{},
		"images/tata": &fstest.MapFile{},
		"images/tutu": &fstest.MapFile{},
	}

	expected := "un vrai de vrai"
	actual, _ := SelectMeme(mapFS)

	if expected != actual {
		t.Error("test failed, expected un vrai de vrai")
	}
}
