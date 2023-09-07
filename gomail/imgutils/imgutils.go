package imgutils

import (
	"log"
	"os"
	"net/http"
	"encoding/base64"
)

func ImgToBase64(fpath string) string{
	
	// Read file into a byte slice
	bytes, err :=os.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Convert bytes to string and append to mime type string
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)

	return base64Encoding
}
