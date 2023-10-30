package emailutils

import (
	"bytes"
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"os"
)

func GetMessageString(
	subject string, replyTo string, from string, body string,
) []byte {
	return []byte(
		"Reply-To: " + replyTo +
			"\r\n" + "From: " +
			from + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
			body + "\r\n",
	)
}

func ParseTemplate(templateFileName string, data interface{}) string {
	tplt, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	if err = tplt.Execute(buf, data); err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func GetEmailListFromCsv(fpath string) []string {

	// Email list is in CSV format
	f, err := os.Open(fpath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	emailList := []string{}

	// Skip header
	csvReader.Read()

	for {
		rec, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		emailList = append(emailList, rec[1])
	}

	return emailList

}
