package zipper

import (
	"archive/zip"
	"io"
	"os"
)

func Compress(fpaths []string, zpath string) error {
	zfile, err := os.Create(zpath)
	if err != nil {
		return err
	}
	defer zfile.Close()

	zw := zip.NewWriter(zfile)
	defer zw.Close()

	for _, fpath := range fpaths {
		f, err := os.Open(fpath)
		if err != nil {
			return err
		}

		info, err := f.Stat()
		if err != nil {
			return err
		}

		h, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		h.Method = zip.Deflate
		writer, err := zw.CreateHeader(h)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, f)
		if err != nil {
			return err
		}
	}

	return nil
}
