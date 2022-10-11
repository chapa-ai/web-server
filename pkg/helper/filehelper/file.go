package filehelper

import (
	"io"
	"mime/multipart"
	"os"
)

func CopyFileByFile(src io.Reader, dstPath string) error {
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func CopyFileHeader2Str(file *multipart.FileHeader, dstPath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	return CopyFileByFile(src, dstPath)
}
