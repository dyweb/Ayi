package mail

import (
	"fmt"
	"io/ioutil"

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
	sender := getMailgunConfigString(common.Sender)
	email := c.Args().Get(0)
	subject := c.Args().Get(1)
	fileName := c.Args().Get(2)
	
	
	// validate the args and config of mailgun
	if validateConfig(domain, prikey, pubkey) == false {
		fmt.Println("Mailgun error: config needed.")
		cli.ShowCommandHelp(c, "send")
		return
	}
	if validateArgs(email, subject, fileName) == false {
		fmt.Println("Mailgun error: wrong args.")
		cli.ShowCommandHelp(c, "send")
		return
	}
	
	gun := mailgun.NewMailgun(domain, prikey, pubkey)
	m := mailgun.NewMessage(sender, subject, "nil", fmt.Sprintf("Web stuff <%s>", email))
	m.SetHtml(readHTMLFile(fileName))
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

func validateArgs(email string, subject string, fileName string) bool {
	fmt.Println(subject, fileName)
	if email == "" || fileName == "" || subject == "" {
		return false
	}
	return true
}

func readHTMLFile(fileName string) string {
	

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Errorf("Disk I/O error: %s", err)
	}
	return fmt.Sprintf("%s", b)
}