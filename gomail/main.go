// https://articles.wesionary.team/sending-emails-with-go-golang-using-smtp-gmail-and-oauth2-185ee12ab306
// https://medium.com/@dhanushgopinath/sending-html-emails-using-templates-in-golang-9e953ca32f3d
// https://www.geeksforgeeks.org/sending-email-using-smtp-in-golang/
// https://stackoverflow.com/questions/55393965/how-to-send-a-html-templated-body-as-an-email-through-gmail-go-sdk
// https://freshman.tech/snippets/go/image-to-base64/
// https://www.geeksforgeeks.org/email-template-using-html-and-css/

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
 
    // Its the default port of smtp server
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
    
        // handling the errors
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    
        fmt.Printf("Successfully sent mail to: %s\n", email)
        time.Sleep(1 * time.Second)
    }    
}