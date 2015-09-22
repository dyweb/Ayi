package configs

// TODO: ref https://golang.org/src/net/hosts.go

import (
	"fmt"
	"regexp"
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
	// TODO: handle wrong config that does not match
	r, _ := regexp.Compile("([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})\\s*(\\S*)$")
	m := r.FindStringSubmatch(s)
	return Host{ip: m[1], name: m[2]}
}