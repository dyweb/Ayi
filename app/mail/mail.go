package mail

import (
	"errors"
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/dyweb/Ayi/common"
	"github.com/mailgun/mailgun-go"
	"github.com/spf13/viper"
)

// SendMailToWebStuff will send the mail to web stuff group
func SendMailToWebStuff(c *cli.Context) {
	domain := getMailgunConfigString(common.MailgunDomain)
	prikey := getMailgunConfigString(common.MailgunPrikey)
	pubkey := getMailgunConfigString(common.MailgunPubkey)

	if validateConfig(domain, prikey, pubkey) == false {
		panic(errors.New("Mailgun error: config needed."))
	}
	gun := mailgun.NewMailgun(domain, prikey, pubkey)
	m := mailgun.NewMessage("Sender <sender@example.com>", "Subject", "Message Body", "Recipient <recipient@example.com>")
	response, id, _ := gun.Send(m)
	fmt.Printf("Response ID: %s\n", id)
	fmt.Printf("Message from server: %s\n", response)
}

func getMailgunConfigString(key string) string {
	return viper.GetString(fmt.Sprintf("%s.%s", common.Mailgun, key))
}

func validateConfig(domain string, prikey string, pubkey string) bool {
	if domain == "" || prikey == "" || pubkey == "" {
		return false
	}
	return true
}
