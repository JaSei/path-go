package path_test

import (
	"encoding/json"
	"github.com/jasei/path-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FileSource struct {
	Path string `json:"path"`
	Size int    `json:"size"`
}

type FileInfo struct {
	FileId  string       `json:"file_id"`
	Sources []FileSource `json:"sources"`
}

var expected = FileInfo{
	FileId: "01ba4719c80b6fe911b091a7c05124b64eeece964e09c058ef8f9805daca546b",
	Sources: []FileSource{
		FileSource{Path: "c:\\tmp\\empty_file", Size: 0},
		FileSource{Path: "/tmp/empty_file", Size: 0},
	},
}

func TestLoadJsonViaReader(t *testing.T) {
	path, err := path.NewPath("example.json")
	assert.Nil(t, err)

	reader, err := path.OpenReader()
	defer func() {
		assert.Nil(t, path.Close())
	}()
	assert.Nil(t, err)
	assert.NotNil(t, reader)

	decodedJson := new(FileInfo)

	err = json.NewDecoder(reader).Decode(decodedJson)
	if !assert.Nil(t, err) {
		t.Log(err)
	}

	assert.Equal(t, &expected, decodedJson)
}

func TestLoadJsonViaSlurp(t *testing.T) {
	path, err := path.NewPath("example.json")
	assert.Nil(t, err)

	jsonBytes, err := path.Slurp()
	assert.Nil(t, err)

	decodedJson := new(FileInfo)
	json.Unmarshal(jsonBytes, decodedJson)

	assert.Equal(t, &expected, decodedJson)

}