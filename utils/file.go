package tyutils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ZipDir(dir, zipFile string, fn func(name string)) error {
	zfile, err := os.Create(zipFile)
	if err != nil {
		return err
	}

	zwriter := zip.NewWriter(zfile)
	defer zwriter.Close()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			path = strings.Replace(path, "\\", "/", -1)
			var fDest io.Writer
			fDest, err = zwriter.Create(path[len(dir)+1:])
			if err != nil {
				return err
			}

			if fn != nil {
				fn(path)
			}

			var fSrc *os.File
			fSrc, err = os.Open(path)
			if err != nil {
				return err
			}
			defer fSrc.Close()

			var n int64
			n, err = io.Copy(fDest, fSrc)
			if err != nil {
				return err
			}
			if n < 0 {

			}
		}
		return nil
	})

	return err
}

func UnzipDir(zipFile, dir string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, f := range reader.File {
		err = func() error {
			path := dir + string(filepath.Separator) + f.Name
			err = os.MkdirAll(filepath.Dir(path), 0755)
			if err != nil {
				return err
			}

			var fDest *os.File
			fDest, err = os.Create(path)
			if err != nil {
				return err
			}
			defer fDest.Close()

			var fSrc io.ReadCloser
			fSrc, err = f.Open()
			if err != nil {
				return err
			}
			defer fSrc.Close()

			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				return err
			}
			return nil
		}()

		if err != nil {

		}
	}

	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
