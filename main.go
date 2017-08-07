package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func hasSuffixInArray(s string, arr []string) bool {
	for _, suf := range arr {
		if strings.HasSuffix(strings.ToLower(s), strings.ToLower(suf)) {
			return true
		}
	}
	return false
}

func main() {

	if len(os.Args) == 1 {
		log.Fatal(`
            No Zip file was provided.
            Pass the path to a Zip-formatted file as an argument to use the program.

            Usage:
                ./image-extractor [zipFileName]

            Examples:
                ./image-extractor some-zip-file.zip
                ./image-extractor some-word-doc.docx`)
	}

	zipFileName := os.Args[1]
	newZipFileName := "media-" + zipFileName + ".zip"

	r, err := zip.OpenReader(zipFileName)
	errCheck(err)
	defer r.Close()

	newZip, err := os.Create(newZipFileName)
	errCheck(err)
	defer newZip.Close()

	w := zip.NewWriter(newZip)
	defer w.Close()

	for _, file := range r.File {

		fname := file.FileHeader.Name

		if !hasSuffixInArray(fname, []string{".png", ".jpg", ".jpeg", ".gif"}) {
			continue
		}

		f, err := w.Create(path.Base(fname))
		errCheck(err)

		rc, err := file.Open()
		errCheck(err)
		defer rc.Close()

		_, err = io.Copy(f, rc)
		errCheck(err)

	}

}
