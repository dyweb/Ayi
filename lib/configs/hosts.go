package configs

import (
	"fmt"
	"io/ioutil"
)

type Host struct {
	// TODO: must use upper case ?
	Ip   string
	Name string
}

func AddLocalHost() {
	fmt.Println("Add localhost! ")
}

func ReadHostsFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//func parseHosts(s string) []Host {
//	return
//}

//func parseHost(s string) Host {
//	return Host{
//		Ip: "a",
//		Name: "b"
//	}
//}