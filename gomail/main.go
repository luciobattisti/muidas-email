package main
 
import (
    "fmt"
    "net/smtp"
    "os"
    "flag"
    "time"
    emailutils "github.com/luciobattisti/emailutils"
)

// Main function
func main() {

    // Parse args
    htmlTemplatePtr := flag.String("html-template", "gotemplates/default.html", "HTML Template File")
    emailFpathPtr := flag.String("newsletter", "../data/-emails.csv", "Email Newsletter File")
    from := flag.String("from", "", "Email sender")
    password := flag.String("token", "", "Email token")

    flag.Parse()

    // Get subscribers' emails
    toList := emailutils.GetEmailListFromCsv(*emailFpathPtr)
 
    // Specify host
    host := "smtp.gmail.com"
 
    // This is the default port of smtp server
    port := "587"
 
    // Authenticate
    auth := smtp.PlainAuth("", *from, *password, host)

    // Send emails
    for index, email := range toList {

        var data interface {}
        template := emailutils.ParseTemplate(*htmlTemplatePtr, data)
        msg := emailutils.GetMessageString(
            "Muidas Newsletter Test",
            *from,
            *from,
            email, 
            template,
        )

        // We can't send strings directly in mail,
        // strings need to be converted into slice bytes
        body := []byte(msg)
    
        // SendMail uses TLS connection to send the mail
        // The email is sent to all address in the toList,
        // the body should be of type bytes, not strings
        // This returns error if any occurred.
        err := smtp.SendMail(
            host+":"+port, 
            auth, 
            *from, 
            toList[index:index+1], 
            body,
        )
    
        // Handling the errors
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    
        fmt.Printf("Successfully sent mail to: %s\n", email)
        time.Sleep(1 * time.Second)
    }    
}