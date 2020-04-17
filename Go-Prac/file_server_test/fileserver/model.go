package fileserver

import "io"

type FileServer interface {
	Login() error
	Upload(path string, fileName string, f io.Reader) (err error)
	Delete(path string) (ok bool, err error)
	GetFileList(path string) (file FileInfo, err error)
	ZQBUpload(path string, fileName string, image []byte) (fullpath string, err error)
}

type Config struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
