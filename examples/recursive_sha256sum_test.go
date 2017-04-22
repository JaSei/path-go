package path_test

import (
	"crypto"
	"fmt"
	pathutils "github.com/jasei/path-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisitRecursiveAndHashAllFiles(t *testing.T) {
	path, err := pathutils.NewPath("/tmp")
	assert.Nil(t, err)

	path.Visit(
		func(path *pathutils.Path) {
			if path.IsDir() {
				return
			}

			hash, err := path.Crypto(crypto.SHA256)

			if err == nil {
				fmt.Printf("%s\t%s\n", hash.HexSum(), path.String())
			} else {
				fmt.Println(err)
			}
		},
		pathutils.VisitOpt{Recurse: true},
	)
}