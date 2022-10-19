package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func ResizeImage(src string, dst string) error {
	fmt.Println("Resizing:", src, dst)
	if _, err := os.Stat(dst); err == nil {
		if !conf.Force {
			return errors.New("File already exists: " + dst)
		} else {
			os.Remove(dst)
		}
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return errors.New(err.Error() + src)
	}
	defer srcFile.Close()

	ext := strings.ToLower(filepath.Ext(src))
	if format, ok := decoders[ext]; ok {
		srcImg, err := format.Decoder(srcFile)
		if err != nil {
			return errors.New(err.Error() + srcFile.Name())
		}

		var newW uint = conf.Size
		var newH uint = conf.Size
		size := srcImg.Bounds().Size()
		if size.Y > size.X {
			newH = 0
		} else {
			newW = 0
		}

		dstImg := resize.Resize(newW, newH, srcImg, resize.Lanczos2)

		dstFile, err := os.Create(dst)
		if err != nil {
			return errors.New("Destination error: " + err.Error())
		}
		defer dstFile.Close()
		return format.Encoder(dstFile, dstImg)
	}
	return errors.New("Format not supported: " + ext)
}

func ProcessDirecotry(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, e := range dir {
		relPath, _ := filepath.Rel(filepath.Join(conf.Destination, ".."), path)
		if e.IsDir() {
			os.MkdirAll(filepath.Join(conf.Destination, relPath, e.Name()), os.ModePerm)
			ProcessDirecotry(filepath.Join(path, e.Name()))
		} else {
			os.MkdirAll(filepath.Join(conf.Destination, relPath), os.ModePerm)
			err := ResizeImage(filepath.Join(path, e.Name()), filepath.Join(conf.Destination, relPath, e.Name())+".jpg")
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
