package file

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path"
)

func GetFileSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)
	return len(content), err
}

func GetFileExt(filename string) string {
	return path.Ext(filename)
}

func CheckFileNotExist(src string) bool {
	_, err := os.Stat(src)
	return errors.Is(err, fs.ErrNotExist)
}

func CheckFilePermission(src string) bool {
	_, err := os.Stat(src)
	return errors.Is(err, fs.ErrPermission)
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func IsNotExistMkDir(src string) error {
	if CheckFileNotExist(src) {
		return MkDir(src)
	}
	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd error:%v", err)
	}

	src := dir + "/" + filePath
	fmt.Println("MustOpen:", src)
	if CheckFilePermission(src) {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	if err = IsNotExistMkDir(src); err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}

	return f, nil
}
