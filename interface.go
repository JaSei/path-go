package pathutil

import (
	"io"
	"os"

	"github.com/JaSei/hashutil-go"
)

type VisitFunc func(path Path)

type VisitOpt struct {
	Recurse bool
}

type LinesFunc func(string)

type Path interface {
	String() string
	Canonpath() string
	Basename() string

	Chdir() (Path, error)
	Rename(string) (Path, error)

	Stat() (os.FileInfo, error)

	IsDir() bool
	Exists() bool
	IsFile() bool
	IsRegularFile() bool

	Remove() error
	RemoveTree() error

	Visit(VisitFunc, VisitOpt)
	CopyFile(string) (Path, error)

	CopyFrom(io.Reader) (int64, error)

	CryptoMd5() (hashutil.Md5, error)
	CryptoSha1() (hashutil.Sha1, error)
	CryptoSha256() (hashutil.Sha256, error)
	CryptoSha384() (hashutil.Sha384, error)
	CryptoSha512() (hashutil.Sha512, error)

	MakePath() error
	MakePathMode(os.FileMode) error

	OpenReader() (ReadSeekCloser, error)
	OpenWriter() (*os.File, error)
	OpenWriterAppend() (*os.File, error)

	Slurp() (string, error)
	SlurpBytes() ([]byte, error)

	Spew(string) error
	SpewBytes([]byte) error

	Lines() ([]string, error)
	LinesWalker(LinesFunc) error

	Child(...string) (Path, error)
	Children() ([]Path, error)

	Parent() Path

	Append(string) error
	AppendBytes([]byte) error
}
