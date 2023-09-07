package imgutils

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestImgToBase64(t *testing.T){

	imgBase64 :=ImgToBase64("data/sample-clouds-400x300.jpeg")
    
    outFpath := "data/base64-encoding.txt"
    if err := os.WriteFile(outFpath, []byte(imgBase64), 0666); err != nil {
        log.Fatal(err)
    }

    imgBytes, err := os.ReadFile(outFpath)
    if err != nil {
        log.Fatal(err)
    }

    imgBase64FromFile := string(imgBytes)

    if imgBase64 != imgBase64FromFile {
        t.Errorf(fmt.Sprintf("got %s, wanted %s", imgBase64FromFile, imgBase64))
    }
}


