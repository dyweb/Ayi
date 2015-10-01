package config

// TODO: ref https://golang.org/src/net/hosts.go

import (
	"fmt"
	"regexp"
	"bufio"
	"os"
	"log"
	"errors"
	"github.com/crackcomm/go-clitable"
)

type Host struct {
	// must use upper case ? no.
	ip   string
	name string
	line int
	// TODO: add line, useful for remove
}

func (host Host) Print() {
	table := clitable.New([]string{"ip", "host", "line"})
	table.AddRow(map[string]interface{}{
		"ip":host.ip,
		"host":host.name,
		"line":host.line,
	})
}

func PrintHosts(hosts []Host) {
	table := clitable.New([]string{"ip", "host", "line"})
	for i := 0; i < len(hosts); i++ {
		table.AddRow(map[string]interface{}{
			"ip": hosts[i].ip,
			"host":hosts[i].name,
			"line":hosts[i].line,
		})
	}
	table.Print()
}

func ParseHosts() []Host {
	// TODO: error handling
	hostsFile, _ := getHostFile()
	return parseHostsFile(hostsFile)
}

func AddDomainToLocalhost(domain string) (bool, error) {
	fmt.Println("Add localhost! ")
	return false, nil
}

func addHostToFile(hostsFile string, ip string, name string) (bool, error) {
	hosts := parseHostsFile(hostsFile)
	for i := 0; i < len(hosts); i++ {
		if hosts[i].name == name {
			return false, errors.New("name " + name + " already exists ")
		}
	}

	return true, nil
}

func removeHostFromFile(hostsFile string, name string) (bool, error) {
	hosts := parseHostsFile(hostsFile)
	for i := 0; i < len(hosts); i++ {
		if hosts[i].name == name {
			// TODO: real remove
			return true, nil
		}
	}
	return false, errors.New("name " + name + " does not exists ")
}

func getHostFile() (string, error) {
	// TODO: support for win
	return "/etc/hosts", nil
}

func parseHostsFile(hostsFile string) []Host {
	file, err := os.Open(hostsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hosts := make([]Host, 0)
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		host, err := parseHost(line, lineNumber)
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

func parseHost(s string, lineNumber int) (Host, error) {
	s = removeComment(s)
	r, _ := regexp.Compile("([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})\\s*(\\S*)\\s*$")
	m := r.FindStringSubmatch(s)
	if len(m) == 3 {
		return Host{ip: m[1], name: m[2], line:lineNumber}, nil
	}
	return Host{}, errors.New("invalid host config")
}

// copied from net/hosts.go
func removeComment(line string) string {
	if i := byteIndex(line, '#'); i >= 0 {
		// Discard comments.
		line = line[0:i]
	}
	return line
}

// copied from net/parser.go
func byteIndex(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}
