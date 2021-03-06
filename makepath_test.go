package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakePath(t *testing.T) {
	tempdir, err := NewTempDir()
	assert.NoError(t, err)

	defer func() {
		assert.NoError(t, tempdir.RemoveTree())
	}()

	newPath, err := New(tempdir, "a/b/c")
	assert.NoError(t, err)

	assert.False(t, newPath.Exists())

	assert.NoError(t, newPath.MakePath())

	assert.True(t, newPath.IsDir())
}
