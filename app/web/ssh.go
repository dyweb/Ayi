// +build !windows

package web

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"

	gossh "golang.org/x/crypto/ssh"

	"github.com/dyweb/gommon/errors"
	"github.com/gliderlabs/ssh"
	"github.com/kr/pty"
	"github.com/spf13/cobra"
)

// ssh.go creates a server
// https://github.com/dyweb/Ayi/issues/82
// TODO: if we just call bash, user need to su to make it a login shell, which is different from normal sshd I guess

// TODO: just copy from https://github.com/gliderlabs/ssh/blob/master/_examples/ssh-pty/pty.go
func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

// based on https://github.com/golang/crypto/blob/master/ssh/example_test.go
func loadAuthorizedKeys(f string) (map[string]bool, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, errors.Wrap(err, "error read file")
	}
	keys := make(map[string]bool)
	i := 0
	for len(b) > 0 {
		pub, _, _, rest, err := gossh.ParseAuthorizedKey(b)
		if err != nil {
			return keys, errors.Wrapf(err, "error parse key %d", i)
		}
		keys[string(pub.Marshal())] = true
		b = rest
		i++
	}
	return keys, nil
}

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
			log.Infof("use private key %s", hostKeyFile)
			log.Infof("use authorized keys %s", authorizedKeysFile)

			// TODO: we only support one user because only the authorized keys of the user start ayi web sshd is used
			authorizedKeys, err := loadAuthorizedKeys(authorizedKeysFile)
			if err != nil {
				log.Fatalf("error load authorized keys %s", err)
				return
			}

			ssh.Handle(func(s ssh.Session) {
				start := time.Now()
				log.Infof("connected from %s as %s", s.RemoteAddr().String(), s.User())
				defer func() {
					log.Infof("disconnected from %s after %s", s.RemoteAddr().String(), time.Now().Sub(start))
				}()
				//authorizedKey := gossh.MarshalAuthorizedKey(s.PublicKey())
				//io.WriteString(s, fmt.Sprintf("public key used by %s:\n", s.User()))

				cmd := exec.Command("bash")
				ptyReq, winCh, isPty := s.Pty()
				if isPty {
					cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
					f, err := pty.Start(cmd)
					if err != nil {
						panic(err)
					}
					go func() {
						for win := range winCh {
							setWinsize(f, win.Width, win.Height)
						}
					}()
					go func() {
						io.Copy(f, s) // stdin
					}()
					io.Copy(s, f) // stdout
				} else {
					// FIXME: only interactive session is allowed, you can't do things like rsync over ssh
					io.WriteString(s, "No PTY requested.\n")
					s.Exit(1)
				}
				//s.Write(authorizedKey)
			})

			publicKeyOption := ssh.PublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
				// TODO: ssh.KeyEquals is more robust against timing attack because it use constant time compare
				// regardless of the length of the content
				if authorizedKeys[string(key.Marshal())] {
					log.Infof("public key auth success for %s", ctx.User())
					return true
				}
				log.Warnf("public key auth failed for %s", ctx.User())
				return false
			})
			// avoid generate one every time
			hostKey := ssh.HostKeyFile(hostKeyFile)

			port := 2222
			if a.port != 0 {
				port = a.port
			}

			addr := fmt.Sprintf(":%d", port)
			log.Infof("starting ssh server on %s", addr)
			log.Fatal(ssh.ListenAndServe(addr, nil, publicKeyOption, hostKey))
		},
	}
	cmd.Flags().StringVar(&hostKeyFile, "hostKey", hostKeyFile, "private key without passphrase for server")
	cmd.Flags().StringVar(&authorizedKeysFile, "authorizedKeys", authorizedKeysFile, "public key saved on server by client")
	return &cmd
}
