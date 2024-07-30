package mypack

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Translator(logFileName, archDir string) error {

	// checking the extension
	if !strings.HasSuffix(logFileName, ".log") {
		return fmt.Errorf("%s — this is not a log file.", logFileName)
	}

	// creating an archive name
	fileInfo, err := os.Stat(logFileName)
	if err != nil {
		return err
	}
	archiveName := fmt.Sprintf("%s_%d.tar.gz", strings.TrimSuffix(logFileName, ".log"), fileInfo.ModTime().Unix())
	archiveName = filepath.Join(archDir, archiveName)

	// creating tar file
	tarFile, err := os.Create(archiveName)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	// creating writers
	gzipWriter := gzip.NewWriter(tarFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// opening a logfile
	logFile, err := os.Open(logFileName)
	if err != nil {
		return err
	}
	defer logFile.Close()

	// getting information about the logfile for the tar header
	stat, err := logFile.Stat()
	if err != nil {
		return err
	}
	//creating header
	header, err := tar.FileInfoHeader(stat, "")
	if err != nil {
		return err
	}
	err = tarWriter.WriteHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(tarWriter, logFile)
	if err != nil {
		return err
	}

	return nil
}
