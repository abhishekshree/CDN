package utils

import (
	"archive/zip"
	"io"
	"os"
)

func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open("cdn/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		return err
	}
	if _, err := io.Copy(wr, file); err != nil {
		return err
	}
	return nil
}

func ZipFiles(files []string, outfile string) (string, error) {
	uuid, err := GenerateUUID()
	if err != nil {
		return "", err
	}
	x := uuid + "__" + outfile
	x = "zip/" + x

	out, err := os.Create(x)

	if err != nil {
		return "", err
	}
	defer out.Close()

	zipw := zip.NewWriter(out)
	defer zipw.Close()

	for _, filename := range files {
		if err := appendFiles(filename, zipw); err != nil {
			os.Remove(x)
			return "", err
		}
	}

	return uuid + "__" + outfile, nil
}
