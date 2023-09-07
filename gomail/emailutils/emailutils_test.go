package emailutils

import (
	"testing"
	"reflect"
	"fmt"
)

func TestGetEmailsFromCsv(t * testing.T) {
	
	fpath := "data/sample-email-list.csv"
	got := GetEmailListFromCsv(fpath)

	wanted := []string{
		"pippo@gmail.com",
		"pluto@gmail.com",
		"paperino@yahoo.com",
		"qui@gmail.com",
		"quo@gmail.com",
		"qua@yahoo.com",
		"paperoga@yahoo.com",
		"archimede@gmail.com",
		"topolino@aol.com",
		"ziopaperone@gmail.com",
	}

	if ! reflect.DeepEqual(got, wanted) {
		t.Errorf(fmt.Sprintf("got %v, wanted %v", got, wanted))
	}





}