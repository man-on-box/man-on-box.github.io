package util

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func CreateFile(fileName string) *os.File {
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Could not create directory: %v", err)
	}

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}
	return f
}

func CopyFile(src string, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func CopyDir(src string, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func RemoveDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return
	}
	if err := os.RemoveAll(dir); err != nil {
		log.Fatalf("Could not remove directory: %v", err)
	}
}
