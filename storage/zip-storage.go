package storage

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

type ZipStorage struct {
	*Storage
}

func NewZipStorage(filename string) *ZipStorage {
	return &ZipStorage{
		&Storage{filename: filename},
	}
}

func (z *ZipStorage) Save(data []byte) error {
	f, err := os.Create(z.GetFileName())

	if err != nil {
		return err
	}

	defer func() error {
		if err := f.Close(); err != nil {
			return err
		}

		return nil
	}()

	zw := zip.NewWriter(f)
	defer zw.Close()

	w, err := zw.Create("data")

	if err != nil {
		return err
	}

	_, err = w.Write(data)

	return err
}

func (z *ZipStorage) Load() ([]byte, error) {
	r, err := zip.OpenReader(z.GetFileName())

	if err != nil {
		return nil, err
	}

	defer r.Close()

	if len(r.File) == 0 {
		return nil, errors.New("архив пуст")
	}

	file := r.File[0]
	rc, _ := file.Open()
	defer rc.Close()

	return io.ReadAll(rc)
}
