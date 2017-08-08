# Image Extractor

A simple program that takes a Zip file (or Word document), and creates a new Zip file containing its images.

## Usage

Pass the path to a Zip-formatted file as an argument to use the program.
Currently there are no options. It will output a new Zip file with the same name as the original archive, but with a "media-" prefix.

Example:
`./go-image-extractor some-word-doc.docx`

will output a file named `media-some-word-doc.docx.zip`
