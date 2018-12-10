package web

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"

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

func (a *App) sshdCommand() *cobra.Command {
	hostKeyFile := filepath.Join(os.Getenv("HOME"), ".ssh/id_rsa")

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
			ssh.Handle(func(s ssh.Session) {
				start := time.Now()
				log.Infof("connected from % as %s", s.RemoteAddr(), s.User())
				defer func() {
					log.Info("disconnected from %s after %s", s.RemoteAddr(), time.Now().Sub(start))
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
					io.WriteString(s, "No PTY requested.\n")
					s.Exit(1)
				}
				//s.Write(authorizedKey)
			})

			// FIXME: this allow everyone to access
			publicKeyOption := ssh.PublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
				return true // allow all keys, or use ssh.KeysEqual() to compare against known keys
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
	cmd.Flags().StringVar(&hostKeyFile, "hostKey", filepath.Join(os.Getenv("HOME"), ".ssh/id_rsa"), "private key without passphrase")
	return &cmd
}
