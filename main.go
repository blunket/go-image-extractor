package main

import (
	"archive/zip"
	"io/ioutil"
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
		defer rc.Close()

		filecontents, err := ioutil.ReadAll(rc)

		_, err = f.Write(filecontents)
		errCheck(err)
	}

}
