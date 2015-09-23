package configs

// TODO: ref https://golang.org/src/net/hosts.go

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"bufio"
	"os"
	"log"
	"errors"
)

type Host struct {
	// TODO: must use upper case ?
	ip   string
	name string
}

func (host Host) Print() {
	fmt.Println("ip: " + host.ip)
	fmt.Println("name: " + host.name)
}

func PrintHosts(hosts []Host) {
	for i := 0; i < len(hosts); i++ {
		hosts[i].Print()
	}
}

func ParseHosts() []Host {
	// TODO: support for win
	hostsFile := "/etc/hosts"
	return parseHostsFile(hostsFile)
}

func AddDomainToLocalhost(domain string) (bool, error) {
	fmt.Println("Add localhost! ")
	return false, nil
}

func readHostsFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func parseHostsFile(hostsFile string) []Host {
	file, err := os.Open(hostsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hosts := make([]Host, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		host, err := parseHost(line)
		if err != nil {
			// TODO: log?
			//			log.Println(line)
			//			log.Fatal(err)
		}else {
			hosts = append(hosts, host)
		}
	}
	return hosts
}

func parseHost(s string) (Host, error) {
	s = removeComment(s)
	r, _ := regexp.Compile("([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})\\s*(\\S*)\\s*$")
	m := r.FindStringSubmatch(s)
	if len(m) == 3 {
		return Host{ip: m[1], name: m[2]}, nil
	}
	return Host{}, errors.New("invalid host config")
}

func removeComment(line string) string {
	if i := byteIndex(line, '#'); i >= 0 {
		// Discard comments.
		line = line[0:i]
	}
	return line
}

func byteIndex(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}
