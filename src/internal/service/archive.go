package service

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ArchiveFile(file string, archiveDir string) error {

	fileInfo, err := os.Stat(file)
	if err != nil {
		return errors.New("файл не существует")
	}

	// получаем время модификации файла
	mtime := fileInfo.ModTime().Unix()

	// создаем имя архива
	fileName := strings.TrimSuffix(file, filepath.Ext(file))
	archiveName := fmt.Sprintf("%s_%d.tar.gz", filepath.Base(fileName), mtime)
	archivePath := filepath.Join(archiveDir, archiveName)

	// создаем файл-архив
	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return errors.New("ошибка при создании файла")
	}
	defer archiveFile.Close()

	//Создание gzip-архива
	gzipWriter := gzip.NewWriter(archiveFile)
	defer gzipWriter.Close()

	// Создаем tar-архив
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	f, err := os.Open(file)
	if err != nil {
		return errors.New("ошибка при открытии файла")
	}
	defer f.Close()

	// Записываем файл в tar-архив
	header := &tar.Header{
		Name: filepath.Base(file),
		Mode: int64(fileInfo.Mode()),
		Size: fileInfo.Size(),
	}
	if err := tarWriter.WriteHeader(header); err != nil {
		return errors.New("ошибка записи header в архив")
	}

	// Копируем содержимое файла в tar-архив
	if _, err := io.Copy(tarWriter, f); err != nil {
		return fmt.Errorf("ошибка копирования информации из файла: %v %w", file, err)
	}

	return nil

}
