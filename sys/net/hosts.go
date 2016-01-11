package net

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/crackcomm/go-clitable"
	"github.com/dyweb/Ayi/util"
	"github.com/go-errors/errors"
)

// Host has ip, name, line
// some impl and util functions are from https://golang.org/src/net/hosts.go
type Host struct {
	// this not for external use
	ip   string
	name string
	line int
}

// Print print a host object as a table row
func (host Host) Print() {
	table := clitable.New([]string{"ip", "host", "line"})
	table.AddRow(map[string]interface{}{
		"ip":   host.ip,
		"host": host.name,
		"line": host.line,
	})
}

// PrintHosts print a list of host as table
func PrintHosts(hosts []Host) {
	table := clitable.New([]string{"ip", "host", "line"})
	for i := 0; i < len(hosts); i++ {
		table.AddRow(map[string]interface{}{
			"ip":   hosts[i].ip,
			"host": hosts[i].name,
			"line": hosts[i].line,
		})
	}
	table.Print()
}

// ParseHosts parse system host file
func ParseHosts() []Host {
	// TODO: error handling
	hostsFile, _ := getHostFile()
	return parseHostsFile(hostsFile)
}

// AddDomainToIP add a record of domain to system host file
func AddDomainToIP(domain string, ip string) (bool, error) {
	hostFile, _ := getHostFile()
	added, err := addHostToFile(hostFile, ip, domain)
	if added {
		return true, nil
	}
	return false, errors.Wrap(err, 1)
}

// RemoveDomain remove a record of domain form system host file
func RemoveDomain(domain string) (bool, error) {
	hostFile, _ := getHostFile()
	removed, err := removeHostFromFile(hostFile, domain)
	if removed {
		return true, nil
	}
	// TODO: 0 or 1
	return false, errors.Wrap(err, 1)
}

func getHostFile() (string, error) {
	// TODO: support for win
	return "/etc/hosts", nil
}

// TODO: support ipv6
func parseHostsFile(hostsFile string) []Host {
	file, err := os.Open(hostsFile)
	if err != nil {
		// TODO: return error instead of shutdown here
		log.Fatal(err)
	}
	defer file.Close()

	var hosts []Host
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
			// Just ignore it ....
		} else {
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
		return Host{ip: m[1], name: m[2], line: lineNumber}, nil
	}
	return Host{}, errors.New("invalid host config")
}

func addHostToFile(hostsFile string, ip string, name string) (bool, error) {
	hosts := parseHostsFile(hostsFile)
	for i := 0; i < len(hosts); i++ {
		if hosts[i].name == name {
			return false, errors.New("name " + name + " already exists ")
		}
	}
	file, err := os.OpenFile(hostsFile, os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close()
	if err != nil {
		return false, errors.Wrap(err, 1)
	}
	// TODO: need to check the error here
	file.WriteString(ip + "	" + name + "\n")
	file.Sync()
	return true, nil
}

func removeHostFromFile(hostsFile string, name string) (bool, error) {
	hosts := parseHostsFile(hostsFile)
	for i := 0; i < len(hosts); i++ {
		if hosts[i].name == name {
			err := util.RemoveLine(hostsFile, hosts[i].line)
			if err != nil {
				return false, errors.Wrap(err, 1)
			}
			return true, nil
		}
	}
	return false, errors.New("name " + name + " does not exists ")
}

// start of util functions

// copied from net/hosts.go
// TODO: this should be put into util too
func removeComment(line string) string {
	if i := util.ByteIndex(line, '#'); i >= 0 {
		// Discard comments.
		line = line[0:i]
	}
	return line
}

// end of util functions
