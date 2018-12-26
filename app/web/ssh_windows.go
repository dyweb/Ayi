package web

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func (a *App) sshdCommand() *cobra.Command {
	hostKeyFile := filepath.Join(os.Getenv("HOME"), ".ssh/id_rsa")
	authorizedKeysFile := filepath.Join(os.Getenv("HOME"), ".ssh/authorized_keys")

	cmd := cobra.Command{
		Use:   "sshd",
		Short: "start ssh server",
		Long:  "Start ssh server",
		Example: `
Server, give a host key without passphrase

Ayi web sshd --hostKey ~/.ssh/id_rsa_test --port 2222

Client 

ssh localhost -p 2222

For generate key pair on server

ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Fatal("not supported on windows")
		},
	}
	cmd.Flags().StringVar(&hostKeyFile, "hostKey", hostKeyFile, "private key without passphrase for server")
	cmd.Flags().StringVar(&authorizedKeysFile, "authorizedKeys", authorizedKeysFile, "public key saved on server by client")
	return &cmd
}
