package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("sh", "-c", "go test -v -cover $(glide novendor)")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	fmt.Println(cmd.Run())
}
