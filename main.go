package main

import (
	"archive/zip"
	"fmt"
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
	s = strings.ToLower(s)
	for _, suf := range arr {
		if strings.HasSuffix(s, strings.ToLower(suf)) {
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
                ./go-image-extractor [zipFileName]

            Examples:
                ./go-image-extractor some-zip-file.zip
                ./go-image-extractor some-word-doc.docx`)
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

	// store list of files written in case of duplicate names
	m := make(map[string]int)

	for _, file := range r.File {

		fname := path.Base(file.Name)

		if !hasSuffixInArray(fname, []string{".png", ".jpg", ".jpeg", ".gif"}) {
			continue
		}

		for m[fname] != 0 {
			suffix := path.Ext(fname)
			m[fname] = m[fname] + 1
			trimmed := strings.TrimSuffix(fname, suffix)
			fname = fmt.Sprintf("%s%d%s", trimmed, m[fname], suffix)
			fmt.Println("warning: conflicting filename changed")
		}

		f, err := w.Create(fname)
		errCheck(err)
		m[fname] = 1

		rc, err := file.Open()
		errCheck(err)
		defer rc.Close()

		_, err = io.Copy(f, rc)
		errCheck(err)

	}

}
