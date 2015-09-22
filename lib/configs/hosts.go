package configs

// TODO: ref https://golang.org/src/net/hosts.go

import (
	"fmt"
	"io/ioutil"
)

type Host struct {
	// TODO: must use upper case ?
	ip   string
	name string

}

func AddLocalHost() {
	fmt.Println("Add localhost! ")
}

func readHostsFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//func parseHosts(s string) []Host {
//	return
//}

func parseHost(s string) Host {
	// TODO: implement
	return Host{ip: "127.0.0.1", name: "localhost"}
}