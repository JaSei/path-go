package pathutil_test

import (
	"github.com/JaSei/pathutil-go"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func wc(path *pathutil.Path) (lines, words, chars int) {
	path.LinesFunc(func(line string) {
		lines++

		w := strings.Split(line, " ")
		for _, i := range w {
			if len(i) > 0 {
				words++
			}
		}

		chars = len(line) + chars + len("\n") //+newline
	})

	return
}

func Test(t *testing.T) {
	path, err := pathutil.NewPath("../LICENSE")

	assert.Nil(t, err)

	lines, words, chars := wc(path)

	//wc LICENSE
	//21  169 1066 LICENSE
	assert.Equal(t, 21, lines, "count of lines")
	assert.Equal(t, 169, words, "count of words")
	assert.Equal(t, 1066, chars, "count of chars")
}